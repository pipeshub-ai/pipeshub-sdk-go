# UpdateConnectorAuthRequest

Request to update authentication config (clears OAuth tokens, requires re-auth)


## Fields

| Field                                                                            | Type                                                                             | Required                                                                         | Description                                                                      |
| -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- |
| `Auth`                                                                           | [components.ConnectorAuthConfig](../../models/components/connectorauthconfig.md) | :heavy_check_mark:                                                               | Authentication configuration for a connector instance                            |
| `BaseURL`                                                                        | **string*                                                                        | :heavy_minus_sign:                                                               | Base URL (if changed with auth update)                                           |