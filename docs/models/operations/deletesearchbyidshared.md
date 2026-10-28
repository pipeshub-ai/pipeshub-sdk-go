# DeleteSearchByIDShared

Additional `isShared` filter. `'true'` / `'1'` and `'false'` / `'0'`,
case-insensitive; anything else is refused with `400`. The row is only deleted if
it matches.


## Example Usage

```go
import (
	"github.com/pipeshub-ai/pipeshub-sdk-go/v2/models/operations"
)

value := operations.DeleteSearchByIDSharedTrue
```


## Values

| Name                          | Value                         |
| ----------------------------- | ----------------------------- |
| `DeleteSearchByIDSharedTrue`  | true                          |
| `DeleteSearchByIDSharedFalse` | false                         |
| `DeleteSearchByIDSharedOne`   | 1                             |
| `DeleteSearchByIDSharedZero`  | 0                             |