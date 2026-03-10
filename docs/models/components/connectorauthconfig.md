# ConnectorAuthConfig

Authentication configuration for a connector instance


## Fields

| Field                                                          | Type                                                           | Required                                                       | Description                                                    | Example                                                        |
| -------------------------------------------------------------- | -------------------------------------------------------------- | -------------------------------------------------------------- | -------------------------------------------------------------- | -------------------------------------------------------------- |
| `Values`                                                       | map[string]*any*                                               | :heavy_minus_sign:                                             | Authentication values (keys depend on connector's auth schema) | {<br/>"apiKey": "sk-xxxxx",<br/>"baseUrl": "https://api.example.com"<br/>} |
| `OauthConfigID`                                                | **string*                                                      | :heavy_minus_sign:                                             | ID of admin-created OAuth configuration to use                 | oauth_config_123                                               |
| `CustomValues`                                                 | map[string]*any*                                               | :heavy_minus_sign:                                             | Custom authentication values specific to the connector         |                                                                |