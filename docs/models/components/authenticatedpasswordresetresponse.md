# AuthenticatedPasswordResetResponse

Response after authenticated user changes password (new access token issued)


## Fields

| Field                                      | Type                                       | Required                                   | Description                                | Example                                    |
| ------------------------------------------ | ------------------------------------------ | ------------------------------------------ | ------------------------------------------ | ------------------------------------------ |
| `Data`                                     | `string`                                   | :heavy_check_mark:                         | N/A                                        | password reset                             |
| `AccessToken`                              | `string`                                   | :heavy_check_mark:                         | New JWT access token after password change |                                            |