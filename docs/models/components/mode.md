# Mode

Budgeted reads only. Which kind of read this was.

## Example Usage

```go
import (
	"github.com/pipeshub-ai/pipeshub-sdk-go/v2/models/components"
)

value := components.ModeRecord

// Open enum: custom values can be created with a direct type cast
custom := components.Mode("custom_value")
```


## Values

| Name         | Value        |
| ------------ | ------------ |
| `ModeRecord` | record       |
| `ModeThread` | thread       |