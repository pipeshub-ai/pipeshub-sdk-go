# AddMessageRequest

Request body for adding a message to an existing conversation


## Fields

| Field                                                     | Type                                                      | Required                                                  | Description                                               | Example                                                   |
| --------------------------------------------------------- | --------------------------------------------------------- | --------------------------------------------------------- | --------------------------------------------------------- | --------------------------------------------------------- |
| `Query`                                                   | *string*                                                  | :heavy_check_mark:                                        | The follow-up question or message content                 | Can you elaborate on the revenue trends?                  |
| `Filters`                                                 | [*components.Filters](../../models/components/filters.md) | :heavy_minus_sign:                                        | N/A                                                       |                                                           |
| `ModelKey`                                                | **string*                                                 | :heavy_minus_sign:                                        | Override the model for this specific message              |                                                           |
| `ModelName`                                               | **string*                                                 | :heavy_minus_sign:                                        | Display name of the model                                 |                                                           |
| `ChatMode`                                                | **string*                                                 | :heavy_minus_sign:                                        | Chat mode for this message                                |                                                           |