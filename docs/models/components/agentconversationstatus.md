# AgentConversationStatus

Same values as `Conversation.status`.

## Example Usage

```go
import (
	"github.com/pipeshub-ai/pipeshub-sdk-go/models/components"
)

value := components.AgentConversationStatusNone

// Open enum: custom values can be created with a direct type cast
custom := components.AgentConversationStatus("custom_value")
```


## Values

| Name                                | Value                               |
| ----------------------------------- | ----------------------------------- |
| `AgentConversationStatusNone`       | None                                |
| `AgentConversationStatusInprogress` | Inprogress                          |
| `AgentConversationStatusComplete`   | Complete                            |
| `AgentConversationStatusFailed`     | Failed                              |
| `AgentConversationStatusStopped`    | Stopped                             |