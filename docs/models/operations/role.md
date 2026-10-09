# Role

## Example Usage

```go
import (
	"github.com/pipeshub-ai/pipeshub-sdk-go/models/operations"
)

value := operations.RoleOwner

// Open enum: custom values can be created with a direct type cast
custom := operations.Role("custom_value")
```


## Values

| Name         | Value        |
| ------------ | ------------ |
| `RoleOwner`  | owner        |
| `RoleEditor` | editor       |
| `RoleViewer` | viewer       |