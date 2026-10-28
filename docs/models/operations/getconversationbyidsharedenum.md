# GetConversationByIDSharedEnum

Filters on the conversation's `isShared`; a conversation that
does not match is answered `404`. Read case-insensitively;
any other value is refused with `400`.


## Example Usage

```go
import (
	"github.com/pipeshub-ai/pipeshub-sdk-go/models/operations"
)

value := operations.GetConversationByIDSharedEnumTrue
```


## Values

| Name                                 | Value                                |
| ------------------------------------ | ------------------------------------ |
| `GetConversationByIDSharedEnumTrue`  | true                                 |
| `GetConversationByIDSharedEnumFalse` | false                                |
| `GetConversationByIDSharedEnumOne`   | 1                                    |
| `GetConversationByIDSharedEnumZero`  | 0                                    |