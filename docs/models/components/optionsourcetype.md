# OptionSourceType

`dynamic` options come from GET /connectors/{connectorId}/filters/{filterKey}/options

## Example Usage

```go
import (
	"github.com/pipeshub-ai/pipeshub-sdk-go/models/components"
)

value := components.OptionSourceTypeManual

// Open enum: custom values can be created with a direct type cast
custom := components.OptionSourceType("custom_value")
```


## Values

| Name                      | Value                     |
| ------------------------- | ------------------------- |
| `OptionSourceTypeManual`  | manual                    |
| `OptionSourceTypeStatic`  | static                    |
| `OptionSourceTypeDynamic` | dynamic                   |