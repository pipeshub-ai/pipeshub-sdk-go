# OAuthConfigScopes

OAuth scopes copied from the registry when the configuration was created, by use case


## Fields

| Field                                     | Type                                      | Required                                  | Description                               |
| ----------------------------------------- | ----------------------------------------- | ----------------------------------------- | ----------------------------------------- |
| `PersonalSync`                            | []`string`                                | :heavy_minus_sign:                        | Scopes requested for a personal connector |
| `TeamSync`                                | []`string`                                | :heavy_minus_sign:                        | Scopes requested for a team connector     |
| `Agent`                                   | []`string`                                | :heavy_minus_sign:                        | Scopes requested for agent tools          |