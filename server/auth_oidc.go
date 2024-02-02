package server

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/lestrrat-go/jwx/v2/jwk"
	"github.com/lestrrat-go/jwx/v2/jwt"
	"github.com/ohsu-comp-bio/funnel/config"
	"golang.org/x/net/context"
)

// JSON structure of the OIDC configuration (only some fields)
type OidcRemoteConfig struct {
	Issuer           string `json:"issuer"`
	UserinfoEndpoint string `json:"userinfo_endpoint"`
	JwksURI          string `json:"jwks_uri"`
}

// OIDC configuration structure used for validating input from request.
type OidcConfig struct {
	local  config.OidcAuth
	remote OidcRemoteConfig
	jwks   jwk.Cache
}

func initOidcConfig(config config.OidcAuth) *OidcConfig {
	if config.ServiceConfigUrl == "" ||
		config.ClientId == "" ||
		config.ClientSecret == "" {
		return nil
	}

	result := OidcConfig{local: config}
	result.initConfig()
	return &result
}

func (c *OidcConfig) initConfig() {
	c.remote = OidcRemoteConfig{}
	parsedUrl := validateUrl(c.local.ServiceConfigUrl)
	err := json.Unmarshal(fetchJson(parsedUrl), &c.remote)
	if err != nil {
		log.Fatalf("Failed to parse the configuration (JSON) of the OIDC "+
			"service: %s", err)
	}

	c.initJwks()
}

func (c *OidcConfig) initJwks() {
	jwksUrl := c.remote.JwksURI
	ctx := context.Background()

	// Define JWKS cache:
	c.jwks = *jwk.NewCache(ctx)
	c.jwks.Register(jwksUrl, jwk.WithMinRefreshInterval(1*time.Hour))

	// Init JWKS cache:
	ctx2, _ := context.WithTimeout(ctx, 10*time.Second)
	_, err := c.jwks.Refresh(ctx2, jwksUrl)

	if err != nil {
		log.Fatalf("Failed to fetch JWKS (%s) of the OIDC service (%s).",
			jwksUrl, c.local.ServiceConfigUrl, err)
	}
}

func (c *OidcConfig) ParseJwt(jwtString string) *jwt.Token {
	keySet, err := c.jwks.Get(context.Background(), c.remote.JwksURI)
	if err != nil {
		log.Println("Failed to retrieve JWKS key-set.", err)
		return nil
	}

	token, err := jwt.ParseString(
		jwtString,
		jwt.WithVerify(true),
		jwt.WithKeySet(keySet),
		jwt.WithIssuer(c.remote.Issuer),
	)

	// If audience is required, it must be in the token.
	if c.local.RequireAudience != "" {
		found := false
		for _, value := range token.Audience() {
			if value == c.local.RequireAudience {
				found = true
				break
			}
		}
		if !found {
			return nil
		}
	}

	// If scope is required, it must be in the token.
	if c.local.RequireScope != "" {
		value, found := token.Get("scope")
		if found {
			found = false
			for _, value := range strings.Split(value.(string), " ") {
				if value == c.local.RequireScope {
					found = true
					break
				}
			}
		}
		if !found {
			return nil
		}
	}

	return &token
}

func validateUrl(providedUrl string) *url.URL {
	parsedUrl, err := url.ParseRequestURI(providedUrl)
	if err != nil {
		log.Fatalf("OIDC configuration URL (%s) could not be parsed.", parsedUrl, err)
	} else if parsedUrl.Scheme == "" || parsedUrl.Host == "" {
		log.Fatalf("OIDC configuration URL (%s) is not absolute.", parsedUrl)
	}
	return parsedUrl
}

func fetchJson(url *url.URL) []byte {
	res, err := http.Get(url.String())

	if err != nil {
		log.Fatal("OIDC service configuration could not be loaded", err)
	} else if res.StatusCode != 200 {
		log.Fatalf("OIDC service configuration could not be loaded (HTTP "+
			" response status: %d)", res.StatusCode)
	} else if res.Body == nil {
		log.Fatal("OIDC service configuration could not be loaded (empty " +
			"response)")
	}

	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		log.Fatal("Failed to read the body of the OIDC configuration response", err)
	}

	return body
}
