# RecordOrigin

Source of the record:
- UPLOAD: Manually uploaded via API/UI
- CONNECTOR: Synced from external connector


## Example Usage

```go
import (
	"github.com/pipeshub-ai/pipeshub-sdk-go/models/components"
)

value := components.RecordOriginUpload

// Open enum: custom values can be created with a direct type cast
custom := components.RecordOrigin("custom_value")
```


## Values

| Name                    | Value                   |
| ----------------------- | ----------------------- |
| `RecordOriginUpload`    | UPLOAD                  |
| `RecordOriginConnector` | CONNECTOR               |