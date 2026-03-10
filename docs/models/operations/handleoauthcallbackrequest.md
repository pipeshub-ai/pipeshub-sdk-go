# HandleOAuthCallbackRequest


## Fields

| Field                                   | Type                                    | Required                                | Description                             |
| --------------------------------------- | --------------------------------------- | --------------------------------------- | --------------------------------------- |
| `Code`                                  | **string*                               | :heavy_minus_sign:                      | Authorization code from provider        |
| `State`                                 | **string*                               | :heavy_minus_sign:                      | State parameter (contains connector ID) |
| `Error`                                 | **string*                               | :heavy_minus_sign:                      | Error code if authorization failed      |
| `BaseURL`                               | **string*                               | :heavy_minus_sign:                      | Base URL for redirect                   |