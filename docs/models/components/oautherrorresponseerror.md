# OAuthErrorResponseError

Error code. Common values:
- `invalid_request` - Missing or invalid parameter
- `invalid_client` - Client authentication failed
- `invalid_grant` - Invalid authorization code or refresh token
- `unauthorized_client` - Client not authorized for this grant type
- `unsupported_grant_type` - Grant type not supported
- `invalid_scope` - Requested scope is invalid or exceeds allowed
- `access_denied` - User denied authorization
- `invalid_target` - The requested subject cannot be acted as
  (RFC 8693). Returned by the token-exchange grant for every
  refusal about the person asked for, with one wording, so a
  service token cannot tell "no such person" from "not on your
  list". The real cause is in the delegation log.


## Example Usage

```go
import (
	"github.com/pipeshub-ai/pipeshub-sdk-go/models/components"
)

value := components.OAuthErrorResponseErrorInvalidRequest

// Open enum: custom values can be created with a direct type cast
custom := components.OAuthErrorResponseError("custom_value")
```


## Values

| Name                                          | Value                                         |
| --------------------------------------------- | --------------------------------------------- |
| `OAuthErrorResponseErrorInvalidRequest`       | invalid_request                               |
| `OAuthErrorResponseErrorInvalidClient`        | invalid_client                                |
| `OAuthErrorResponseErrorInvalidGrant`         | invalid_grant                                 |
| `OAuthErrorResponseErrorInvalidTarget`        | invalid_target                                |
| `OAuthErrorResponseErrorUnauthorizedClient`   | unauthorized_client                           |
| `OAuthErrorResponseErrorUnsupportedGrantType` | unsupported_grant_type                        |
| `OAuthErrorResponseErrorInvalidScope`         | invalid_scope                                 |
| `OAuthErrorResponseErrorAccessDenied`         | access_denied                                 |
| `OAuthErrorResponseErrorServerError`          | server_error                                  |
| `OAuthErrorResponseErrorAuthorizationPending` | authorization_pending                         |
| `OAuthErrorResponseErrorSlowDown`             | slow_down                                     |
| `OAuthErrorResponseErrorExpiredToken`         | expired_token                                 |