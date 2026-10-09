# MessageFeedbackCategory

## Example Usage

```go
import (
	"github.com/pipeshub-ai/pipeshub-sdk-go/models/components"
)

value := components.MessageFeedbackCategoryIncorrectInformation

// Open enum: custom values can be created with a direct type cast
custom := components.MessageFeedbackCategory("custom_value")
```


## Values

| Name                                           | Value                                          |
| ---------------------------------------------- | ---------------------------------------------- |
| `MessageFeedbackCategoryIncorrectInformation`  | incorrect_information                          |
| `MessageFeedbackCategoryMissingInformation`    | missing_information                            |
| `MessageFeedbackCategoryIrrelevantInformation` | irrelevant_information                         |
| `MessageFeedbackCategoryUnclearExplanation`    | unclear_explanation                            |
| `MessageFeedbackCategoryPoorCitations`         | poor_citations                                 |
| `MessageFeedbackCategoryExcellentAnswer`       | excellent_answer                               |
| `MessageFeedbackCategoryHelpfulCitations`      | helpful_citations                              |
| `MessageFeedbackCategoryWellExplained`         | well_explained                                 |
| `MessageFeedbackCategoryOther`                 | other                                          |