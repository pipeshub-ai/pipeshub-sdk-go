# ConnectorRegistryScope

Scope determines visibility and access control for connectors:<br>
<ul>
<li><code>team</code> - Available to all users in the organization (admin-only creation)</li>
<li><code>personal</code> - Private to the creating user only</li>
</ul>


## Example Usage

```go
import (
	"github.com/pipeshub-ai/pipeshub-sdk-go/models/operations"
)

value := operations.ConnectorRegistryScopeTeam
```


## Values

| Name                             | Value                            |
| -------------------------------- | -------------------------------- |
| `ConnectorRegistryScopeTeam`     | team                             |
| `ConnectorRegistryScopePersonal` | personal                         |