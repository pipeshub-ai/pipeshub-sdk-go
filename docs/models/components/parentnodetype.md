# ParentNodeType

Parent node type; opened as `/nodes/{nodeType}/{id}`.

## Example Usage

```go
import (
	"github.com/pipeshub-ai/pipeshub-sdk-go/v2/models/components"
)

value := components.ParentNodeTypeApp

// Open enum: custom values can be created with a direct type cast
custom := components.ParentNodeType("custom_value")
```


## Values

| Name                        | Value                       |
| --------------------------- | --------------------------- |
| `ParentNodeTypeApp`         | app                         |
| `ParentNodeTypeRecordGroup` | recordGroup                 |
| `ParentNodeTypeFolder`      | folder                      |
| `ParentNodeTypeRecord`      | record                      |