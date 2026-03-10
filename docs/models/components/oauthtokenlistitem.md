# OAuthTokenListItem

Information about an issued token


## Fields

| Field                                                         | Type                                                          | Required                                                      | Description                                                   |
| ------------------------------------------------------------- | ------------------------------------------------------------- | ------------------------------------------------------------- | ------------------------------------------------------------- |
| `ID`                                                          | **string*                                                     | :heavy_minus_sign:                                            | Token ID                                                      |
| `TokenType`                                                   | [*components.TokenType](../../models/components/tokentype.md) | :heavy_minus_sign:                                            | Type of token                                                 |
| `UserID`                                                      | **string*                                                     | :heavy_minus_sign:                                            | User ID (if user-based token)                                 |
| `Scopes`                                                      | []*string*                                                    | :heavy_minus_sign:                                            | Granted scopes                                                |
| `CreatedAt`                                                   | [*time.Time](https://pkg.go.dev/time#Time)                    | :heavy_minus_sign:                                            | Token creation time                                           |
| `ExpiresAt`                                                   | [*time.Time](https://pkg.go.dev/time#Time)                    | :heavy_minus_sign:                                            | Token expiration time                                         |
| `IsRevoked`                                                   | **bool*                                                       | :heavy_minus_sign:                                            | Whether token has been revoked                                |