# OauthConfigurations

## Overview

### Available Operations

* [Update](#update) - Update OAuth config
* [Delete](#delete) - Delete OAuth configuration

## Update

Update OAuth config

### Example Usage

<!-- UsageSnippet language="go" operationID="updateToolsetOAuthConfig" method="put" path="/toolsets/oauth-configs/{toolsetType}/{oauthConfigId}" -->
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

    res, err := s.OauthConfigurations.Update(ctx, "<value>", "<id>", map[string]any{
        "key": "<value>",
        "key1": "<value>",
        "key2": "<value>",
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

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `toolsetType`                                            | *string*                                                 | :heavy_check_mark:                                       | N/A                                                      |
| `oauthConfigID`                                          | *string*                                                 | :heavy_check_mark:                                       | N/A                                                      |
| `body`                                                   | map[string]*any*                                         | :heavy_check_mark:                                       | N/A                                                      |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.UpdateToolsetOAuthConfigResponse](../../models/operations/updatetoolsetoauthconfigresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## Delete

Delete an OAuth configuration.<br><br>
<b>Warning:</b><br>
Cannot delete if the configuration is used by active connectors.
Disable or delete dependent connectors first.


### Example Usage

<!-- UsageSnippet language="go" operationID="deleteOAuthConfig" method="delete" path="/oauth/{connectorType}/{configId}" -->
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

    res, err := s.OauthConfigurations.Delete(ctx, "<value>", "<id>")
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
| `connectorType`                                          | *string*                                                 | :heavy_check_mark:                                       | N/A                                                      |
| `configID`                                               | *string*                                                 | :heavy_check_mark:                                       | N/A                                                      |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.DeleteOAuthConfigResponse](../../models/operations/deleteoauthconfigresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |