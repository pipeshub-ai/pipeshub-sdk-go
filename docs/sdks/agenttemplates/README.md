# AgentTemplates

## Overview

### Available Operations

* [List](#list) - List agent templates
* [Create](#create) - Create agent template
* [Get](#get) - Get agent template
* [Update](#update) - Update agent template
* [Delete](#delete) - Delete agent template

## List

Retrieve all available agent templates.<br><br>
<b>Overview:</b><br>
Agent templates provide pre-configured starting points for creating
custom AI agents. Templates include system prompts, recommended tools,
and configuration schemas.<br><br>
<b>Template Types:</b><br>
<ul>
<li>Public templates available to all organization users</li>
<li>Private templates created by individual users</li>
</ul>


### Example Usage

<!-- UsageSnippet language="go" operationID="listAgentTemplates" method="get" path="/agents/template" -->
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

    res, err := s.AgentTemplates.List(ctx)
    if err != nil {
        log.Fatal(err)
    }
    if res.AgentTemplates != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.ListAgentTemplatesResponse](../../models/operations/listagenttemplatesresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## Create

Create a new reusable agent template.<br><br>
<b>Overview:</b><br>
Templates define the base configuration for agents including
system prompts, tool recommendations, and customization options.<br><br>
<b>Template Components:</b><br>
<ul>
<li><b>System prompt:</b> Default instructions for agents</li>
<li><b>Recommended tools:</b> Suggested tool integrations</li>
<li><b>Config schema:</b> JSON Schema for customization options</li>
</ul>


### Example Usage

<!-- UsageSnippet language="go" operationID="createAgentTemplate" method="post" path="/agents/template" -->
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

    res, err := s.AgentTemplates.Create(ctx, operations.CreateAgentTemplateRequest{
        Name: "Customer Support Agent",
        Category: pipeshub.Pointer("Support"),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.AgentTemplate != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                      | Type                                                                                           | Required                                                                                       | Description                                                                                    |
| ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- |
| `ctx`                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                          | :heavy_check_mark:                                                                             | The context to use for the request.                                                            |
| `request`                                                                                      | [operations.CreateAgentTemplateRequest](../../models/operations/createagenttemplaterequest.md) | :heavy_check_mark:                                                                             | The request object to use for the request.                                                     |
| `opts`                                                                                         | [][operations.Option](../../models/operations/option.md)                                       | :heavy_minus_sign:                                                                             | The options for this request.                                                                  |

### Response

**[*operations.CreateAgentTemplateResponse](../../models/operations/createagenttemplateresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## Get

Retrieve a specific agent template by ID.

### Example Usage

<!-- UsageSnippet language="go" operationID="getAgentTemplate" method="get" path="/agents/template/{templateId}" -->
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

    res, err := s.AgentTemplates.Get(ctx, "<id>")
    if err != nil {
        log.Fatal(err)
    }
    if res.AgentTemplate != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `templateID`                                             | *string*                                                 | :heavy_check_mark:                                       | Template identifier                                      |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.GetAgentTemplateResponse](../../models/operations/getagenttemplateresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## Update

Update an existing agent template.<br><br>
<b>Permissions:</b><br>
Only the template creator can update it.


### Example Usage

<!-- UsageSnippet language="go" operationID="updateAgentTemplate" method="put" path="/agents/template/{templateId}" -->
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

    res, err := s.AgentTemplates.Update(ctx, "<id>", operations.UpdateAgentTemplateRequestBody{})
    if err != nil {
        log.Fatal(err)
    }
    if res.AgentTemplate != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                              | Type                                                                                                   | Required                                                                                               | Description                                                                                            |
| ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                                  | :heavy_check_mark:                                                                                     | The context to use for the request.                                                                    |
| `templateID`                                                                                           | *string*                                                                                               | :heavy_check_mark:                                                                                     | N/A                                                                                                    |
| `body`                                                                                                 | [operations.UpdateAgentTemplateRequestBody](../../models/operations/updateagenttemplaterequestbody.md) | :heavy_check_mark:                                                                                     | Request body for Update agent template                                                                 |
| `opts`                                                                                                 | [][operations.Option](../../models/operations/option.md)                                               | :heavy_minus_sign:                                                                                     | The options for this request.                                                                          |

### Response

**[*operations.UpdateAgentTemplateResponse](../../models/operations/updateagenttemplateresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## Delete

Delete an agent template.<br><br>
<b>Note:</b><br>
Existing agents created from this template are not affected.


### Example Usage

<!-- UsageSnippet language="go" operationID="deleteAgentTemplate" method="delete" path="/agents/template/{templateId}" -->
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

    res, err := s.AgentTemplates.Delete(ctx, "<id>")
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
| `templateID`                                             | *string*                                                 | :heavy_check_mark:                                       | N/A                                                      |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.DeleteAgentTemplateResponse](../../models/operations/deleteagenttemplateresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |