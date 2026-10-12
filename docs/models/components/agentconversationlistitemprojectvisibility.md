# AgentConversationListItemProjectVisibility

Only meaningful when `projectId` is set. `project` exposes the
conversation to every member of the linked project.


## Example Usage

```go
import (
	"github.com/pipeshub-ai/pipeshub-sdk-go/models/components"
)

value := components.AgentConversationListItemProjectVisibilityPrivate

// Open enum: custom values can be created with a direct type cast
custom := components.AgentConversationListItemProjectVisibility("custom_value")
```


## Values

| Name                                                | Value                                               |
| --------------------------------------------------- | --------------------------------------------------- |
| `AgentConversationListItemProjectVisibilityPrivate` | private                                             |
| `AgentConversationListItemProjectVisibilityProject` | project                                             |