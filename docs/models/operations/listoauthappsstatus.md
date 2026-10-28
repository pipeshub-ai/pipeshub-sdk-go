# ListOAuthAppsStatus

Filter by status. `revoked` always yields an empty page, because an app only becomes revoked by being deleted and deleted apps are never listed.

## Example Usage

```go
import (
	"github.com/pipeshub-ai/pipeshub-sdk-go/v2/models/operations"
)

value := operations.ListOAuthAppsStatusActive
```


## Values

| Name                           | Value                          |
| ------------------------------ | ------------------------------ |
| `ListOAuthAppsStatusActive`    | active                         |
| `ListOAuthAppsStatusSuspended` | suspended                      |
| `ListOAuthAppsStatusRevoked`   | revoked                        |