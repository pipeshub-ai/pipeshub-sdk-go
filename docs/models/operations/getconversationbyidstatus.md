# GetConversationByIDStatus

## Example Usage

```go
import (
	"github.com/pipeshub-ai/pipeshub-sdk-go/models/operations"
)

value := operations.GetConversationByIDStatusNone

// Open enum: custom values can be created with a direct type cast
custom := operations.GetConversationByIDStatus("custom_value")
```


## Values

| Name                                  | Value                                 |
| ------------------------------------- | ------------------------------------- |
| `GetConversationByIDStatusNone`       | None                                  |
| `GetConversationByIDStatusInprogress` | Inprogress                            |
| `GetConversationByIDStatusComplete`   | Complete                              |
| `GetConversationByIDStatusFailed`     | Failed                                |
| `GetConversationByIDStatusStopped`    | Stopped                               |