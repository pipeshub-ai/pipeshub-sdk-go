# AgentListItemWebSearch

Web-search provider attachment for this agent, or `null` when none is attached.

For `GET /agents`, the response formatter always emits `provider`.
It may also emit `providerKey` and `providerLabel` when those values
were present on the stored attachment. It does not emit `iconPath`
on this response path.



## Fields

| Field              | Type               | Required           | Description        |
| ------------------ | ------------------ | ------------------ | ------------------ |
| `Provider`         | **string*          | :heavy_minus_sign: | N/A                |
| `ProviderKey`      | **string*          | :heavy_minus_sign: | N/A                |
| `ProviderLabel`    | **string*          | :heavy_minus_sign: | N/A                |