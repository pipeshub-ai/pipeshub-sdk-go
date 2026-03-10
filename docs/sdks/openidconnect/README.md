# OpenIDConnect

## Overview

### Available Operations

* [UserInfo](#userinfo) - Get authenticated user information
* [OauthAuthorizationServerMetadata](#oauthauthorizationservermetadata) - OAuth 2.0 Authorization Server Metadata
* [Jwks](#jwks) - JSON Web Key Set
* [GetProtectedResourceMetadata](#getprotectedresourcemetadata) - OAuth Protected Resource Metadata
* [GetConfiguration](#getconfiguration) - OpenID Connect Discovery

## UserInfo

OpenID Connect UserInfo Endpoint.
<br><br>
Returns claims about the authenticated user. Requires a valid access token
with the `openid` scope.
<br><br>
<b>Available Claims:</b><br>
- `user_id` - User identifier<br>
- `name`, `given_name`, `family_name` - Name claims (with `profile` scope)<br>
- `email`, `email_verified` - Email claims (with `email` scope)
<br><br>
<b>Authentication:</b><br>
Pass the access token as a Bearer token: `Authorization: Bearer {access_token}`


### Example Usage

<!-- UsageSnippet language="go" operationID="oauthUserInfo" method="get" path="/oauth2/userinfo" -->
```go
package main

import(
	"context"
	"os"
	"github.com/pipeshub-ai/pipeshub-sdk-go/models/components"
	pipeshub "github.com/pipeshub-ai/pipeshub-sdk-go"
	"log"
)

func main() {
    ctx := context.Background()

    s := pipeshub.New(
        pipeshub.WithSecurity(components.Security{
            BearerAuth: pipeshub.Pointer(os.Getenv("PIPESHUB_BEARER_AUTH")),
        }),
    )

    res, err := s.OpenIDConnect.UserInfo(ctx)
    if err != nil {
        log.Fatal(err)
    }
    if res.OAuthUserInfoResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.OauthUserInfoResponse](../../models/operations/oauthuserinforesponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## OauthAuthorizationServerMetadata

OAuth 2.0 Authorization Server Metadata Endpoint (RFC 8414).
<br><br>
Returns the same metadata as the OpenID Connect Discovery endpoint
but at the RFC 8414 standard path. MCP clients like Claude Code use
this endpoint for discovery instead of openid-configuration.
<br><br>
<b>Note:</b> This endpoint does not require authentication.


### Example Usage

<!-- UsageSnippet language="go" operationID="oauthAuthorizationServerMetadata" method="get" path="/.well-known/oauth-authorization-server" -->
```go
package main

import(
	"context"
	pipeshub "github.com/pipeshub-ai/pipeshub-sdk-go"
	"log"
)

func main() {
    ctx := context.Background()

    s := pipeshub.New()

    res, err := s.OpenIDConnect.OauthAuthorizationServerMetadata(ctx)
    if err != nil {
        log.Fatal(err)
    }
    if res.OpenIDConfiguration != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.OauthAuthorizationServerMetadataResponse](../../models/operations/oauthauthorizationservermetadataresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## Jwks

JSON Web Key Set Endpoint (RFC 7517).
<br><br>
Returns the public keys used to verify JWT signatures for ID tokens
and access tokens.
<br><br>
<b>Use Cases:</b><br>
- Verifying ID token signatures<br>
- Verifying access token signatures (if JWT-based)
<br><br>
<b>Note:</b><br>
- For HS256 (symmetric) signing, this returns empty keys<br>
- For RS256 (asymmetric) signing, returns public RSA keys<br>
- Keys should be cached with appropriate TTL


### Example Usage

<!-- UsageSnippet language="go" operationID="jwks" method="get" path="/.well-known/jwks.json" -->
```go
package main

import(
	"context"
	pipeshub "github.com/pipeshub-ai/pipeshub-sdk-go"
	"log"
)

func main() {
    ctx := context.Background()

    s := pipeshub.New()

    res, err := s.OpenIDConnect.Jwks(ctx)
    if err != nil {
        log.Fatal(err)
    }
    if res.Jwks != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.JwksResponse](../../models/operations/jwksresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## GetProtectedResourceMetadata

OAuth Protected Resource Metadata Endpoint (RFC 9728).
<br><br>
Returns metadata about the protected resource including the resource
identifier, authorization servers, supported scopes, and bearer token methods.
<br><br>
<b>Use Cases:</b><br>
- Discovering which authorization server to use for this resource<br>
- Determining supported scopes and bearer token methods<br>
- MCP client auto-configuration
<br><br>
<b>Note:</b> This endpoint does not require authentication.


### Example Usage

<!-- UsageSnippet language="go" operationID="oauthProtectedResource" method="get" path="/.well-known/oauth-protected-resource/mcp" -->
```go
package main

import(
	"context"
	pipeshub "github.com/pipeshub-ai/pipeshub-sdk-go"
	"log"
)

func main() {
    ctx := context.Background()

    s := pipeshub.New()

    res, err := s.OpenIDConnect.GetProtectedResourceMetadata(ctx)
    if err != nil {
        log.Fatal(err)
    }
    if res.OAuthProtectedResourceMetadata != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.OauthProtectedResourceResponse](../../models/operations/oauthprotectedresourceresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## GetConfiguration

OpenID Connect Discovery Endpoint (RFC 8414).
<br><br>
Returns metadata about the OAuth/OIDC authorization server including
endpoint URLs, supported features, and capabilities.
<br><br>
<b>Use Cases:</b><br>
- Automatic client configuration<br>
- Discovering supported features<br>
- Getting endpoint URLs without hardcoding
<br><br>
<b>Note:</b> This endpoint does not require authentication.


### Example Usage

<!-- UsageSnippet language="go" operationID="openidConfiguration" method="get" path="/.well-known/openid-configuration" -->
```go
package main

import(
	"context"
	pipeshub "github.com/pipeshub-ai/pipeshub-sdk-go"
	"log"
)

func main() {
    ctx := context.Background()

    s := pipeshub.New()

    res, err := s.OpenIDConnect.GetConfiguration(ctx)
    if err != nil {
        log.Fatal(err)
    }
    if res.OpenIDConfiguration != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.OpenidConfigurationResponse](../../models/operations/openidconfigurationresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |