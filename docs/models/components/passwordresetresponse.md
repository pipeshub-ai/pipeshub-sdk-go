# PasswordResetResponse

Response after successful password reset


## Fields

| Field                                                               | Type                                                                | Required                                                            | Description                                                         | Example                                                             |
| ------------------------------------------------------------------- | ------------------------------------------------------------------- | ------------------------------------------------------------------- | ------------------------------------------------------------------- | ------------------------------------------------------------------- |
| `Data`                                                              | **string*                                                           | :heavy_minus_sign:                                                  | N/A                                                                 | password reset                                                      |
| `AccessToken`                                                       | **string*                                                           | :heavy_minus_sign:                                                  | New JWT access token (since password change invalidates old tokens) |                                                                     |