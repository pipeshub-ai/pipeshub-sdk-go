# Kind

Whether this account is a person who signs in (`human`) or a
machine identity that automation authenticates as (`service`).
Absent on records written before service accounts existed, which
are all human.


## Example Usage

```go
import (
	"github.com/pipeshub-ai/pipeshub-sdk-go/models/components"
)

value := components.KindHuman

// Open enum: custom values can be created with a direct type cast
custom := components.Kind("custom_value")
```


## Values

| Name          | Value         |
| ------------- | ------------- |
| `KindHuman`   | human         |
| `KindService` | service       |