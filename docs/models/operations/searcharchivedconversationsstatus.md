# SearchArchivedConversationsStatus

## Example Usage

```go
import (
	"github.com/pipeshub-ai/pipeshub-sdk-go/v2/models/operations"
)

value := operations.SearchArchivedConversationsStatusNone

// Open enum: custom values can be created with a direct type cast
custom := operations.SearchArchivedConversationsStatus("custom_value")
```


## Values

| Name                                          | Value                                         |
| --------------------------------------------- | --------------------------------------------- |
| `SearchArchivedConversationsStatusNone`       | None                                          |
| `SearchArchivedConversationsStatusInprogress` | Inprogress                                    |
| `SearchArchivedConversationsStatusComplete`   | Complete                                      |
| `SearchArchivedConversationsStatusFailed`     | Failed                                        |
| `SearchArchivedConversationsStatusStopped`    | Stopped                                       |