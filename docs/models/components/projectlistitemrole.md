# ProjectListItemRole

Caller's effective role on this project (owner, explicit
member role, or `viewer` via `visibility: org`).


## Example Usage

```go
import (
	"github.com/pipeshub-ai/pipeshub-sdk-go/models/components"
)

value := components.ProjectListItemRoleOwner

// Open enum: custom values can be created with a direct type cast
custom := components.ProjectListItemRole("custom_value")
```


## Values

| Name                        | Value                       |
| --------------------------- | --------------------------- |
| `ProjectListItemRoleOwner`  | owner                       |
| `ProjectListItemRoleEditor` | editor                      |
| `ProjectListItemRoleViewer` | viewer                      |