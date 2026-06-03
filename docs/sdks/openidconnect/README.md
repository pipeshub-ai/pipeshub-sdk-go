# OpenIDConnect

## Overview

OpenID Connect 1.0 endpoints for identity federation and discovery.

**Discovery:**
- `/.well-known/openid-configuration` - Authorization server metadata
- `/.well-known/oauth-authorization-server` - Authorization server metadata (RFC 8414)
- `/.well-known/oauth-protected-resource/mcp` - Protected resource metadata (RFC 9728)
- `/.well-known/jwks.json` - Public keys for token verification

**UserInfo:**
- `/oauth2/userinfo` - Get authenticated user's profile information

**Supported Claims:**
- `user_id` - User identifier
- `email`, `email_verified` - Email information
- `name`, `given_name`, `family_name` - Name information


### Available Operations

* [OauthUserInfo](#oauthuserinfo) - Get authenticated user information

## OauthUserInfo

OpenID Connect UserInfo Endpoint.

Returns claims about the authenticated user. Requires a valid access token
with the `openid` scope.

**Available Claims:**
- `user_id` - User identifier
- `name`, `given_name`, `family_name` - Name claims (with `profile` scope)
- `email`, `email_verified` - Email claims (with `email` scope)

**Authentication:**
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

    res, err := s.OpenIDConnect.OauthUserInfo(ctx)
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