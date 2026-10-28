# ParentType

Type of the parent node whose children to retrieve.

Must be one of: `app`, `recordGroup`, `folder`, `record` (case-sensitive).
Any other value returns a 400 error, and so does a type that does not
match what `parentId` is (for example `folder` with a knowledge base id).


## Example Usage

```go
import (
	"github.com/pipeshub-ai/pipeshub-sdk-go/v2/models/operations"
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