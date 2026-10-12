# AgentConversationListItemStatus

## Example Usage

```go
import (
	"github.com/pipeshub-ai/pipeshub-sdk-go/models/components"
)

value := components.AgentConversationListItemStatusNone

// Open enum: custom values can be created with a direct type cast
custom := components.AgentConversationListItemStatus("custom_value")
```


## Values

| Name                                        | Value                                       |
| ------------------------------------------- | ------------------------------------------- |
| `AgentConversationListItemStatusNone`       | None                                        |
| `AgentConversationListItemStatusInprogress` | Inprogress                                  |
| `AgentConversationListItemStatusComplete`   | Complete                                    |
| `AgentConversationListItemStatusFailed`     | Failed                                      |
| `AgentConversationListItemStatusStopped`    | Stopped                                     |