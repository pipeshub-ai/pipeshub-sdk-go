# MessageContentFormat

Format of the content for rendering

## Example Usage

```go
import (
	"github.com/pipeshub-ai/pipeshub-sdk-go/models/components"
)

value := components.MessageContentFormatMarkdown

// Open enum: custom values can be created with a direct type cast
custom := components.MessageContentFormat("custom_value")
```


## Values

| Name                           | Value                          |
| ------------------------------ | ------------------------------ |
| `MessageContentFormatMarkdown` | MARKDOWN                       |
| `MessageContentFormatJSON`     | JSON                           |
| `MessageContentFormatHTML`     | HTML                           |