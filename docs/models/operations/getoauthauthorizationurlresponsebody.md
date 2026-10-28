# GetOAuthAuthorizationURLResponseBody

Provider consent URL


## Fields

| Field                                              | Type                                               | Required                                           | Description                                        |
| -------------------------------------------------- | -------------------------------------------------- | -------------------------------------------------- | -------------------------------------------------- |
| `Success`                                          | `bool`                                             | :heavy_check_mark:                                 | N/A                                                |
| `AuthorizationURL`                                 | `string`                                           | :heavy_check_mark:                                 | Open in a browser to give consent                  |
| `State`                                            | `string`                                           | :heavy_check_mark:                                 | Pass back with the provider's code to the callback |