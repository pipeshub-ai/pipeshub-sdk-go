# ErrorEnum

Error code. Common values:
- `invalid_request` - Missing or invalid parameter
- `invalid_client` - Client authentication failed
- `invalid_grant` - Invalid authorization code or refresh token
- `unauthorized_client` - Client not authorized for this grant type
- `unsupported_grant_type` - Grant type not supported
- `invalid_scope` - Requested scope is invalid or exceeds allowed
- `access_denied` - User denied authorization



## Values

| Name                            | Value                           |
| ------------------------------- | ------------------------------- |
| `ErrorEnumInvalidRequest`       | invalid_request                 |
| `ErrorEnumInvalidClient`        | invalid_client                  |
| `ErrorEnumInvalidGrant`         | invalid_grant                   |
| `ErrorEnumUnauthorizedClient`   | unauthorized_client             |
| `ErrorEnumUnsupportedGrantType` | unsupported_grant_type          |
| `ErrorEnumInvalidScope`         | invalid_scope                   |
| `ErrorEnumAccessDenied`         | access_denied                   |
| `ErrorEnumServerError`          | server_error                    |