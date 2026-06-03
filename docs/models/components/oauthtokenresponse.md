# OAuthTokenResponse

OAuth 2.0 Token Response (RFC 6749 Section 5.1).
Contains the access token and optional refresh/ID tokens.



## Fields

| Field                                                       | Type                                                        | Required                                                    | Description                                                 | Example                                                     |
| ----------------------------------------------------------- | ----------------------------------------------------------- | ----------------------------------------------------------- | ----------------------------------------------------------- | ----------------------------------------------------------- |
| `AccessToken`                                               | **string*                                                   | :heavy_minus_sign:                                          | The access token for API requests                           |                                                             |
| `TokenType`                                                 | **string*                                                   | :heavy_minus_sign:                                          | Token type (always "Bearer")                                | Bearer                                                      |
| `ExpiresIn`                                                 | **int64*                                                    | :heavy_minus_sign:                                          | Access token lifetime in seconds                            | 3600                                                        |
| `RefreshToken`                                              | **string*                                                   | :heavy_minus_sign:                                          | Refresh token for obtaining new access tokens               |                                                             |
| `Scope`                                                     | **string*                                                   | :heavy_minus_sign:                                          | Granted scopes (may differ from requested)                  |                                                             |
| `IDToken`                                                   | **string*                                                   | :heavy_minus_sign:                                          | OpenID Connect ID token (JWT) if openid scope was requested |                                                             |