# ConnectorControl

## Overview

### Available Operations

* [Toggle](#toggle) - Toggle connector sync or agent

## Toggle

Enable or disable a connector for sync or agent functionality.<br><br>
<b>Toggle Types:</b><br>
<ul>
<li><code>sync</code> - Enable/disable data synchronization</li>
<li><code>agent</code> - Enable/disable AI agent integration</li>
</ul>
<b>Prerequisites for Enabling:</b><br>
<ul>
<li>Connector must be configured (<code>isConfigured: true</code>)</li>
<li>For OAuth connectors: Must be authenticated (<code>isAuthenticated: true</code>)</li>
<li>For agent: Connector must support agent (<code>supportsAgent: true</code>)</li>
</ul>
<b>Permissions:</b><br>
<ul>
<li>Team scope: Requires admin</li>
<li>Personal scope: Only creator can toggle</li>
</ul>


### Example Usage

<!-- UsageSnippet language="go" operationID="toggleConnector" method="post" path="/connectors/{connectorId}/toggle" -->
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

    res, err := s.ConnectorControl.Toggle(ctx, "<id>", components.ConnectorToggleRequest{
        Type: components.ConnectorToggleRequestTypeSync,
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.Object != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                              | Type                                                                                   | Required                                                                               | Description                                                                            |
| -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- |
| `ctx`                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                  | :heavy_check_mark:                                                                     | The context to use for the request.                                                    |
| `connectorID`                                                                          | *string*                                                                               | :heavy_check_mark:                                                                     | N/A                                                                                    |
| `body`                                                                                 | [components.ConnectorToggleRequest](../../models/components/connectortogglerequest.md) | :heavy_check_mark:                                                                     | Request payload                                                                        |
| `opts`                                                                                 | [][operations.Option](../../models/operations/option.md)                               | :heavy_minus_sign:                                                                     | The options for this request.                                                          |

### Response

**[*operations.ToggleConnectorResponse](../../models/operations/toggleconnectorresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |