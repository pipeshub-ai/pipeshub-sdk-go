# QueryParamSource

`owned` — owner list (`userId` filter only).
`shared` — explicit share grant list (`isShared` + `sharedWith`). Includes agent
conversations shared with the caller; those rows carry `agentKey`.
Defaults to `owned` when omitted.


## Example Usage

```go
import (
	"github.com/pipeshub-ai/pipeshub-sdk-go/v2/models/operations"
)

value := operations.QueryParamSourceOwned
```


## Values

| Name                     | Value                    |
| ------------------------ | ------------------------ |
| `QueryParamSourceOwned`  | owned                    |
| `QueryParamSourceShared` | shared                   |