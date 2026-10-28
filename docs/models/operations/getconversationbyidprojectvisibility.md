# GetConversationByIDProjectVisibility

Who in the project sees the conversation. Absent when never set.

## Example Usage

```go
import (
	"github.com/pipeshub-ai/pipeshub-sdk-go/v2/models/operations"
)

value := operations.GetConversationByIDProjectVisibilityPrivate

// Open enum: custom values can be created with a direct type cast
custom := operations.GetConversationByIDProjectVisibility("custom_value")
```


## Values

| Name                                          | Value                                         |
| --------------------------------------------- | --------------------------------------------- |
| `GetConversationByIDProjectVisibilityPrivate` | private                                       |
| `GetConversationByIDProjectVisibilityProject` | project                                       |