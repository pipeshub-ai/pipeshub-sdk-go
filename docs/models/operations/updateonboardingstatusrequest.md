# UpdateOnboardingStatusRequest

Request payload


## Fields

| Field                                                  | Type                                                   | Required                                               | Description                                            | Example                                                |
| ------------------------------------------------------ | ------------------------------------------------------ | ------------------------------------------------------ | ------------------------------------------------------ | ------------------------------------------------------ |
| `StepID`                                               | *string*                                               | :heavy_check_mark:                                     | ID of the step to update                               | invite_team                                            |
| `Action`                                               | [operations.Action](../../models/operations/action.md) | :heavy_check_mark:                                     | Action to perform on the step                          | complete                                               |