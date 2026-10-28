# GetAgentConversationByIDMessageType

Filter messages by type. An empty value is read as not sent.

## Example Usage

```go
import (
	"github.com/pipeshub-ai/pipeshub-sdk-go/v2/models/operations"
)

value := operations.GetAgentConversationByIDMessageTypeUnknown
```


## Values

| Name                                             | Value                                            |
| ------------------------------------------------ | ------------------------------------------------ |
| `GetAgentConversationByIDMessageTypeUnknown`     |                                                  |
| `GetAgentConversationByIDMessageTypeUserQuery`   | user_query                                       |
| `GetAgentConversationByIDMessageTypeBotResponse` | bot_response                                     |
| `GetAgentConversationByIDMessageTypeError`       | error                                            |
| `GetAgentConversationByIDMessageTypeFeedback`    | feedback                                         |
| `GetAgentConversationByIDMessageTypeSystem`      | system                                           |