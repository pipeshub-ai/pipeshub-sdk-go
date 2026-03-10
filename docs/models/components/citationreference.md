# CitationReference

Reference to a source document cited in a response


## Fields

| Field                                            | Type                                             | Required                                         | Description                                      |
| ------------------------------------------------ | ------------------------------------------------ | ------------------------------------------------ | ------------------------------------------------ |
| `CitationID`                                     | **string*                                        | :heavy_minus_sign:                               | ID of the citation record                        |
| `RelevanceScore`                                 | **float64*                                       | :heavy_minus_sign:                               | How relevant this citation is to the query (0-1) |
| `Excerpt`                                        | **string*                                        | :heavy_minus_sign:                               | Relevant excerpt from the source document        |
| `Context`                                        | **string*                                        | :heavy_minus_sign:                               | Additional context around the citation           |