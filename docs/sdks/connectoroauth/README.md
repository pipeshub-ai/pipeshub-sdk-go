# ConnectorOAuth

## Overview

OAuth 2.0 authorization flow for connectors requiring user consent

### Available Operations

* [GetOAuthAuthorizationURL](#getoauthauthorizationurl) - Get OAuth authorization URL
* [HandleOAuthCallback](#handleoauthcallback) - OAuth callback handler

## GetOAuthAuthorizationURL

Start the OAuth sign-in for a connector instance.<br><br>
<b>Flow:</b><br>
<ol>
<li>Call this endpoint to get <code>authorizationUrl</code> and <code>state</code></li>
<li>A person opens the URL and gives consent at the provider</li>
<li>The provider redirects to <code>{baseUrl}/{redirectUri}</code> with <code>code</code> and <code>state</code></li>
<li>Pass both to <code>GET /connectors/oauth/callback</code> with a bearer token</li>
</ol>
The redirect URI is fixed when the connector is created (its <code>baseUrl</code>) and must be
registered with the provider.


### Example Usage

<!-- UsageSnippet language="go" operationID="getOAuthAuthorizationUrl" method="get" path="/connectors/{connectorId}/oauth/authorize" -->
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

    res, err := s.ConnectorOAuth.GetOAuthAuthorizationURL(ctx, "<id>", nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.Object != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                                                                                                                                          | Type                                                                                                                                                                                                                                                               | Required                                                                                                                                                                                                                                                           | Description                                                                                                                                                                                                                                                        |
| ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                                                                                                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                                                                                                                                                                                              | :heavy_check_mark:                                                                                                                                                                                                                                                 | The context to use for the request.                                                                                                                                                                                                                                |
| `connectorID`                                                                                                                                                                                                                                                      | `string`                                                                                                                                                                                                                                                           | :heavy_check_mark:                                                                                                                                                                                                                                                 | Connector instance ID                                                                                                                                                                                                                                              |
| `desktop`                                                                                                                                                                                                                                                          | `*string`                                                                                                                                                                                                                                                          | :heavy_minus_sign:                                                                                                                                                                                                                                                 | `true` when the sign-in starts in the desktop app: the `state` sent to the provider is<br/>marked so the web callback page hands `code` and `state` to the app over a `pipeshub://`<br/>link instead of finishing the sign-in in the browser. Any other value is ignored.<br/> |
| `opts`                                                                                                                                                                                                                                                             | [][operations.Option](../../models/operations/option.md)                                                                                                                                                                                                           | :heavy_minus_sign:                                                                                                                                                                                                                                                 | The options for this request.                                                                                                                                                                                                                                      |

### Response

**[*operations.GetOAuthAuthorizationURLResponse](../../models/operations/getoauthauthorizationurlresponse.md), error**

### Errors

| Error Type              | Status Code             | Content Type            |
| ----------------------- | ----------------------- | ----------------------- |
| apierrors.ErrorResponse | 400, 401, 403, 404      | application/json        |
| apierrors.ErrorResponse | 500                     | application/json        |
| apierrors.APIError      | 4XX, 5XX                | \*/\*                   |

## HandleOAuthCallback

Finish the OAuth sign-in: exchange the provider's <code>code</code> for tokens and store
them. Call it with the <code>code</code> and <code>state</code> the provider sent to the redirect URI,
as the connector's creator or an admin.<br><br>
<b>Result:</b><br>
Always JSON. <code>success</code> is true when the connector is signed in; otherwise
<code>error</code> and <code>errorMessage</code> say why. Provider tokens are never returned.


### Example Usage

<!-- UsageSnippet language="go" operationID="handleOAuthCallback" method="get" path="/connectors/oauth/callback" -->
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

    res, err := s.ConnectorOAuth.HandleOAuthCallback(ctx, "<value>", "Arizona", nil, nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.Object != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                           | Type                                                                                                                                | Required                                                                                                                            | Description                                                                                                                         |
| ----------------------------------------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                               | [context.Context](https://pkg.go.dev/context#Context)                                                                               | :heavy_check_mark:                                                                                                                  | The context to use for the request.                                                                                                 |
| `code`                                                                                                                              | `string`                                                                                                                            | :heavy_check_mark:                                                                                                                  | Authorization code from provider. A missing, empty or blank value is<br/>refused with `400`, even when the provider only sent `error`.<br/> |
| `state`                                                                                                                             | `string`                                                                                                                            | :heavy_check_mark:                                                                                                                  | State parameter (contains connector ID). A missing, empty or blank value is refused with `400`.                                     |
| `error_`                                                                                                                            | `*string`                                                                                                                           | :heavy_minus_sign:                                                                                                                  | Error code if authorization failed                                                                                                  |
| `baseURL_`                                                                                                                          | `*string`                                                                                                                           | :heavy_minus_sign:                                                                                                                  | Origin prefixed to the returned `redirectUrl`                                                                                       |
| `opts`                                                                                                                              | [][operations.Option](../../models/operations/option.md)                                                                            | :heavy_minus_sign:                                                                                                                  | The options for this request.                                                                                                       |

### Response

**[*operations.HandleOAuthCallbackResponse](../../models/operations/handleoauthcallbackresponse.md), error**

### Errors

| Error Type              | Status Code             | Content Type            |
| ----------------------- | ----------------------- | ----------------------- |
| apierrors.ErrorResponse | 400, 401, 403           | application/json        |
| apierrors.ErrorResponse | 500                     | application/json        |
| apierrors.APIError      | 4XX, 5XX                | \*/\*                   |