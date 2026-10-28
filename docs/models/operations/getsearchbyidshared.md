# GetSearchByIDShared

Return the search only if its `isShared` matches. `'true'` / `'1'` and `'false'` /
`'0'`, case-insensitive; anything else is refused with `400`.


## Example Usage

```go
import (
	"github.com/pipeshub-ai/pipeshub-sdk-go/v2/models/operations"
)

value := operations.GetSearchByIDSharedTrue
```


## Values

| Name                       | Value                      |
| -------------------------- | -------------------------- |
| `GetSearchByIDSharedTrue`  | true                       |
| `GetSearchByIDSharedFalse` | false                      |
| `GetSearchByIDSharedOne`   | 1                          |
| `GetSearchByIDSharedZero`  | 0                          |