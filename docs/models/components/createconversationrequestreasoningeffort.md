# CreateConversationRequestReasoningEffort

How hard a reasoning model should think for this turn. Saved in
the conversation's `modelInfo`.


## Example Usage

```go
import (
	"github.com/pipeshub-ai/pipeshub-sdk-go/v2/models/components"
)

value := components.CreateConversationRequestReasoningEffortNone
```


## Values

| Name                                             | Value                                            |
| ------------------------------------------------ | ------------------------------------------------ |
| `CreateConversationRequestReasoningEffortNone`   | none                                             |
| `CreateConversationRequestReasoningEffortLow`    | low                                              |
| `CreateConversationRequestReasoningEffortMedium` | medium                                           |
| `CreateConversationRequestReasoningEffortHigh`   | high                                             |
| `CreateConversationRequestReasoningEffortMax`    | max                                              |