# ChatSharing

Owner-controlled default for new conversations created in this
project. `private` keeps new chats visible to their own owner
only; `members` exposes them to every project member
(`projectVisibility: project`). A conversation's own
`projectVisibility` can override this default per-chat.


## Example Usage

```go
import (
	"github.com/pipeshub-ai/pipeshub-sdk-go/models/operations"
)

value := operations.ChatSharingPrivate

// Open enum: custom values can be created with a direct type cast
custom := operations.ChatSharing("custom_value")
```


## Values

| Name                 | Value                |
| -------------------- | -------------------- |
| `ChatSharingPrivate` | private              |
| `ChatSharingMembers` | members              |