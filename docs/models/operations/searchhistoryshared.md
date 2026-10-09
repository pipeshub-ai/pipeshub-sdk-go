# SearchHistoryShared

Filter results by their shared status. Accepted values are
`'true'` / `'1'` (return only shared searches) and
`'false'` / `'0'` (exclude shared searches). Matching is
case-insensitive and surrounding whitespace is trimmed.


## Example Usage

```go
import (
	"github.com/pipeshub-ai/pipeshub-sdk-go/models/operations"
)

value := operations.SearchHistorySharedTrue
```


## Values

| Name                       | Value                      |
| -------------------------- | -------------------------- |
| `SearchHistorySharedTrue`  | true                       |
| `SearchHistorySharedFalse` | false                      |
| `SearchHistorySharedOne`   | 1                          |
| `SearchHistorySharedZero`  | 0                          |