# ProjectMemberRole

`viewer` can read the project and its `project`-visible
conversations. `editor` can additionally update project
metadata, instructions, scope, tools, and files. Only the owner
can manage members or change `visibility`/`chatSharing`.


## Example Usage

```go
import (
	"github.com/pipeshub-ai/pipeshub-sdk-go/models/components"
)

value := components.ProjectMemberRoleViewer

// Open enum: custom values can be created with a direct type cast
custom := components.ProjectMemberRole("custom_value")
```


## Values

| Name                      | Value                     |
| ------------------------- | ------------------------- |
| `ProjectMemberRoleViewer` | viewer                    |
| `ProjectMemberRoleEditor` | editor                    |