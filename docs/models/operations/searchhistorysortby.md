# SearchHistorySortBy

Field used to sort results. Any value other than `createdAt`,
`lastActivityAt`, or `title` is treated as `lastActivityAt`.


## Example Usage

```go
import (
	"github.com/pipeshub-ai/pipeshub-sdk-go/models/operations"
)

value := operations.SearchHistorySortByCreatedAt
```


## Values

| Name                                | Value                               |
| ----------------------------------- | ----------------------------------- |
| `SearchHistorySortByCreatedAt`      | createdAt                           |
| `SearchHistorySortByLastActivityAt` | lastActivityAt                      |
| `SearchHistorySortByTitle`          | title                               |