# GetOAuthRegistryResponseBody

OAuth connector types retrieved


## Fields

| Field                                                                                    | Type                                                                                     | Required                                                                                 | Description                                                                              |
| ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- |
| `Success`                                                                                | `bool`                                                                                   | :heavy_check_mark:                                                                       | N/A                                                                                      |
| `Connectors`                                                                             | [][components.OAuthRegistryEntry](../../models/components/oauthregistryentry.md)         | :heavy_check_mark:                                                                       | N/A                                                                                      |
| `Pagination`                                                                             | [components.OAuthRegistryPagination](../../models/components/oauthregistrypagination.md) | :heavy_check_mark:                                                                       | Pagination block of `GET /oauth/registry` and `GET /oauth/{connectorType}`               |