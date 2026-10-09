# MetadataCode

Machine-readable error code mapped from the Zod issue code.

## Example Usage

```go
import (
	"github.com/pipeshub-ai/pipeshub-sdk-go/models/operations"
)

value := operations.MetadataCodeInvalidType

// Open enum: custom values can be created with a direct type cast
custom := operations.MetadataCode("custom_value")
```


## Values

| Name                               | Value                              |
| ---------------------------------- | ---------------------------------- |
| `MetadataCodeInvalidType`          | INVALID_TYPE                       |
| `MetadataCodeInvalidLiteral`       | INVALID_LITERAL                    |
| `MetadataCodeInvalidEnum`          | INVALID_ENUM                       |
| `MetadataCodeInvalidUnion`         | INVALID_UNION                      |
| `MetadataCodeInvalidDiscriminator` | INVALID_DISCRIMINATOR              |
| `MetadataCodeInvalidArguments`     | INVALID_ARGUMENTS                  |
| `MetadataCodeTooSmall`             | TOO_SMALL                          |
| `MetadataCodeTooBig`               | TOO_BIG                            |