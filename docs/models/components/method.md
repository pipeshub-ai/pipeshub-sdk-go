# Method

Authentication method to use. It must be one the session's current step allows;
any other value, known or not, is answered `400`.


## Example Usage

```go
import (
	"github.com/pipeshub-ai/pipeshub-sdk-go/models/components"
)

value := components.MethodOtp
```


## Values

| Name              | Value             |
| ----------------- | ----------------- |
| `MethodOtp`       | otp               |
| `MethodPassword`  | password          |
| `MethodGoogle`    | google            |
| `MethodMicrosoft` | microsoft         |