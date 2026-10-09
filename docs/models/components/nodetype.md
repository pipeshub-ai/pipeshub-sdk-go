# NodeType

Type of the node (app, recordGroup, folder, or record).

## Example Usage

```go
import (
	"github.com/pipeshub-ai/pipeshub-sdk-go/models/components"
)

value := components.NodeTypeApp

// Open enum: custom values can be created with a direct type cast
custom := components.NodeType("custom_value")
```


## Values

| Name                  | Value                 |
| --------------------- | --------------------- |
| `NodeTypeApp`         | app                   |
| `NodeTypeRecordGroup` | recordGroup           |
| `NodeTypeFolder`      | folder                |
| `NodeTypeRecord`      | record                |