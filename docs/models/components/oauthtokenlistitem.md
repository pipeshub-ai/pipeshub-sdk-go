# OAuthTokenListItem

Information about an issued token (one element returned by `listTokensForApp`
in `oauth_token.service.ts`). `userId` is omitted for client-credentials access
tokens; all other fields are always populated.



## Fields

| Field                                                        | Type                                                         | Required                                                     | Description                                                  |
| ------------------------------------------------------------ | ------------------------------------------------------------ | ------------------------------------------------------------ | ------------------------------------------------------------ |
| `ID`                                                         | *string*                                                     | :heavy_check_mark:                                           | Token ID                                                     |
| `TokenType`                                                  | [components.TokenType](../../models/components/tokentype.md) | :heavy_check_mark:                                           | Type of token                                                |
| `UserID`                                                     | **string*                                                    | :heavy_minus_sign:                                           | User ID (omitted for client-credentials access tokens)       |
| `Scopes`                                                     | []*string*                                                   | :heavy_check_mark:                                           | Granted scopes                                               |
| `CreatedAt`                                                  | [time.Time](https://pkg.go.dev/time#Time)                    | :heavy_check_mark:                                           | Token creation time                                          |
| `ExpiresAt`                                                  | [time.Time](https://pkg.go.dev/time#Time)                    | :heavy_check_mark:                                           | Token expiration time                                        |
| `IsRevoked`                                                  | *bool*                                                       | :heavy_check_mark:                                           | Whether token has been revoked                               |