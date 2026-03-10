# AiModelsProviders

## Overview

### Available Operations

* [GetByType](#getbytype) - Get models by type
* [GetAvailableModels](#getavailablemodels) - Get available models for selection
* [Add](#add) - Add new AI model provider
* [Update](#update) - Update AI model provider
* [Delete](#delete) - Delete AI model provider
* [SetDefault](#setdefault) - Set default AI model

## GetByType

Get all configured models of a specific type.

### Example Usage

<!-- UsageSnippet language="go" operationID="getModelsByType" method="get" path="/configurationManager/ai-models/{modelType}" -->
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

    res, err := s.AiModelsProviders.GetByType(ctx, components.ModelTypeEmbedding)
    if err != nil {
        log.Fatal(err)
    }
    if res.Object != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                    | Type                                                         | Required                                                     | Description                                                  |
| ------------------------------------------------------------ | ------------------------------------------------------------ | ------------------------------------------------------------ | ------------------------------------------------------------ |
| `ctx`                                                        | [context.Context](https://pkg.go.dev/context#Context)        | :heavy_check_mark:                                           | The context to use for the request.                          |
| `modelType`                                                  | [components.ModelType](../../models/components/modeltype.md) | :heavy_check_mark:                                           | Type of AI model                                             |
| `opts`                                                       | [][operations.Option](../../models/operations/option.md)     | :heavy_minus_sign:                                           | The options for this request.                                |

### Response

**[*operations.GetModelsByTypeResponse](../../models/operations/getmodelsbytyperesponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## GetAvailableModels

Get available models in a flattened format for UI selection dropdowns.

### Example Usage

<!-- UsageSnippet language="go" operationID="getAvailableModelsByType" method="get" path="/configurationManager/ai-models/available/{modelType}" -->
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

    res, err := s.AiModelsProviders.GetAvailableModels(ctx, components.ModelTypeLlm)
    if err != nil {
        log.Fatal(err)
    }
    if res.Object != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                    | Type                                                         | Required                                                     | Description                                                  |
| ------------------------------------------------------------ | ------------------------------------------------------------ | ------------------------------------------------------------ | ------------------------------------------------------------ |
| `ctx`                                                        | [context.Context](https://pkg.go.dev/context#Context)        | :heavy_check_mark:                                           | The context to use for the request.                          |
| `modelType`                                                  | [components.ModelType](../../models/components/modeltype.md) | :heavy_check_mark:                                           | Type of AI model                                             |
| `opts`                                                       | [][operations.Option](../../models/operations/option.md)     | :heavy_minus_sign:                                           | The options for this request.                                |

### Response

**[*operations.GetAvailableModelsByTypeResponse](../../models/operations/getavailablemodelsbytyperesponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## Add

Add a new AI model provider configuration. Performs a health check before saving to verify connectivity. Supported providers: openai, anthropic, azure-openai, aws-bedrock, google-vertex, ollama, huggingface.

### Example Usage

<!-- UsageSnippet language="go" operationID="addAIModelProvider" method="post" path="/configurationManager/ai-models/providers" -->
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

    res, err := s.AiModelsProviders.Add(ctx, components.AddAIModelProviderRequest{
        ModelType: components.ModelTypeEmbedding,
        Provider: "openai",
        Configuration: components.AddAIModelProviderRequestConfiguration{
            Model: pipeshub.Pointer("gpt-4"),
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.AIModelProviderResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                    | Type                                                                                         | Required                                                                                     | Description                                                                                  |
| -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- |
| `ctx`                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                        | :heavy_check_mark:                                                                           | The context to use for the request.                                                          |
| `request`                                                                                    | [components.AddAIModelProviderRequest](../../models/components/addaimodelproviderrequest.md) | :heavy_check_mark:                                                                           | The request object to use for the request.                                                   |
| `opts`                                                                                       | [][operations.Option](../../models/operations/option.md)                                     | :heavy_minus_sign:                                                                           | The options for this request.                                                                |

### Response

**[*operations.AddAIModelProviderResponse](../../models/operations/addaimodelproviderresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## Update

Update an existing AI model provider configuration.

### Example Usage

<!-- UsageSnippet language="go" operationID="updateAIModelProvider" method="put" path="/configurationManager/ai-models/providers/{modelType}/{modelKey}" -->
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

    res, err := s.AiModelsProviders.Update(ctx, components.ModelTypeReasoning, "<value>", components.UpdateAIModelProviderRequest{
        Provider: "<value>",
        Configuration: components.UpdateAIModelProviderRequestConfiguration{},
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

| Parameter                                                                                          | Type                                                                                               | Required                                                                                           | Description                                                                                        |
| -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                              | :heavy_check_mark:                                                                                 | The context to use for the request.                                                                |
| `modelType`                                                                                        | [components.ModelType](../../models/components/modeltype.md)                                       | :heavy_check_mark:                                                                                 | Type of AI model                                                                                   |
| `modelKey`                                                                                         | *string*                                                                                           | :heavy_check_mark:                                                                                 | Unique model key (UUID)                                                                            |
| `body`                                                                                             | [components.UpdateAIModelProviderRequest](../../models/components/updateaimodelproviderrequest.md) | :heavy_check_mark:                                                                                 | Request payload                                                                                    |
| `opts`                                                                                             | [][operations.Option](../../models/operations/option.md)                                           | :heavy_minus_sign:                                                                                 | The options for this request.                                                                      |

### Response

**[*operations.UpdateAIModelProviderResponse](../../models/operations/updateaimodelproviderresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## Delete

Remove an AI model provider configuration. Cannot delete the default model if it's the only one.

### Example Usage

<!-- UsageSnippet language="go" operationID="deleteAIModelProvider" method="delete" path="/configurationManager/ai-models/providers/{modelType}/{modelKey}" -->
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

    res, err := s.AiModelsProviders.Delete(ctx, components.ModelTypeReasoning, "<value>")
    if err != nil {
        log.Fatal(err)
    }
    if res != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                    | Type                                                         | Required                                                     | Description                                                  |
| ------------------------------------------------------------ | ------------------------------------------------------------ | ------------------------------------------------------------ | ------------------------------------------------------------ |
| `ctx`                                                        | [context.Context](https://pkg.go.dev/context#Context)        | :heavy_check_mark:                                           | The context to use for the request.                          |
| `modelType`                                                  | [components.ModelType](../../models/components/modeltype.md) | :heavy_check_mark:                                           | Type of AI model                                             |
| `modelKey`                                                   | *string*                                                     | :heavy_check_mark:                                           | N/A                                                          |
| `opts`                                                       | [][operations.Option](../../models/operations/option.md)     | :heavy_minus_sign:                                           | The options for this request.                                |

### Response

**[*operations.DeleteAIModelProviderResponse](../../models/operations/deleteaimodelproviderresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## SetDefault

Set a model as the default for its type.

### Example Usage

<!-- UsageSnippet language="go" operationID="setDefaultAIModel" method="put" path="/configurationManager/ai-models/default/{modelType}/{modelKey}" -->
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

    res, err := s.AiModelsProviders.SetDefault(ctx, components.ModelTypeOcr, "<value>")
    if err != nil {
        log.Fatal(err)
    }
    if res != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                    | Type                                                         | Required                                                     | Description                                                  |
| ------------------------------------------------------------ | ------------------------------------------------------------ | ------------------------------------------------------------ | ------------------------------------------------------------ |
| `ctx`                                                        | [context.Context](https://pkg.go.dev/context#Context)        | :heavy_check_mark:                                           | The context to use for the request.                          |
| `modelType`                                                  | [components.ModelType](../../models/components/modeltype.md) | :heavy_check_mark:                                           | Type of AI model                                             |
| `modelKey`                                                   | *string*                                                     | :heavy_check_mark:                                           | N/A                                                          |
| `opts`                                                       | [][operations.Option](../../models/operations/option.md)     | :heavy_minus_sign:                                           | The options for this request.                                |

### Response

**[*operations.SetDefaultAIModelResponse](../../models/operations/setdefaultaimodelresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |