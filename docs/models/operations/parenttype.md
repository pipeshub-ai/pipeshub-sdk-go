# ParentType

Type of the parent node whose children to retrieve.

Must be one of: `app`, `recordGroup`, `folder`, `record`.
Any other value returns a 400 error.


## Example Usage

```go
import (
	"github.com/pipeshub-ai/pipeshub-sdk-go/models/operations"
)

value := operations.ParentTypeApp
```


## Values

| Name                    | Value                   |
| ----------------------- | ----------------------- |
| `ParentTypeApp`         | app                     |
| `ParentTypeRecordGroup` | recordGroup             |
| `ParentTypeFolder`      | folder                  |
| `ParentTypeRecord`      | record                  |