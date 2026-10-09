# GetArchivedConversationsProjectVisibility

Only meaningful when `projectId` is set. `private` (default)
keeps the conversation visible to its owner only; `project`
exposes it to every member of the linked project. See
`PATCH /conversations/{conversationId}/project-visibility`.


## Example Usage

```go
import (
	"github.com/pipeshub-ai/pipeshub-sdk-go/models/operations"
)

value := operations.GetArchivedConversationsProjectVisibilityPrivate

// Open enum: custom values can be created with a direct type cast
custom := operations.GetArchivedConversationsProjectVisibility("custom_value")
```


## Values

| Name                                               | Value                                              |
| -------------------------------------------------- | -------------------------------------------------- |
| `GetArchivedConversationsProjectVisibilityPrivate` | private                                            |
| `GetArchivedConversationsProjectVisibilityProject` | project                                            |