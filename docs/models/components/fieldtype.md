# FieldType

## Example Usage

```go
import (
	"github.com/pipeshub-ai/pipeshub-sdk-go/models/components"
)

value := components.FieldTypeText

// Open enum: custom values can be created with a direct type cast
custom := components.FieldType("custom_value")
```


## Values

| Name                | Value               |
| ------------------- | ------------------- |
| `FieldTypeText`     | TEXT                |
| `FieldTypePassword` | PASSWORD            |
| `FieldTypeURL`      | URL                 |
| `FieldTypeSelect`   | SELECT              |