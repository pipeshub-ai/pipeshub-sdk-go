# OAuthProtectedResourceMetadata

OAuth Protected Resource Metadata (RFC 9728).
Describes the protected resource, its authorization servers, supported scopes, and bearer token methods.



## Fields

| Field                                                         | Type                                                          | Required                                                      | Description                                                   | Example                                                       |
| ------------------------------------------------------------- | ------------------------------------------------------------- | ------------------------------------------------------------- | ------------------------------------------------------------- | ------------------------------------------------------------- |
| `Resource`                                                    | *string*                                                      | :heavy_check_mark:                                            | Protected resource identifier                                 |                                                               |
| `AuthorizationServers`                                        | []*string*                                                    | :heavy_check_mark:                                            | Authorization servers that can issue tokens for this resource |                                                               |
| `ScopesSupported`                                             | []*string*                                                    | :heavy_minus_sign:                                            | OAuth scopes supported by this resource                       | [<br/>"read:records",<br/>"write:records",<br/>"admin:connectors"<br/>] |
| `BearerMethodsSupported`                                      | []*string*                                                    | :heavy_minus_sign:                                            | Methods supported for sending bearer tokens                   | [<br/>"header"<br/>]                                          |
| `ResourceDocumentation`                                       | **string*                                                     | :heavy_minus_sign:                                            | URL to human-readable documentation for the resource          |                                                               |