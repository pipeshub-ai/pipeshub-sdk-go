# ConversationMessageStreamRequestReasoningEffort

How hard a reasoning model should think for this turn. Saved in
the conversation's `modelInfo`.


## Example Usage

```go
import (
	"github.com/pipeshub-ai/pipeshub-sdk-go/v2/models/components"
)

value := components.ConversationMessageStreamRequestReasoningEffortNone
```


## Values

| Name                                                    | Value                                                   |
| ------------------------------------------------------- | ------------------------------------------------------- |
| `ConversationMessageStreamRequestReasoningEffortNone`   | none                                                    |
| `ConversationMessageStreamRequestReasoningEffortLow`    | low                                                     |
| `ConversationMessageStreamRequestReasoningEffortMedium` | medium                                                  |
| `ConversationMessageStreamRequestReasoningEffortHigh`   | high                                                    |
| `ConversationMessageStreamRequestReasoningEffortMax`    | max                                                     |