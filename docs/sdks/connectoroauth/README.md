# ConnectorOAuth

## Overview

### Available Operations

* [Authorize](#authorize) - Get OAuth authorization URL
* [HandleCallback](#handlecallback) - OAuth callback handler
* [~~ExchangeLegacyToken~~](#exchangelegacytoken) - Exchange Google authorization code for tokens :warning: **Deprecated**

## Authorize

Generate an OAuth authorization URL to start the OAuth flow.<br><br>
<b>Flow:</b><br>
<ol>
<li>Call this endpoint to get the authorization URL</li>
<li>Redirect user's browser to the URL</li>
<li>User authenticates with the provider</li>
<li>Provider redirects to callback with authorization code</li>
<li>Callback exchanges code for tokens automatically</li>
</ol>
<b>State Parameter:</b><br>
The response includes a <code>state</code> value that encodes the
connector ID. This is validated in the callback.


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

    res, err := s.ConnectorOAuth.Authorize(ctx, "<id>", nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.Object != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `connectorID`                                            | *string*                                                 | :heavy_check_mark:                                       | N/A                                                      |
| `instanceBaseURL`                                        | **string*                                                | :heavy_minus_sign:                                       | Base URL for self-hosted instances                       |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.GetOAuthAuthorizationURLResponse](../../models/operations/getoauthauthorizationurlresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## HandleCallback

Handle the OAuth callback from the identity provider.<br><br>
<b>Note:</b><br>
This endpoint is called by the OAuth provider after user authentication.
The state parameter contains the encoded connector ID.<br><br>
<b>Success:</b><br>
On success, tokens are stored and the connector becomes authenticated.
User is redirected to the frontend success page.<br><br>
<b>Error:</b><br>
If the provider returns an error (e.g., user denied access),
user is redirected with error information.


### Example Usage

<!-- UsageSnippet language="go" operationID="handleOAuthCallback" method="get" path="/connectors/oauth/callback" -->
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

    res, err := s.ConnectorOAuth.HandleCallback(ctx, nil, nil, nil, nil)
    if err != nil {
        log.Fatal(err)
    }
    if res != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `code`                                                   | **string*                                                | :heavy_minus_sign:                                       | Authorization code from provider                         |
| `state`                                                  | **string*                                                | :heavy_minus_sign:                                       | State parameter (contains connector ID)                  |
| `error_`                                                 | **string*                                                | :heavy_minus_sign:                                       | Error code if authorization failed                       |
| `instanceBaseURL`                                        | **string*                                                | :heavy_minus_sign:                                       | Base URL for redirect                                    |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.HandleOAuthCallbackResponse](../../models/operations/handleoauthcallbackresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## ~~ExchangeLegacyToken~~

<b>⚠️ Deprecated:</b> Legacy Google Workspace token exchange endpoint. Use the generic
OAuth flow via <code>/connectors/{connectorId}/oauth/authorize</code> instead.<br><br>
<b>Overview:</b><br>
Exchanges a Google OAuth authorization code for access and refresh tokens,
stores the credentials, and enables the Google Workspace connector.<br><br>
<b>What Happens:</b><br>
<ol>
<li>Retrieves Google Workspace OAuth config (client ID/secret)</li>
<li>Exchanges the authorization code for tokens via Google's token endpoint</li>
<li>Verifies the ID token</li>
<li>Stores access and refresh tokens in configuration manager</li>
<li>Creates or enables the Google Workspace connector</li>
<li>Publishes an AppEnabledEvent for the sync service</li>
</ol>
<b>Admin Only:</b><br>
Requires organization admin privileges.


> :warning: **DEPRECATED**: This will be removed in a future release, please migrate away from it as soon as possible.

### Example Usage

<!-- UsageSnippet language="go" operationID="getTokenFromCode" method="post" path="/connectors/getTokenFromCode" -->
```go
package main

import(
	"context"
	"os"
	"github.com/pipeshub-ai/pipeshub-sdk-go/models/components"
	pipeshub "github.com/pipeshub-ai/pipeshub-sdk-go"
	"github.com/pipeshub-ai/pipeshub-sdk-go/models/operations"
	"log"
)

func main() {
    ctx := context.Background()

    s := pipeshub.New(
        pipeshub.WithSecurity(components.Security{
            BearerAuth: pipeshub.Pointer(os.Getenv("PIPESHUB_BEARER_AUTH")),
        }),
    )

    res, err := s.ConnectorOAuth.ExchangeLegacyToken(ctx, operations.GetTokenFromCodeRequest{
        TempCode: "<value>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.TwoHundredApplicationJSONObject != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                | Type                                                                                     | Required                                                                                 | Description                                                                              |
| ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- |
| `ctx`                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                    | :heavy_check_mark:                                                                       | The context to use for the request.                                                      |
| `request`                                                                                | [operations.GetTokenFromCodeRequest](../../models/operations/gettokenfromcoderequest.md) | :heavy_check_mark:                                                                       | The request object to use for the request.                                               |
| `opts`                                                                                   | [][operations.Option](../../models/operations/option.md)                                 | :heavy_minus_sign:                                                                       | The options for this request.                                                            |

### Response

**[*operations.GetTokenFromCodeResponse](../../models/operations/gettokenfromcoderesponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |