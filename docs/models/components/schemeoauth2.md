# SchemeOauth2

OAuth 2.0 authentication with fine-grained scopes.
Supports authorization_code (with PKCE) and client_credentials flows.
OAuth tokens are Bearer JWTs — use the same Authorization header as regular tokens.



## Fields

| Field              | Type               | Required           | Description        |
| ------------------ | ------------------ | ------------------ | ------------------ |
| `ClientID`         | *string*           | :heavy_check_mark: | N/A                |
| `ClientSecret`     | *string*           | :heavy_check_mark: | N/A                |
| `TokenURL`         | *string*           | :heavy_check_mark: | N/A                |