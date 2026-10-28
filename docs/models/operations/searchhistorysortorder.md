# SearchHistorySortOrder

Sort direction applied to `sortBy`. Only the exact, case-sensitive value `asc` sorts
ascending; any other value (`ASC`, `sideways`, ...) is accepted and sorts descending.
The raw value is echoed in `filters.applied.values.sortOrder`.


## Example Usage

```go
import (
	"github.com/pipeshub-ai/pipeshub-sdk-go/models/operations"
)

value := operations.SearchHistorySortOrderAsc
```


## Values

| Name                         | Value                        |
| ---------------------------- | ---------------------------- |
| `SearchHistorySortOrderAsc`  | asc                          |
| `SearchHistorySortOrderDesc` | desc                         |