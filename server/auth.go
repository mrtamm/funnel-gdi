package server

import (
	"encoding/base64"
	"strings"

	"github.com/ohsu-comp-bio/funnel/config"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

var (
	errMissingMetadata    = status.Errorf(codes.InvalidArgument, "Missing metadata in the context")
	errTokenRequired      = status.Errorf(codes.Unauthenticated, "Bearer/Basic authorization token missing")
	errInvalidBasicToken  = status.Errorf(codes.Unauthenticated, "Invalid Basic authorization token")
	errInvalidBearerToken = status.Errorf(codes.Unauthenticated, "Invalid Bearer authorization token")
)

// Return a new interceptor function that authorizes RPCs.
func newAuthInterceptor(creds []config.BasicCredential, oidc config.OidcAuth) grpc.UnaryServerInterceptor {
	basicCreds := initBasicCredsMap(creds)
	oidcConfig := initOidcConfig(oidc)
	requireAuth := len(basicCreds) > 0 || oidcConfig != nil

	// Return a function that is the interceptor.
	return func(
		ctx context.Context,
		req interface{},
		info *grpc.UnaryServerInfo,
		handler grpc.UnaryHandler) (interface{}, error) {

		if !requireAuth {
			return handler(ctx, req)
		}

		md, ok := metadata.FromIncomingContext(ctx)
		if !ok {
			return nil, errMissingMetadata
		}

		values := md["authorization"]
		if len(values) < 1 {
			return nil, errTokenRequired
		}

		authorized := false
		auth_err := errTokenRequired

		if strings.HasPrefix(values[0], "Basic ") {
			auth_err = errInvalidBasicToken
			authorized = basicCreds[values[0]]
		} else if oidcConfig != nil && strings.HasPrefix(values[0], "Bearer ") {
			auth_err = errInvalidBearerToken
			jwtString := strings.TrimPrefix(values[0], "Bearer ")
			jwt := oidcConfig.ParseJwt(jwtString)
			authorized = jwt != nil
		}

		if !authorized {
			return nil, auth_err
		}

		return handler(ctx, req)
	}
}

func initBasicCredsMap(creds []config.BasicCredential) map[string]bool {
	basicCreds := make(map[string]bool)
	for _, cred := range creds {
		credBytes := []byte(cred.User + ":" + cred.Password)
		fullValue := "Basic " + base64.StdEncoding.EncodeToString(credBytes)
		basicCreds[fullValue] = true
	}

	return basicCreds
}
