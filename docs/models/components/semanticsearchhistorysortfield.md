# SemanticSearchHistorySortField

Used for `available.sorting.{sortBy,sortOrder}` and
`available.sortingMessages.{sortBy,sortOrder}`. The `applied` flag
is present on `sorting.*` and absent on `sortingMessages.*`, so it
is optional here.



## Fields

| Field              | Type               | Required           | Description        |
| ------------------ | ------------------ | ------------------ | ------------------ |
| `Values`           | []`string`         | :heavy_check_mark: | N/A                |
| `Default`          | `string`           | :heavy_check_mark: | N/A                |
| `Description`      | `string`           | :heavy_check_mark: | N/A                |
| `Current`          | `string`           | :heavy_check_mark: | N/A                |
| `Applied`          | `*bool`            | :heavy_minus_sign: | N/A                |