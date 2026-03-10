# RefreshTokenResponse

Response with new access token


## Fields

| Field                                                         | Type                                                          | Required                                                      | Description                                                   |
| ------------------------------------------------------------- | ------------------------------------------------------------- | ------------------------------------------------------------- | ------------------------------------------------------------- |
| `User`                                                        | [*components.UserBasic](../../models/components/userbasic.md) | :heavy_minus_sign:                                            | Basic user information                                        |
| `AccessToken`                                                 | **string*                                                     | :heavy_minus_sign:                                            | New JWT access token (1 hour expiry)                          |