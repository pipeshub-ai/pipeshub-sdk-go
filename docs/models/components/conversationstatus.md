# ConversationStatus

Current status of the conversation:
- `None` — no activity yet
- `Inprogress` — AI is processing
- `Complete` — response ready
- `Failed` — error occurred
- `Stopped` — cancelled, or the client disconnected mid-answer;
  the last message keeps the partial answer


## Example Usage

```go
import (
	"github.com/pipeshub-ai/pipeshub-sdk-go/models/components"
)

value := components.ConversationStatusNone

// Open enum: custom values can be created with a direct type cast
custom := components.ConversationStatus("custom_value")
```


## Values

| Name                           | Value                          |
| ------------------------------ | ------------------------------ |
| `ConversationStatusNone`       | None                           |
| `ConversationStatusInprogress` | Inprogress                     |
| `ConversationStatusComplete`   | Complete                       |
| `ConversationStatusFailed`     | Failed                         |
| `ConversationStatusStopped`    | Stopped                        |