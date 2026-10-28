# RedactedSource

Stands in for a source the viewer of a shared chat or search cannot read:
a citation, reference, filter node or search hit with no id, name, link or
text. Only the owner of a conversation or search sees every source.



## Fields

| Field                                                  | Type                                                   | Required                                               | Description                                            |
| ------------------------------------------------------ | ------------------------------------------------------ | ------------------------------------------------------ | ------------------------------------------------------ |
| `Redacted`                                             | `bool`                                                 | :heavy_check_mark:                                     | N/A                                                    |
| `Reason`                                               | [components.Reason](../../models/components/reason.md) | :heavy_check_mark:                                     | N/A                                                    |
| `Name`                                                 | `*string`                                              | :heavy_minus_sign:                                     | A neutral label (`Restricted source`).                 |