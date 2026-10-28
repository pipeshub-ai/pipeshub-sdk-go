# UnarchiveSearchShared

Additional `isShared` filter. `'true'` / `'1'` and `'false'` / `'0'`,
case-insensitive; anything else is refused with `400`. The search is only unarchived
if it matches, otherwise `404`.


## Example Usage

```go
import (
	"github.com/pipeshub-ai/pipeshub-sdk-go/v2/models/operations"
)

value := operations.UnarchiveSearchSharedTrue
```


## Values

| Name                         | Value                        |
| ---------------------------- | ---------------------------- |
| `UnarchiveSearchSharedTrue`  | true                         |
| `UnarchiveSearchSharedFalse` | false                        |
| `UnarchiveSearchSharedOne`   | 1                            |
| `UnarchiveSearchSharedZero`  | 0                            |