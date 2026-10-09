# UploadStreamSSEEventEvent

## Example Usage

```go
import (
	"github.com/pipeshub-ai/pipeshub-sdk-go/models/components"
)

value := components.UploadStreamSSEEventEventFileSucceeded

// Open enum: custom values can be created with a direct type cast
custom := components.UploadStreamSSEEventEvent("custom_value")
```


## Values

| Name                                     | Value                                    |
| ---------------------------------------- | ---------------------------------------- |
| `UploadStreamSSEEventEventFileSucceeded` | file:succeeded                           |
| `UploadStreamSSEEventEventFileFailed`    | file:failed                              |
| `UploadStreamSSEEventEventDone`          | done                                     |
| `UploadStreamSSEEventEventError`         | error                                    |