# ConnectorStoredConfigFilters

Saved filter selections


## Fields

| Field                                                                                | Type                                                                                 | Required                                                                             | Description                                                                          |
| ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ |
| `Sync`                                                                               | map[string]`any`                                                                     | :heavy_minus_sign:                                                                   | What to fetch from the source. Selections are under `values`, keyed by filter name.  |
| `Indexing`                                                                           | map[string]`any`                                                                     | :heavy_minus_sign:                                                                   | Which fetched records to index. Selections are under `values`, keyed by filter name. |
| `AdditionalProperties`                                                               | map[string]`any`                                                                     | :heavy_minus_sign:                                                                   | N/A                                                                                  |