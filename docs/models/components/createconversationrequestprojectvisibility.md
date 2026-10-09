# CreateConversationRequestProjectVisibility

Only meaningful together with `projectId`. Overrides the
project's default sharing behavior for this one conversation:
`private` keeps it visible to the owner only; `project` exposes
it to every project member. Defaults from the project's
`chatSharing` setting when omitted.


## Example Usage

```go
import (
	"github.com/pipeshub-ai/pipeshub-sdk-go/models/components"
)

value := components.CreateConversationRequestProjectVisibilityPrivate
```


## Values

| Name                                                | Value                                               |
| --------------------------------------------------- | --------------------------------------------------- |
| `CreateConversationRequestProjectVisibilityPrivate` | private                                             |
| `CreateConversationRequestProjectVisibilityProject` | project                                             |