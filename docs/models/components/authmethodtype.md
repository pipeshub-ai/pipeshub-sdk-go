# AuthMethodType

Type of authentication method:
- `password`: Email/password authentication
- `otp`: One-time password via email (6-digit, expires in 10 minutes)
- `google`: Google OAuth 2.0
- `microsoft`: Microsoft OAuth 2.0
- `azureAd`: Azure Active Directory
- `samlSso`: SAML 2.0 Single Sign-On
- `oauth`: Generic OAuth 2.0 provider


## Example Usage

```go
import (
	"github.com/pipeshub-ai/pipeshub-sdk-go/models/components"
)

value := components.AuthMethodTypeSamlSso

// Open enum: custom values can be created with a direct type cast
custom := components.AuthMethodType("custom_value")
```


## Values

| Name                      | Value                     |
| ------------------------- | ------------------------- |
| `AuthMethodTypeSamlSso`   | samlSso                   |
| `AuthMethodTypeOtp`       | otp                       |
| `AuthMethodTypePassword`  | password                  |
| `AuthMethodTypeGoogle`    | google                    |
| `AuthMethodTypeMicrosoft` | microsoft                 |
| `AuthMethodTypeAzureAd`   | azureAd                   |
| `AuthMethodTypeOauth`     | oauth                     |