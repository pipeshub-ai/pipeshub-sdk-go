# HandleOAuthCallbackResponseBody

Result of exchanging the code. `success` is false when the provider refused it.


## Fields

| Field                                      | Type                                       | Required                                   | Description                                |
| ------------------------------------------ | ------------------------------------------ | ------------------------------------------ | ------------------------------------------ |
| `Success`                                  | `bool`                                     | :heavy_check_mark:                         | N/A                                        |
| `RedirectURL`                              | `*string`                                  | :heavy_minus_sign:                         | Where the web app would send the user next |
| `Error`                                    | `*string`                                  | :heavy_minus_sign:                         | N/A                                        |
| `ErrorMessage`                             | `*string`                                  | :heavy_minus_sign:                         | N/A                                        |