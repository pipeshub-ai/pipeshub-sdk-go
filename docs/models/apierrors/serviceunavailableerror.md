# ServiceUnavailableError

The database could not give this listing the memory it needs right now, because other requests were using
it (nothing is wrong with the request). Retry after `Retry-After` seconds; the web app retries by itself.
The body is the gateway's own busy message, not the service's wording.



## Fields

| Field                                                                                                                                    | Type                                                                                                                                     | Required                                                                                                                                 | Description                                                                                                                              |
| ---------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------- |
| `Error`                                                                                                                                  | [operations.GetKnowledgeHubRootNodesServiceUnavailableError](../../models/operations/getknowledgehubrootnodesserviceunavailableerror.md) | :heavy_check_mark:                                                                                                                       | N/A                                                                                                                                      |
| `HTTPMeta`                                                                                                                               | [components.HTTPMetadata](../../models/components/httpmetadata.md)                                                                       | :heavy_check_mark:                                                                                                                       | N/A                                                                                                                                      |