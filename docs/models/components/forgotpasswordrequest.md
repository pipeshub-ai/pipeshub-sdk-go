# ForgotPasswordRequest

Request to send password reset email


## Fields

| Field                                         | Type                                          | Required                                      | Description                                   |
| --------------------------------------------- | --------------------------------------------- | --------------------------------------------- | --------------------------------------------- |
| `Email`                                       | *string*                                      | :heavy_check_mark:                            | Email address to send reset link to           |
| `CfTurnstileResponse`                         | **string*                                     | :heavy_minus_sign:                            | Cloudflare Turnstile CAPTCHA token (optional) |