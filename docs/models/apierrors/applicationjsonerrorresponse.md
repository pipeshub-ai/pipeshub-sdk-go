# ApplicationJSONErrorResponse

Standard JSON error envelope from `ErrorMiddleware` for `BaseError` subclasses (`error.middleware.ts`).
Returned for most API 4xx errors (unauthorized, forbidden, not found, validation failures, etc.).



## Fields

| Field                                                                                                        | Type                                                                                                         | Required                                                                                                     | Description                                                                                                  |
| ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ |
| `Error`                                                                                                      | [components.ApplicationJSONErrorResponseError](../../models/components/applicationjsonerrorresponseerror.md) | :heavy_check_mark:                                                                                           | N/A                                                                                                          |
| `HTTPMeta`                                                                                                   | [components.HTTPMetadata](../../models/components/httpmetadata.md)                                           | :heavy_check_mark:                                                                                           | N/A                                                                                                          |