# PrincipalType

Whether `principalId` names a user or a team. Both grant the same `role`, mirroring Collection sharing.

## Example Usage

```go
import (
	"github.com/pipeshub-ai/pipeshub-sdk-go/models/components"
)

value := components.PrincipalTypeUser

// Open enum: custom values can be created with a direct type cast
custom := components.PrincipalType("custom_value")
```


## Values

| Name                | Value               |
| ------------------- | ------------------- |
| `PrincipalTypeUser` | user                |
| `PrincipalTypeTeam` | team                |