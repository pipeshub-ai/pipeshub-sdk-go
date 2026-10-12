# ConversationAccessLevel

Computed per request. The requester's effective access level:
their entry in `sharedWith`, or `read` by default.


## Example Usage

```go
import (
	"github.com/pipeshub-ai/pipeshub-sdk-go/models/components"
)

value := components.ConversationAccessLevelRead

// Open enum: custom values can be created with a direct type cast
custom := components.ConversationAccessLevel("custom_value")
```


## Values

| Name                           | Value                          |
| ------------------------------ | ------------------------------ |
| `ConversationAccessLevelRead`  | read                           |
| `ConversationAccessLevelWrite` | write                          |