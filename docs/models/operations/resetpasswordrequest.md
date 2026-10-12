# ResetPasswordRequest

Request payload


## Fields

| Field                                                                                  | Type                                                                                   | Required                                                                               | Description                                                                            |
| -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- |
| `CurrentPassword`                                                                      | `string`                                                                               | :heavy_check_mark:                                                                     | N/A                                                                                    |
| `NewPassword`                                                                          | `string`                                                                               | :heavy_check_mark:                                                                     | N/A                                                                                    |
| `CfTurnstileResponse`                                                                  | `*string`                                                                              | :heavy_minus_sign:                                                                     | Cloudflare Turnstile CAPTCHA token (required when Turnstile is configured server-side) |