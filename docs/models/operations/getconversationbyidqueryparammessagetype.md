# GetConversationByIDQueryParamMessageType

Checked against the listed values (`400` otherwise) but not
applied: every message type is returned.


## Example Usage

```go
import (
	"github.com/pipeshub-ai/pipeshub-sdk-go/v2/models/operations"
)

value := operations.GetConversationByIDQueryParamMessageTypeUserQuery
```


## Values

| Name                                                  | Value                                                 |
| ----------------------------------------------------- | ----------------------------------------------------- |
| `GetConversationByIDQueryParamMessageTypeUserQuery`   | user_query                                            |
| `GetConversationByIDQueryParamMessageTypeBotResponse` | bot_response                                          |
| `GetConversationByIDQueryParamMessageTypeError`       | error                                                 |
| `GetConversationByIDQueryParamMessageTypeFeedback`    | feedback                                              |
| `GetConversationByIDQueryParamMessageTypeSystem`      | system                                                |
| `GetConversationByIDQueryParamMessageTypeToolCall`    | tool_call                                             |