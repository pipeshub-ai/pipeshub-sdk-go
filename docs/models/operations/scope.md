# Scope

`mine` — owned only. `shared` — projects the caller is a member
of. `all` — owned, member, and `visibility: org` projects.
Defaults to `mine`.


## Example Usage

```go
import (
	"github.com/pipeshub-ai/pipeshub-sdk-go/models/operations"
)

value := operations.ScopeMine
```


## Values

| Name          | Value         |
| ------------- | ------------- |
| `ScopeMine`   | mine          |
| `ScopeShared` | shared        |
| `ScopeAll`    | all           |