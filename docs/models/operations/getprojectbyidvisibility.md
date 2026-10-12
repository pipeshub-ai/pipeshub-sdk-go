# GetProjectByIDVisibility

`org` makes the project (and, per `chatSharing`, its
`project`-visible conversations) readable by every member of the
organization, without adding them to `members[]`. Also grants
the organization's synthetic all-members team `READER` access
on the linked hidden Collection.


## Example Usage

```go
import (
	"github.com/pipeshub-ai/pipeshub-sdk-go/models/operations"
)

value := operations.GetProjectByIDVisibilityPrivate

// Open enum: custom values can be created with a direct type cast
custom := operations.GetProjectByIDVisibility("custom_value")
```


## Values

| Name                              | Value                             |
| --------------------------------- | --------------------------------- |
| `GetProjectByIDVisibilityPrivate` | private                           |
| `GetProjectByIDVisibilityOrg`     | org                               |