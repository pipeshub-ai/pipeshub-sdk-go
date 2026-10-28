# WebSearch

## Overview

Manage web search providers (DuckDuckGo, Serper, Tavily, Exa) and settings for internet search.

### Available Operations

* [GetWebSearchProviders](#getwebsearchproviders) - Get all web search providers

## GetWebSearchProviders

Retrieve all configured web search providers and current web search settings.

**Authentication:** Session JWT or OAuth 2.0 access token via `Authorization: Bearer`.
OAuth tokens must include the `config:read` scope. Admin role is not required.

**API keys:** for anyone who isn't an org admin, each provider's `configuration.apiKey`
comes back as the placeholder `****************`, and so does `configuration.baseUrl`
(admins always get the stored address). Admins get the stored key, unless the
server hides secrets from everyone (`HIDE_SECRET_CONFIG=true`). When updating a provider,
sending the placeholder back keeps the stored key.

The built-in `duckduckgo` entry (never stored, empty `configuration`) is always listed
first; it is the default when no stored provider is. With nothing stored, `message` is
`No web search providers found` and the settings are the defaults. Of the query
string only `reveal` is read.


### Example Usage

<!-- UsageSnippet language="go" operationID="getWebSearchProviders" method="get" path="/configurationManager/web-search" -->
```go
package main

import(
	"context"
	"os"
	"github.com/pipeshub-ai/pipeshub-sdk-go/v2/models/components"
	pipeshub "github.com/pipeshub-ai/pipeshub-sdk-go/v2"
	"log"
)

func main() {
    ctx := context.Background()

    s := pipeshub.New(
        pipeshub.WithSecurity(components.Security{
            BearerAuth: pipeshub.Pointer(os.Getenv("PIPESHUB_BEARER_AUTH")),
        }),
    )

    res, err := s.WebSearch.GetWebSearchProviders(ctx, nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.WebSearchProvidersResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                                                                                                                         | Type                                                                                                                                                                                                                                              | Required                                                                                                                                                                                                                                          | Description                                                                                                                                                                                                                                       |
| ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                                                                                                                             | [context.Context](https://pkg.go.dev/context#Context)                                                                                                                                                                                             | :heavy_check_mark:                                                                                                                                                                                                                                | The context to use for the request.                                                                                                                                                                                                               |
| `reveal`                                                                                                                                                                                                                                          | [*components.RevealStoredSecrets](../../models/components/revealstoredsecrets.md)                                                                                                                                                                 | :heavy_minus_sign:                                                                                                                                                                                                                                | Set to `true` to receive stored secrets instead of placeholders. Honoured only for an<br/>admin's own configuration where the deployment allows it (see `GET<br/>/configurationManager/secretReveal`); ignored for OAuth-app tokens and any other value.<br/> |
| `opts`                                                                                                                                                                                                                                            | [][operations.Option](../../models/operations/option.md)                                                                                                                                                                                          | :heavy_minus_sign:                                                                                                                                                                                                                                | The options for this request.                                                                                                                                                                                                                     |

### Response

**[*operations.GetWebSearchProvidersResponse](../../models/operations/getwebsearchprovidersresponse.md), error**

### Errors

| Error Type              | Status Code             | Content Type            |
| ----------------------- | ----------------------- | ----------------------- |
| apierrors.ErrorResponse | 400, 401, 403           | application/json        |
| apierrors.ErrorResponse | 500                     | application/json        |
| apierrors.APIError      | 4XX, 5XX                | \*/\*                   |