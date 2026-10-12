# GetArchivedConversationsShared


## Fields

| Field                                          | Type                                           | Required                                       | Description                                    |
| ---------------------------------------------- | ---------------------------------------------- | ---------------------------------------------- | ---------------------------------------------- |
| `Values`                                       | []`string`                                     | :heavy_minus_sign:                             | Allowed values for the `shared` filter         |
| `Description`                                  | `*string`                                      | :heavy_minus_sign:                             | N/A                                            |
| `Current`                                      | optionalnullable.OptionalNullable[`string`]    | :heavy_minus_sign:                             | Current value supplied by the caller, or null  |
| `Applied`                                      | `*bool`                                        | :heavy_minus_sign:                             | Whether this filter was applied on the request |