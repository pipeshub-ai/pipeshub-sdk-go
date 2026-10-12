# GetArchivedConversationsStatus

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
	"github.com/pipeshub-ai/pipeshub-sdk-go/models/operations"
)

value := operations.GetArchivedConversationsStatusNone

// Open enum: custom values can be created with a direct type cast
custom := operations.GetArchivedConversationsStatus("custom_value")
```


## Values

| Name                                       | Value                                      |
| ------------------------------------------ | ------------------------------------------ |
| `GetArchivedConversationsStatusNone`       | None                                       |
| `GetArchivedConversationsStatusInprogress` | Inprogress                                 |
| `GetArchivedConversationsStatusComplete`   | Complete                                   |
| `GetArchivedConversationsStatusFailed`     | Failed                                     |
| `GetArchivedConversationsStatusStopped`    | Stopped                                    |