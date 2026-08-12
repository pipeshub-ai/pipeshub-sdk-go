# AdminPatListResponse

Response body for `GET /personal-access-tokens/admin` (`adminListTokens`).
Paginated — unlike the self-service `ListPatResponse` — since an org
can have far more active tokens than a fixed-window cap's worth.



## Fields

| Field                                                                                                  | Type                                                                                                   | Required                                                                                               | Description                                                                                            |
| ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ |
| `Data`                                                                                                 | [][components.AdminPatListItem](../../models/components/adminpatlistitem.md)                           | :heavy_check_mark:                                                                                     | N/A                                                                                                    |
| `Pagination`                                                                                           | [components.AdminPatListResponsePagination](../../models/components/adminpatlistresponsepagination.md) | :heavy_check_mark:                                                                                     | N/A                                                                                                    |