# MessagePartStatus

## Example Usage

```go
import (
	"github.com/pipeshub-ai/pipeshub-sdk-go/models/components"
)

value := components.MessagePartStatusRunning

// Open enum: custom values can be created with a direct type cast
custom := components.MessagePartStatus("custom_value")
```


## Values

| Name                         | Value                        |
| ---------------------------- | ---------------------------- |
| `MessagePartStatusRunning`   | running                      |
| `MessagePartStatusCompleted` | completed                    |
| `MessagePartStatusFailed`    | failed                       |
| `MessagePartStatusBlocked`   | blocked                      |