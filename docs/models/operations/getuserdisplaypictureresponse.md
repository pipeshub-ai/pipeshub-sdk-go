# GetUserDisplayPictureResponse


## Fields

| Field                                                              | Type                                                               | Required                                                           | Description                                                        |
| ------------------------------------------------------------------ | ------------------------------------------------------------------ | ------------------------------------------------------------------ | ------------------------------------------------------------------ |
| `HTTPMeta`                                                         | [components.HTTPMetadata](../../models/components/httpmetadata.md) | :heavy_check_mark:                                                 | N/A                                                                |
| `TwoHundredImageJpegResponseStream`                                | *io.ReadCloser*                                                    | :heavy_minus_sign:                                                 | Display picture retrieved successfully                             |
| `TwoHundredImagePngResponseStream`                                 | *io.ReadCloser*                                                    | :heavy_minus_sign:                                                 | Display picture retrieved successfully                             |