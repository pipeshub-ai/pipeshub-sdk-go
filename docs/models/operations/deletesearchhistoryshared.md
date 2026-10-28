# DeleteSearchHistoryShared

Restrict the deletion to rows with this `isShared` value. `'true'` / `'1'` and
`'false'` / `'0'`, case-insensitive; anything else is refused with `400`.


## Example Usage

```go
import (
	"github.com/pipeshub-ai/pipeshub-sdk-go/v2/models/operations"
)

value := operations.DeleteSearchHistorySharedTrue
```


## Values

| Name                             | Value                            |
| -------------------------------- | -------------------------------- |
| `DeleteSearchHistorySharedTrue`  | true                             |
| `DeleteSearchHistorySharedFalse` | false                            |
| `DeleteSearchHistorySharedOne`   | 1                                |
| `DeleteSearchHistorySharedZero`  | 0                                |