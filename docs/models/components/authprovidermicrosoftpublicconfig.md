# AuthProviderMicrosoftPublicConfig

Public Microsoft OAuth settings returned to clients


## Fields

| Field                                                                                  | Type                                                                                   | Required                                                                               | Description                                                                            |
| -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- |
| `TenantID`                                                                             | `*string`                                                                              | :heavy_minus_sign:                                                                     | Microsoft tenant ID                                                                    |
| `ClientID`                                                                             | `*string`                                                                              | :heavy_minus_sign:                                                                     | Microsoft OAuth client ID                                                              |
| `Authority`                                                                            | `*string`                                                                              | :heavy_minus_sign:                                                                     | Sign-in authority URL, built on save as `https://login.microsoftonline.com/{tenantId}` |
| `EnableJit`                                                                            | `*bool`                                                                                | :heavy_minus_sign:                                                                     | Whether just-in-time user provisioning is enabled for Microsoft                        |