# ~~ToolsetConfigurations~~

> [!WARNING]
> This SDK is **DEPRECATED**

## Overview

### Available Operations

* [~~Save~~](#save) - Save toolset configuration :warning: **Deprecated**

## ~~Save~~

Save or update toolset configuration (deprecated - use PUT)

> :warning: **DEPRECATED**: This will be removed in a future release, please migrate away from it as soon as possible.

### Example Usage

<!-- UsageSnippet language="go" operationID="saveToolsetConfig" method="post" path="/toolsets/{toolsetId}/config" -->
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

    res, err := s.ToolsetConfigurations.Save(ctx, "<id>", operations.SaveToolsetConfigRequestBody{
        Auth: operations.SaveToolsetConfigAuth{},
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
| `toolsetID`                                                                                        | *string*                                                                                           | :heavy_check_mark:                                                                                 | N/A                                                                                                |
| `body`                                                                                             | [operations.SaveToolsetConfigRequestBody](../../models/operations/savetoolsetconfigrequestbody.md) | :heavy_check_mark:                                                                                 | N/A                                                                                                |
| `opts`                                                                                             | [][operations.Option](../../models/operations/option.md)                                           | :heavy_minus_sign:                                                                                 | The options for this request.                                                                      |

### Response

**[*operations.SaveToolsetConfigResponse](../../models/operations/savetoolsetconfigresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |