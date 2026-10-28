# AuthMethodType

Type of authentication method:
- `password`: Email/password authentication
- `otp`: One-time password via email (6-digit, expires in 10 minutes)
- `google`: Google OAuth 2.0
- `microsoft`: Microsoft OAuth 2.0
- `github`: accepted and stored, but no sign-in flow implements it:
  `/userAccount/initAuth` never lists it and `/userAccount/authenticate` answers `400`

`samlSso`, `oauth` (removed in 1.0) and `azureAd` (now `microsoft`) are refused when a
policy is saved. A stored policy that still names one is read without it (`azureAd`
reads as `microsoft`), with `password` added to the first step.


## Example Usage

```go
import (
	"github.com/pipeshub-ai/pipeshub-sdk-go/models/components"
)

value := components.AuthMethodTypeOtp

// Open enum: custom values can be created with a direct type cast
custom := components.AuthMethodType("custom_value")
```


## Values

| Name                      | Value                     |
| ------------------------- | ------------------------- |
| `AuthMethodTypeOtp`       | otp                       |
| `AuthMethodTypePassword`  | password                  |
| `AuthMethodTypeGoogle`    | google                    |
| `AuthMethodTypeMicrosoft` | microsoft                 |
| `AuthMethodTypeGithub`    | github                    |