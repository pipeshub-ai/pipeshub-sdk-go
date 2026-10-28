# ConnectorInstanceListScope

Scope determines visibility and access control for connectors:<br>
<ul>
<li><code>team</code> - Available to all users in the organization (admin-only creation)</li>
<li><code>personal</code> - Private to the creating user only</li>
</ul>


## Example Usage

```go
import (
	"github.com/pipeshub-ai/pipeshub-sdk-go/v2/models/operations"
)

value := operations.ConnectorInstanceListScopeTeam
```


## Values

| Name                                 | Value                                |
| ------------------------------------ | ------------------------------------ |
| `ConnectorInstanceListScopeTeam`     | team                                 |
| `ConnectorInstanceListScopePersonal` | personal                             |