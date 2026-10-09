# AgentConversationListItemAccessLevel

Computed per request from `sharedWith`; defaults to `read` when no
explicit share grant is attached to the serialized row.


## Example Usage

```go
import (
	"github.com/pipeshub-ai/pipeshub-sdk-go/models/components"
)

value := components.AgentConversationListItemAccessLevelRead

// Open enum: custom values can be created with a direct type cast
custom := components.AgentConversationListItemAccessLevel("custom_value")
```


## Values

| Name                                        | Value                                       |
| ------------------------------------------- | ------------------------------------------- |
| `AgentConversationListItemAccessLevelRead`  | read                                        |
| `AgentConversationListItemAccessLevelWrite` | write                                       |