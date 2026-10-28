# ListOAuthConfigsByTypeResponseBody

Saved OAuth apps


## Fields

| Field                                                                                     | Type                                                                                      | Required                                                                                  | Description                                                                               |
| ----------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------- |
| `Success`                                                                                 | `bool`                                                                                    | :heavy_check_mark:                                                                        | N/A                                                                                       |
| `OauthConfigs`                                                                            | [][components.OAuthConfigRecord](../../models/components/oauthconfigrecord.md)            | :heavy_check_mark:                                                                        | N/A                                                                                       |
| `Pagination`                                                                              | [*components.OAuthRegistryPagination](../../models/components/oauthregistrypagination.md) | :heavy_minus_sign:                                                                        | Pagination block of `GET /oauth/registry` and `GET /oauth/{connectorType}`                |