# AgentConversationDetailMessageMessageType

## Example Usage

```go
import (
	"github.com/pipeshub-ai/pipeshub-sdk-go/models/components"
)

value := components.AgentConversationDetailMessageMessageTypeUserQuery

// Open enum: custom values can be created with a direct type cast
custom := components.AgentConversationDetailMessageMessageType("custom_value")
```


## Values

| Name                                                   | Value                                                  |
| ------------------------------------------------------ | ------------------------------------------------------ |
| `AgentConversationDetailMessageMessageTypeUserQuery`   | user_query                                             |
| `AgentConversationDetailMessageMessageTypeBotResponse` | bot_response                                           |
| `AgentConversationDetailMessageMessageTypeError`       | error                                                  |
| `AgentConversationDetailMessageMessageTypeFeedback`    | feedback                                               |
| `AgentConversationDetailMessageMessageTypeSystem`      | system                                                 |
| `AgentConversationDetailMessageMessageTypeToolCall`    | tool_call                                              |