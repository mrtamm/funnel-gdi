package server

import (
	"encoding/base64"
	"net/http"
	"strings"

	"github.com/ohsu-comp-bio/funnel/config"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

type Authentication struct {
	basic  map[string]string
	admins map[string]bool
	oidc   *OidcConfig
}

// Extracted info about the current user, which is exposed through Context.
type UserInfo struct {
	// Public users are non-authenticated, in case Funnel configuration does
	// not require OIDC nor Basic authentication.
	IsPublic bool
	// Administrator is a Basic-authentication user with `Admin: true` property
	// in the configuration file.
	IsAdmin bool
	// Username of an authenticated user (subject field from JWT).
	Username string
	// In case of OIDC authentication, the provided Bearer token, which can be
	// used when requesting task input data.
	Token string
}

// Context key type for storing UserInfo.
// Note: UserInfo is not in the context when the system internally requests data.
type userInfoContextKey string

var (
	errMissingMetadata    = status.Errorf(codes.InvalidArgument, "Missing metadata in the context")
	errTokenRequired      = status.Errorf(codes.Unauthenticated, "Basic/Bearer authorization token missing")
	errInvalidBasicToken  = status.Errorf(codes.Unauthenticated, "Basic-authentication failed")
	errInvalidBearerToken = status.Errorf(codes.Unauthenticated, "Bearer authorization token not accepted")
	publicUserInfo        = UserInfo{IsPublic: true, IsAdmin: false, Username: ""}
	UserInfoKey           = userInfoContextKey("user-info")
)

func NewAuthentication(creds []config.BasicCredential, oidc config.OidcAuth) *Authentication {
	basicCreds := make(map[string]string)
	adminUsers := make(map[string]bool)

	for _, cred := range creds {
		credBytes := []byte(cred.User + ":" + cred.Password)
		fullValue := "Basic " + base64.StdEncoding.EncodeToString(credBytes)
		basicCreds[fullValue] = cred.User
		if cred.Admin {
			adminUsers[cred.User] = true
		}
	}

	return &Authentication{
		basic:  basicCreds,
		admins: adminUsers,
		oidc:   initOidcConfig(oidc),
	}
}

// Return a new gRPC interceptor function that authorizes RPCs.
func (a *Authentication) Interceptor(
	ctx context.Context,
	req interface{},
	info *grpc.UnaryServerInfo,
	handler grpc.UnaryHandler) (interface{}, error) {

	// Case when authentication is not required:
	if len(a.basic) == 0 && a.oidc == nil {
		ctx = context.WithValue(ctx, UserInfoKey, &publicUserInfo)
		return handler(ctx, req)
	}

	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, errMissingMetadata
	}

	values := md["authorization"]
	if len(values) == 0 {
		return nil, errTokenRequired
	}

	authorized := false
	authErr := errTokenRequired

	if strings.HasPrefix(values[0], "Basic ") {
		authErr = errInvalidBasicToken
		headerValue := values[0]
		username := a.basic[headerValue]
		isAdmin := a.admins[username]
		authorized = username != ""

		if authorized {
			ctx = context.WithValue(ctx, UserInfoKey, &UserInfo{Username: username, IsAdmin: isAdmin})
		}
	} else if a.oidc != nil && strings.HasPrefix(values[0], "Bearer ") {
		authErr = errInvalidBearerToken
		jwtString := strings.TrimPrefix(values[0], "Bearer ")
		subject := a.oidc.ParseJwtSubject(jwtString)
		authorized = subject != ""
		if authorized {
			ctx = context.WithValue(ctx, UserInfoKey, &UserInfo{Username: subject, Token: jwtString})
		}
	}

	if !authorized {
		return nil, authErr
	}

	return handler(ctx, req)
}

// HTTP request handler for the /login endpoint. Initiates user authentication
// flow based on the configuration (OIDC, Basic, none).
func (a *Authentication) LoginHandler(w http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodGet {
		http.Error(w, "Only GET method is supported.", http.StatusMethodNotAllowed)
	} else if a.oidc != nil {
		a.oidc.HandleAuthCode(w, req)
	} else if len(a.basic) > 0 {
		a.handleBasicAuth(w, req)
	} else {
		http.Redirect(w, req, "/", http.StatusSeeOther)
	}
}

// HTTP request handler for the /login/token endpoint. In case of OIDC enabled,
// prints the JWT from the sent cookie. In all other cases, an empty HTTP 200
// response.
func (a *Authentication) EchoTokenHandler(w http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodGet {
		http.Error(w, "Only GET method is supported.", http.StatusMethodNotAllowed)
	} else if a.oidc != nil {
		a.oidc.EchoTokenHandler(w, req)
	} else {
		w.Header().Set("Content-Type", "text/plain; charset=utf-8")
		w.Header().Set("Content-Length", "0")
		w.WriteHeader(http.StatusOK)
	}
}

func (a *Authentication) handleBasicAuth(w http.ResponseWriter, req *http.Request) {
	// Check if provided value in the header is valid:
	if a.basic[req.Header.Get("Authorization")] == "" {
		http.Redirect(w, req, "/", http.StatusSeeOther)
	} else {
		w.Header().Set("WWW-Authenticate", "Basic realm=Funnel")
		msg := "User authentication is required (Basic authentication with " +
			"username and password)"
		http.Error(w, msg, http.StatusUnauthorized)
	}
}

// Reports whether current user can access data with the specified owner.
// Admin-user can access everything.
// Non-authenticated users can access data without ownership.
// Authenticated users can access only the data that was created by the same user.
func (u *UserInfo) IsAccessible(dataOwner string) bool {
	result := u.IsAdmin || u.IsPublic && dataOwner == "" || !u.IsPublic && dataOwner == u.Username
	return result
}
