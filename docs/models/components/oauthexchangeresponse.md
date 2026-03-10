# OAuthExchangeResponse

OAuth token response


## Fields

| Field                   | Type                    | Required                | Description             | Example                 |
| ----------------------- | ----------------------- | ----------------------- | ----------------------- | ----------------------- |
| `AccessToken`           | **string*               | :heavy_minus_sign:      | OAuth access token      |                         |
| `IDToken`               | **string*               | :heavy_minus_sign:      | OAuth ID token (JWT)    |                         |
| `TokenType`             | **string*               | :heavy_minus_sign:      | N/A                     | Bearer                  |
| `ExpiresIn`             | **int64*                | :heavy_minus_sign:      | Token expiry in seconds |                         |