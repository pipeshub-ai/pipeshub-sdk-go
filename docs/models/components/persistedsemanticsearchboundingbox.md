# PersistedSemanticSearchBoundingBox

Bounding box subdocument embedded in persisted citation metadata.
`boundingBoxSchema` does not set `_id: false`, so Mongoose auto-injects an `_id`.



## Fields

| Field              | Type               | Required           | Description        |
| ------------------ | ------------------ | ------------------ | ------------------ |
| `ID`               | `string`           | :heavy_check_mark: | N/A                |
| `X`                | `float64`          | :heavy_check_mark: | N/A                |
| `Y`                | `float64`          | :heavy_check_mark: | N/A                |