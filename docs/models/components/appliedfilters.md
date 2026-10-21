# AppliedFilters

Rich filter state selected by the user, used for display and persistence only.
This mirrors the active selection shown in the UI and is distinct from the
machine-readable `filters` field used for retrieval scoping.



## Fields

| Field                                                                          | Type                                                                           | Required                                                                       | Description                                                                    |
| ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ |
| `Apps`                                                                         | [][components.AppliedFilterNode](../../models/components/appliedfilternode.md) | :heavy_minus_sign:                                                             | Applied app/connector filter nodes                                             |
| `Kb`                                                                           | [][components.AppliedFilterNode](../../models/components/appliedfilternode.md) | :heavy_minus_sign:                                                             | Applied knowledge-base filter nodes                                            |
| `RecordGroups`                                                                 | [][components.AppliedFilterNode](../../models/components/appliedfilternode.md) | :heavy_minus_sign:                                                             | Applied record-group filter nodes                                              |
| `Records`                                                                      | [][components.AppliedFilterNode](../../models/components/appliedfilternode.md) | :heavy_minus_sign:                                                             | Applied folder and record filter nodes (each with everything under it)         |
| `RecordsExact`                                                                 | [][components.AppliedFilterNode](../../models/components/appliedfilternode.md) | :heavy_minus_sign:                                                             | Applied record filter nodes selected alone                                     |