# ArchiveSearchShared

Additional `isShared` filter. `'true'` / `'1'` and `'false'` / `'0'`,
case-insensitive; anything else is refused with `400`. The search is only archived
if it matches, otherwise `404`.


## Example Usage

```go
import (
	"github.com/pipeshub-ai/pipeshub-sdk-go/v2/models/operations"
)

value := operations.ArchiveSearchSharedTrue
```


## Values

| Name                       | Value                      |
| -------------------------- | -------------------------- |
| `ArchiveSearchSharedTrue`  | true                       |
| `ArchiveSearchSharedFalse` | false                      |
| `ArchiveSearchSharedOne`   | 1                          |
| `ArchiveSearchSharedZero`  | 0                          |