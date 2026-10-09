# AgentRegenerateRequestProtocol

AG-UI is the only supported wire protocol. When present must be
`"agui"`. Omitting the field is equivalent — the server always
uses the AG-UI vocabulary. Kept in the schema for backward
compatibility with callers that already send it.


## Example Usage

```go
import (
	"github.com/pipeshub-ai/pipeshub-sdk-go/models/components"
)

value := components.AgentRegenerateRequestProtocolAgui
```


## Values

| Name                                 | Value                                |
| ------------------------------------ | ------------------------------------ |
| `AgentRegenerateRequestProtocolAgui` | agui                                 |