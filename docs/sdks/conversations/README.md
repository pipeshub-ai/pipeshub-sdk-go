# Conversations

## Overview

AI-powered conversational chat management with citations and follow-up questions

### Available Operations

* [Create](#create) - Create a new AI conversation
* [Stream](#stream) - Create conversation with streaming response
* [List](#list) - List all conversations
* [ListArchives](#listarchives) - List archived conversations
* [Get](#get) - Get conversation by ID
* [Delete](#delete) - Delete conversation
* [AddMessage](#addmessage) - Add message to conversation
* [AddMessageStream](#addmessagestream) - Add message with streaming response
* [Share](#share) - Share conversation with users
* [UpdateTitle](#updatetitle) - Update conversation title
* [Archive](#archive) - Archive conversation
* [Unarchive](#unarchive) - Unarchive conversation
* [Regenerate](#regenerate) - Regenerate AI response
* [SubmitFeedback](#submitfeedback) - Submit feedback on AI response
* [Unshare](#unshare) - Unshare a conversation

## Create

Start a new conversation with PipesHub's AI assistant.<br><br>
<b>Overview:</b><br>
This endpoint creates a new conversation session and processes the initial query.
The AI searches your organization's knowledge bases for relevant information and
generates a response with citations to source documents.<br><br>
<b>How It Works:</b><br>
<ol>
<li>Your query is analyzed and converted to semantic embeddings</li>
<li>Relevant content is retrieved from indexed knowledge bases</li>
<li>The AI generates a response using the retrieved context</li>
<li>Citations link back to source documents for verification</li>
<li>Follow-up questions are suggested based on the conversation</li>
</ol>
<b>Filtering Options:</b><br>
<ul>
<li><b>recordIds:</b> Limit search to specific documents</li>
<li><b>filters.apps:</b> Search only specific connector apps</li>
<li><b>filters.kb:</b> Search only specific knowledge bases</li>
</ul>
<b>Model Selection:</b><br>
Use <code>modelKey</code> to select different AI models configured for your organization.
Each model may have different capabilities, speed, and accuracy trade-offs.


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

    res, err := s.Conversations.Create(ctx, components.CreateConversationRequest{
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
    if res.Conversation != nil {
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

    res, err := s.Conversations.Create(ctx, components.CreateConversationRequest{
        Query: "What is our company's vacation policy?",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.Conversation != nil {
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

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## Stream

Start a new conversation with real-time streaming response using Server-Sent Events (SSE).<br><br>
<b>Overview:</b><br>
This endpoint works like <code>/conversations/create</code> but streams the AI response
in real-time as it's generated, providing a more interactive user experience.<br><br>
<b>SSE Event Types:</b><br>
<ul>
<li><code>connected</code> - Connection established, processing started</li>
<li><code>chunk</code> - Partial response text (stream these to show typing effect)</li>
<li><code>citation</code> - Citation reference found during generation</li>
<li><code>complete</code> - Final message with full response, citations, and follow-up questions</li>
<li><code>error</code> - Error occurred during processing</li>
</ul>
<b>Client Implementation:</b><br>
<code>
const eventSource = new EventSource('/conversations/stream');<br>
eventSource.onmessage = (event) => {<br>
&nbsp;&nbsp;const data = JSON.parse(event.data);<br>
&nbsp;&nbsp;// Handle different event types<br>
};
</code><br><br>
<b>Error Handling:</b><br>
If an error occurs mid-stream, an <code>error</code> event is sent and the stream closes.
The conversation is marked as FAILED with the error reason stored.


### Example Usage

<!-- UsageSnippet language="go" operationID="streamChat" method="post" path="/conversations/stream" -->
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

    res, err := s.Conversations.Stream(ctx, components.CreateConversationRequest{
        Query: "What are the key findings from our Q4 financial report?",
        RecordIds: []string{
            "507f1f77bcf86cd799439011",
            "507f1f77bcf86cd799439012",
        },
        ModelKey: pipeshub.Pointer("gpt-4-turbo"),
        ModelName: pipeshub.Pointer("GPT-4 Turbo"),
        ChatMode: pipeshub.Pointer("balanced"),
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

| Parameter                                                                                    | Type                                                                                         | Required                                                                                     | Description                                                                                  |
| -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- |
| `ctx`                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                        | :heavy_check_mark:                                                                           | The context to use for the request.                                                          |
| `request`                                                                                    | [components.CreateConversationRequest](../../models/components/createconversationrequest.md) | :heavy_check_mark:                                                                           | The request object to use for the request.                                                   |
| `opts`                                                                                       | [][operations.Option](../../models/operations/option.md)                                     | :heavy_minus_sign:                                                                           | The options for this request.                                                                |

### Response

**[*operations.StreamChatResponse](../../models/operations/streamchatresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## List

Retrieve all conversations for the authenticated user.<br><br>
<b>Overview:</b><br>
Returns a list of all conversations owned by or shared with the current user.
Conversations are returned with their messages, status, and metadata.<br><br>
<b>Filtering:</b><br>
<ul>
<li>Only non-archived conversations are returned by default</li>
<li>Use <code>/conversations/show/archives</code> for archived conversations</li>
</ul>
<b>Sorting:</b><br>
Conversations are sorted by last activity timestamp (most recent first).


### Example Usage

<!-- UsageSnippet language="go" operationID="getAllConversations" method="get" path="/conversations" -->
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

    res, err := s.Conversations.List(ctx)
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
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.GetAllConversationsResponse](../../models/operations/getallconversationsresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## ListArchives

Retrieve all archived conversations for the authenticated user.<br><br>
<b>Overview:</b><br>
Archived conversations are hidden from the main list but preserved for reference.
This endpoint returns only conversations where <code>isArchived: true</code>.<br><br>
<b>Unarchiving:</b><br>
Use <code>PATCH /conversations/{id}/unarchive</code> to restore a conversation
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
	"log"
)

func main() {
    ctx := context.Background()

    s := pipeshub.New(
        pipeshub.WithSecurity(components.Security{
            BearerAuth: pipeshub.Pointer(os.Getenv("PIPESHUB_BEARER_AUTH")),
        }),
    )

    res, err := s.Conversations.ListArchives(ctx)
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
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.GetArchivedConversationsResponse](../../models/operations/getarchivedconversationsresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## Get

Retrieve a specific conversation with its full message history.<br><br>
<b>Overview:</b><br>
Returns the complete conversation including all messages, citations,
feedback, and metadata. Messages can be paginated for long conversations.<br><br>
<b>Message Pagination:</b><br>
For conversations with many messages, use pagination parameters:
<ul>
<li><code>page</code>: Page number (default: 1)</li>
<li><code>limit</code>: Messages per page (default: 10)</li>
<li><code>sortBy</code>: Sort field (default: createdAt)</li>
<li><code>sortOrder</code>: 'asc' or 'desc' (default: desc)</li>
</ul>
<b>Access Control:</b><br>
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

    res, err := s.Conversations.Get(ctx, operations.GetConversationByIDRequest{
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

## Delete

Delete a conversation by its ID.<br><br>
<b>Overview:</b><br>
Performs a soft delete by setting <code>isDeleted: true</code>.
The conversation is removed from listings but preserved in the database.<br><br>
<b>Permissions:</b><br>
Only the conversation owner (initiator) can delete it.
Shared users cannot delete conversations.


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

    res, err := s.Conversations.Delete(ctx, "<value>")
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
| `conversationID`                                         | *string*                                                 | :heavy_check_mark:                                       | Unique conversation identifier                           |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.DeleteConversationByIDResponse](../../models/operations/deleteconversationbyidresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## AddMessage

Add a follow-up message to an existing conversation.<br><br>
<b>Overview:</b><br>
Continues an existing conversation by adding a new user query.
The AI maintains context from previous messages when generating the response.<br><br>
<b>Context Handling:</b><br>
<ul>
<li>Previous messages provide context for the new query</li>
<li>Citations from earlier messages may be referenced</li>
<li>The AI can refer back to previous topics discussed</li>
</ul>
<b>Model Override:</b><br>
You can specify a different model for this message using <code>modelKey</code>.
This allows switching models mid-conversation if needed.


### Example Usage

<!-- UsageSnippet language="go" operationID="addMessage" method="post" path="/conversations/{conversationId}/messages" -->
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

    res, err := s.Conversations.AddMessage(ctx, "<value>", components.AddMessageRequest{
        Query: "Can you elaborate on the revenue trends?",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.Conversation != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                    | Type                                                                         | Required                                                                     | Description                                                                  |
| ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- |
| `ctx`                                                                        | [context.Context](https://pkg.go.dev/context#Context)                        | :heavy_check_mark:                                                           | The context to use for the request.                                          |
| `conversationID`                                                             | *string*                                                                     | :heavy_check_mark:                                                           | Unique conversation identifier                                               |
| `body`                                                                       | [components.AddMessageRequest](../../models/components/addmessagerequest.md) | :heavy_check_mark:                                                           | Request payload                                                              |
| `opts`                                                                       | [][operations.Option](../../models/operations/option.md)                     | :heavy_minus_sign:                                                           | The options for this request.                                                |

### Response

**[*operations.AddMessageResponse](../../models/operations/addmessageresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## AddMessageStream

Add a follow-up message to an existing conversation with real-time SSE streaming.<br><br>
<b>Overview:</b><br>
Same as <code>POST /conversations/{id}/messages</code> but with streaming response.
Provides real-time feedback as the AI generates its response.<br><br>
<b>SSE Events:</b><br>
See <code>/conversations/stream</code> for event type documentation.


### Example Usage

<!-- UsageSnippet language="go" operationID="addMessageStream" method="post" path="/conversations/{conversationId}/messages/stream" -->
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

    res, err := s.Conversations.AddMessageStream(ctx, "<value>", components.AddMessageRequest{
        Query: "Can you elaborate on the revenue trends?",
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

| Parameter                                                                    | Type                                                                         | Required                                                                     | Description                                                                  |
| ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- |
| `ctx`                                                                        | [context.Context](https://pkg.go.dev/context#Context)                        | :heavy_check_mark:                                                           | The context to use for the request.                                          |
| `conversationID`                                                             | *string*                                                                     | :heavy_check_mark:                                                           | N/A                                                                          |
| `body`                                                                       | [components.AddMessageRequest](../../models/components/addmessagerequest.md) | :heavy_check_mark:                                                           | Request payload                                                              |
| `opts`                                                                       | [][operations.Option](../../models/operations/option.md)                     | :heavy_minus_sign:                                                           | The options for this request.                                                |

### Response

**[*operations.AddMessageStreamResponse](../../models/operations/addmessagestreamresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## Share

Share a conversation with other users in your organization.<br><br>
<b>Overview:</b><br>
Allows the conversation owner to grant access to other users.
Shared users can view the conversation and optionally add messages.<br><br>
<b>Access Levels:</b><br>
<ul>
<li><code>read</code> - Can view conversation and messages (default)</li>
<li><code>write</code> - Can view and add new messages</li>
</ul>
<b>Permissions:</b><br>
Only the conversation initiator (owner) can share. Users must belong
to the same organization.


### Example Usage

<!-- UsageSnippet language="go" operationID="shareConversation" method="post" path="/conversations/{conversationId}/share" -->
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

    res, err := s.Conversations.Share(ctx, "<value>", components.ShareRequest{
        UserIds: []string{
            "507f1f77bcf86cd799439011",
        },
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.Conversation != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                          | Type                                                               | Required                                                           | Description                                                        |
| ------------------------------------------------------------------ | ------------------------------------------------------------------ | ------------------------------------------------------------------ | ------------------------------------------------------------------ |
| `ctx`                                                              | [context.Context](https://pkg.go.dev/context#Context)              | :heavy_check_mark:                                                 | The context to use for the request.                                |
| `conversationID`                                                   | *string*                                                           | :heavy_check_mark:                                                 | N/A                                                                |
| `body`                                                             | [components.ShareRequest](../../models/components/sharerequest.md) | :heavy_check_mark:                                                 | Request payload                                                    |
| `opts`                                                             | [][operations.Option](../../models/operations/option.md)           | :heavy_minus_sign:                                                 | The options for this request.                                      |

### Response

**[*operations.ShareConversationResponse](../../models/operations/shareconversationresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## UpdateTitle

Update the title of a conversation.<br><br>
<b>Overview:</b><br>
Conversation titles are auto-generated from the first query by default.
Use this endpoint to set a custom, more descriptive title.<br><br>
<b>Title Limits:</b><br>
<ul>
<li>Minimum: 1 character</li>
<li>Maximum: 200 characters</li>
</ul>


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

    res, err := s.Conversations.UpdateTitle(ctx, "<value>", operations.UpdateConversationTitleRequestBody{
        Title: "Q4 Sales Analysis Discussion",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.Conversation != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                      | Type                                                                                                           | Required                                                                                                       | Description                                                                                                    |
| -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                                          | :heavy_check_mark:                                                                                             | The context to use for the request.                                                                            |
| `conversationID`                                                                                               | *string*                                                                                                       | :heavy_check_mark:                                                                                             | N/A                                                                                                            |
| `body`                                                                                                         | [operations.UpdateConversationTitleRequestBody](../../models/operations/updateconversationtitlerequestbody.md) | :heavy_check_mark:                                                                                             | Request payload                                                                                                |
| `opts`                                                                                                         | [][operations.Option](../../models/operations/option.md)                                                       | :heavy_minus_sign:                                                                                             | The options for this request.                                                                                  |

### Response

**[*operations.UpdateConversationTitleResponse](../../models/operations/updateconversationtitleresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## Archive

Archive a conversation to hide it from the main list.<br><br>
<b>Overview:</b><br>
Archived conversations are preserved but hidden from the default conversation list.
Use archiving to clean up your workspace without permanently deleting conversations.<br><br>
<b>Retrieval:</b><br>
View archived conversations using <code>GET /conversations/show/archives</code>.


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

    res, err := s.Conversations.Archive(ctx, "<value>")
    if err != nil {
        log.Fatal(err)
    }
    if res.Conversation != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `conversationID`                                         | *string*                                                 | :heavy_check_mark:                                       | N/A                                                      |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.ArchiveConversationResponse](../../models/operations/archiveconversationresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## Unarchive

Restore an archived conversation to the active list.<br><br>
<b>Overview:</b><br>
Removes the archived flag, making the conversation visible in the main list again.


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

    res, err := s.Conversations.Unarchive(ctx, "<value>")
    if err != nil {
        log.Fatal(err)
    }
    if res.Conversation != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `conversationID`                                         | *string*                                                 | :heavy_check_mark:                                       | N/A                                                      |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.UnarchiveConversationResponse](../../models/operations/unarchiveconversationresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## Regenerate

Regenerate the AI response for a specific message.<br><br>
<b>Overview:</b><br>
If you're not satisfied with an AI response, use this endpoint to generate
a new answer. The AI will re-process the original query and may produce
a different response.<br><br>
<b>Use Cases:</b><br>
<ul>
<li>Response was incomplete or unclear</li>
<li>Want to try a different AI model</li>
<li>New documents have been indexed since original response</li>
</ul>
<b>Model Override:</b><br>
Specify <code>modelKey</code> to use a different model for regeneration.


### Example Usage

<!-- UsageSnippet language="go" operationID="regenerateAnswer" method="post" path="/conversations/{conversationId}/message/{messageId}/regenerate" -->
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

    res, err := s.Conversations.Regenerate(ctx, "<value>", "<value>", nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.Conversation != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                         | Type                                                                                              | Required                                                                                          | Description                                                                                       |
| ------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                             | [context.Context](https://pkg.go.dev/context#Context)                                             | :heavy_check_mark:                                                                                | The context to use for the request.                                                               |
| `conversationID`                                                                                  | *string*                                                                                          | :heavy_check_mark:                                                                                | N/A                                                                                               |
| `messageID`                                                                                       | *string*                                                                                          | :heavy_check_mark:                                                                                | ID of the message to regenerate response for                                                      |
| `body`                                                                                            | [*operations.RegenerateAnswerRequestBody](../../models/operations/regenerateanswerrequestbody.md) | :heavy_minus_sign:                                                                                | Request payload                                                                                   |
| `opts`                                                                                            | [][operations.Option](../../models/operations/option.md)                                          | :heavy_minus_sign:                                                                                | The options for this request.                                                                     |

### Response

**[*operations.RegenerateAnswerResponse](../../models/operations/regenerateanswerresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## SubmitFeedback

Provide feedback on an AI-generated response.<br><br>
<b>Overview:</b><br>
Feedback helps improve AI response quality over time. You can rate
various aspects of the response and provide detailed comments.<br><br>
<b>Feedback Options:</b><br>
<ul>
<li><b>isHelpful:</b> Overall thumbs up/down</li>
<li><b>ratings:</b> 1-5 scale for accuracy, relevance, completeness, clarity</li>
<li><b>categories:</b> Issue categories (incorrect info, too verbose, etc.)</li>
<li><b>comments:</b> Free-text positive/negative feedback and suggestions</li>
<li><b>citationFeedback:</b> Rate individual citations</li>
</ul>
<b>Restrictions:</b><br>
Feedback can only be submitted on <code>bot_response</code> messages,
not on user queries or system messages.


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

    res, err := s.Conversations.SubmitFeedback(ctx, "<value>", "<value>", components.MessageFeedback{})
    if err != nil {
        log.Fatal(err)
    }
    if res.Conversation != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                | Type                                                                     | Required                                                                 | Description                                                              |
| ------------------------------------------------------------------------ | ------------------------------------------------------------------------ | ------------------------------------------------------------------------ | ------------------------------------------------------------------------ |
| `ctx`                                                                    | [context.Context](https://pkg.go.dev/context#Context)                    | :heavy_check_mark:                                                       | The context to use for the request.                                      |
| `conversationID`                                                         | *string*                                                                 | :heavy_check_mark:                                                       | N/A                                                                      |
| `messageID`                                                              | *string*                                                                 | :heavy_check_mark:                                                       | N/A                                                                      |
| `body`                                                                   | [components.MessageFeedback](../../models/components/messagefeedback.md) | :heavy_check_mark:                                                       | Request payload                                                          |
| `opts`                                                                   | [][operations.Option](../../models/operations/option.md)                 | :heavy_minus_sign:                                                       | The options for this request.                                            |

### Response

**[*operations.UpdateMessageFeedbackResponse](../../models/operations/updatemessagefeedbackresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## Unshare

Revoke sharing for a conversation, making it private again.


### Example Usage

<!-- UsageSnippet language="go" operationID="unshareConversationById" method="post" path="/conversations/{conversationId}/unshare" -->
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

    res, err := s.Conversations.Unshare(ctx, "<value>", nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.Object != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                       | Type                                                                                                            | Required                                                                                                        | Description                                                                                                     |
| --------------------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                           | [context.Context](https://pkg.go.dev/context#Context)                                                           | :heavy_check_mark:                                                                                              | The context to use for the request.                                                                             |
| `conversationID`                                                                                                | *string*                                                                                                        | :heavy_check_mark:                                                                                              | N/A                                                                                                             |
| `body`                                                                                                          | [*operations.UnshareConversationByIDRequestBody](../../models/operations/unshareconversationbyidrequestbody.md) | :heavy_minus_sign:                                                                                              | Request payload                                                                                                 |
| `opts`                                                                                                          | [][operations.Option](../../models/operations/option.md)                                                        | :heavy_minus_sign:                                                                                              | The options for this request.                                                                                   |

### Response

**[*operations.UnshareConversationByIDResponse](../../models/operations/unshareconversationbyidresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |