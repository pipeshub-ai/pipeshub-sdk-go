# RegenerateRequestReasoningEffort

Reasoning effort for this turn, for models that support it. Stored on
the conversation's model info and forwarded to the AI backend.


## Example Usage

```go
import (
	"github.com/pipeshub-ai/pipeshub-sdk-go/v2/models/components"
)

value := components.RegenerateRequestReasoningEffortNone
```


## Values

| Name                                     | Value                                    |
| ---------------------------------------- | ---------------------------------------- |
| `RegenerateRequestReasoningEffortNone`   | none                                     |
| `RegenerateRequestReasoningEffortLow`    | low                                      |
| `RegenerateRequestReasoningEffortMedium` | medium                                   |
| `RegenerateRequestReasoningEffortHigh`   | high                                     |
| `RegenerateRequestReasoningEffortMax`    | max                                      |