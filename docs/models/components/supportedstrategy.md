# SupportedStrategy

## Example Usage

```go
import (
	"github.com/pipeshub-ai/pipeshub-sdk-go/v2/models/components"
)

value := components.SupportedStrategyManual

// Open enum: custom values can be created with a direct type cast
custom := components.SupportedStrategy("custom_value")
```


## Values

| Name                         | Value                        |
| ---------------------------- | ---------------------------- |
| `SupportedStrategyManual`    | MANUAL                       |
| `SupportedStrategyScheduled` | SCHEDULED                    |
| `SupportedStrategyWebhook`   | WEBHOOK                      |