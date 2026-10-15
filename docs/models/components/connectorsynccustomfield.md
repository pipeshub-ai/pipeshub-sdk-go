# ConnectorSyncCustomField

A connector-specific sync setting. Send its value in `config.sync` under `name`.


## Fields

| Field              | Type               | Required           | Description        |
| ------------------ | ------------------ | ------------------ | ------------------ |
| `Name`             | `*string`          | :heavy_minus_sign: | N/A                |
| `DisplayName`      | `*string`          | :heavy_minus_sign: | N/A                |
| `Description`      | `*string`          | :heavy_minus_sign: | N/A                |
| `FieldType`        | `*string`          | :heavy_minus_sign: | N/A                |
| `Required`         | `*bool`            | :heavy_minus_sign: | N/A                |
| `DefaultValue`     | `any`              | :heavy_minus_sign: | N/A                |
| `Validation`       | map[string]`any`   | :heavy_minus_sign: | N/A                |
| `IsSecret`         | `*bool`            | :heavy_minus_sign: | N/A                |
| `NonEditable`      | `*bool`            | :heavy_minus_sign: | N/A                |
| `Options`          | []`string`         | :heavy_minus_sign: | N/A                |