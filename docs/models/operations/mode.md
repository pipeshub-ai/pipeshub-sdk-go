# Mode

`record` (default) reads this record. `thread` reads every record of the Slack thread this message belongs to, oldest first; a record outside any thread answers `422`.

## Example Usage

```go
import (
	"github.com/pipeshub-ai/pipeshub-sdk-go/models/operations"
)

value := operations.ModeRecord
```


## Values

| Name         | Value        |
| ------------ | ------------ |
| `ModeRecord` | record       |
| `ModeThread` | thread       |