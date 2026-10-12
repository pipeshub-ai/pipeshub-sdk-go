# AllowedMethod

## Example Usage

```go
import (
	"github.com/pipeshub-ai/pipeshub-sdk-go/models/components"
)

value := components.AllowedMethodSamlSso

// Open enum: custom values can be created with a direct type cast
custom := components.AllowedMethod("custom_value")
```


## Values

| Name                     | Value                    |
| ------------------------ | ------------------------ |
| `AllowedMethodSamlSso`   | samlSso                  |
| `AllowedMethodOtp`       | otp                      |
| `AllowedMethodPassword`  | password                 |
| `AllowedMethodGoogle`    | google                   |
| `AllowedMethodMicrosoft` | microsoft                |
| `AllowedMethodAzureAd`   | azureAd                  |
| `AllowedMethodOauth`     | oauth                    |