# OAuthBearerTokenErrorError

- `invalid_request` - no `Authorization: Bearer` header was sent
- `invalid_token` - the token is malformed, expired, revoked, or its user no longer exists
- `insufficient_scope` - the token is valid but lacks a required scope
- `server_error` - the scope check itself failed


## Example Usage

```go
import (
	"github.com/pipeshub-ai/pipeshub-sdk-go/models/components"
)

value := components.OAuthBearerTokenErrorErrorInvalidRequest

// Open enum: custom values can be created with a direct type cast
custom := components.OAuthBearerTokenErrorError("custom_value")
```


## Values

| Name                                          | Value                                         |
| --------------------------------------------- | --------------------------------------------- |
| `OAuthBearerTokenErrorErrorInvalidRequest`    | invalid_request                               |
| `OAuthBearerTokenErrorErrorInvalidToken`      | invalid_token                                 |
| `OAuthBearerTokenErrorErrorInsufficientScope` | insufficient_scope                            |
| `OAuthBearerTokenErrorErrorServerError`       | server_error                                  |