# GetAllConversationsSourceResponse

Echoes the requested `source` query value.

## Example Usage

```go
import (
	"github.com/pipeshub-ai/pipeshub-sdk-go/models/operations"
)

value := operations.GetAllConversationsSourceResponseOwned

// Open enum: custom values can be created with a direct type cast
custom := operations.GetAllConversationsSourceResponse("custom_value")
```


## Values

| Name                                      | Value                                     |
| ----------------------------------------- | ----------------------------------------- |
| `GetAllConversationsSourceResponseOwned`  | owned                                     |
| `GetAllConversationsSourceResponseShared` | shared                                    |