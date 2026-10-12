# UpdateConversationTitleStatus

Current status of the conversation:
- `None` — no activity yet
- `Inprogress` — AI is processing
- `Complete` — response ready
- `Failed` — error occurred
- `Stopped` — cancelled, or the client disconnected mid-answer


## Example Usage

```go
import (
	"github.com/pipeshub-ai/pipeshub-sdk-go/models/operations"
)

value := operations.UpdateConversationTitleStatusNone

// Open enum: custom values can be created with a direct type cast
custom := operations.UpdateConversationTitleStatus("custom_value")
```


## Values

| Name                                      | Value                                     |
| ----------------------------------------- | ----------------------------------------- |
| `UpdateConversationTitleStatusNone`       | None                                      |
| `UpdateConversationTitleStatusInprogress` | Inprogress                                |
| `UpdateConversationTitleStatusComplete`   | Complete                                  |
| `UpdateConversationTitleStatusFailed`     | Failed                                    |
| `UpdateConversationTitleStatusStopped`    | Stopped                                   |