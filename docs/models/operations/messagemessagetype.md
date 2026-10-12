# MessageMessageType

## Example Usage

```go
import (
	"github.com/pipeshub-ai/pipeshub-sdk-go/models/operations"
)

value := operations.MessageMessageTypeUserQuery

// Open enum: custom values can be created with a direct type cast
custom := operations.MessageMessageType("custom_value")
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