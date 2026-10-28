# PermissionModel

How records of this connector get their per-user visibility:
`RECORD_LEVEL` (per-record permissions from the source) or
`APP_LEVEL` (everyone who can reach the connector sees every record).


## Example Usage

```go
import (
	"github.com/pipeshub-ai/pipeshub-sdk-go/v2/models/components"
)

value := components.PermissionModelRecordLevel

// Open enum: custom values can be created with a direct type cast
custom := components.PermissionModel("custom_value")
```


## Values

| Name                         | Value                        |
| ---------------------------- | ---------------------------- |
| `PermissionModelRecordLevel` | RECORD_LEVEL                 |
| `PermissionModelAppLevel`    | APP_LEVEL                    |