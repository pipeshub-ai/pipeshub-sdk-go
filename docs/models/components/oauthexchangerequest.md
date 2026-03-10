# OAuthExchangeRequest

Request to exchange OAuth authorization code for tokens


## Fields

| Field                                      | Type                                       | Required                                   | Description                                |
| ------------------------------------------ | ------------------------------------------ | ------------------------------------------ | ------------------------------------------ |
| `Code`                                     | *string*                                   | :heavy_check_mark:                         | OAuth authorization code                   |
| `Email`                                    | *string*                                   | :heavy_check_mark:                         | User email                                 |
| `Provider`                                 | *string*                                   | :heavy_check_mark:                         | OAuth provider name                        |
| `RedirectURI`                              | *string*                                   | :heavy_check_mark:                         | Redirect URI used in authorization request |