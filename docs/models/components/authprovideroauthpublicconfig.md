# AuthProviderOAuthPublicConfig

Public generic OAuth provider settings returned to clients


## Fields

| Field                                                               | Type                                                                | Required                                                            | Description                                                         |
| ------------------------------------------------------------------- | ------------------------------------------------------------------- | ------------------------------------------------------------------- | ------------------------------------------------------------------- |
| `ProviderName`                                                      | `string`                                                            | :heavy_check_mark:                                                  | Custom OAuth provider display name                                  |
| `ClientID`                                                          | `string`                                                            | :heavy_check_mark:                                                  | OAuth client ID                                                     |
| `TokenEndpoint`                                                     | `string`                                                            | :heavy_check_mark:                                                  | OAuth token endpoint URL                                            |
| `AuthorizationURL`                                                  | `string`                                                            | :heavy_check_mark:                                                  | OAuth authorization URL                                             |
| `ClientSecret`                                                      | `*string`                                                           | :heavy_minus_sign:                                                  | Client secret (omitted when stripped for public responses)          |
| `UserInfoEndpoint`                                                  | `*string`                                                           | :heavy_minus_sign:                                                  | UserInfo endpoint URL                                               |
| `Scope`                                                             | `*string`                                                           | :heavy_minus_sign:                                                  | Default OAuth scopes                                                |
| `EnableJit`                                                         | `*bool`                                                             | :heavy_minus_sign:                                                  | Whether just-in-time user provisioning is enabled for this provider |