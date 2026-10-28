# ChatAttachmentRefSource

How the attachment was produced: a file upload, or text pasted into the
composer and sent as a file. Stored on the saved message as sent.


## Example Usage

```go
import (
	"github.com/pipeshub-ai/pipeshub-sdk-go/models/components"
)

value := components.ChatAttachmentRefSourceUpload

// Open enum: custom values can be created with a direct type cast
custom := components.ChatAttachmentRefSource("custom_value")
```


## Values

| Name                               | Value                              |
| ---------------------------------- | ---------------------------------- |
| `ChatAttachmentRefSourceUpload`    | upload                             |
| `ChatAttachmentRefSourcePasteText` | paste-text                         |