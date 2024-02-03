---
title: OAuth2
menu:
  main:
    parent: Security
    weight: 10
---
# OAuth2

By default, a Funnel server allows open access to its API endpoints, but in
addition to Basic authentication it can also be configured to require a valid
JWT in the request.

Funnel itself does not redirect users to perform the login.
It just validates that the presented token is issued by a trusted service
(specified in the configuration file) and the token has not expired.

Optionally, Funnel can also validate the scope and audience claims to contain
specific values.

To enable JWT authentication, specify `OidcAuth` section in your config file:

```yaml
Server:
  OidcAuth:
    # URL of the OIDC service configuration:
    ServiceConfigUrl: "https://my.oidc.service/.well-known/openid-configuration"
    # Optional: if specified, this scope value must be in the token:
    RequireScope: funnel-id
    # Optional: if specified, this audience value must be in the token:
    RequireAudience: tes-api
```

Make sure to properly protect the configuration file so that it's not readable 
by everyone:

```bash
$ chmod 600 funnel.config.yml
```

To use the token, set the `FUNNEL_SERVER_BEARER` environment variable:

```bash
$ export FUNNEL_SERVER_BEARER=eyJraWQiOiJyc2ExIiwi...
$ funnel task list
```
