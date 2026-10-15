# StopConnectorSyncStatus

The status the connector is actually left in.

## Example Usage

```go
import (
	"github.com/pipeshub-ai/pipeshub-sdk-go/models/operations"
)

value := operations.StopConnectorSyncStatusIdle

// Open enum: custom values can be created with a direct type cast
custom := operations.StopConnectorSyncStatus("custom_value")
```


## Values

| Name                                 | Value                                |
| ------------------------------------ | ------------------------------------ |
| `StopConnectorSyncStatusIdle`        | IDLE                                 |
| `StopConnectorSyncStatusSyncing`     | SYNCING                              |
| `StopConnectorSyncStatusFullSyncing` | FULL_SYNCING                         |
| `StopConnectorSyncStatusQueued`      | QUEUED                               |
| `StopConnectorSyncStatusDeleting`    | DELETING                             |