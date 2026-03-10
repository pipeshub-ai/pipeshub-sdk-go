# GetOrganizationLogoResponseBody

Logo URL retrieved successfully


## Fields

| Field                                                                        | Type                                                                         | Required                                                                     | Description                                                                  | Example                                                                      |
| ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- |
| `Success`                                                                    | **bool*                                                                      | :heavy_minus_sign:                                                           | N/A                                                                          | true                                                                         |
| `HasLogo`                                                                    | **bool*                                                                      | :heavy_minus_sign:                                                           | Whether organization has a custom logo                                       | true                                                                         |
| `LogoURL`                                                                    | **string*                                                                    | :heavy_minus_sign:                                                           | Signed URL to access the logo (null if no logo)                              | https://storage.pipeshub.com/org/507f1f77bcf86cd799439011/logo.png?token=... |