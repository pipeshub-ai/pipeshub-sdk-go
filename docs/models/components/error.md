# Error

Error code. Common values:
- `invalid_request` - Missing or invalid parameter
- `invalid_client` - Client authentication failed
- `invalid_grant` - Invalid authorization code or refresh token
- `unauthorized_client` - Client not authorized for this grant type
- `unsupported_grant_type` - Grant type not supported
- `invalid_scope` - Requested scope is invalid or exceeds allowed
- `access_denied` - User denied authorization



## Values

| Name                        | Value                       |
| --------------------------- | --------------------------- |
| `ErrorInvalidRequest`       | invalid_request             |
| `ErrorInvalidClient`        | invalid_client              |
| `ErrorInvalidGrant`         | invalid_grant               |
| `ErrorUnauthorizedClient`   | unauthorized_client         |
| `ErrorUnsupportedGrantType` | unsupported_grant_type      |
| `ErrorInvalidScope`         | invalid_scope               |
| `ErrorAccessDenied`         | access_denied               |
| `ErrorServerError`          | server_error                |