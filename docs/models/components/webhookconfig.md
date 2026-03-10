# WebhookConfig

Configuration for webhook-based sync


## Fields

| Field                                               | Type                                                | Required                                            | Description                                         | Example                                             |
| --------------------------------------------------- | --------------------------------------------------- | --------------------------------------------------- | --------------------------------------------------- | --------------------------------------------------- |
| `WebhookURL`                                        | **string*                                           | :heavy_minus_sign:                                  | URL to receive webhook events (auto-generated)      |                                                     |
| `Events`                                            | []*string*                                          | :heavy_minus_sign:                                  | Subscribed event types                              | [<br/>"file.created",<br/>"file.modified",<br/>"file.deleted"<br/>] |