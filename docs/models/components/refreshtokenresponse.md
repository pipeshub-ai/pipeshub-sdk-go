# RefreshTokenResponse

Response with new access token


## Fields

| Field                                                                               | Type                                                                                | Required                                                                            | Description                                                                         |
| ----------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------- |
| `User`                                                                              | [components.RefreshTokenUser](../../models/components/refreshtokenuser.md)          | :heavy_check_mark:                                                                  | User record returned with a refreshed access token                                  |
| `AccessToken`                                                                       | *string*                                                                            | :heavy_check_mark:                                                                  | New JWT access token (24 hour default expiry, configurable via ACCESS_TOKEN_EXPIRY) |