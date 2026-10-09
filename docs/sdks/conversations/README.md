# Conversations

## Overview

AI-powered conversational chat management with citations and follow-up questions

### Available Operations

* [CreateConversation](#createconversation) - Create conversation (non-streaming)
* [StreamChat](#streamchat) - Create conversation with streaming response
* [GetAllConversations](#getallconversations) - List all conversations
* [GetArchivedConversations](#getarchivedconversations) - List archived conversations
* [SearchArchivedConversations](#searcharchivedconversations) - Search archived conversations
* [GetConversationByID](#getconversationbyid) - Get conversation by ID
* [DeleteConversationByID](#deleteconversationbyid) - Delete conversation
* [AddMessage](#addmessage) - Add message (non-streaming)
* [AddMessageStream](#addmessagestream) - Add message to a conversation with streaming response
* [UpdateConversationTitle](#updateconversationtitle) - Update conversation title
* [ArchiveConversation](#archiveconversation) - Archive conversation
* [UnarchiveConversation](#unarchiveconversation) - Unarchive conversation
* [RegenerateAnswer](#regenerateanswer) - Regenerate AI response
* [CancelConversationStream](#cancelconversationstream) - Cancel an in-flight chat stream
* [UpdateMessageFeedback](#updatemessagefeedback) - Submit feedback on AI response
* [SetConversationProject](#setconversationproject) - Link or unlink a conversation to a project
* [SetConversationProjectVisibility](#setconversationprojectvisibility) - Override a conversation's project visibility
* [GetProjectConversations](#getprojectconversations) - List a project's conversations

## CreateConversation

Start a new assistant conversation and wait for the complete answer.
The JSON counterpart of `POST /conversations/stream`, for API, SDK
and automation callers that do not consume SSE.

**How a turn runs**

1. The user's message is saved (in its own short transaction on a
   replica set).
2. The AI backend runs the same agent-loop pipeline as the `/stream`
   route and returns only its final result. No transaction is held
   during this call, and the call is never retried.
3. The answer, citations and status are saved exactly as the
   streaming route saves them, and the updated conversation is
   returned.

Every failure after step 1 is persisted: the conversation ends with
status `Failed`, a `failReason`, and an `error` message, and the
response carries `X-Conversation-Id` so the caller can fetch it.
A 4xx from the AI backend (for example no model configured) keeps
its status and user-facing message; other failures return 500 with
a generic message.

The response arrives only when the whole answer is ready, which can
take minutes for agent runs. Allow a generous client and proxy
timeout, or use the `/stream` variant for interactive clients.

**Modes**

`chatMode: agent` (or `agent:<mode>`) runs the universal agent and
honours `tools` / `agentCapabilities`; `internal_search` and
`web_search` run the search assistant and ignore `tools`. Omitted,
it defaults to internal search.


### Example Usage: filtered

<!-- UsageSnippet language="go" operationID="createConversation" method="post" path="/conversations/create" example="filtered" -->
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

    res, err := s.Conversations.CreateConversation(ctx, components.CreateConversationRequest{
        Query: "Summarize the Q4 sales report",
        Filters: &components.Filters{
            Kb: []string{
                "550e8400-e29b-41d4-a716-446655440000",
            },
        },
        ModelKey: pipeshub.Pointer("gpt-4-turbo"),
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.CreateConversationResponse != nil {
        // handle response
    }
}
```
### Example Usage: simple

<!-- UsageSnippet language="go" operationID="createConversation" method="post" path="/conversations/create" example="simple" -->
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

    res, err := s.Conversations.CreateConversation(ctx, components.CreateConversationRequest{
        Query: "What is our company's vacation policy?",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.CreateConversationResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                    | Type                                                                                         | Required                                                                                     | Description                                                                                  |
| -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- |
| `ctx`                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                        | :heavy_check_mark:                                                                           | The context to use for the request.                                                          |
| `request`                                                                                    | [components.CreateConversationRequest](../../models/components/createconversationrequest.md) | :heavy_check_mark:                                                                           | The request object to use for the request.                                                   |
| `opts`                                                                                       | [][operations.Option](../../models/operations/option.md)                                     | :heavy_minus_sign:                                                                           | The options for this request.                                                                |

### Response

**[*operations.CreateConversationResponse](../../models/operations/createconversationresponse.md), error**

### Errors

| Error Type              | Status Code             | Content Type            |
| ----------------------- | ----------------------- | ----------------------- |
| apierrors.ErrorResponse | 400                     | application/json        |
| apierrors.ErrorResponse | 424                     | application/json        |
| apierrors.ErrorResponse | 500                     | application/json        |
| apierrors.APIError      | 4XX, 5XX                | \*/\*                   |

## StreamChat

Start a new conversation and stream the AI response over Server-Sent
Events (SSE). Behaves like `POST /conversations` but emits tokens,
tool activity, and status updates incrementally instead of returning
a single JSON response at the end.

**Lifecycle**

1. The server validates `query`, persists an in-progress
   conversation, then opens the SSE stream with HTTP `200`.
2. A `CUSTOM` event named `conversation_created` is emitted
   immediately with the new `conversationId` so the client can link
   the stream (sidebar, parallel tabs, deep links) without an extra
   request.
3. AI-backend events stream through (token chunks, tool calls,
   status, etc.).
4. On success a single root `RUN_FINISHED` event is emitted carrying
   the full persisted conversation in `result`.
5. On failure a root `RUN_ERROR` event is emitted and the
   conversation is marked FAILED before the stream closes.

**Event vocabulary**

AG-UI is the sole wire protocol. See `ConversationStreamSSEEvent`
for the full event enum and payload guidance.

Clients should ignore unknown event names rather than treating them
as errors.

**Agent mode**

When `chatMode` is `agent`,
the optional `tools` list restricts which tools the agent may
invoke for this turn. Outside agent modes the `tools` field is
ignored.


### Example Usage

<!-- UsageSnippet language="go" operationID="streamChat" method="post" path="/conversations/stream" -->
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

    res, err := s.Conversations.StreamChat(ctx, components.ConversationStreamRequest{
        Query: "What are the key findings from our Q4 financial report?",
        RecordIds: []string{
            "507f1f77bcf86cd799439011",
            "507f1f77bcf86cd799439012",
        },
        ModelKey: pipeshub.Pointer("gpt-4-turbo"),
        ModelName: pipeshub.Pointer("GPT-4 Turbo"),
        ModelFriendlyName: pipeshub.Pointer("GPT-4 Turbo"),
        ChatMode: components.ConversationStreamRequestChatModeWebSearch,
        Timezone: pipeshub.Pointer("America/New_York"),
        CurrentTime: types.MustNewTimeFromString("2026-04-12T16:00:00+05:30"),
        Tools: []string{
            "jira.create_issue",
            "confluence.search_content",
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.ConversationStreamSSEEvent != nil {
        defer res.ConversationStreamSSEEvent.Close()

        for res.ConversationStreamSSEEvent.Next() {
            event := res.ConversationStreamSSEEvent.Value()
            log.Print(event)
            // Handle the event
	      }
    }
}
```

### Parameters

| Parameter                                                                                    | Type                                                                                         | Required                                                                                     | Description                                                                                  |
| -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- |
| `ctx`                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                        | :heavy_check_mark:                                                                           | The context to use for the request.                                                          |
| `request`                                                                                    | [components.ConversationStreamRequest](../../models/components/conversationstreamrequest.md) | :heavy_check_mark:                                                                           | The request object to use for the request.                                                   |
| `opts`                                                                                       | [][operations.Option](../../models/operations/option.md)                                     | :heavy_minus_sign:                                                                           | The options for this request.                                                                |

### Response

**[*operations.StreamChatResponse](../../models/operations/streamchatresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## GetAllConversations

Retrieve paginated conversations for the authenticated user.

**Overview:**

Use the optional `source` query parameter to choose which list to return:
`owned` — only conversations you own (`userId` matches the current user).
`shared` — conversations where you have recipient access
(`isShared` and your user appears in `sharedWith`), without the owner-only branch.
Defaults to `owned` when omitted. Each call returns one list; call twice if you need both.

**Filtering:**

- Only non-archived conversations are returned by default
- Use `/conversations/show/archives` for archived conversations

**Sorting:**

Conversations are sorted by last activity timestamp (most recent first) by default.


### Example Usage

<!-- UsageSnippet language="go" operationID="getAllConversations" method="get" path="/conversations" -->
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

    res, err := s.Conversations.GetAllConversations(ctx, operations.GetAllConversationsRequest{})
    if err != nil {
        log.Fatal(err)
    }
    if res.Object != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                      | Type                                                                                           | Required                                                                                       | Description                                                                                    |
| ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- |
| `ctx`                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                          | :heavy_check_mark:                                                                             | The context to use for the request.                                                            |
| `request`                                                                                      | [operations.GetAllConversationsRequest](../../models/operations/getallconversationsrequest.md) | :heavy_check_mark:                                                                             | The request object to use for the request.                                                     |
| `opts`                                                                                         | [][operations.Option](../../models/operations/option.md)                                       | :heavy_minus_sign:                                                                             | The options for this request.                                                                  |

### Response

**[*operations.GetAllConversationsResponse](../../models/operations/getallconversationsresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## GetArchivedConversations

Retrieve all archived conversations for the authenticated user.

**Overview:**

Archived conversations are hidden from the main list but preserved for reference.
This endpoint returns only conversations where `isArchived: true` and `archivedBy`
is set. Results include conversations the caller owns and those shared with them.

**Filtering and sorting:**

Results can be narrowed using `search`, `shared`, `startDate`, `endDate`, and
`conversationId`. Sorting is controlled by `sortBy` and `sortOrder`. Pagination
is controlled by `page` and `limit`.

**Unarchiving:**

Use `PATCH /conversations/{conversationId}/unarchive` to restore a conversation
to the active list.


### Example Usage

<!-- UsageSnippet language="go" operationID="getArchivedConversations" method="get" path="/conversations/show/archives" -->
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

    res, err := s.Conversations.GetArchivedConversations(ctx, operations.GetArchivedConversationsRequest{})
    if err != nil {
        log.Fatal(err)
    }
    if res.Object != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                | Type                                                                                                     | Required                                                                                                 | Description                                                                                              |
| -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                                    | :heavy_check_mark:                                                                                       | The context to use for the request.                                                                      |
| `request`                                                                                                | [operations.GetArchivedConversationsRequest](../../models/operations/getarchivedconversationsrequest.md) | :heavy_check_mark:                                                                                       | The request object to use for the request.                                                               |
| `opts`                                                                                                   | [][operations.Option](../../models/operations/option.md)                                                 | :heavy_minus_sign:                                                                                       | The options for this request.                                                                            |

### Response

**[*operations.GetArchivedConversationsResponse](../../models/operations/getarchivedconversationsresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## SearchArchivedConversations

Search across all archived conversations (assistant and agent) for the authenticated user.

**Overview:**

Performs a case-insensitive substring match against conversation titles and message content
across both assistant (`Conversation`) and agent (`AgentConversation`) archived collections.
Results are merged server-side and sorted by `lastActivityAt` descending.

**Search parameter:**

The `search` query parameter is required, must be a non-empty string, and is capped at
1000 characters. Requests that omit it or exceed the cap return `400`.

**Pagination:**

Results are paginated using `page` and `limit`. The response includes a `pagination`
block with total counts and a `summary` block that breaks matches down by source.

**Item shape:**

Each item is a conversation list entry (no `messages` payload — that field is omitted
for performance) tagged with `source`, plus computed `isOwner`, `accessLevel`,
`archivedAt`, and `archivedBy`. `agentKey` is present only when `source` is `agent`.


### Example Usage

<!-- UsageSnippet language="go" operationID="searchArchivedConversations" method="get" path="/conversations/show/archives/search" -->
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

    res, err := s.Conversations.SearchArchivedConversations(ctx, "<value>", pipeshub.Pointer[int64](1), pipeshub.Pointer[int64](20))
    if err != nil {
        log.Fatal(err)
    }
    if res.Object != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                             | Type                                                                                  | Required                                                                              | Description                                                                           |
| ------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------- |
| `ctx`                                                                                 | [context.Context](https://pkg.go.dev/context#Context)                                 | :heavy_check_mark:                                                                    | The context to use for the request.                                                   |
| `search`                                                                              | `string`                                                                              | :heavy_check_mark:                                                                    | Search term to match against conversation titles and message content (max 1000 chars) |
| `page`                                                                                | `*int64`                                                                              | :heavy_minus_sign:                                                                    | Page number (1-indexed)                                                               |
| `limit`                                                                               | `*int64`                                                                              | :heavy_minus_sign:                                                                    | Items per page                                                                        |
| `opts`                                                                                | [][operations.Option](../../models/operations/option.md)                              | :heavy_minus_sign:                                                                    | The options for this request.                                                         |

### Response

**[*operations.SearchArchivedConversationsResponse](../../models/operations/searcharchivedconversationsresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## GetConversationByID

Retrieve a specific conversation with its full message history.

**Overview:**

Returns the complete conversation including all messages, citations,
feedback, and metadata. Messages can be paginated for long conversations.

**Message Pagination:**

For conversations with many messages, use pagination parameters:

- `page`: Page number (default: 1)
- `limit`: Messages per page (default: 10)
- `sortBy`: Sort field (default: createdAt)
- `sortOrder`: 'asc' or 'desc' (default: desc)

**Access Control:**

Users can access conversations they own or that have been shared with them.


### Example Usage

<!-- UsageSnippet language="go" operationID="getConversationById" method="get" path="/conversations/{conversationId}" -->
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

    res, err := s.Conversations.GetConversationByID(ctx, operations.GetConversationByIDRequest{
        ConversationID: "507f1f77bcf86cd799439011",
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

| Parameter                                                                                      | Type                                                                                           | Required                                                                                       | Description                                                                                    |
| ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- |
| `ctx`                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                          | :heavy_check_mark:                                                                             | The context to use for the request.                                                            |
| `request`                                                                                      | [operations.GetConversationByIDRequest](../../models/operations/getconversationbyidrequest.md) | :heavy_check_mark:                                                                             | The request object to use for the request.                                                     |
| `opts`                                                                                         | [][operations.Option](../../models/operations/option.md)                                       | :heavy_minus_sign:                                                                             | The options for this request.                                                                  |

### Response

**[*operations.GetConversationByIDResponse](../../models/operations/getconversationbyidresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## DeleteConversationByID

Delete a conversation by its ID.

**Overview:**

Performs a soft delete by setting `isDeleted: true`. The conversation is
removed from listings but preserved in the database. All citations
referenced by messages in the conversation are also soft-deleted.

**Permissions:**

The conversation initiator can always delete. Users the conversation has
been shared with may delete it only when their `sharedWith.accessLevel`
is `write`.


### Example Usage

<!-- UsageSnippet language="go" operationID="deleteConversationById" method="delete" path="/conversations/{conversationId}" -->
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

    res, err := s.Conversations.DeleteConversationByID(ctx, "<value>")
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
| `conversationID`                                         | `string`                                                 | :heavy_check_mark:                                       | Unique conversation identifier                           |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.DeleteConversationByIDResponse](../../models/operations/deleteconversationbyidresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## AddMessage

Ask a follow-up in an existing assistant conversation and wait for
the complete answer. The JSON counterpart of
`POST /conversations/{conversationId}/messages/stream`. Earlier turns
are sent to the model as history; project context comes from the
conversation, never from the request.

**How a turn runs**

1. The user's message is saved (in its own short transaction on a
   replica set).
2. The AI backend runs the same agent-loop pipeline as the `/stream`
   route and returns only its final result. No transaction is held
   during this call, and the call is never retried.
3. The answer, citations and status are saved exactly as the
   streaming route saves them, and the updated conversation is
   returned.

Every failure after step 1 is persisted: the conversation ends with
status `Failed`, a `failReason`, and an `error` message, and the
response carries `X-Conversation-Id` so the caller can fetch it.
A 4xx from the AI backend (for example no model configured) keeps
its status and user-facing message; other failures return 500 with
a generic message.

The response arrives only when the whole answer is ready, which can
take minutes for agent runs. Allow a generous client and proxy
timeout, or use the `/stream` variant for interactive clients.


### Example Usage

<!-- UsageSnippet language="go" operationID="addMessage" method="post" path="/conversations/{conversationId}/messages" -->
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

    res, err := s.Conversations.AddMessage(ctx, "<value>", components.AddMessageRequest{
        Query: "Can you elaborate on the revenue trends?",
        Timezone: pipeshub.Pointer("America/New_York"),
        CurrentTime: types.MustNewTimeFromString("2026-04-12T16:00:00+05:30"),
        Tools: []string{
            "jira.create_issue",
            "confluence.search_content",
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.AddMessageResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                    | Type                                                                         | Required                                                                     | Description                                                                  |
| ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- |
| `ctx`                                                                        | [context.Context](https://pkg.go.dev/context#Context)                        | :heavy_check_mark:                                                           | The context to use for the request.                                          |
| `conversationID`                                                             | `string`                                                                     | :heavy_check_mark:                                                           | N/A                                                                          |
| `body`                                                                       | [components.AddMessageRequest](../../models/components/addmessagerequest.md) | :heavy_check_mark:                                                           | The follow-up question, with optional scope and model overrides.             |
| `opts`                                                                       | [][operations.Option](../../models/operations/option.md)                     | :heavy_minus_sign:                                                           | The options for this request.                                                |

### Response

**[*operations.AddMessageResponse](../../models/operations/addmessageresponse.md), error**

### Errors

| Error Type              | Status Code             | Content Type            |
| ----------------------- | ----------------------- | ----------------------- |
| apierrors.ErrorResponse | 400, 404                | application/json        |
| apierrors.ErrorResponse | 424                     | application/json        |
| apierrors.ErrorResponse | 500                     | application/json        |
| apierrors.APIError      | 4XX, 5XX                | \*/\*                   |

## AddMessageStream

Add a follow-up message to an existing conversation and stream the
assistant's response over Server-Sent Events.

Functionally equivalent to `POST /conversations/{conversationId}/messages`
but the response is delivered as an SSE stream so clients can render
the answer incrementally.

AG-UI is the sole wire protocol. The vocabulary is described by
`ConversationMessageStreamSSEEvent`; it is the same event set as
`/conversations/stream`, while the terminal result reflects an
existing conversation.


### Example Usage

<!-- UsageSnippet language="go" operationID="addMessageStream" method="post" path="/conversations/{conversationId}/messages/stream" -->
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

    res, err := s.Conversations.AddMessageStream(ctx, "<value>", components.ConversationMessageStreamRequest{
        Query: "Can you elaborate on the revenue trends?",
        ChatMode: components.ConversationMessageStreamRequestChatModeInternalSearch,
        Timezone: pipeshub.Pointer("America/New_York"),
        CurrentTime: types.MustNewTimeFromString("2026-04-12T16:00:00+05:30"),
        Tools: []string{
            "jira.create_issue",
            "confluence.search_content",
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.ConversationMessageStreamSSEEvent != nil {
        defer res.ConversationMessageStreamSSEEvent.Close()

        for res.ConversationMessageStreamSSEEvent.Next() {
            event := res.ConversationMessageStreamSSEEvent.Value()
            log.Print(event)
            // Handle the event
	      }
    }
}
```

### Parameters

| Parameter                                                                                                                     | Type                                                                                                                          | Required                                                                                                                      | Description                                                                                                                   |
| ----------------------------------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                         | [context.Context](https://pkg.go.dev/context#Context)                                                                         | :heavy_check_mark:                                                                                                            | The context to use for the request.                                                                                           |
| `conversationID`                                                                                                              | `string`                                                                                                                      | :heavy_check_mark:                                                                                                            | Identifier of the conversation to append the message to. The<br/>conversation must belong to the caller and must not be deleted.<br/> |
| `body`                                                                                                                        | [components.ConversationMessageStreamRequest](../../models/components/conversationmessagestreamrequest.md)                    | :heavy_check_mark:                                                                                                            | Request payload                                                                                                               |
| `opts`                                                                                                                        | [][operations.Option](../../models/operations/option.md)                                                                      | :heavy_minus_sign:                                                                                                            | The options for this request.                                                                                                 |

### Response

**[*operations.AddMessageStreamResponse](../../models/operations/addmessagestreamresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## UpdateConversationTitle

Update the title of a conversation.

**Overview:**

Conversation titles are auto-generated from the first query by default.
Use this endpoint to set a custom, more descriptive title.

**Title limits:**

- Minimum: 1 character
- Maximum: 200 characters

**Permissions:**

The conversation must exist, belong to the calling user's organization,
be owned by the caller (matched on `userId`), and not be soft-deleted.


### Example Usage

<!-- UsageSnippet language="go" operationID="updateConversationTitle" method="patch" path="/conversations/{conversationId}/title" -->
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

    res, err := s.Conversations.UpdateConversationTitle(ctx, "<value>", operations.UpdateConversationTitleRequestBody{
        Title: "Q4 Sales Analysis Discussion",
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

| Parameter                                                                                                      | Type                                                                                                           | Required                                                                                                       | Description                                                                                                    |
| -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                                          | :heavy_check_mark:                                                                                             | The context to use for the request.                                                                            |
| `conversationID`                                                                                               | `string`                                                                                                       | :heavy_check_mark:                                                                                             | Unique conversation identifier                                                                                 |
| `body`                                                                                                         | [operations.UpdateConversationTitleRequestBody](../../models/operations/updateconversationtitlerequestbody.md) | :heavy_check_mark:                                                                                             | Request payload                                                                                                |
| `opts`                                                                                                         | [][operations.Option](../../models/operations/option.md)                                                       | :heavy_minus_sign:                                                                                             | The options for this request.                                                                                  |

### Response

**[*operations.UpdateConversationTitleResponse](../../models/operations/updateconversationtitleresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## ArchiveConversation

Archive a conversation to hide it from the main list.

**Overview:**

Archived conversations are preserved but hidden from the default conversation list.
Use archiving to clean up your workspace without permanently deleting conversations.

**Access:**

The caller must be the conversation's initiator, or be listed in `sharedWith`
with `accessLevel: write`. Already-archived conversations return `400`.

**Retrieval:**

View archived conversations using `GET /conversations/show/archives`.
Restore one with `PATCH /conversations/{conversationId}/unarchive`.


### Example Usage

<!-- UsageSnippet language="go" operationID="archiveConversation" method="patch" path="/conversations/{conversationId}/archive" -->
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

    res, err := s.Conversations.ArchiveConversation(ctx, "<value>")
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
| `conversationID`                                         | `string`                                                 | :heavy_check_mark:                                       | Conversation identifier                                  |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.ArchiveConversationResponse](../../models/operations/archiveconversationresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## UnarchiveConversation

Restore an archived conversation.

- Path params: `conversationId`
- Query params: none
- Body: none


### Example Usage

<!-- UsageSnippet language="go" operationID="unarchiveConversation" method="patch" path="/conversations/{conversationId}/unarchive" -->
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

    res, err := s.Conversations.UnarchiveConversation(ctx, "<value>")
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
| `conversationID`                                         | `string`                                                 | :heavy_check_mark:                                       | Conversation identifier                                  |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.UnarchiveConversationResponse](../../models/operations/unarchiveconversationresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## RegenerateAnswer

Regenerate the AI response for a specific message and stream the new
answer over Server-Sent Events.

**Overview:**

If you're not satisfied with an AI response, use this endpoint to generate
a new answer. The original user query is re-processed and a new bot
response replaces the previous one in place.

**Constraints:**

- Only the *last* message of the conversation can be regenerated.
- The target message must be of type `bot_response`.

**Use Cases:**

- Response was incomplete or unclear
- Want to try a different AI model
- New documents have been indexed since original response

**Model Override:**

Specify `modelKey` to use a different model for regeneration.

**Streaming:**

The response is delivered as an AG-UI `text/event-stream` stream.
Routing still depends on `chatMode`: `internal_search` and
`web_search` use the assistant backend, while `agent` uses the
universal agent loop. See `SSEEvent` for the event vocabulary.


### Example Usage

<!-- UsageSnippet language="go" operationID="regenerateAnswer" method="post" path="/conversations/{conversationId}/message/{messageId}/regenerate" -->
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

    res, err := s.Conversations.RegenerateAnswer(ctx, "<value>", "<value>", &components.RegenerateRequest{
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
    if res.SSEEvent != nil {
        defer res.SSEEvent.Close()

        for res.SSEEvent.Next() {
            event := res.SSEEvent.Value()
            log.Print(event)
            // Handle the event
	      }
    }
}
```

### Parameters

| Parameter                                                                     | Type                                                                          | Required                                                                      | Description                                                                   |
| ----------------------------------------------------------------------------- | ----------------------------------------------------------------------------- | ----------------------------------------------------------------------------- | ----------------------------------------------------------------------------- |
| `ctx`                                                                         | [context.Context](https://pkg.go.dev/context#Context)                         | :heavy_check_mark:                                                            | The context to use for the request.                                           |
| `conversationID`                                                              | `string`                                                                      | :heavy_check_mark:                                                            | N/A                                                                           |
| `messageID`                                                                   | `string`                                                                      | :heavy_check_mark:                                                            | ID of the message to regenerate response for                                  |
| `body`                                                                        | [*components.RegenerateRequest](../../models/components/regeneraterequest.md) | :heavy_minus_sign:                                                            | Request payload                                                               |
| `opts`                                                                        | [][operations.Option](../../models/operations/option.md)                      | :heavy_minus_sign:                                                            | The options for this request.                                                 |

### Response

**[*operations.RegenerateAnswerResponse](../../models/operations/regenerateanswerresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## CancelConversationStream

Cooperatively stop a `POST /conversations/stream` or
`POST /conversations/{conversationId}/messages/stream` run that is
still generating, using the `runId` sent when that stream started.

This is a synchronous JSON ack, not another SSE stream. The
cancelled run's own stream (if still connected) receives a terminal
frame with a `stopped` status and whatever partial answer had
already generated; nothing further is delivered here.

`{ cancelled: false }` — not an error — covers a `runId` that
already finished or was never registered; the caller only needs to
know the stream is not running anymore, not why.

`runId` must belong to a run started on THIS `conversationId` — a
`runId` that exists but was registered under a different
conversation (even one owned by the same caller) is rejected with
`403`, same as a `runId` owned by a different user/org.


### Example Usage

<!-- UsageSnippet language="go" operationID="cancelConversationStream" method="post" path="/conversations/{conversationId}/cancel" -->
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

    res, err := s.Conversations.CancelConversationStream(ctx, "<value>", operations.CancelConversationStreamRequestBody{
        RunID: "62326da7-dd7e-4ecc-9f64-2af864c12ca2",
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

| Parameter                                                                                                        | Type                                                                                                             | Required                                                                                                         | Description                                                                                                      |
| ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                                            | :heavy_check_mark:                                                                                               | The context to use for the request.                                                                              |
| `conversationID`                                                                                                 | `string`                                                                                                         | :heavy_check_mark:                                                                                               | N/A                                                                                                              |
| `body`                                                                                                           | [operations.CancelConversationStreamRequestBody](../../models/operations/cancelconversationstreamrequestbody.md) | :heavy_check_mark:                                                                                               | N/A                                                                                                              |
| `opts`                                                                                                           | [][operations.Option](../../models/operations/option.md)                                                         | :heavy_minus_sign:                                                                                               | The options for this request.                                                                                    |

### Response

**[*operations.CancelConversationStreamResponse](../../models/operations/cancelconversationstreamresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## UpdateMessageFeedback

Append a feedback entry to a bot-response message.

**Overview**

Feedback helps improve AI response quality over time. You can record an
overall helpfulness signal, issue categories, and free-text comments.
Each call appends a new entry to the message; previous entries are
preserved.

**Feedback options**

- `isHelpful` — overall thumbs up/down.
- `categories` — issue or positive categories from a fixed list.
- `comments` — free-text `positive` and `negative`.

**Restrictions**

Feedback can only be submitted on `bot_response` messages — user
queries and system messages are rejected with `400`.


### Example Usage

<!-- UsageSnippet language="go" operationID="updateMessageFeedback" method="post" path="/conversations/{conversationId}/message/{messageId}/feedback" -->
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

    res, err := s.Conversations.UpdateMessageFeedback(ctx, "<value>", "<value>", components.MessageFeedbackSubmitRequest{})
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
| `conversationID`                                                                                   | `string`                                                                                           | :heavy_check_mark:                                                                                 | Unique conversation identifier.                                                                    |
| `messageID`                                                                                        | `string`                                                                                           | :heavy_check_mark:                                                                                 | Identifier of the bot-response message being rated.                                                |
| `body`                                                                                             | [components.MessageFeedbackSubmitRequest](../../models/components/messagefeedbacksubmitrequest.md) | :heavy_check_mark:                                                                                 | Request payload                                                                                    |
| `opts`                                                                                             | [][operations.Option](../../models/operations/option.md)                                           | :heavy_minus_sign:                                                                                 | The options for this request.                                                                      |

### Response

**[*operations.UpdateMessageFeedbackResponse](../../models/operations/updatemessagefeedbackresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## SetConversationProject

Set (`projectId: <id>`) or clear (`projectId: null`) the project this
conversation belongs to. Initiator-only.

**Access:**

The caller must be the conversation's initiator. Linking to a
non-null `projectId` also requires at least viewer access to that
project (`404` if not visible to the caller — never `403`, to avoid
leaking project existence across an org boundary).

**Visibility on link:**

When linking, `projectVisibility` defaults from the project's
`chatSharing` setting (`members` → `project`, otherwise `private`)
unless the conversation was already `project`-visible, in which case
that is preserved. Use
`PATCH /conversations/{conversationId}/project-visibility` to
override it explicitly. Unlinking (`projectId: null`) always clears
both `projectId` and `projectVisibility`.


### Example Usage

<!-- UsageSnippet language="go" operationID="setConversationProject" method="put" path="/conversations/{conversationId}/project" -->
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

    res, err := s.Conversations.SetConversationProject(ctx, "<value>", operations.SetConversationProjectRequestBody{
        ProjectID: pipeshub.Pointer("<value>"),
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

| Parameter                                                                                                    | Type                                                                                                         | Required                                                                                                     | Description                                                                                                  |
| ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                                        | :heavy_check_mark:                                                                                           | The context to use for the request.                                                                          |
| `conversationID`                                                                                             | `string`                                                                                                     | :heavy_check_mark:                                                                                           | Unique conversation identifier                                                                               |
| `body`                                                                                                       | [operations.SetConversationProjectRequestBody](../../models/operations/setconversationprojectrequestbody.md) | :heavy_check_mark:                                                                                           | N/A                                                                                                          |
| `opts`                                                                                                       | [][operations.Option](../../models/operations/option.md)                                                     | :heavy_minus_sign:                                                                                           | The options for this request.                                                                                |

### Response

**[*operations.SetConversationProjectResponse](../../models/operations/setconversationprojectresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## SetConversationProjectVisibility

Explicitly set whether a project-linked conversation is visible to
other members of that project (`project`) or only to its owner
(`private`). Initiator-only. Requires the conversation to already be
linked to a project via
`PUT /conversations/{conversationId}/project`.


### Example Usage

<!-- UsageSnippet language="go" operationID="setConversationProjectVisibility" method="patch" path="/conversations/{conversationId}/project-visibility" -->
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

    res, err := s.Conversations.SetConversationProjectVisibility(ctx, "<value>", operations.SetConversationProjectVisibilityRequestBody{
        Visibility: operations.SetConversationProjectVisibilityVisibilityProject,
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

| Parameter                                                                                                                        | Type                                                                                                                             | Required                                                                                                                         | Description                                                                                                                      |
| -------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                                                            | :heavy_check_mark:                                                                                                               | The context to use for the request.                                                                                              |
| `conversationID`                                                                                                                 | `string`                                                                                                                         | :heavy_check_mark:                                                                                                               | Unique conversation identifier                                                                                                   |
| `body`                                                                                                                           | [operations.SetConversationProjectVisibilityRequestBody](../../models/operations/setconversationprojectvisibilityrequestbody.md) | :heavy_check_mark:                                                                                                               | N/A                                                                                                                              |
| `opts`                                                                                                                           | [][operations.Option](../../models/operations/option.md)                                                                         | :heavy_minus_sign:                                                                                                               | The options for this request.                                                                                                    |

### Response

**[*operations.SetConversationProjectVisibilityResponse](../../models/operations/setconversationprojectvisibilityresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## GetProjectConversations

Requires viewer access to the project. Returns both chat and agent
sessions (`chatSessions`, discriminated by `sessionType`/`agentKey`)
that the caller may see: rows they own, plus rows with
`projectVisibility: project`. Access to the project is asserted
first, so a private conversation belonging to a *different* project
member never leaks through this endpoint.


### Example Usage

<!-- UsageSnippet language="go" operationID="getProjectConversations" method="get" path="/projects/{projectId}/conversations" -->
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

    res, err := s.Conversations.GetProjectConversations(ctx, "<value>", pipeshub.Pointer[int64](1), pipeshub.Pointer[int64](20))
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
| `projectID`                                              | `string`                                                 | :heavy_check_mark:                                       | N/A                                                      |
| `page`                                                   | `*int64`                                                 | :heavy_minus_sign:                                       | N/A                                                      |
| `limit`                                                  | `*int64`                                                 | :heavy_minus_sign:                                       | N/A                                                      |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.GetProjectConversationsResponse](../../models/operations/getprojectconversationsresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |