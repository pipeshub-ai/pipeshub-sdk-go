# AuthenticateFinalResponse

All authentication steps complete; JWT tokens returned


## Fields

| Field                             | Type                              | Required                          | Description                       | Example                           |
| --------------------------------- | --------------------------------- | --------------------------------- | --------------------------------- | --------------------------------- |
| `Message`                         | `string`                          | :heavy_check_mark:                | Success message                   | Fully authenticated               |
| `AccessToken`                     | `string`                          | :heavy_check_mark:                | JWT access token (1 hour expiry)  |                                   |
| `RefreshToken`                    | `string`                          | :heavy_check_mark:                | JWT refresh token (7 days expiry) |                                   |