# Role

Organization role stored on the user document (`admin` or `member`)

## Example Usage

```go
import (
	"github.com/pipeshub-ai/pipeshub-sdk-go/models/components"
)

value := components.RoleAdmin

// Open enum: custom values can be created with a direct type cast
custom := components.Role("custom_value")
```


## Values

| Name         | Value        |
| ------------ | ------------ |
| `RoleAdmin`  | admin        |
| `RoleMember` | member       |