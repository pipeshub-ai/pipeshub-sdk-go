# GenericOAuthConfig

Generic OAuth 2.0 provider configuration


## Fields

| Field                               | Type                                | Required                            | Description                         | Example                             |
| ----------------------------------- | ----------------------------------- | ----------------------------------- | ----------------------------------- | ----------------------------------- |
| `ProviderName`                      | **string*                           | :heavy_minus_sign:                  | Display name for the OAuth provider | Custom OAuth Provider               |
| `ClientID`                          | **string*                           | :heavy_minus_sign:                  | OAuth client ID                     |                                     |
| `ClientSecret`                      | **string*                           | :heavy_minus_sign:                  | OAuth client secret                 |                                     |
| `AuthorizationURL`                  | **string*                           | :heavy_minus_sign:                  | Authorization endpoint URL          |                                     |
| `TokenEndpoint`                     | **string*                           | :heavy_minus_sign:                  | Token endpoint URL                  |                                     |
| `UserInfoEndpoint`                  | **string*                           | :heavy_minus_sign:                  | User info endpoint URL              |                                     |
| `Scope`                             | **string*                           | :heavy_minus_sign:                  | OAuth scopes to request             | openid profile email                |
| `RedirectURI`                       | **string*                           | :heavy_minus_sign:                  | OAuth redirect URI                  |                                     |