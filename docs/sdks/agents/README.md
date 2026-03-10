# Agents

## Overview

Custom AI agents with specialized capabilities and tool integrations

### Available Operations

* [List](#list) - List agents
* [Create](#create) - Create agent
* [ListTools](#listtools) - List available tools
* [Get](#get) - Get agent
* [Update](#update) - Update agent
* [Delete](#delete) - Delete agent
* [GetPermissions](#getpermissions) - Get agent permissions
* [UpdatePermissions](#updatepermissions) - Update agent permissions
* [Share](#share) - Share agent
* [Unshare](#unshare) - Unshare an agent

## List

Retrieve all agents available to the authenticated user.<br><br>
<b>Overview:</b><br>
Returns agents created by the user plus public agents in the organization.
Each agent has unique capabilities defined by its tools and knowledge scope.


### Example Usage

<!-- UsageSnippet language="go" operationID="listAgents" method="get" path="/agents" -->
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

    res, err := s.Agents.List(ctx)
    if err != nil {
        log.Fatal(err)
    }
    if res.Agents != nil {
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

**[*operations.ListAgentsResponse](../../models/operations/listagentsresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## Create

Create a new custom AI agent.<br><br>
<b>Overview:</b><br>
Agents are specialized AI assistants configured for specific tasks.
They can have custom system prompts, access to specific tools, and
be limited to certain knowledge bases.<br><br>
<b>Agent Configuration:</b><br>
<ul>
<li><b>System prompt:</b> Instructions that define agent behavior</li>
<li><b>Tools:</b> Capabilities like web search, code execution, etc.</li>
<li><b>Knowledge bases:</b> Data sources the agent can access</li>
<li><b>Model config:</b> AI model settings (temperature, max tokens)</li>
</ul>
<b>Use Cases:</b><br>
<ul>
<li>Customer support bot with product knowledge</li>
<li>Code review assistant with repository access</li>
<li>HR assistant with policy documents</li>
</ul>


### Example Usage

<!-- UsageSnippet language="go" operationID="createAgent" method="post" path="/agents/create" -->
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

    res, err := s.Agents.Create(ctx, operations.CreateAgentRequest{
        Name: "Product Support Agent",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.Agent != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                      | Type                                                                           | Required                                                                       | Description                                                                    |
| ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ |
| `ctx`                                                                          | [context.Context](https://pkg.go.dev/context#Context)                          | :heavy_check_mark:                                                             | The context to use for the request.                                            |
| `request`                                                                      | [operations.CreateAgentRequest](../../models/operations/createagentrequest.md) | :heavy_check_mark:                                                             | The request object to use for the request.                                     |
| `opts`                                                                         | [][operations.Option](../../models/operations/option.md)                       | :heavy_minus_sign:                                                             | The options for this request.                                                  |

### Response

**[*operations.CreateAgentResponse](../../models/operations/createagentresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## ListTools

Get all tools that can be assigned to agents.<br><br>
<b>Overview:</b><br>
Tools extend agent capabilities beyond basic Q&A. Each tool
has specific inputs and outputs defined by its schema.<br><br>
<b>Common Tools:</b><br>
<ul>
<li><b>web-search:</b> Search the internet</li>
<li><b>code-interpreter:</b> Execute code snippets</li>
<li><b>file-reader:</b> Read uploaded files</li>
<li><b>api-caller:</b> Make external API requests</li>
</ul>


### Example Usage

<!-- UsageSnippet language="go" operationID="listAgentTools" method="get" path="/agents/tools/list" -->
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

    res, err := s.Agents.ListTools(ctx)
    if err != nil {
        log.Fatal(err)
    }
    if res.AgentTools != nil {
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

**[*operations.ListAgentToolsResponse](../../models/operations/listagenttoolsresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## Get

Retrieve agent details by its unique key.

### Example Usage

<!-- UsageSnippet language="go" operationID="getAgent" method="get" path="/agents/{agentKey}" -->
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

    res, err := s.Agents.Get(ctx, "customer-support-agent")
    if err != nil {
        log.Fatal(err)
    }
    if res.Agent != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              | Example                                                  |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |                                                          |
| `agentKey`                                               | *string*                                                 | :heavy_check_mark:                                       | Unique agent identifier                                  | customer-support-agent                                   |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |                                                          |

### Response

**[*operations.GetAgentResponse](../../models/operations/getagentresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## Update

Update an existing agent's configuration.<br><br>
<b>Permissions:</b><br>
Only the agent creator can update it.


### Example Usage

<!-- UsageSnippet language="go" operationID="updateAgent" method="put" path="/agents/{agentKey}" -->
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

    res, err := s.Agents.Update(ctx, "<value>", operations.UpdateAgentRequestBody{})
    if err != nil {
        log.Fatal(err)
    }
    if res.Agent != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                              | Type                                                                                   | Required                                                                               | Description                                                                            |
| -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- |
| `ctx`                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                  | :heavy_check_mark:                                                                     | The context to use for the request.                                                    |
| `agentKey`                                                                             | *string*                                                                               | :heavy_check_mark:                                                                     | N/A                                                                                    |
| `body`                                                                                 | [operations.UpdateAgentRequestBody](../../models/operations/updateagentrequestbody.md) | :heavy_check_mark:                                                                     | Request body for Update agent                                                          |
| `opts`                                                                                 | [][operations.Option](../../models/operations/option.md)                               | :heavy_minus_sign:                                                                     | The options for this request.                                                          |

### Response

**[*operations.UpdateAgentResponse](../../models/operations/updateagentresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## Delete

Delete an agent.<br><br>
<b>Warning:</b><br>
All conversations with this agent will become inaccessible.


### Example Usage

<!-- UsageSnippet language="go" operationID="deleteAgent" method="delete" path="/agents/{agentKey}" -->
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

    res, err := s.Agents.Delete(ctx, "<value>")
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
| `agentKey`                                               | *string*                                                 | :heavy_check_mark:                                       | N/A                                                      |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.DeleteAgentResponse](../../models/operations/deleteagentresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## GetPermissions

Get the current permission configuration for an agent.

### Example Usage

<!-- UsageSnippet language="go" operationID="getAgentPermissions" method="get" path="/agents/{agentKey}/permissions" -->
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

    res, err := s.Agents.GetPermissions(ctx, "<value>")
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
| `agentKey`                                               | *string*                                                 | :heavy_check_mark:                                       | N/A                                                      |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.GetAgentPermissionsResponse](../../models/operations/getagentpermissionsresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## UpdatePermissions

Update who can access and use the agent.

### Example Usage

<!-- UsageSnippet language="go" operationID="updateAgentPermissions" method="put" path="/agents/{agentKey}/permissions" -->
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

    res, err := s.Agents.UpdatePermissions(ctx, "<value>", operations.UpdateAgentPermissionsRequestBody{})
    if err != nil {
        log.Fatal(err)
    }
    if res != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                    | Type                                                                                                         | Required                                                                                                     | Description                                                                                                  |
| ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                                        | :heavy_check_mark:                                                                                           | The context to use for the request.                                                                          |
| `agentKey`                                                                                                   | *string*                                                                                                     | :heavy_check_mark:                                                                                           | N/A                                                                                                          |
| `body`                                                                                                       | [operations.UpdateAgentPermissionsRequestBody](../../models/operations/updateagentpermissionsrequestbody.md) | :heavy_check_mark:                                                                                           | Request body for Update agent permissions                                                                    |
| `opts`                                                                                                       | [][operations.Option](../../models/operations/option.md)                                                     | :heavy_minus_sign:                                                                                           | The options for this request.                                                                                |

### Response

**[*operations.UpdateAgentPermissionsResponse](../../models/operations/updateagentpermissionsresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## Share

Share an agent with specific users.

### Example Usage

<!-- UsageSnippet language="go" operationID="shareAgent" method="post" path="/agents/{agentKey}/share" -->
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

    res, err := s.Agents.Share(ctx, "<value>", components.ShareRequest{
        UserIds: []string{
            "507f1f77bcf86cd799439011",
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.Agent != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                          | Type                                                               | Required                                                           | Description                                                        |
| ------------------------------------------------------------------ | ------------------------------------------------------------------ | ------------------------------------------------------------------ | ------------------------------------------------------------------ |
| `ctx`                                                              | [context.Context](https://pkg.go.dev/context#Context)              | :heavy_check_mark:                                                 | The context to use for the request.                                |
| `agentKey`                                                         | *string*                                                           | :heavy_check_mark:                                                 | N/A                                                                |
| `body`                                                             | [components.ShareRequest](../../models/components/sharerequest.md) | :heavy_check_mark:                                                 | Request body for Share agent                                       |
| `opts`                                                             | [][operations.Option](../../models/operations/option.md)           | :heavy_minus_sign:                                                 | The options for this request.                                      |

### Response

**[*operations.ShareAgentResponse](../../models/operations/shareagentresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## Unshare

Revoke sharing for an agent, removing access for specified users or teams.


### Example Usage

<!-- UsageSnippet language="go" operationID="unshareAgent" method="post" path="/agents/{agentKey}/unshare" -->
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

    res, err := s.Agents.Unshare(ctx, "<value>", operations.UnshareAgentRequestBody{})
    if err != nil {
        log.Fatal(err)
    }
    if res.Object != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                | Type                                                                                     | Required                                                                                 | Description                                                                              |
| ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- |
| `ctx`                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                    | :heavy_check_mark:                                                                       | The context to use for the request.                                                      |
| `agentKey`                                                                               | *string*                                                                                 | :heavy_check_mark:                                                                       | N/A                                                                                      |
| `body`                                                                                   | [operations.UnshareAgentRequestBody](../../models/operations/unshareagentrequestbody.md) | :heavy_check_mark:                                                                       | Request payload                                                                          |
| `opts`                                                                                   | [][operations.Option](../../models/operations/option.md)                                 | :heavy_minus_sign:                                                                       | The options for this request.                                                            |

### Response

**[*operations.UnshareAgentResponse](../../models/operations/unshareagentresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |