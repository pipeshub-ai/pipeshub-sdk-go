# UpdateAuthMethodResponse

Response after updating organization authentication methods


## Fields

| Field                                                        | Type                                                         | Required                                                     | Description                                                  | Example                                                      |
| ------------------------------------------------------------ | ------------------------------------------------------------ | ------------------------------------------------------------ | ------------------------------------------------------------ | ------------------------------------------------------------ |
| `Message`                                                    | `string`                                                     | :heavy_check_mark:                                           | N/A                                                          | Auth method updated                                          |
| `AuthMethod`                                                 | [][components.AuthStep](../../models/components/authstep.md) | :heavy_check_mark:                                           | Updated authentication steps (same shape as request body)    |                                                              |