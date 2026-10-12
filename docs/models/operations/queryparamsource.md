# QueryParamSource

`owned` — owner list (`userId` filter only).
`shared` — explicit share grant list (`isShared` + `sharedWith`).
Defaults to `owned` when omitted.


## Example Usage

```go
import (
	"github.com/pipeshub-ai/pipeshub-sdk-go/models/operations"
)

value := operations.QueryParamSourceOwned
```


## Values

| Name                     | Value                    |
| ------------------------ | ------------------------ |
| `QueryParamSourceOwned`  | owned                    |
| `QueryParamSourceShared` | shared                   |