# AuthProviders

Settings of the external sign-in providers among the allowed methods, keyed by provider.
A provider is listed whenever its settings can be read, which is also the case when
nothing has been saved for it yet: its entry is then an empty object. Only `google`
and `microsoft` have a sign-in flow.



## Fields

| Field                                                                                                         | Type                                                                                                          | Required                                                                                                      | Description                                                                                                   |
| ------------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------- |
| `Google`                                                                                                      | [*components.AuthProviderGooglePublicConfig](../../models/components/authprovidergooglepublicconfig.md)       | :heavy_minus_sign:                                                                                            | Public Google OAuth settings returned to clients                                                              |
| `Microsoft`                                                                                                   | [*components.AuthProviderMicrosoftPublicConfig](../../models/components/authprovidermicrosoftpublicconfig.md) | :heavy_minus_sign:                                                                                            | Public Microsoft OAuth settings returned to clients                                                           |