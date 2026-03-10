# AuthConfigurations

## Overview

### Available Operations

* [SetOAuth](#setoauth) - Configure generic OAuth provider

## SetOAuth

Set up a custom OAuth 2.0 authentication provider.

### Example Usage

<!-- UsageSnippet language="go" operationID="setOAuthConfig" method="post" path="/configurationManager/authConfig/oauth" -->
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

    res, err := s.AuthConfigurations.SetOAuth(ctx, components.GenericOAuthConfig{
        ProviderName: pipeshub.Pointer("Custom OAuth Provider"),
        Scope: pipeshub.Pointer("openid profile email"),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                      | Type                                                                           | Required                                                                       | Description                                                                    |
| ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ |
| `ctx`                                                                          | [context.Context](https://pkg.go.dev/context#Context)                          | :heavy_check_mark:                                                             | The context to use for the request.                                            |
| `request`                                                                      | [components.GenericOAuthConfig](../../models/components/genericoauthconfig.md) | :heavy_check_mark:                                                             | The request object to use for the request.                                     |
| `opts`                                                                         | [][operations.Option](../../models/operations/option.md)                       | :heavy_minus_sign:                                                             | The options for this request.                                                  |

### Response

**[*operations.SetOAuthConfigResponse](../../models/operations/setoauthconfigresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |