# SamlSignInCallbackRequest

SAML response from Identity Provider


## Fields

| Field                                 | Type                                  | Required                              | Description                           |
| ------------------------------------- | ------------------------------------- | ------------------------------------- | ------------------------------------- |
| `SAMLResponse`                        | **string*                             | :heavy_minus_sign:                    | Base64-encoded SAML response          |
| `RelayState`                          | **string*                             | :heavy_minus_sign:                    | Relay state from the original request |