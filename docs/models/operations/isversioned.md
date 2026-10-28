# IsVersioned

Enable version tracking. `true`/`1` and `false`/`0`/empty are read as
booleans; any other text is answered `500` (the validator throws a plain
error).


## Example Usage

```go
import (
	"github.com/pipeshub-ai/pipeshub-sdk-go/models/operations"
)

value := operations.IsVersionedTrue
```


## Values

| Name                 | Value                |
| -------------------- | -------------------- |
| `IsVersionedTrue`    | true                 |
| `IsVersionedFalse`   | false                |
| `IsVersionedOne`     | 1                    |
| `IsVersionedZero`    | 0                    |
| `IsVersionedUnknown` |                      |