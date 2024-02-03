package server

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
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
	if config.ServiceConfigUrl == "" {
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
		fmt.Printf("[ERROR] Failed to parse the configuration (JSON) of the "+
			"OIDC service: %s\n", err)
		os.Exit(1)
	}

	c.initJwks()
}

func (c *OidcConfig) initJwks() {
	jwksUrl := c.remote.JwksURI
	ctx := context.Background()

	// Define JWKS cache:
	c.jwks = *jwk.NewCache(ctx)
	c.jwks.Register(jwksUrl, jwk.WithMinRefreshInterval(15*time.Minute))

	// Init JWKS cache:
	ctx2, _ := context.WithTimeout(ctx, 10*time.Second)
	_, err := c.jwks.Refresh(ctx2, jwksUrl)

	if err != nil {
		fmt.Printf("[ERROR] Failed to fetch JWKS (%s) of the OIDC service "+
			"(%s): %s\n", jwksUrl, c.local.ServiceConfigUrl, err)
		os.Exit(1)
	}
}

func (c *OidcConfig) ParseJwt(jwtString string) *jwt.Token {
	keySet, err := c.jwks.Get(context.Background(), c.remote.JwksURI)
	if err != nil {
		fmt.Printf("[WARN] Failed to retrieve JWKS key-set: %s", err)
		return nil
	}

	token, err := jwt.ParseString(
		jwtString,
		jwt.WithVerify(true),
		jwt.WithKeySet(keySet),
		jwt.WithIssuer(c.remote.Issuer),
	)

	if err != nil {
		fmt.Printf("[WARN] Provided JWT is not valid: %s.\n", err)
		return nil
	}

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
			fmt.Printf("[WARN] Audience [%s] not found in %v.",
				c.local.RequireAudience, token.Audience())
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
			fmt.Printf("[WARN] Scope [%s] not found in [%s]",
				c.local.RequireScope, value)
			return nil
		}
	}

	return &token
}

func validateUrl(providedUrl string) *url.URL {
	parsedUrl, err := url.ParseRequestURI(providedUrl)
	if err != nil {
		fmt.Printf("[ERROR] OIDC configuration URL (%s) could not be "+
			"parsed: %s\n", parsedUrl, err)
		os.Exit(1)
	} else if parsedUrl.Scheme == "" || parsedUrl.Host == "" {
		fmt.Printf("[ERROR] OIDC configuration URL (%s) is not absolute.",
			parsedUrl)
		os.Exit(1)
	}
	return parsedUrl
}

func fetchJson(url *url.URL) []byte {
	res, err := http.Get(url.String())

	if err != nil {
		fmt.Printf("[ERROR] OIDC service configuration (%s) could not be "+
			"loaded: %s.\n", url.String(), err)
		os.Exit(1)
	} else if res.StatusCode != 200 {
		fmt.Printf("[ERROR] OIDC service configuration (%s) could not be "+
			"loaded (HTTP response status: %d).", url.String(), res.StatusCode)
		os.Exit(1)
	} else if res.Body == nil {
		fmt.Printf("[ERROR] OIDC service configuration (%s) could not be "+
			"loaded (empty response).\n", url.String())
		os.Exit(1)
	}

	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Printf("[ERROR] Failed to read the body of the OIDC "+
			"configuration (%s) response: %s\n", url.String(), err)
		os.Exit(1)
	}

	return body
}
