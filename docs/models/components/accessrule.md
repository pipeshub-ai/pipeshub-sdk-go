# AccessRule

How this record's own permissions combine with its ancestors'.
OPEN: an accessible parent, or a direct grant, is enough.
STRICT: every ancestor up to the connector must be accessible too.
RESTRICTED: strict, and a grant on the record itself is required.
Absent on records written before the field existed; read as OPEN.


## Example Usage

```go
import (
	"github.com/pipeshub-ai/pipeshub-sdk-go/v2/models/components"
)

value := components.AccessRuleOpen

// Open enum: custom values can be created with a direct type cast
custom := components.AccessRule("custom_value")
```


## Values

| Name                   | Value                  |
| ---------------------- | ---------------------- |
| `AccessRuleOpen`       | OPEN                   |
| `AccessRuleStrict`     | STRICT                 |
| `AccessRuleRestricted` | RESTRICTED             |