# OAuthAppResponseStatus

App status

## Example Usage

```go
import (
	"github.com/pipeshub-ai/pipeshub-sdk-go/models/components"
)

value := components.OAuthAppResponseStatusActive

// Open enum: custom values can be created with a direct type cast
custom := components.OAuthAppResponseStatus("custom_value")
```


## Values

| Name                              | Value                             |
| --------------------------------- | --------------------------------- |
| `OAuthAppResponseStatusActive`    | active                            |
| `OAuthAppResponseStatusSuspended` | suspended                         |
| `OAuthAppResponseStatusRevoked`   | revoked                           |