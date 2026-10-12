# AuthProviderAzureAdPublicConfig

Public Azure AD OAuth settings returned to clients


## Fields

| Field                                                          | Type                                                           | Required                                                       | Description                                                    |
| -------------------------------------------------------------- | -------------------------------------------------------------- | -------------------------------------------------------------- | -------------------------------------------------------------- |
| `TenantID`                                                     | `*string`                                                      | :heavy_minus_sign:                                             | Azure AD tenant ID                                             |
| `ClientID`                                                     | `*string`                                                      | :heavy_minus_sign:                                             | Azure AD client ID                                             |
| `EnableJit`                                                    | `*bool`                                                        | :heavy_minus_sign:                                             | Whether just-in-time user provisioning is enabled for Azure AD |