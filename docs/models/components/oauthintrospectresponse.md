# OAuthIntrospectResponse

OAuth 2.0 Token Introspection Response (RFC 7662).
Contains token metadata if active, or just `active: false` if not.



## Fields

| Field                                   | Type                                    | Required                                | Description                             |
| --------------------------------------- | --------------------------------------- | --------------------------------------- | --------------------------------------- |
| `Active`                                | *bool*                                  | :heavy_check_mark:                      | Whether the token is currently active   |
| `Scope`                                 | **string*                               | :heavy_minus_sign:                      | Scopes granted to the token             |
| `ClientID`                              | **string*                               | :heavy_minus_sign:                      | Client ID the token was issued to       |
| `Username`                              | **string*                               | :heavy_minus_sign:                      | User identifier (if user-based token)   |
| `TokenType`                             | **string*                               | :heavy_minus_sign:                      | Token type                              |
| `Exp`                                   | **int64*                                | :heavy_minus_sign:                      | Token expiration timestamp (Unix epoch) |
| `Iat`                                   | **int64*                                | :heavy_minus_sign:                      | Token issuance timestamp (Unix epoch)   |
| `Nbf`                                   | **int64*                                | :heavy_minus_sign:                      | Token not-before timestamp (Unix epoch) |
| `UserID`                                | **string*                               | :heavy_minus_sign:                      | User ID                                 |
| `Aud`                                   | **string*                               | :heavy_minus_sign:                      | Audience (client ID)                    |
| `Iss`                                   | **string*                               | :heavy_minus_sign:                      | Issuer URL                              |
| `Jti`                                   | **string*                               | :heavy_minus_sign:                      | Unique token identifier                 |