# MessageFeedbackSource

Origin of the feedback. Always present in responses (server applies the default `user`).

## Example Usage

```go
import (
	"github.com/pipeshub-ai/pipeshub-sdk-go/v2/models/components"
)

value := components.MessageFeedbackSourceUser

// Open enum: custom values can be created with a direct type cast
custom := components.MessageFeedbackSource("custom_value")
```


## Values

| Name                          | Value                         |
| ----------------------------- | ----------------------------- |
| `MessageFeedbackSourceUser`   | user                          |
| `MessageFeedbackSourceSystem` | system                        |
| `MessageFeedbackSourceAdmin`  | admin                         |
| `MessageFeedbackSourceAuto`   | auto                          |