# MessageMessageType

Type of message:
- `user_query` - User's question or input
- `bot_response` - AI-generated response
- `error` - Error message from the system
- `feedback` - User feedback on a response
- `system` - System notification or status
- `tool_call` - Tool invocation turn; details are on `tools`


## Example Usage

```go
import (
	"github.com/pipeshub-ai/pipeshub-sdk-go/models/components"
)

value := components.MessageMessageTypeUserQuery

// Open enum: custom values can be created with a direct type cast
custom := components.MessageMessageType("custom_value")
```


## Values

| Name                            | Value                           |
| ------------------------------- | ------------------------------- |
| `MessageMessageTypeUserQuery`   | user_query                      |
| `MessageMessageTypeBotResponse` | bot_response                    |
| `MessageMessageTypeError`       | error                           |
| `MessageMessageTypeFeedback`    | feedback                        |
| `MessageMessageTypeSystem`      | system                          |
| `MessageMessageTypeToolCall`    | tool_call                       |