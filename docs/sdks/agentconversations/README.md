# AgentConversations

## Overview

### Available Operations

* [List](#list) - List agent conversations
* [CreateConversation](#createconversation) - Create agent conversation
* [Stream](#stream) - Create agent conversation with streaming
* [Get](#get) - Get agent conversation
* [Delete](#delete) - Delete agent conversation
* [AddMessage](#addmessage) - Add message to agent conversation
* [StreamMessage](#streammessage) - Add message with streaming
* [RegenerateAnswer](#regenerateanswer) - Regenerate agent response

## List

Get all conversations with a specific agent.<br><br>
<b>Overview:</b><br>
Returns conversations the user has had with this particular agent.
Agent conversations maintain the agent's context and capabilities.


### Example Usage

<!-- UsageSnippet language="go" operationID="listAgentConversations" method="get" path="/agents/{agentKey}/conversations" -->
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

    res, err := s.AgentConversations.List(ctx, "<value>")
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
| `agentKey`                                               | *string*                                                 | :heavy_check_mark:                                       | Agent identifier                                         |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.ListAgentConversationsResponse](../../models/operations/listagentconversationsresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## CreateConversation

Start a new conversation with an agent.<br><br>
<b>Overview:</b><br>
Creates a conversation using the agent's configuration including
its system prompt, tools, and knowledge base access.


### Example Usage

<!-- UsageSnippet language="go" operationID="createAgentConversation" method="post" path="/agents/{agentKey}/conversations" -->
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

    res, err := s.AgentConversations.CreateConversation(ctx, "<value>", components.CreateConversationRequest{
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
    if res.AgentConversation != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                    | Type                                                                                         | Required                                                                                     | Description                                                                                  |
| -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- |
| `ctx`                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                        | :heavy_check_mark:                                                                           | The context to use for the request.                                                          |
| `agentKey`                                                                                   | *string*                                                                                     | :heavy_check_mark:                                                                           | N/A                                                                                          |
| `body`                                                                                       | [components.CreateConversationRequest](../../models/components/createconversationrequest.md) | :heavy_check_mark:                                                                           | Request payload                                                                              |
| `opts`                                                                                       | [][operations.Option](../../models/operations/option.md)                                     | :heavy_minus_sign:                                                                           | The options for this request.                                                                |

### Response

**[*operations.CreateAgentConversationResponse](../../models/operations/createagentconversationresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## Stream

Start a new agent conversation with SSE streaming response.<br><br>
<b>Overview:</b><br>
Same as POST /agents/{agentKey}/conversations but with real-time streaming.


### Example Usage

<!-- UsageSnippet language="go" operationID="streamAgentConversation" method="post" path="/agents/{agentKey}/conversations/stream" -->
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

    res, err := s.AgentConversations.Stream(ctx, "<value>", components.CreateConversationRequest{
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
| `agentKey`                                                                                   | *string*                                                                                     | :heavy_check_mark:                                                                           | N/A                                                                                          |
| `body`                                                                                       | [components.CreateConversationRequest](../../models/components/createconversationrequest.md) | :heavy_check_mark:                                                                           | Request body for Create agent conversation with streaming                                    |
| `opts`                                                                                       | [][operations.Option](../../models/operations/option.md)                                     | :heavy_minus_sign:                                                                           | The options for this request.                                                                |

### Response

**[*operations.StreamAgentConversationResponse](../../models/operations/streamagentconversationresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## Get

Retrieve a specific agent conversation by ID.

### Example Usage

<!-- UsageSnippet language="go" operationID="getAgentConversation" method="get" path="/agents/{agentKey}/conversations/{conversationId}" -->
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

    res, err := s.AgentConversations.Get(ctx, "<value>", "<value>")
    if err != nil {
        log.Fatal(err)
    }
    if res.AgentConversation != nil {
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

**[*operations.GetAgentConversationResponse](../../models/operations/getagentconversationresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## Delete

Delete a conversation with an agent.

### Example Usage

<!-- UsageSnippet language="go" operationID="deleteAgentConversation" method="delete" path="/agents/{agentKey}/conversations/{conversationId}" -->
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

    res, err := s.AgentConversations.Delete(ctx, "<value>", "<value>")
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
| `conversationID`                                         | *string*                                                 | :heavy_check_mark:                                       | N/A                                                      |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.DeleteAgentConversationResponse](../../models/operations/deleteagentconversationresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## AddMessage

Add a follow-up message to an agent conversation.

### Example Usage

<!-- UsageSnippet language="go" operationID="addAgentMessage" method="post" path="/agents/{agentKey}/conversations/{conversationId}/messages" -->
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

    res, err := s.AgentConversations.AddMessage(ctx, "<value>", "<value>", components.AddMessageRequest{
        Query: "Can you elaborate on the revenue trends?",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.AgentConversation != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                    | Type                                                                         | Required                                                                     | Description                                                                  |
| ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- |
| `ctx`                                                                        | [context.Context](https://pkg.go.dev/context#Context)                        | :heavy_check_mark:                                                           | The context to use for the request.                                          |
| `agentKey`                                                                   | *string*                                                                     | :heavy_check_mark:                                                           | N/A                                                                          |
| `conversationID`                                                             | *string*                                                                     | :heavy_check_mark:                                                           | N/A                                                                          |
| `body`                                                                       | [components.AddMessageRequest](../../models/components/addmessagerequest.md) | :heavy_check_mark:                                                           | Request payload                                                              |
| `opts`                                                                       | [][operations.Option](../../models/operations/option.md)                     | :heavy_minus_sign:                                                           | The options for this request.                                                |

### Response

**[*operations.AddAgentMessageResponse](../../models/operations/addagentmessageresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## StreamMessage

Add a message to agent conversation with SSE streaming response.

### Example Usage

<!-- UsageSnippet language="go" operationID="streamAgentMessage" method="post" path="/agents/{agentKey}/conversations/{conversationId}/messages/stream" -->
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

    res, err := s.AgentConversations.StreamMessage(ctx, "<value>", "<value>", components.AddMessageRequest{
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
| `agentKey`                                                                   | *string*                                                                     | :heavy_check_mark:                                                           | N/A                                                                          |
| `conversationID`                                                             | *string*                                                                     | :heavy_check_mark:                                                           | N/A                                                                          |
| `body`                                                                       | [components.AddMessageRequest](../../models/components/addmessagerequest.md) | :heavy_check_mark:                                                           | Request payload                                                              |
| `opts`                                                                       | [][operations.Option](../../models/operations/option.md)                     | :heavy_minus_sign:                                                           | The options for this request.                                                |

### Response

**[*operations.StreamAgentMessageResponse](../../models/operations/streamagentmessageresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## RegenerateAnswer

Regenerate the agent's response for a specific message.<br><br>
<b>Overview:</b><br>
Similar to conversation regeneration but uses the agent's configuration.


### Example Usage

<!-- UsageSnippet language="go" operationID="regenerateAgentAnswer" method="post" path="/agents/{agentKey}/conversations/{conversationId}/message/{messageId}/regenerate" -->
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

    res, err := s.AgentConversations.RegenerateAnswer(ctx, "<value>", "<value>", "<value>", nil)
    if err != nil {
        log.Fatal(err)
    }
    if res.AgentConversation != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                   | Type                                                                                                        | Required                                                                                                    | Description                                                                                                 |
| ----------------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                       | [context.Context](https://pkg.go.dev/context#Context)                                                       | :heavy_check_mark:                                                                                          | The context to use for the request.                                                                         |
| `agentKey`                                                                                                  | *string*                                                                                                    | :heavy_check_mark:                                                                                          | N/A                                                                                                         |
| `conversationID`                                                                                            | *string*                                                                                                    | :heavy_check_mark:                                                                                          | N/A                                                                                                         |
| `messageID`                                                                                                 | *string*                                                                                                    | :heavy_check_mark:                                                                                          | N/A                                                                                                         |
| `body`                                                                                                      | [*operations.RegenerateAgentAnswerRequestBody](../../models/operations/regenerateagentanswerrequestbody.md) | :heavy_minus_sign:                                                                                          | Request payload                                                                                             |
| `opts`                                                                                                      | [][operations.Option](../../models/operations/option.md)                                                    | :heavy_minus_sign:                                                                                          | The options for this request.                                                                               |

### Response

**[*operations.RegenerateAgentAnswerResponse](../../models/operations/regenerateagentanswerresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |