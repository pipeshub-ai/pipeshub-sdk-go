# ListAgentConversationsIsArchived

Optional archived flag applied to the `sharedWithMeConversations`
branch before the route-level non-archived guard is enforced.
Accepted values are `true` and `false`.


## Example Usage

```go
import (
	"github.com/pipeshub-ai/pipeshub-sdk-go/models/operations"
)

value := operations.ListAgentConversationsIsArchivedTrue
```


## Values

| Name                                    | Value                                   |
| --------------------------------------- | --------------------------------------- |
| `ListAgentConversationsIsArchivedTrue`  | true                                    |
| `ListAgentConversationsIsArchivedFalse` | false                                   |