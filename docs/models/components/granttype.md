# GrantType

OAuth grant type:
- `authorization_code`: Exchange auth code for tokens
- `client_credentials`: Machine-to-machine auth
- `refresh_token`: Get new access token using refresh token
- `urn:ietf:params:oauth:grant-type:device_code`: Poll for the tokens of a device
  authorization (RFC 8628) started with `POST /oauth2/device_authorization`
- `urn:ietf:params:oauth:grant-type:token-exchange`: Service-account
  delegation (RFC 8693). Exchanges a service token for a
  short-lived token that acts as one named person. Enterprise
  edition; kept in this enum because the request shape is the same
  wherever the grant is offered.


## Example Usage

```go
import (
	"github.com/pipeshub-ai/pipeshub-sdk-go/models/components"
)

value := components.GrantTypeAuthorizationCode
```


## Values

| Name                                                | Value                                               |
| --------------------------------------------------- | --------------------------------------------------- |
| `GrantTypeAuthorizationCode`                        | authorization_code                                  |
| `GrantTypeClientCredentials`                        | client_credentials                                  |
| `GrantTypeRefreshToken`                             | refresh_token                                       |
| `GrantTypeUrnIetfParamsOauthGrantTypeDeviceCode`    | urn:ietf:params:oauth:grant-type:device_code        |
| `GrantTypeUrnIetfParamsOauthGrantTypeTokenExchange` | urn:ietf:params:oauth:grant-type:token-exchange     |