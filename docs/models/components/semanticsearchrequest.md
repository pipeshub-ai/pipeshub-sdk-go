# SemanticSearchRequest

Request body for performing semantic search across the enterprise knowledge base.<br><br>
<b>How Semantic Search Works:</b><br>
<ol>
<li>Query is converted to vector embeddings</li>
<li>Similar content is found using vector similarity</li>
<li>Results are ranked by relevance score</li>
<li>Matching chunks with metadata are returned</li>
</ol>
<b>Filtering:</b><br>
Use filters to narrow search scope to specific apps or knowledge bases.



## Fields

| Field                                                                                       | Type                                                                                        | Required                                                                                    | Description                                                                                 | Example                                                                                     |
| ------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------- |
| `Query`                                                                                     | *string*                                                                                    | :heavy_check_mark:                                                                          | Natural language search query. The system understands<br/>semantic meaning, not just keywords.<br/> | employee onboarding procedures                                                              |
| `Filters`                                                                                   | [*components.Filters](../../models/components/filters.md)                                   | :heavy_minus_sign:                                                                          | N/A                                                                                         |                                                                                             |
| `Limit`                                                                                     | **int64*                                                                                    | :heavy_minus_sign:                                                                          | Maximum number of results to return                                                         |                                                                                             |
| `ModelKey`                                                                                  | **string*                                                                                   | :heavy_minus_sign:                                                                          | AI model to use for embeddings                                                              |                                                                                             |
| `ModelName`                                                                                 | **string*                                                                                   | :heavy_minus_sign:                                                                          | Display name of the model                                                                   |                                                                                             |
| `ChatMode`                                                                                  | **string*                                                                                   | :heavy_minus_sign:                                                                          | Processing mode configuration                                                               |                                                                                             |