# OAuthRevokeRequest

OAuth 2.0 Token Revocation Request (RFC 7009).
Revokes an access or refresh token.



## Fields

| Field                                                                                                     | Type                                                                                                      | Required                                                                                                  | Description                                                                                               |
| --------------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------------- |
| `Token`                                                                                                   | *string*                                                                                                  | :heavy_check_mark:                                                                                        | The token to revoke                                                                                       |
| `TokenTypeHint`                                                                                           | [*components.OAuthRevokeRequestTokenTypeHint](../../models/components/oauthrevokerequesttokentypehint.md) | :heavy_minus_sign:                                                                                        | Hint about token type (optional, improves performance)                                                    |
| `ClientID`                                                                                                | *string*                                                                                                  | :heavy_check_mark:                                                                                        | Client ID                                                                                                 |
| `ClientSecret`                                                                                            | **string*                                                                                                 | :heavy_minus_sign:                                                                                        | Client secret                                                                                             |