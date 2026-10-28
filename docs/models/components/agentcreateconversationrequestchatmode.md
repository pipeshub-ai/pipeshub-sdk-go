# AgentCreateConversationRequestChatMode

Execution mode. Scoped agent conversations support only `quick`.
Required on the `/stream` route; optional on the non-streaming
route.


## Example Usage

```go
import (
	"github.com/pipeshub-ai/pipeshub-sdk-go/v2/models/components"
)

value := components.AgentCreateConversationRequestChatModeQuick
```


## Values

| Name                                          | Value                                         |
| --------------------------------------------- | --------------------------------------------- |
| `AgentCreateConversationRequestChatModeQuick` | quick                                         |