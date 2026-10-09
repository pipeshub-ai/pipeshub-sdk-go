# ModelType

Type of AI model

## Example Usage

```go
import (
	"github.com/pipeshub-ai/pipeshub-sdk-go/models/components"
)

value := components.ModelTypeLlm

// Open enum: custom values can be created with a direct type cast
custom := components.ModelType("custom_value")
```


## Values

| Name                       | Value                      |
| -------------------------- | -------------------------- |
| `ModelTypeLlm`             | llm                        |
| `ModelTypeEmbedding`       | embedding                  |
| `ModelTypeOcr`             | ocr                        |
| `ModelTypeSlm`             | slm                        |
| `ModelTypeReasoning`       | reasoning                  |
| `ModelTypeMultiModal`      | multiModal                 |
| `ModelTypeImageGeneration` | imageGeneration            |
| `ModelTypeTts`             | tts                        |
| `ModelTypeStt`             | stt                        |