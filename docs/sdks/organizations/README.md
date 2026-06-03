# Organizations

## Overview

Organization management operations

### Available Operations

* [GetCurrentOrganization](#getcurrentorganization) - Get current organization

## GetCurrentOrganization

Retrieve details about the authenticated user's organization.

**Overview:**

This endpoint returns the organization document for the current user's org, including profile data and configuration.

**Response Includes:**

- Organization profile (registeredName, shortName, contactEmail, domain)
- Account type
- Onboarding status
- Permanent address
- Creation and modification timestamps

**Use Cases:**

- Organization profile pages
- Settings and configuration screens


### Example Usage

<!-- UsageSnippet language="go" operationID="getCurrentOrganization" method="get" path="/org" -->
```go
package main

import(
	"context"
	"os"
	"github.com/pipeshub-ai/pipeshub-sdk-go/models/components"
	pipeshub "github.com/pipeshub-ai/pipeshub-sdk-go"
	"log"
)

func main() {
    ctx := context.Background()

    s := pipeshub.New(
        pipeshub.WithSecurity(components.Security{
            BearerAuth: pipeshub.Pointer(os.Getenv("PIPESHUB_BEARER_AUTH")),
        }),
    )

    res, err := s.Organizations.GetCurrentOrganization(ctx)
    if err != nil {
        log.Fatal(err)
    }
    if res.Organization != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.GetCurrentOrganizationResponse](../../models/operations/getcurrentorganizationresponse.md), error**

### Errors

| Error Type              | Status Code             | Content Type            |
| ----------------------- | ----------------------- | ----------------------- |
| apierrors.ErrorResponse | 401, 404                | application/json        |
| apierrors.ErrorResponse | 500                     | application/json        |
| apierrors.APIError      | 4XX, 5XX                | \*/\*                   |