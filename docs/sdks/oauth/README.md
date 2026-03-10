# Oauth

## Overview

### Available Operations

* [ExchangeCode](#exchangecode) - Exchange OAuth authorization code for tokens

## ExchangeCode

Exchange an OAuth authorization code for access and ID tokens.
Used after the OAuth authorization flow redirects back to the application.
<br><br>
<b>Supported Providers:</b><br>
- Generic OAuth 2.0 providers configured in org settings
<br><br>
<b>Flow:</b><br>
1. User is redirected to OAuth provider's authorization URL<br>
2. User authorizes and is redirected back with a code<br>
3. This endpoint exchanges the code for tokens<br>
4. Tokens can then be used with the <code>/userAccount/authenticate</code> endpoint


### Example Usage

<!-- UsageSnippet language="go" operationID="exchangeOAuthCode" method="post" path="/userAccount/oauth/exchange" -->
```go
package main

import(
	"context"
	pipeshub "github.com/pipeshub-ai/pipeshub-sdk-go"
	"github.com/pipeshub-ai/pipeshub-sdk-go/models/components"
	"log"
)

func main() {
    ctx := context.Background()

    s := pipeshub.New()

    res, err := s.Oauth.ExchangeCode(ctx, components.OAuthExchangeRequest{
        Code: "<value>",
        Email: "Jason2@gmail.com",
        Provider: "<value>",
        RedirectURI: "https://enlightened-developing.info",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.OAuthExchangeResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                          | Type                                                                               | Required                                                                           | Description                                                                        |
| ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- |
| `ctx`                                                                              | [context.Context](https://pkg.go.dev/context#Context)                              | :heavy_check_mark:                                                                 | The context to use for the request.                                                |
| `request`                                                                          | [components.OAuthExchangeRequest](../../models/components/oauthexchangerequest.md) | :heavy_check_mark:                                                                 | The request object to use for the request.                                         |
| `opts`                                                                             | [][operations.Option](../../models/operations/option.md)                           | :heavy_minus_sign:                                                                 | The options for this request.                                                      |

### Response

**[*operations.ExchangeOAuthCodeResponse](../../models/operations/exchangeoauthcoderesponse.md), error**

### Errors

| Error Type          | Status Code         | Content Type        |
| ------------------- | ------------------- | ------------------- |
| apierrors.AuthError | 400                 | application/json    |
| apierrors.APIError  | 4XX, 5XX            | \*/\*               |