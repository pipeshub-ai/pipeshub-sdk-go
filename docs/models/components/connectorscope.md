# ConnectorScope

Scope determines visibility and access control for connectors:<br>
<ul>
<li><code>team</code> - Available to all users in the organization (admin-only creation)</li>
<li><code>personal</code> - Private to the creating user only</li>
</ul>


## Example Usage

```go
import (
	"github.com/pipeshub-ai/pipeshub-sdk-go/v2/models/components"
)

value := components.ConnectorScopeTeam

// Open enum: custom values can be created with a direct type cast
custom := components.ConnectorScope("custom_value")
```


## Values

| Name                     | Value                    |
| ------------------------ | ------------------------ |
| `ConnectorScopeTeam`     | team                     |
| `ConnectorScopePersonal` | personal                 |