# Order

Budgeted record reads only. `relevance` when `query` chose which blocks to show; the blocks are then not contiguous. `extract` when `extract=true` was asked and the reader model answered: `content` is the passages it copied out (told to copy word for word; see `extract.unverifiedLines`), not a page of the record. An extract is one response, so `complete` is true and `nextStartBlock` null; whether the reader saw the whole record is `extract.sourceComplete`.

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
| `OrderExtract`   | extract          |