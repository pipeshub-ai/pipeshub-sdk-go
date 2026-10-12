# GrantType

OAuth grant type:
- `authorization_code`: Exchange auth code for tokens
- `client_credentials`: Machine-to-machine auth
- `refresh_token`: Get new access token using refresh token


## Example Usage

```go
import (
	"github.com/pipeshub-ai/pipeshub-sdk-go/models/components"
)

value := components.GrantTypeAuthorizationCode
```


## Values

| Name                         | Value                        |
| ---------------------------- | ---------------------------- |
| `GrantTypeAuthorizationCode` | authorization_code           |
| `GrantTypeClientCredentials` | client_credentials           |
| `GrantTypeRefreshToken`      | refresh_token                |