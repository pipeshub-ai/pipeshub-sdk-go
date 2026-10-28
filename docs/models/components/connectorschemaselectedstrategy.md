# ConnectorSchemaSelectedStrategy

Strategy preselected in the form

## Example Usage

```go
import (
	"github.com/pipeshub-ai/pipeshub-sdk-go/v2/models/components"
)

value := components.ConnectorSchemaSelectedStrategyManual

// Open enum: custom values can be created with a direct type cast
custom := components.ConnectorSchemaSelectedStrategy("custom_value")
```


## Values

| Name                                       | Value                                      |
| ------------------------------------------ | ------------------------------------------ |
| `ConnectorSchemaSelectedStrategyManual`    | MANUAL                                     |
| `ConnectorSchemaSelectedStrategyScheduled` | SCHEDULED                                  |
| `ConnectorSchemaSelectedStrategyWebhook`   | WEBHOOK                                    |