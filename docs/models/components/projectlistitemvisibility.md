# ProjectListItemVisibility

`org` makes the project (and, per `chatSharing`, its
`project`-visible conversations) readable by every member of the
organization, without adding them to `members[]`. Also grants
the organization's synthetic all-members team `READER` access
on the linked hidden Collection.


## Example Usage

```go
import (
	"github.com/pipeshub-ai/pipeshub-sdk-go/models/components"
)

value := components.ProjectListItemVisibilityPrivate

// Open enum: custom values can be created with a direct type cast
custom := components.ProjectListItemVisibility("custom_value")
```


## Values

| Name                               | Value                              |
| ---------------------------------- | ---------------------------------- |
| `ProjectListItemVisibilityPrivate` | private                            |
| `ProjectListItemVisibilityOrg`     | org                                |