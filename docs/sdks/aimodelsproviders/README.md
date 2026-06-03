# AIModelsProviders

## Overview

Manage individual AI model providers - add, update, delete, and set defaults.

### Available Operations

* [GetAvailableModelsByType](#getavailablemodelsbytype) - Get available models by type

## GetAvailableModelsByType

Returns a **flattened list** of individual AI models of the requested type,
suitable for use in selection dropdowns and model-picker UIs.

Each provider configuration entry may specify multiple comma-separated model
names; this endpoint expands those into one object per model name so callers
receive a flat, enumerable collection.

**Flattening rules:**
- Only the **first** model in a multi-model provider entry is marked
  `isDefault: true`; all subsequent models from the same entry get `false`.
- `modelFriendlyName` is included **only** when the provider entry contains
  exactly one model name (not a comma-separated list).
- When no providers of the requested type are configured the endpoint still
  returns HTTP **200** with an empty `models` array — this is **not** an error.

**Access control:** requires a valid bearer token. For OAuth tokens the
`config:read` scope must be present; regular JWT bearer tokens pass through
without scope enforcement.


### Example Usage: no_models_configured

<!-- UsageSnippet language="go" operationID="getAvailableModelsByType" method="get" path="/configurationManager/ai-models/available/{modelType}" example="no_models_configured" -->
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

    res, err := s.AIModelsProviders.GetAvailableModelsByType(ctx, components.ModelTypeLlm)
    if err != nil {
        log.Fatal(err)
    }
    if res.Object != nil {
        // handle response
    }
}
```
### Example Usage: two_llm_models

<!-- UsageSnippet language="go" operationID="getAvailableModelsByType" method="get" path="/configurationManager/ai-models/available/{modelType}" example="two_llm_models" -->
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

    res, err := s.AIModelsProviders.GetAvailableModelsByType(ctx, components.ModelTypeEmbedding)
    if err != nil {
        log.Fatal(err)
    }
    if res.Object != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                         | Type                                                                                                                                              | Required                                                                                                                                          | Description                                                                                                                                       |
| ------------------------------------------------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                             | [context.Context](https://pkg.go.dev/context#Context)                                                                                             | :heavy_check_mark:                                                                                                                                | The context to use for the request.                                                                                                               |
| `modelType`                                                                                                                                       | [components.ModelType](../../models/components/modeltype.md)                                                                                      | :heavy_check_mark:                                                                                                                                | Category of AI model to retrieve.<br/><br/>Must be one of: `llm`, `embedding`, `ocr`, `slm`, `reasoning`, `multiModal`, `imageGeneration`, `tts`, `stt`.<br/> |
| `opts`                                                                                                                                            | [][operations.Option](../../models/operations/option.md)                                                                                          | :heavy_minus_sign:                                                                                                                                | The options for this request.                                                                                                                     |

### Response

**[*operations.GetAvailableModelsByTypeResponse](../../models/operations/getavailablemodelsbytyperesponse.md), error**

### Errors

| Error Type                                            | Status Code                                           | Content Type                                          |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| apierrors.GetAvailableModelsByTypeBadRequestError     | 400                                                   | application/json                                      |
| apierrors.GetKnowledgeHubRootNodesUnauthorizedError   | 401                                                   | application/json                                      |
| apierrors.GetAvailableModelsByTypeForbiddenError      | 403                                                   | application/json                                      |
| apierrors.GetKnowledgeHubRootNodesInternalServerError | 500                                                   | application/json                                      |
| apierrors.APIError                                    | 4XX, 5XX                                              | \*/\*                                                 |