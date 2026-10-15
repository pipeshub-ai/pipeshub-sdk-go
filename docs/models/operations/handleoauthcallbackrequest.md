# HandleOAuthCallbackRequest


## Fields

| Field                                   | Type                                    | Required                                | Description                             |
| --------------------------------------- | --------------------------------------- | --------------------------------------- | --------------------------------------- |
| `Code`                                  | `string`                                | :heavy_check_mark:                      | Authorization code from provider        |
| `State`                                 | `string`                                | :heavy_check_mark:                      | State parameter (contains connector ID) |
| `Error`                                 | `*string`                               | :heavy_minus_sign:                      | Error code if authorization failed      |
| `BaseURL`                               | `*string`                               | :heavy_minus_sign:                      | Base URL for redirect                   |