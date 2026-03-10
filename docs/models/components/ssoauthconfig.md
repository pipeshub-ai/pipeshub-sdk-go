# SSOAuthConfig

SAML SSO authentication configuration


## Fields

| Field                                                   | Type                                                    | Required                                                | Description                                             | Example                                                 |
| ------------------------------------------------------- | ------------------------------------------------------- | ------------------------------------------------------- | ------------------------------------------------------- | ------------------------------------------------------- |
| `EntryPoint`                                            | **string*                                               | :heavy_minus_sign:                                      | Identity provider SSO URL                               | https://idp.example.com/sso/saml                        |
| `Certificate`                                           | **string*                                               | :heavy_minus_sign:                                      | X.509 certificate for signature validation (PEM format) |                                                         |
| `EmailKey`                                              | **string*                                               | :heavy_minus_sign:                                      | SAML attribute name for user email                      | email                                                   |