# StoredAgentConversationStatus

## Example Usage

```go
import (
	"github.com/pipeshub-ai/pipeshub-sdk-go/models/components"
)

value := components.StoredAgentConversationStatusNone

// Open enum: custom values can be created with a direct type cast
custom := components.StoredAgentConversationStatus("custom_value")
```


## Values

| Name                                      | Value                                     |
| ----------------------------------------- | ----------------------------------------- |
| `StoredAgentConversationStatusNone`       | None                                      |
| `StoredAgentConversationStatusInprogress` | Inprogress                                |
| `StoredAgentConversationStatusComplete`   | Complete                                  |
| `StoredAgentConversationStatusFailed`     | Failed                                    |
| `StoredAgentConversationStatusStopped`    | Stopped                                   |