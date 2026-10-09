# AuthenticateRequest

Request to authenticate using specified method.
**Credential format varies by method:**
- `password`: `{ password: "string" }`
- `otp`: `{ otp: "123456" }` (6-digit code)
- `google`: `"google-id-token-string"`
- `microsoft`: `{ accessToken: "...", idToken: "..." }`
- `azureAd`: `{ accessToken: "...", idToken: "..." }`
- `oauth`: `{ accessToken: "...", idToken: "..." }`
- `samlSso`: not accepted by `/userAccount/authenticate`, which answers `400`. SAML sign-in runs as a browser redirect: send the browser to `/saml/signIn` instead



## Fields

| Field                                                                | Type                                                                 | Required                                                             | Description                                                          |
| -------------------------------------------------------------------- | -------------------------------------------------------------------- | -------------------------------------------------------------------- | -------------------------------------------------------------------- |
| `Method`                                                             | [components.Method](../../models/components/method.md)               | :heavy_check_mark:                                                   | Authentication method to use                                         |
| `Credentials`                                                        | [components.Credentials](../../models/components/credentials.md)     | :heavy_check_mark:                                                   | Credentials based on the authentication method                       |
| `Email`                                                              | `*string`                                                            | :heavy_minus_sign:                                                   | Optional email for verification (used with some OAuth methods)       |
| `CfTurnstileResponse`                                                | `*string`                                                            | :heavy_minus_sign:                                                   | Cloudflare Turnstile CAPTCHA token (optional, if CAPTCHA is enabled) |