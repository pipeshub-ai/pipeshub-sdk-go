# MicrosoftCredentials

Credentials for `microsoft` sign-in. Only `idToken` is read: it is
verified against Microsoft's signing keys, with the saved client ID as audience.
Without it the sign-in answers `400`.



## Fields

| Field                    | Type                     | Required                 | Description              |
| ------------------------ | ------------------------ | ------------------------ | ------------------------ |
| `IDToken`                | `string`                 | :heavy_check_mark:       | Microsoft ID token (JWT) |
| `AccessToken`            | `*string`                | :heavy_minus_sign:       | Accepted but not read    |