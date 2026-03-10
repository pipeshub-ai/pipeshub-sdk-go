# OtpGenerateRequest

Request to generate and send OTP


## Fields

| Field                                         | Type                                          | Required                                      | Description                                   |
| --------------------------------------------- | --------------------------------------------- | --------------------------------------------- | --------------------------------------------- |
| `Email`                                       | *string*                                      | :heavy_check_mark:                            | Email address to send OTP to                  |
| `CfTurnstileResponse`                         | **string*                                     | :heavy_minus_sign:                            | Cloudflare Turnstile CAPTCHA token (optional) |