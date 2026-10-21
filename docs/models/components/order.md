# Order

Budgeted record reads only. `relevance` when `query` chose which blocks to show; the blocks are then not contiguous.

## Example Usage

```go
import (
	"github.com/pipeshub-ai/pipeshub-sdk-go/models/components"
)

value := components.OrderInOrder

// Open enum: custom values can be created with a direct type cast
custom := components.Order("custom_value")
```


## Values

| Name             | Value            |
| ---------------- | ---------------- |
| `OrderInOrder`   | in_order         |
| `OrderRelevance` | relevance        |