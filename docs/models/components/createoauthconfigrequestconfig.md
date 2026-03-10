# CreateOAuthConfigRequestConfig

OAuth application credentials


## Fields

| Field                                                               | Type                                                                | Required                                                            | Description                                                         | Example                                                             |
| ------------------------------------------------------------------- | ------------------------------------------------------------------- | ------------------------------------------------------------------- | ------------------------------------------------------------------- | ------------------------------------------------------------------- |
| `ClientID`                                                          | *string*                                                            | :heavy_check_mark:                                                  | OAuth client ID from your OAuth application                         | 123456789-abc.apps.googleusercontent.com                            |
| `ClientSecret`                                                      | *string*                                                            | :heavy_check_mark:                                                  | OAuth client secret (stored encrypted, never returned in responses) | GOCSPX-xxxxxxxxxxxxx                                                |
| `TenantID`                                                          | **string*                                                           | :heavy_minus_sign:                                                  | Azure tenant ID (required only for Microsoft connectors)            | 12345678-1234-1234-1234-123456789abc                                |