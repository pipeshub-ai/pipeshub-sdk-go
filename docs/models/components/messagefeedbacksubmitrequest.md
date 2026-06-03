# MessageFeedbackSubmitRequest

Gateway request body for submitting message feedback (Zod
`feedbackBodySchema`). All fields are optional; an empty object is
accepted. Matches the first-party chat UI payload shape.



## Fields

| Field                                                                                                                | Type                                                                                                                 | Required                                                                                                             | Description                                                                                                          |
| -------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------- |
| `IsHelpful`                                                                                                          | **bool*                                                                                                              | :heavy_minus_sign:                                                                                                   | Overall helpfulness signal (thumbs up/down).                                                                         |
| `Categories`                                                                                                         | [][components.MessageFeedbackSubmitRequestCategory](../../models/components/messagefeedbacksubmitrequestcategory.md) | :heavy_minus_sign:                                                                                                   | Issue or positive categories that apply to the response.                                                             |
| `Comments`                                                                                                           | [*components.MessageFeedbackSubmitRequestComments](../../models/components/messagefeedbacksubmitrequestcomments.md)  | :heavy_minus_sign:                                                                                                   | Free-text comments grouped by sentiment.                                                                             |