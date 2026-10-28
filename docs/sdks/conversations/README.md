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
* [PinConversation](#pinconversation) - Pin conversation
* [UnpinConversation](#unpinconversation) - Unpin conversation
* [RegenerateAnswer](#regenerateanswer) - Regenerate AI response
* [ForkConversation](#forkconversation) - Fork conversation from a message
* [CancelConversationStream](#cancelconversationstream) - Cancel an in-flight chat stream
* [UpdateMessageFeedback](#updatemessagefeedback) - Submit feedback on AI response

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

`chatMode: agent` runs the universal agent and honours `tools` /
`agentCapabilities`; `internal_search` and `web_search` run the
search assistant and ignore `tools`. The legacy values `deep`,
`quick`, `verification` and `auto` are accepted and run as internal
search. Omitted, it defaults to internal search.


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
| apierrors.ErrorResponse | 400, 401, 403, 404, 409 | application/json        |
| apierrors.ErrorResponse | 424                     | application/json        |
| apierrors.ErrorResponse | 500, 503, 504           | application/json        |
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

| Error Type              | Status Code             | Content Type            |
| ----------------------- | ----------------------- | ----------------------- |
| apierrors.ErrorResponse | 400, 401, 403, 404, 409 | application/json        |
| apierrors.ErrorResponse | 500                     | application/json        |
| apierrors.APIError      | 4XX, 5XX                | \*/\*                   |

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
With `source=owned`, conversations the caller pinned
(`PATCH /conversations/{conversationId}/pin`) come first, in the same order,
whatever `sortBy` is. `source=shared` is never pinned-first.

Query parameters not listed here are ignored, not refused.


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

| Error Type              | Status Code             | Content Type            |
| ----------------------- | ----------------------- | ----------------------- |
| apierrors.ErrorResponse | 400, 401, 403           | application/json        |
| apierrors.ErrorResponse | 500                     | application/json        |
| apierrors.APIError      | 4XX, 5XX                | \*/\*                   |

## GetArchivedConversations

Retrieve all archived conversations for the authenticated user.

**Overview:**

Archived conversations are hidden from the main list but preserved for reference.
This endpoint returns only conversations where `isArchived: true` and `archivedBy`
is set. Results include conversations the caller owns and those shared with them
(`isShared` and the caller listed in `sharedWith`). Deleted and agent
conversations are never listed. Each item carries its `messages`.

**Filtering and sorting:**

Results can be narrowed using `search`, `shared`, `startDate`, `endDate`, and
`conversationId`. Sorting is controlled by `sortBy` and `sortOrder`. Pagination
is controlled by `page` and `limit`.

Query parameters not listed here are dropped by the validator and
ignored, including `projectId`, which `GET /conversations` accepts.

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

| Error Type              | Status Code             | Content Type            |
| ----------------------- | ----------------------- | ----------------------- |
| apierrors.ErrorResponse | 400, 401, 403           | application/json        |
| apierrors.ErrorResponse | 500                     | application/json        |
| apierrors.APIError      | 4XX, 5XX                | \*/\*                   |

## SearchArchivedConversations

Search across all archived conversations (assistant and agent) for the authenticated user.

**Overview:**

Performs a case-insensitive substring match against conversation titles and message content
across both assistant (`Conversation`) and agent (`AgentConversation`) archived collections.
Results are merged server-side and sorted by `lastActivityAt` descending.

**Search parameter:**

The `search` query parameter is required, must contain a non-space character, and is
capped at 1000 characters; it is trimmed before use. Requests that omit it, send only
spaces or exceed the cap return `400 VALIDATION_ERROR`. A value holding markup or a
printf-style format specifier (such as `%s`) is refused by the controller with
`400 HTTP_BAD_REQUEST`.

Query parameters not listed here are dropped by the validator and ignored.

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

    res, err := s.Conversations.SearchArchivedConversations(ctx, "<value>", pipeshub.Pointer[float64](1.0), pipeshub.Pointer[float64](20.0))
    if err != nil {
        log.Fatal(err)
    }
    if res.Object != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                        | Type                                                                                                                                             | Required                                                                                                                                         | Description                                                                                                                                      |
| ------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                                                                            | :heavy_check_mark:                                                                                                                               | The context to use for the request.                                                                                                              |
| `search`                                                                                                                                         | `string`                                                                                                                                         | :heavy_check_mark:                                                                                                                               | Search term to match against conversation titles and message content (max 1000 chars)                                                            |
| `page`                                                                                                                                           | `*float64`                                                                                                                                       | :heavy_minus_sign:                                                                                                                               | Page number (1-indexed). The validator accepts any number, so a fraction such<br/>as `1.5` is not refused; it is truncated (`1.5` reads as page 1).<br/> |
| `limit`                                                                                                                                          | `*float64`                                                                                                                                       | :heavy_minus_sign:                                                                                                                               | Items per page. A fraction is not refused; it is truncated (`5.5` reads as 5).<br/>                                                              |
| `opts`                                                                                                                                           | [][operations.Option](../../models/operations/option.md)                                                                                         | :heavy_minus_sign:                                                                                                                               | The options for this request.                                                                                                                    |

### Response

**[*operations.SearchArchivedConversationsResponse](../../models/operations/searcharchivedconversationsresponse.md), error**

### Errors

| Error Type              | Status Code             | Content Type            |
| ----------------------- | ----------------------- | ----------------------- |
| apierrors.ErrorResponse | 400, 401, 403           | application/json        |
| apierrors.ErrorResponse | 500                     | application/json        |
| apierrors.APIError      | 4XX, 5XX                | \*/\*                   |

## GetConversationByID

Retrieve a specific conversation with its full message history.

**Overview:**

Returns the complete conversation including all messages, citations,
feedback, and metadata. Messages can be paginated for long conversations.

**Message Pagination:**

For conversations with many messages, use pagination parameters:

- `page`: Page number (default: 1)
- `limit`: Messages per page (default: 20)
- `sortBy`: Sort field (default: createdAt)
- `sortOrder`: 'asc' or 'desc' (default: desc)

**Access Control:**

Users can access conversations they own or that have been shared with them.

Query parameters not listed here are ignored, not refused.


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

| Error Type              | Status Code             | Content Type            |
| ----------------------- | ----------------------- | ----------------------- |
| apierrors.ErrorResponse | 400, 401, 403, 404      | application/json        |
| apierrors.ErrorResponse | 500                     | application/json        |
| apierrors.APIError      | 4XX, 5XX                | \*/\*                   |

## DeleteConversationByID

Delete a conversation by its ID.

**Overview:**

Performs a soft delete by setting `isDeleted: true`. The conversation is
removed from listings but preserved in the database. All citations
referenced by messages in the conversation are also soft-deleted.

**Permissions:**

Only the conversation's owner can delete it. The lookup also allows a
user the conversation is shared with at `write` access, but it
requires the caller to be the owner as well, so such a user is
answered `404`.

**Artifacts:** before the conversation is marked deleted, every artifact
the caller created in it is deleted for good (every version and stored
file), and so are older stored outputs of the conversation. If that cannot
be done the conversation is left in place and the error is returned
(for example 503); repeating the request finishes the cleanup.


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

| Error Type              | Status Code             | Content Type            |
| ----------------------- | ----------------------- | ----------------------- |
| apierrors.ErrorResponse | 400, 401, 403, 404      | application/json        |
| apierrors.ErrorResponse | 500                     | application/json        |
| apierrors.APIError      | 4XX, 5XX                | \*/\*                   |

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

| Parameter                                                                                                                                                                                                     | Type                                                                                                                                                                                                          | Required                                                                                                                                                                                                      | Description                                                                                                                                                                                                   |
| ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                                                                                         | [context.Context](https://pkg.go.dev/context#Context)                                                                                                                                                         | :heavy_check_mark:                                                                                                                                                                                            | The context to use for the request.                                                                                                                                                                           |
| `conversationID`                                                                                                                                                                                              | `string`                                                                                                                                                                                                      | :heavy_check_mark:                                                                                                                                                                                            | N/A                                                                                                                                                                                                           |
| `body`                                                                                                                                                                                                        | [components.AddMessageRequest](../../models/components/addmessagerequest.md)                                                                                                                                  | :heavy_check_mark:                                                                                                                                                                                            | The follow-up question, with optional scope and model overrides. A `query` of only<br/>whitespace is refused with `400` `HTTP_BAD_REQUEST` ("Query is required"), after the<br/>validator and before any model runs.<br/> |
| `opts`                                                                                                                                                                                                        | [][operations.Option](../../models/operations/option.md)                                                                                                                                                      | :heavy_minus_sign:                                                                                                                                                                                            | The options for this request.                                                                                                                                                                                 |

### Response

**[*operations.AddMessageResponse](../../models/operations/addmessageresponse.md), error**

### Errors

| Error Type              | Status Code             | Content Type            |
| ----------------------- | ----------------------- | ----------------------- |
| apierrors.ErrorResponse | 400, 401, 403, 404      | application/json        |
| apierrors.ErrorResponse | 424                     | application/json        |
| apierrors.ErrorResponse | 500, 503, 504           | application/json        |
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

| Error Type              | Status Code             | Content Type            |
| ----------------------- | ----------------------- | ----------------------- |
| apierrors.ErrorResponse | 400, 401, 403           | application/json        |
| apierrors.ErrorResponse | 500                     | application/json        |
| apierrors.APIError      | 4XX, 5XX                | \*/\*                   |

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
be owned by the caller (matched on `userId`), not be soft-deleted, and
not be an agent conversation. A user it is shared with, even at
`write` access, gets `404`.

Body fields other than `title` are dropped by the validator, not
refused.


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

| Error Type              | Status Code             | Content Type            |
| ----------------------- | ----------------------- | ----------------------- |
| apierrors.ErrorResponse | 400, 401, 403, 404      | application/json        |
| apierrors.ErrorResponse | 500                     | application/json        |
| apierrors.APIError      | 4XX, 5XX                | \*/\*                   |

## ArchiveConversation

Archive a conversation to hide it from the main list.

**Overview:**

Archived conversations are preserved but hidden from the default conversation list.
Use archiving to clean up your workspace without permanently deleting conversations.

**Access:**

Only the conversation's owner can archive it: the lookup filters on
the caller's `userId`. It also has a branch for users the
conversation is shared with at `write` access, but that branch can
never match, so such users get `404` (a known defect). A missing,
deleted or agent conversation is `404` too.
Already-archived conversations return `400`.

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

| Error Type              | Status Code             | Content Type            |
| ----------------------- | ----------------------- | ----------------------- |
| apierrors.ErrorResponse | 400, 401, 403, 404      | application/json        |
| apierrors.ErrorResponse | 500                     | application/json        |
| apierrors.APIError      | 4XX, 5XX                | \*/\*                   |

## UnarchiveConversation

Restore an archived conversation.

- Path params: `conversationId`
- Query params: none
- Body: none

**Access:**

Only the conversation's owner can unarchive it: the lookup filters on
the caller's `userId`. It also has a branch for users the
conversation is shared with at `write` access, but that branch can
never match, so such users get `404` (a known defect). A missing,
deleted or agent conversation is `404` too.
A conversation that is not archived returns `400`.


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

| Error Type              | Status Code             | Content Type            |
| ----------------------- | ----------------------- | ----------------------- |
| apierrors.ErrorResponse | 400, 401, 403, 404      | application/json        |
| apierrors.ErrorResponse | 500                     | application/json        |
| apierrors.APIError      | 4XX, 5XX                | \*/\*                   |

## PinConversation

Pin a conversation so it is listed first in the caller's own conversation list.

**Access:**

Owner only: the caller must be the conversation's `initiator`. People the
conversation is shared with, even with `write` access, cannot pin it, and
their own lists are not affected by the owner's pin.

**Behaviour:**

Idempotent: pinning an already pinned conversation succeeds and changes nothing.
Archived and deleted conversations cannot be pinned. Archiving a pinned
conversation removes its pin. Pinning does not change `lastActivityAt` or
`updatedAt`, so the conversation keeps its place among the other pinned chats.

- Path params: `conversationId`
- Query params: none
- Body: none


### Example Usage

<!-- UsageSnippet language="go" operationID="pinConversation" method="patch" path="/conversations/{conversationId}/pin" -->
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

    res, err := s.Conversations.PinConversation(ctx, "<value>")
    if err != nil {
        log.Fatal(err)
    }
    if res.ConversationPinResponse != nil {
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

**[*operations.PinConversationResponse](../../models/operations/pinconversationresponse.md), error**

### Errors

| Error Type              | Status Code             | Content Type            |
| ----------------------- | ----------------------- | ----------------------- |
| apierrors.ErrorResponse | 400, 401, 403, 404      | application/json        |
| apierrors.ErrorResponse | 500                     | application/json        |
| apierrors.APIError      | 4XX, 5XX                | \*/\*                   |

## UnpinConversation

Remove a conversation's pin so it is listed by recent activity again.

**Access:**

Owner only, as for `PATCH /conversations/{conversationId}/pin`.

**Behaviour:**

Idempotent: unpinning a conversation that is not pinned succeeds and changes nothing.

- Path params: `conversationId`
- Query params: none
- Body: none


### Example Usage

<!-- UsageSnippet language="go" operationID="unpinConversation" method="patch" path="/conversations/{conversationId}/unpin" -->
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

    res, err := s.Conversations.UnpinConversation(ctx, "<value>")
    if err != nil {
        log.Fatal(err)
    }
    if res.ConversationPinResponse != nil {
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

**[*operations.UnpinConversationResponse](../../models/operations/unpinconversationresponse.md), error**

### Errors

| Error Type              | Status Code             | Content Type            |
| ----------------------- | ----------------------- | ----------------------- |
| apierrors.ErrorResponse | 400, 401, 403, 404      | application/json        |
| apierrors.ErrorResponse | 500                     | application/json        |
| apierrors.APIError      | 4XX, 5XX                | \*/\*                   |

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

**Editing the question:**

Send `query` to re-ask the last question with new wording. The answer
is regenerated for the edited question, which replaces the stored one
once the new answer is saved.

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
        Answers: pipeshub.Pointer("User selections:\n1. \"Which region?\" → EU"),
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

| Error Type              | Status Code             | Content Type            |
| ----------------------- | ----------------------- | ----------------------- |
| apierrors.ErrorResponse | 400, 401, 403           | application/json        |
| apierrors.ErrorResponse | 500                     | application/json        |
| apierrors.APIError      | 4XX, 5XX                | \*/\*                   |

## ForkConversation

Copy this conversation into a new private chat owned by the caller,
including every message up to and including `messageId`. The original
conversation is left unchanged.

Only the owner of the conversation can fork it. A shared recipient or a
project member gets 404, because their access to the original's
attachments and artifacts is revoked on unshare. The fork does not
inherit share settings. If the source is linked to a project the caller
can still see, the fork stays in that project but is private to the
caller; otherwise it is created outside the project.


### Example Usage

<!-- UsageSnippet language="go" operationID="forkConversation" method="post" path="/conversations/{conversationId}/message/{messageId}/fork" -->
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

    res, err := s.Conversations.ForkConversation(ctx, "<value>", "<value>")
    if err != nil {
        log.Fatal(err)
    }
    if res.ForkConversationResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `conversationID`                                         | `string`                                                 | :heavy_check_mark:                                       | N/A                                                      |
| `messageID`                                              | `string`                                                 | :heavy_check_mark:                                       | Inclusive last message to copy into the new chat         |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.ForkConversationResponse](../../models/operations/forkconversationresponse.md), error**

### Errors

| Error Type              | Status Code             | Content Type            |
| ----------------------- | ----------------------- | ----------------------- |
| apierrors.ErrorResponse | 400, 401, 403, 404      | application/json        |
| apierrors.ErrorResponse | 500                     | application/json        |
| apierrors.APIError      | 4XX, 5XX                | \*/\*                   |

## CancelConversationStream

Cooperatively stop a `POST /conversations/stream` or
`POST /conversations/{conversationId}/messages/stream` run that is
still generating, using the `runId` sent when that stream started.

This is a synchronous JSON ack, not another SSE stream. The
cancelled run's own stream (if still connected) receives a terminal
frame with a `stopped` status and whatever partial answer had
already generated; nothing further is delivered here.

A `runId` that already finished or was never registered is not an
error either, but it is answered `{ cancelled: true }`, not `false`:
the query service's KV-backed run registry publishes the cancel for
other workers and acknowledges it as cancelled (a known defect).
`{ cancelled: false }` is only returned when that publish fails.

Only the conversation's owner (matched on `userId`) can cancel; any
other caller, or a missing, deleted or agent conversation, is `404`.

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

| Error Type              | Status Code             | Content Type            |
| ----------------------- | ----------------------- | ----------------------- |
| apierrors.ErrorResponse | 400, 401, 403, 404      | application/json        |
| apierrors.ErrorResponse | 500                     | application/json        |
| apierrors.APIError      | 4XX, 5XX                | \*/\*                   |

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

**Who can submit**

The conversation's initiator, or any user it is shared with (at any
access level) while `isShared` is `true`. Agent conversations are
`404` here.

Body fields other than `isHelpful`, `categories` and `comments` are
dropped by the validator, not refused. `metrics` is always computed
by the server.


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

| Error Type              | Status Code             | Content Type            |
| ----------------------- | ----------------------- | ----------------------- |
| apierrors.ErrorResponse | 400, 401, 403, 404      | application/json        |
| apierrors.ErrorResponse | 500                     | application/json        |
| apierrors.APIError      | 4XX, 5XX                | \*/\*                   |