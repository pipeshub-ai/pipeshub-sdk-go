# MessageMetadata


## Fields

| Field                                           | Type                                            | Required                                        | Description                                     |
| ----------------------------------------------- | ----------------------------------------------- | ----------------------------------------------- | ----------------------------------------------- |
| `ProcessingTimeMs`                              | **float64*                                      | :heavy_minus_sign:                              | Time taken to generate response in milliseconds |
| `ModelVersion`                                  | **string*                                       | :heavy_minus_sign:                              | Version of the AI model used                    |
| `AiTransactionID`                               | **string*                                       | :heavy_minus_sign:                              | Transaction ID for tracking in AI backend       |
| `Reason`                                        | **string*                                       | :heavy_minus_sign:                              | Additional context or reasoning                 |