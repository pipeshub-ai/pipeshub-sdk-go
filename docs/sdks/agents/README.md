# Agents

## Overview

Custom AI agents with specialized capabilities and tool integrations

### Available Operations

* [ListAgents](#listagents) - List agents
* [CreateAgent](#createagent) - Create agent
* [GetAgent](#getagent) - Get agent
* [UpdateAgent](#updateagent) - Update agent
* [DeleteAgent](#deleteagent) - Delete agent
* [ListAgentArchivedConversationsGrouped](#listagentarchivedconversationsgrouped) - List archived agent conversations grouped by agent
* [ListAgentConversationArchives](#listagentconversationarchives) - List archived conversations for an agent
* [UploadAgentConversationChatAttachments](#uploadagentconversationchatattachments) - Upload agent chat attachments
* [DeleteAgentConversationChatAttachment](#deleteagentconversationchatattachment) - Delete an agent chat attachment
* [StreamAgentConversation](#streamagentconversation) - Create agent conversation with streaming response
* [StreamAgentConversationMessage](#streamagentconversationmessage) - Add message to agent conversation with streaming response
* [RegenerateAgentConversationMessage](#regenerateagentconversationmessage) - Regenerate agent conversation message
* [UpdateAgentConversationMessageFeedback](#updateagentconversationmessagefeedback) - Submit feedback for an agent message
* [ArchiveAgentConversation](#archiveagentconversation) - Archive an agent conversation
* [UnarchiveAgentConversation](#unarchiveagentconversation) - Unarchive an agent conversation
* [UpdateAgentConversationTitle](#updateagentconversationtitle) - Update agent conversation title
* [DeleteAgentConversationByID](#deleteagentconversationbyid) - Delete an agent conversation
* [GetAgentConversationByID](#getagentconversationbyid) - Get agent conversation by ID
* [ListAgentConversations](#listagentconversations) - List agent conversations

## ListAgents

Retrieve a paginated list of agents available to the authenticated user.

**Overview**

Returns agents accessible through direct, team, or org-level permissions.
Search is performed across agent name, description, and tags. Sorting and
pagination are applied by the AI backend and the resulting envelope is
forwarded unchanged by the Node gateway.

**Gateway contract**

The Node route supports only these query params: `page`, `limit`, `search`,
`sort_by`, and `sort_order`.

The Python backend also understands `isDeleted`, but this gateway route
does not forward it, so it is not part of the public API contract here.


### Example Usage

<!-- UsageSnippet language="go" operationID="listAgents" method="get" path="/agents" example="success" -->
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

    res, err := s.Agents.ListAgents(ctx, operations.ListAgentsRequest{})
    if err != nil {
        log.Fatal(err)
    }
    if res.AgentListResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                    | Type                                                                         | Required                                                                     | Description                                                                  |
| ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- |
| `ctx`                                                                        | [context.Context](https://pkg.go.dev/context#Context)                        | :heavy_check_mark:                                                           | The context to use for the request.                                          |
| `request`                                                                    | [operations.ListAgentsRequest](../../models/operations/listagentsrequest.md) | :heavy_check_mark:                                                           | The request object to use for the request.                                   |
| `opts`                                                                       | [][operations.Option](../../models/operations/option.md)                     | :heavy_minus_sign:                                                           | The options for this request.                                                |

### Response

**[*operations.ListAgentsResponse](../../models/operations/listagentsresponse.md), error**

### Errors

| Error Type              | Status Code             | Content Type            |
| ----------------------- | ----------------------- | ----------------------- |
| apierrors.ErrorResponse | 400, 401                | application/json        |
| apierrors.APIError      | 4XX, 5XX                | \*/\*                   |

## CreateAgent

Create a new custom AI agent.

**Overview:**
Agents are specialized AI assistants configured for specific tasks.
They can have custom system prompts, access to specific tools, and
be limited to certain knowledge bases.

**Agent Configuration:**
- **System prompt:** Instructions that define agent behavior
- **Tools:** Capabilities like web search, code execution, etc.
- **Knowledge bases:** Data sources the agent can access
- **Model config:** AI model settings (temperature, max tokens)

**Use Cases:**
- Customer support bot with product knowledge
- Code review assistant with repository access
- HR assistant with policy documents


### Example Usage

<!-- UsageSnippet language="go" operationID="createAgent" method="post" path="/agents/create" -->
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

    res, err := s.Agents.CreateAgent(ctx, components.AgentCreateRequest{
        Name: "Product Support Agent",
        Models: []components.AgentCreateModelEntryUnion{},
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.AgentCreateResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                      | Type                                                                           | Required                                                                       | Description                                                                    |
| ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ |
| `ctx`                                                                          | [context.Context](https://pkg.go.dev/context#Context)                          | :heavy_check_mark:                                                             | The context to use for the request.                                            |
| `request`                                                                      | [components.AgentCreateRequest](../../models/components/agentcreaterequest.md) | :heavy_check_mark:                                                             | The request object to use for the request.                                     |
| `opts`                                                                         | [][operations.Option](../../models/operations/option.md)                       | :heavy_minus_sign:                                                             | The options for this request.                                                  |

### Response

**[*operations.CreateAgentResponse](../../models/operations/createagentresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## GetAgent

Retrieve agent details by its unique key.

**Gateway not-found behavior:**
Unknown `agentKey`, lookup after soft-delete, and other AI-backend failures
that return 404 from the Python query service are surfaced by the Node
gateway as **HTTP 404** with an `ErrorResponse` body.


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

    res, err := s.Agents.GetAgent(ctx, "customer-support-agent")
    if err != nil {
        log.Fatal(err)
    }
    if res.GetAgentResponse != nil {
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

| Error Type              | Status Code             | Content Type            |
| ----------------------- | ----------------------- | ----------------------- |
| apierrors.ErrorResponse | 400, 401, 403, 404      | application/json        |
| apierrors.ErrorResponse | 500, 503                | application/json        |
| apierrors.APIError      | 4XX, 5XX                | \*/\*                   |

## UpdateAgent

Apply a partial update to an existing agent configuration.

**Gateway contract**

The Node gateway validates the request body via Zod middleware before
forwarding to the Python agent service. The `agentKey` path param and
the request body are both validated. Query parameters are ignored by
the controller.

**Update semantics**

Only fields present in the request body are updated. When `models` is
included, the gateway Zod middleware requires at least one model entry
and at least one object entry with `isReasoning: true`.

**Permissions**

The authenticated user must have `can_edit` on the agent (typically the
owner). Service-account and `shareWithOrg` transitions follow additional
Python business rules.

**Success response**

Returns a lightweight success envelope only. Use
`GET /agents/{agentKey}` to read the persisted agent after an update.


### Example Usage

<!-- UsageSnippet language="go" operationID="updateAgent" method="put" path="/agents/{agentKey}" -->
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

    res, err := s.Agents.UpdateAgent(ctx, "customer-support-agent", components.AgentUpdateRequest{
        Name: pipeshub.Pointer("Renamed Agent"),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.AgentUpdateResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                      | Type                                                                           | Required                                                                       | Description                                                                    | Example                                                                        |
| ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ |
| `ctx`                                                                          | [context.Context](https://pkg.go.dev/context#Context)                          | :heavy_check_mark:                                                             | The context to use for the request.                                            |                                                                                |
| `agentKey`                                                                     | *string*                                                                       | :heavy_check_mark:                                                             | Unique agent identifier                                                        | customer-support-agent                                                         |
| `body`                                                                         | [components.AgentUpdateRequest](../../models/components/agentupdaterequest.md) | :heavy_check_mark:                                                             | Partial agent configuration fields to update                                   |                                                                                |
| `opts`                                                                         | [][operations.Option](../../models/operations/option.md)                       | :heavy_minus_sign:                                                             | The options for this request.                                                  |                                                                                |

### Response

**[*operations.UpdateAgentResponse](../../models/operations/updateagentresponse.md), error**

### Errors

| Error Type              | Status Code             | Content Type            |
| ----------------------- | ----------------------- | ----------------------- |
| apierrors.ErrorResponse | 400, 401, 403, 404      | application/json        |
| apierrors.ErrorResponse | 500                     | application/json        |
| apierrors.APIError      | 4XX, 5XX                | \*/\*                   |

## DeleteAgent

Soft-delete an agent (tombstone) in the graph database.

**Overview:**
The Python query service marks the agent instance deleted inside a transaction.
List and search endpoints exclude tombstoned agents. Toolsets, tools, and
knowledge linked to the agent are not removed by this call.

**Permissions:**
Only the agent owner may delete (`can_delete` on the permission check).

**Warning:**
All conversations with this agent will become inaccessible.

**Gateway not-found behavior:**
Unknown `agentKey`, deleting an already-deleted agent, and `GET /agents/{agentKey}`
after delete return **HTTP 404** with an `ErrorResponse` body.


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

    res, err := s.Agents.DeleteAgent(ctx, "customer-support-agent")
    if err != nil {
        log.Fatal(err)
    }
    if res.AgentDeleteResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                        | Type                                                             | Required                                                         | Description                                                      | Example                                                          |
| ---------------------------------------------------------------- | ---------------------------------------------------------------- | ---------------------------------------------------------------- | ---------------------------------------------------------------- | ---------------------------------------------------------------- |
| `ctx`                                                            | [context.Context](https://pkg.go.dev/context#Context)            | :heavy_check_mark:                                               | The context to use for the request.                              |                                                                  |
| `agentKey`                                                       | *string*                                                         | :heavy_check_mark:                                               | Unique agent identifier (gateway Zod requires non-empty string). | customer-support-agent                                           |
| `opts`                                                           | [][operations.Option](../../models/operations/option.md)         | :heavy_minus_sign:                                               | The options for this request.                                    |                                                                  |

### Response

**[*operations.DeleteAgentResponse](../../models/operations/deleteagentresponse.md), error**

### Errors

| Error Type              | Status Code             | Content Type            |
| ----------------------- | ----------------------- | ----------------------- |
| apierrors.ErrorResponse | 401, 404                | application/json        |
| apierrors.ErrorResponse | 500                     | application/json        |
| apierrors.APIError      | 4XX, 5XX                | \*/\*                   |

## ListAgentArchivedConversationsGrouped

Returns archived agent conversations for the current user, grouped by
`agentKey`, with pagination over agent groups. Excludes conversations
whose agent was soft-deleted upstream.


### Example Usage

<!-- UsageSnippet language="go" operationID="listAgentArchivedConversationsGrouped" method="get" path="/agents/conversations/show/archives" -->
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

    res, err := s.Agents.ListAgentArchivedConversationsGrouped(ctx, pipeshub.Pointer[int64](1), pipeshub.Pointer[int64](5))
    if err != nil {
        log.Fatal(err)
    }
    if res.AgentArchivedGroupsResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `agentPage`                                              | **int64*                                                 | :heavy_minus_sign:                                       | N/A                                                      |
| `agentLimit`                                             | **int64*                                                 | :heavy_minus_sign:                                       | N/A                                                      |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.ListAgentArchivedConversationsGroupedResponse](../../models/operations/listagentarchivedconversationsgroupedresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## ListAgentConversationArchives

Paginated list of archived conversations for the given agent key.

### Example Usage

<!-- UsageSnippet language="go" operationID="listAgentConversationArchives" method="get" path="/agents/{agentKey}/conversations/show/archives" -->
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

    res, err := s.Agents.ListAgentConversationArchives(ctx, operations.ListAgentConversationArchivesRequest{
        AgentKey: "<value>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.AgentArchivedConversationListResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                          | Type                                                                                                               | Required                                                                                                           | Description                                                                                                        |
| ------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                                              | :heavy_check_mark:                                                                                                 | The context to use for the request.                                                                                |
| `request`                                                                                                          | [operations.ListAgentConversationArchivesRequest](../../models/operations/listagentconversationarchivesrequest.md) | :heavy_check_mark:                                                                                                 | The request object to use for the request.                                                                         |
| `opts`                                                                                                             | [][operations.Option](../../models/operations/option.md)                                                           | :heavy_minus_sign:                                                                                                 | The options for this request.                                                                                      |

### Response

**[*operations.ListAgentConversationArchivesResponse](../../models/operations/listagentconversationarchivesresponse.md), error**

### Errors

| Error Type              | Status Code             | Content Type            |
| ----------------------- | ----------------------- | ----------------------- |
| apierrors.ErrorResponse | 400                     | application/json        |
| apierrors.APIError      | 4XX, 5XX                | \*/\*                   |

## UploadAgentConversationChatAttachments

Multipart upload of PDF, JPEG, or PNG files for agent chat. Same limits as assistant
chat (`POST /conversations/attachments/upload`): up to 10 files, 5 MiB each. Proxies to
the AI backend. Optional `conversationId` associates uploads with an existing agent thread.


### Example Usage

<!-- UsageSnippet language="go" operationID="uploadAgentConversationChatAttachments" method="post" path="/agents/{agentKey}/conversations/attachments/upload" -->
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

    res, err := s.Agents.UploadAgentConversationChatAttachments(ctx, "<value>", operations.UploadAgentConversationChatAttachmentsRequestBody{
        Files: []operations.UploadAgentConversationChatAttachmentsFile{},
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.ChatAttachmentUploadResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                    | Type                                                                                                                                         | Required                                                                                                                                     | Description                                                                                                                                  |
| -------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                                                                        | :heavy_check_mark:                                                                                                                           | The context to use for the request.                                                                                                          |
| `agentKey`                                                                                                                                   | *string*                                                                                                                                     | :heavy_check_mark:                                                                                                                           | N/A                                                                                                                                          |
| `body`                                                                                                                                       | [operations.UploadAgentConversationChatAttachmentsRequestBody](../../models/operations/uploadagentconversationchatattachmentsrequestbody.md) | :heavy_check_mark:                                                                                                                           | Multipart form with attachment files and optional `conversationId`.                                                                          |
| `opts`                                                                                                                                       | [][operations.Option](../../models/operations/option.md)                                                                                     | :heavy_minus_sign:                                                                                                                           | The options for this request.                                                                                                                |

### Response

**[*operations.UploadAgentConversationChatAttachmentsResponse](../../models/operations/uploadagentconversationchatattachmentsresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## DeleteAgentConversationChatAttachment

Deletes a previously uploaded attachment by proxying `DELETE` to the query service
(`/api/v1/chat/attachments/{recordId}`). The Node handler always ends the response **without
a JSON body** on success (empty body); the **status code** is the upstream status, or **204**
if none is returned.

On validation failure in the gateway (invalid / blank path params), the response is **400**
with a small JSON error object. Same fire-and-forget semantics as
`DELETE /conversations/attachments/{recordId}` on the client.


### Example Usage

<!-- UsageSnippet language="go" operationID="deleteAgentConversationChatAttachment" method="delete" path="/agents/{agentKey}/conversations/attachments/{recordId}" -->
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

    res, err := s.Agents.DeleteAgentConversationChatAttachment(ctx, "<value>", "<id>")
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
| `agentKey`                                                                     | *string*                                                                       | :heavy_check_mark:                                                             | Agent key path parameter. Must be non-empty.                                   |
| `recordID`                                                                     | *string*                                                                       | :heavy_check_mark:                                                             | Attachment record id (from the upload response). Must be non-blank after trim. |
| `opts`                                                                         | [][operations.Option](../../models/operations/option.md)                       | :heavy_minus_sign:                                                             | The options for this request.                                                  |

### Response

**[*operations.DeleteAgentConversationChatAttachmentResponse](../../models/operations/deleteagentconversationchatattachmentresponse.md), error**

### Errors

| Error Type                                                     | Status Code                                                    | Content Type                                                   |
| -------------------------------------------------------------- | -------------------------------------------------------------- | -------------------------------------------------------------- |
| apierrors.DeleteAgentConversationChatAttachmentBadRequestError | 400                                                            | application/json                                               |
| apierrors.APIError                                             | 4XX, 5XX                                                       | \*/\*                                                          |

## StreamAgentConversation

Start a new conversation with the specified agent and stream the AI
response as Server-Sent Events (SSE). The first user message is saved
and forwarded to the upstream agent backend; subsequent tokens, tool
calls, and lifecycle events are emitted on the open SSE connection.


### Example Usage

<!-- UsageSnippet language="go" operationID="streamAgentConversation" method="post" path="/agents/{agentKey}/conversations/stream" -->
```go
package main

import(
	"context"
	"os"
	"github.com/pipeshub-ai/pipeshub-sdk-go/models/components"
	pipeshub "github.com/pipeshub-ai/pipeshub-sdk-go"
	"github.com/pipeshub-ai/pipeshub-sdk-go/types"
	"log"
)

func main() {
    ctx := context.Background()

    s := pipeshub.New(
        pipeshub.WithSecurity(components.Security{
            BearerAuth: pipeshub.Pointer(os.Getenv("PIPESHUB_BEARER_AUTH")),
        }),
    )

    res, err := s.Agents.StreamAgentConversation(ctx, "<value>", components.AgentStreamCreateConversationRequest{
        Query: "what are some latest tech news?",
        Filters: &components.Filters{
            Apps: []string{
                "2605c882-61d4-4aa2-b480-a68c957c151d",
                "ed6d6cc4-70bd-4838-9aeb-488e910c833a",
                "aeab9ddc-fb9b-47c8-ad98-bd4744e19555",
            },
            Kb: []string{
                "8747da12-4724-4a95-ac92-827b88d79647",
            },
        },
        AppliedFilters: &components.AppliedFilters{
            Apps: []components.AppliedFilterNode{
                components.AppliedFilterNode{
                    ID: pipeshub.Pointer("2605c882-61d4-4aa2-b480-a68c957c151d"),
                    Name: pipeshub.Pointer("US Headlines, abcnews"),
                    NodeType: pipeshub.Pointer("app"),
                    Connector: pipeshub.Pointer("RSS"),
                },
                components.AppliedFilterNode{
                    ID: pipeshub.Pointer("ed6d6cc4-70bd-4838-9aeb-488e910c833a"),
                    Name: pipeshub.Pointer("ABC News RSS"),
                    NodeType: pipeshub.Pointer("app"),
                    Connector: pipeshub.Pointer("RSS"),
                },
                components.AppliedFilterNode{
                    ID: pipeshub.Pointer("aeab9ddc-fb9b-47c8-ad98-bd4744e19555"),
                    Name: pipeshub.Pointer("Hacker news rss"),
                    NodeType: pipeshub.Pointer("app"),
                    Connector: pipeshub.Pointer("RSS"),
                },
            },
            Kb: []components.AppliedFilterNode{
                components.AppliedFilterNode{
                    ID: pipeshub.Pointer("8747da12-4724-4a95-ac92-827b88d79647"),
                    Name: pipeshub.Pointer("Siddhant Ota's Private"),
                    NodeType: pipeshub.Pointer("recordGroup"),
                    Connector: pipeshub.Pointer("KB"),
                },
            },
        },
        ChatMode: components.AgentStreamCreateConversationRequestChatModeAuto.ToPointer(),
        ModelKey: pipeshub.Pointer("5c1832f4-fa19-4167-b913-307fad3a6551"),
        ModelName: pipeshub.Pointer("gpt-5.4-mini"),
        ModelFriendlyName: pipeshub.Pointer("GPT 5.4 mini"),
        Timezone: pipeshub.Pointer("Asia/Kolkata"),
        CurrentTime: types.MustNewTimeFromString("2026-05-19T12:58:01+05:30"),
        Tools: []string{},
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.AgentStreamSSEEvent != nil {
        defer res.AgentStreamSSEEvent.Close()

        for res.AgentStreamSSEEvent.Next() {
            event := res.AgentStreamSSEEvent.Value()
            log.Print(event)
            // Handle the event
	      }
    }
}
```

### Parameters

| Parameter                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                       | Type                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                            | Required                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                        | Description                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                     | Example                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                         |
| ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                           | [context.Context](https://pkg.go.dev/context#Context)                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                           | :heavy_check_mark:                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                              | The context to use for the request.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                             |                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                 |
| `agentKey`                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                      | *string*                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                        | :heavy_check_mark:                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                              | Stable key identifying the agent that owns this conversation.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                   |                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                 |
| `body`                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                          | [components.AgentStreamCreateConversationRequest](../../models/components/agentstreamcreateconversationrequest.md)                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                              | :heavy_check_mark:                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                              | Initial turn payload for the new agent conversation stream.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                     | {<br/>"query": "what are some latest tech news?",<br/>"modelKey": "5c1832f4-fa19-4167-b913-307fad3a6551",<br/>"modelName": "gpt-5.4-mini",<br/>"modelFriendlyName": "GPT 5.4 mini",<br/>"chatMode": "auto",<br/>"timezone": "Asia/Kolkata",<br/>"currentTime": "2026-05-19T12:58:01+05:30",<br/>"tools": [],<br/>"filters": {<br/>"apps": [<br/>"2605c882-61d4-4aa2-b480-a68c957c151d",<br/>"ed6d6cc4-70bd-4838-9aeb-488e910c833a",<br/>"aeab9ddc-fb9b-47c8-ad98-bd4744e19555"<br/>],<br/>"kb": [<br/>"8747da12-4724-4a95-ac92-827b88d79647"<br/>]<br/>},<br/>"appliedFilters": {<br/>"apps": [<br/>{<br/>"id": "2605c882-61d4-4aa2-b480-a68c957c151d",<br/>"name": "US Headlines, abcnews",<br/>"nodeType": "app",<br/>"connector": "RSS"<br/>},<br/>{<br/>"id": "ed6d6cc4-70bd-4838-9aeb-488e910c833a",<br/>"name": "ABC News RSS",<br/>"nodeType": "app",<br/>"connector": "RSS"<br/>},<br/>{<br/>"id": "aeab9ddc-fb9b-47c8-ad98-bd4744e19555",<br/>"name": "Hacker news rss",<br/>"nodeType": "app",<br/>"connector": "RSS"<br/>}<br/>],<br/>"kb": [<br/>{<br/>"id": "8747da12-4724-4a95-ac92-827b88d79647",<br/>"name": "Siddhant Ota's Private",<br/>"nodeType": "recordGroup",<br/>"connector": "KB"<br/>}<br/>]<br/>}<br/>} |
| `opts`                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                          | [][operations.Option](../../models/operations/option.md)                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                        | :heavy_minus_sign:                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                              | The options for this request.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                   |                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                 |

### Response

**[*operations.StreamAgentConversationResponse](../../models/operations/streamagentconversationresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## StreamAgentConversationMessage

Append a user message to an existing agent conversation and stream the
assistant reply over SSE.


### Example Usage

<!-- UsageSnippet language="go" operationID="streamAgentConversationMessage" method="post" path="/agents/{agentKey}/conversations/{conversationId}/messages/stream" -->
```go
package main

import(
	"context"
	"os"
	"github.com/pipeshub-ai/pipeshub-sdk-go/models/components"
	pipeshub "github.com/pipeshub-ai/pipeshub-sdk-go"
	"github.com/pipeshub-ai/pipeshub-sdk-go/types"
	"log"
)

func main() {
    ctx := context.Background()

    s := pipeshub.New(
        pipeshub.WithSecurity(components.Security{
            BearerAuth: pipeshub.Pointer(os.Getenv("PIPESHUB_BEARER_AUTH")),
        }),
    )

    res, err := s.Agents.StreamAgentConversationMessage(ctx, "<value>", "<value>", components.AgentAddMessageStreamRequest{
        Query: "can you elaborate on the latest headlines?",
        Filters: &components.Filters{
            Apps: []string{
                "2605c882-61d4-4aa2-b480-a68c957c151d",
                "ed6d6cc4-70bd-4838-9aeb-488e910c833a",
            },
            Kb: []string{
                "8747da12-4724-4a95-ac92-827b88d79647",
            },
        },
        AppliedFilters: &components.AppliedFilters{
            Apps: []components.AppliedFilterNode{
                components.AppliedFilterNode{
                    ID: pipeshub.Pointer("2605c882-61d4-4aa2-b480-a68c957c151d"),
                    Name: pipeshub.Pointer("US Headlines, abcnews"),
                    NodeType: pipeshub.Pointer("app"),
                    Connector: pipeshub.Pointer("RSS"),
                },
                components.AppliedFilterNode{
                    ID: pipeshub.Pointer("ed6d6cc4-70bd-4838-9aeb-488e910c833a"),
                    Name: pipeshub.Pointer("ABC News RSS"),
                    NodeType: pipeshub.Pointer("app"),
                    Connector: pipeshub.Pointer("RSS"),
                },
            },
            Kb: []components.AppliedFilterNode{
                components.AppliedFilterNode{
                    ID: pipeshub.Pointer("8747da12-4724-4a95-ac92-827b88d79647"),
                    Name: pipeshub.Pointer("Siddhant Ota's Private"),
                    NodeType: pipeshub.Pointer("recordGroup"),
                    Connector: pipeshub.Pointer("KB"),
                },
            },
        },
        ChatMode: components.AgentAddMessageStreamRequestChatModeVerification.ToPointer(),
        ModelKey: pipeshub.Pointer("5c1832f4-fa19-4167-b913-307fad3a6551"),
        ModelName: pipeshub.Pointer("gpt-5.4-mini"),
        ModelFriendlyName: pipeshub.Pointer("GPT 5.4 mini"),
        Timezone: pipeshub.Pointer("Asia/Kolkata"),
        CurrentTime: types.MustNewTimeFromString("2026-05-19T12:58:01+05:30"),
        Tools: []string{},
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.AgentMessageStreamSSEEvent != nil {
        defer res.AgentMessageStreamSSEEvent.Close()

        for res.AgentMessageStreamSSEEvent.Next() {
            event := res.AgentMessageStreamSSEEvent.Value()
            log.Print(event)
            // Handle the event
	      }
    }
}
```

### Parameters

| Parameter                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                              | Type                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                   | Required                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                               | Description                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                            | Example                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                |
| -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                  | :heavy_check_mark:                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                     | The context to use for the request.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                    |                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                        |
| `agentKey`                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                             | *string*                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                               | :heavy_check_mark:                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                     | N/A                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                    |                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                        |
| `conversationID`                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                       | *string*                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                               | :heavy_check_mark:                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                     | N/A                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                    |                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                        |
| `body`                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                 | [components.AgentAddMessageStreamRequest](../../models/components/agentaddmessagestreamrequest.md)                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                     | :heavy_check_mark:                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                     | Follow-up message payload for the agent conversation stream.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                           | {<br/>"query": "can you elaborate on the latest headlines?",<br/>"modelKey": "5c1832f4-fa19-4167-b913-307fad3a6551",<br/>"modelName": "gpt-5.4-mini",<br/>"modelFriendlyName": "GPT 5.4 mini",<br/>"chatMode": "verification",<br/>"timezone": "Asia/Kolkata",<br/>"currentTime": "2026-05-19T12:58:01+05:30",<br/>"tools": [],<br/>"filters": {<br/>"apps": [<br/>"2605c882-61d4-4aa2-b480-a68c957c151d",<br/>"ed6d6cc4-70bd-4838-9aeb-488e910c833a"<br/>],<br/>"kb": [<br/>"8747da12-4724-4a95-ac92-827b88d79647"<br/>]<br/>},<br/>"appliedFilters": {<br/>"apps": [<br/>{<br/>"id": "2605c882-61d4-4aa2-b480-a68c957c151d",<br/>"name": "US Headlines, abcnews",<br/>"nodeType": "app",<br/>"connector": "RSS"<br/>},<br/>{<br/>"id": "ed6d6cc4-70bd-4838-9aeb-488e910c833a",<br/>"name": "ABC News RSS",<br/>"nodeType": "app",<br/>"connector": "RSS"<br/>}<br/>],<br/>"kb": [<br/>{<br/>"id": "8747da12-4724-4a95-ac92-827b88d79647",<br/>"name": "Siddhant Ota's Private",<br/>"nodeType": "recordGroup",<br/>"connector": "KB"<br/>}<br/>]<br/>}<br/>} |
| `opts`                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                 | [][operations.Option](../../models/operations/option.md)                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                               | :heavy_minus_sign:                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                     | The options for this request.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                          |                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                        |

### Response

**[*operations.StreamAgentConversationMessageResponse](../../models/operations/streamagentconversationmessageresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## RegenerateAgentConversationMessage

Regenerate the AI response for a specific message in an agent
conversation and stream the new answer over Server-Sent Events.

**Constraints:**

- Only the last message in the conversation can be regenerated.
- The target message must be of type `bot_response`.

**Request body:**

All request-body fields are optional. When omitted, the server reuses
the original model/context. The body supports:
- `filters`
- `chatMode`
- `modelKey`
- `modelName`
- `modelFriendlyName`
- `timezone`
- `currentTime`
- `tools`

**Streaming behavior:**

The response is delivered as `text/event-stream`. Stable events are
`connected`, `complete`, and `error`. Additional agent/tool lifecycle
events may be forwarded by the backend and should be treated as
informational updates.

Validation failures on params/body are returned as normal HTTP `400`
responses before the stream starts. Valid-shape requests that fail
conversation lookup or regenerate rules are reported as SSE `error`
events after stream initialization.


### Example Usage

<!-- UsageSnippet language="go" operationID="regenerateAgentConversationMessage" method="post" path="/agents/{agentKey}/conversations/{conversationId}/message/{messageId}/regenerate" -->
```go
package main

import(
	"context"
	"os"
	"github.com/pipeshub-ai/pipeshub-sdk-go/models/components"
	pipeshub "github.com/pipeshub-ai/pipeshub-sdk-go"
	"github.com/pipeshub-ai/pipeshub-sdk-go/types"
	"log"
)

func main() {
    ctx := context.Background()

    s := pipeshub.New(
        pipeshub.WithSecurity(components.Security{
            BearerAuth: pipeshub.Pointer(os.Getenv("PIPESHUB_BEARER_AUTH")),
        }),
    )

    res, err := s.Agents.RegenerateAgentConversationMessage(ctx, "<value>", "<value>", "<value>", &components.RegenerateRequest{
        ModelKey: pipeshub.Pointer("05438a37-68f2-4641-a8dc-6c47e63278ca"),
        ModelName: pipeshub.Pointer("gpt-5.4-mini"),
        ModelFriendlyName: pipeshub.Pointer("mini"),
        ChatMode: pipeshub.Pointer("internal_search"),
        Timezone: pipeshub.Pointer("Asia/Calcutta"),
        CurrentTime: types.MustNewTimeFromString("2026-05-11T15:43:21+05:30"),
        Tools: []string{
            "jira.create_issue",
            "confluence.search_content",
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.AgentRegenerateSSEEvent != nil {
        defer res.AgentRegenerateSSEEvent.Close()

        for res.AgentRegenerateSSEEvent.Next() {
            event := res.AgentRegenerateSSEEvent.Value()
            log.Print(event)
            // Handle the event
	      }
    }
}
```

### Parameters

| Parameter                                                                                              | Type                                                                                                   | Required                                                                                               | Description                                                                                            |
| ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                                  | :heavy_check_mark:                                                                                     | The context to use for the request.                                                                    |
| `agentKey`                                                                                             | *string*                                                                                               | :heavy_check_mark:                                                                                     | Stable key identifying the agent that owns this conversation.                                          |
| `conversationID`                                                                                       | *string*                                                                                               | :heavy_check_mark:                                                                                     | ID of the agent conversation containing the target message.                                            |
| `messageID`                                                                                            | *string*                                                                                               | :heavy_check_mark:                                                                                     | ID of the bot-response message to regenerate.                                                          |
| `body`                                                                                                 | [*components.RegenerateRequest](../../models/components/regeneraterequest.md)                          | :heavy_minus_sign:                                                                                     | Optional regeneration payload. All fields are optional and are<br/>validated against `RegenerateRequest`.<br/> |
| `opts`                                                                                                 | [][operations.Option](../../models/operations/option.md)                                               | :heavy_minus_sign:                                                                                     | The options for this request.                                                                          |

### Response

**[*operations.RegenerateAgentConversationMessageResponse](../../models/operations/regenerateagentconversationmessageresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## UpdateAgentConversationMessageFeedback

Append structured feedback to a bot-response message in an agent
conversation. Uses the same request body shape as
`updateMessageFeedback` (helpfulness, categories, comments). Feedback
can only be submitted on `bot_response` messages.


### Example Usage

<!-- UsageSnippet language="go" operationID="updateAgentConversationMessageFeedback" method="post" path="/agents/{agentKey}/conversations/{conversationId}/message/{messageId}/feedback" -->
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

    res, err := s.Agents.UpdateAgentConversationMessageFeedback(ctx, "<value>", "<value>", "<value>", components.MessageFeedbackSubmitRequest{})
    if err != nil {
        log.Fatal(err)
    }
    if res.MessageFeedbackUpdateResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                          | Type                                                                                               | Required                                                                                           | Description                                                                                        |
| -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                              | :heavy_check_mark:                                                                                 | The context to use for the request.                                                                |
| `agentKey`                                                                                         | *string*                                                                                           | :heavy_check_mark:                                                                                 | Unique agent identifier (gateway Zod requires non-empty string).                                   |
| `conversationID`                                                                                   | *string*                                                                                           | :heavy_check_mark:                                                                                 | Unique conversation identifier.                                                                    |
| `messageID`                                                                                        | *string*                                                                                           | :heavy_check_mark:                                                                                 | Identifier of the bot-response message being rated.                                                |
| `body`                                                                                             | [components.MessageFeedbackSubmitRequest](../../models/components/messagefeedbacksubmitrequest.md) | :heavy_check_mark:                                                                                 | Feedback payload for the agent message.                                                            |
| `opts`                                                                                             | [][operations.Option](../../models/operations/option.md)                                           | :heavy_minus_sign:                                                                                 | The options for this request.                                                                      |

### Response

**[*operations.UpdateAgentConversationMessageFeedbackResponse](../../models/operations/updateagentconversationmessagefeedbackresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## ArchiveAgentConversation

Marks the conversation as archived for the authenticated owner.

### Example Usage

<!-- UsageSnippet language="go" operationID="archiveAgentConversation" method="post" path="/agents/{agentKey}/conversations/{conversationId}/archive" -->
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

    res, err := s.Agents.ArchiveAgentConversation(ctx, "<value>", "<value>")
    if err != nil {
        log.Fatal(err)
    }
    if res.AgentConversationArchiveResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `agentKey`                                               | *string*                                                 | :heavy_check_mark:                                       | N/A                                                      |
| `conversationID`                                         | *string*                                                 | :heavy_check_mark:                                       | N/A                                                      |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.ArchiveAgentConversationResponse](../../models/operations/archiveagentconversationresponse.md), error**

### Errors

| Error Type              | Status Code             | Content Type            |
| ----------------------- | ----------------------- | ----------------------- |
| apierrors.ErrorResponse | 400, 401, 404           | application/json        |
| apierrors.APIError      | 4XX, 5XX                | \*/\*                   |

## UnarchiveAgentConversation

Restores an archived agent conversation to the active list.

### Example Usage

<!-- UsageSnippet language="go" operationID="unarchiveAgentConversation" method="post" path="/agents/{agentKey}/conversations/{conversationId}/unarchive" -->
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

    res, err := s.Agents.UnarchiveAgentConversation(ctx, "<value>", "<value>")
    if err != nil {
        log.Fatal(err)
    }
    if res.AgentConversationUnarchiveResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `agentKey`                                               | *string*                                                 | :heavy_check_mark:                                       | N/A                                                      |
| `conversationID`                                         | *string*                                                 | :heavy_check_mark:                                       | N/A                                                      |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.UnarchiveAgentConversationResponse](../../models/operations/unarchiveagentconversationresponse.md), error**

### Errors

| Error Type              | Status Code             | Content Type            |
| ----------------------- | ----------------------- | ----------------------- |
| apierrors.ErrorResponse | 400, 401, 404           | application/json        |
| apierrors.APIError      | 4XX, 5XX                | \*/\*                   |

## UpdateAgentConversationTitle

Updates the display title for an agent conversation owned by the caller.

The controller looks up the conversation by `_id`, `orgId`, `userId`,
`agentKey`, and `isDeleted: false`.

The request body uses the shared title validator (`1..200` chars), and
the controller trims the incoming title before saving it. A whitespace-only
title can therefore still return HTTP 400 even if the raw string is
non-empty.


### Example Usage

<!-- UsageSnippet language="go" operationID="updateAgentConversationTitle" method="patch" path="/agents/{agentKey}/conversations/{conversationId}/title" -->
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

    res, err := s.Agents.UpdateAgentConversationTitle(ctx, "<value>", "<value>", components.ConversationTitleUpdateRequest{
        Title: "ABC News Follow-up",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.AgentConversationTitleUpdateResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                              | Type                                                                                                   | Required                                                                                               | Description                                                                                            |
| ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                                  | :heavy_check_mark:                                                                                     | The context to use for the request.                                                                    |
| `agentKey`                                                                                             | *string*                                                                                               | :heavy_check_mark:                                                                                     | N/A                                                                                                    |
| `conversationID`                                                                                       | *string*                                                                                               | :heavy_check_mark:                                                                                     | N/A                                                                                                    |
| `body`                                                                                                 | [components.ConversationTitleUpdateRequest](../../models/components/conversationtitleupdaterequest.md) | :heavy_check_mark:                                                                                     | New title for the agent conversation.<br/><br/>The server trims the provided string before saving it.<br/> |
| `opts`                                                                                                 | [][operations.Option](../../models/operations/option.md)                                               | :heavy_minus_sign:                                                                                     | The options for this request.                                                                          |

### Response

**[*operations.UpdateAgentConversationTitleResponse](../../models/operations/updateagentconversationtitleresponse.md), error**

### Errors

| Error Type              | Status Code             | Content Type            |
| ----------------------- | ----------------------- | ----------------------- |
| apierrors.ErrorResponse | 400, 401, 404           | application/json        |
| apierrors.APIError      | 4XX, 5XX                | \*/\*                   |

## DeleteAgentConversationByID

Soft-deletes an agent conversation owned by the authenticated user.

The controller scopes the lookup by `_id`, `orgId`, `userId`, and
`agentKey`. If no matching writable conversation is found, the route is
intentionally a no-op and still returns HTTP 200 with `conversation: null`.

This makes the operation idempotent:

- deleting a nonexistent conversation returns success with `null`
- deleting through a different `agentKey` returns success with `null`
- deleting an already deleted conversation returns success with `null`


### Example Usage

<!-- UsageSnippet language="go" operationID="deleteAgentConversationById" method="delete" path="/agents/{agentKey}/conversations/{conversationId}" -->
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

    res, err := s.Agents.DeleteAgentConversationByID(ctx, "<value>", "<value>")
    if err != nil {
        log.Fatal(err)
    }
    if res.AgentConversationDeleteResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `agentKey`                                               | *string*                                                 | :heavy_check_mark:                                       | N/A                                                      |
| `conversationID`                                         | *string*                                                 | :heavy_check_mark:                                       | N/A                                                      |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.DeleteAgentConversationByIDResponse](../../models/operations/deleteagentconversationbyidresponse.md), error**

### Errors

| Error Type              | Status Code             | Content Type            |
| ----------------------- | ----------------------- | ----------------------- |
| apierrors.ErrorResponse | 400, 401                | application/json        |
| apierrors.ErrorResponse | 500                     | application/json        |
| apierrors.APIError      | 4XX, 5XX                | \*/\*                   |

## GetAgentConversationByID

Returns the conversation with paginated/sorted messages and filter metadata.

**Message Pagination:**

Messages are paginated newest-first: `page=1` returns the most recent
batch. Increment `page` to load older batches (used by the infinite-scroll
"load older messages" feature).

- `page`: Page number (default: 1)
- `limit`: Messages per page (default: 20, max: 100)


### Example Usage

<!-- UsageSnippet language="go" operationID="getAgentConversationById" method="get" path="/agents/{agentKey}/conversations/{conversationId}" -->
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

    res, err := s.Agents.GetAgentConversationByID(ctx, operations.GetAgentConversationByIDRequest{
        AgentKey: "<value>",
        ConversationID: "<value>",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.AgentConversationDetailResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                | Type                                                                                                     | Required                                                                                                 | Description                                                                                              |
| -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                                    | :heavy_check_mark:                                                                                       | The context to use for the request.                                                                      |
| `request`                                                                                                | [operations.GetAgentConversationByIDRequest](../../models/operations/getagentconversationbyidrequest.md) | :heavy_check_mark:                                                                                       | The request object to use for the request.                                                               |
| `opts`                                                                                                   | [][operations.Option](../../models/operations/option.md)                                                 | :heavy_minus_sign:                                                                                       | The options for this request.                                                                            |

### Response

**[*operations.GetAgentConversationByIDResponse](../../models/operations/getagentconversationbyidresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## ListAgentConversations

Paginated list of conversations for the agent (owned and shared-with-me),
excluding archived threads.


### Example Usage

<!-- UsageSnippet language="go" operationID="listAgentConversations" method="get" path="/agents/{agentKey}/conversations" -->
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

    res, err := s.Agents.ListAgentConversations(ctx, operations.ListAgentConversationsRequest{
        AgentKey: "<value>",
        StartDate: pipeshub.Pointer("2026-05-26T00:00:00.000Z"),
        EndDate: pipeshub.Pointer("2026-05-27T00:00:00.000Z"),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.AgentConversationListResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                            | Type                                                                                                 | Required                                                                                             | Description                                                                                          |
| ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                | :heavy_check_mark:                                                                                   | The context to use for the request.                                                                  |
| `request`                                                                                            | [operations.ListAgentConversationsRequest](../../models/operations/listagentconversationsrequest.md) | :heavy_check_mark:                                                                                   | The request object to use for the request.                                                           |
| `opts`                                                                                               | [][operations.Option](../../models/operations/option.md)                                             | :heavy_minus_sign:                                                                                   | The options for this request.                                                                        |

### Response

**[*operations.ListAgentConversationsResponse](../../models/operations/listagentconversationsresponse.md), error**

### Errors

| Error Type              | Status Code             | Content Type            |
| ----------------------- | ----------------------- | ----------------------- |
| apierrors.ErrorResponse | 400, 401                | application/json        |
| apierrors.APIError      | 4XX, 5XX                | \*/\*                   |