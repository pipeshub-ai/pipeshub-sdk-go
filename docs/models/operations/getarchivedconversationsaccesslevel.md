# GetArchivedConversationsAccessLevel

Computed per request. The requester's effective access level:
their entry in `sharedWith`, or `read` by default.


## Example Usage

```go
import (
	"github.com/pipeshub-ai/pipeshub-sdk-go/models/operations"
)

value := operations.GetArchivedConversationsAccessLevelRead

// Open enum: custom values can be created with a direct type cast
custom := operations.GetArchivedConversationsAccessLevel("custom_value")
```


## Values

| Name                                       | Value                                      |
| ------------------------------------------ | ------------------------------------------ |
| `GetArchivedConversationsAccessLevelRead`  | read                                       |
| `GetArchivedConversationsAccessLevelWrite` | write                                      |