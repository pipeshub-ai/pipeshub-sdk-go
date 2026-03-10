# SmtpConfiguration

## Overview

### Available Operations

* [CreateOrUpdate](#createorupdate) - Create or update SMTP configuration
* [Get](#get) - Get SMTP configuration

## CreateOrUpdate

Configure SMTP email server for sending system emails including user invitations, notifications, and password resets.

Common SMTP providers and their settings:
- Gmail: host=smtp.gmail.com, port=587 (requires App Password)
- SendGrid: host=smtp.sendgrid.net, port=587
- Amazon SES: host=email-smtp.{region}.amazonaws.com, port=587
- Microsoft 365: host=smtp.office365.com, port=587

Configuration is encrypted before storage.


### Example Usage: amazonSes

<!-- UsageSnippet language="go" operationID="createSMTPConfig" method="post" path="/configurationManager/smtpConfig" example="amazonSes" -->
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

    res, err := s.SMTPConfiguration.CreateOrUpdate(ctx, components.SMTPConfig{
        Host: "email-smtp.us-east-1.amazonaws.com",
        Port: 587,
        Username: pipeshub.Pointer("AKIAIOSFODNN7EXAMPLE"),
        Password: pipeshub.Pointer("your-ses-smtp-password"),
        FromEmail: "noreply@yourcompany.com",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res != nil {
        // handle response
    }
}
```
### Example Usage: gmail

<!-- UsageSnippet language="go" operationID="createSMTPConfig" method="post" path="/configurationManager/smtpConfig" example="gmail" -->
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

    res, err := s.SMTPConfiguration.CreateOrUpdate(ctx, components.SMTPConfig{
        Host: "smtp.gmail.com",
        Port: 587,
        Username: pipeshub.Pointer("notifications@yourcompany.com"),
        Password: pipeshub.Pointer("your-app-password"),
        FromEmail: "noreply@yourcompany.com",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res != nil {
        // handle response
    }
}
```
### Example Usage: office365

<!-- UsageSnippet language="go" operationID="createSMTPConfig" method="post" path="/configurationManager/smtpConfig" example="office365" -->
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

    res, err := s.SMTPConfiguration.CreateOrUpdate(ctx, components.SMTPConfig{
        Host: "smtp.office365.com",
        Port: 587,
        Username: pipeshub.Pointer("notifications@yourcompany.onmicrosoft.com"),
        Password: pipeshub.Pointer("your-password"),
        FromEmail: "notifications@yourcompany.onmicrosoft.com",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res != nil {
        // handle response
    }
}
```
### Example Usage: sendgrid

<!-- UsageSnippet language="go" operationID="createSMTPConfig" method="post" path="/configurationManager/smtpConfig" example="sendgrid" -->
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

    res, err := s.SMTPConfiguration.CreateOrUpdate(ctx, components.SMTPConfig{
        Host: "smtp.sendgrid.net",
        Port: 587,
        Username: pipeshub.Pointer("apikey"),
        Password: pipeshub.Pointer("SG.your-sendgrid-api-key"),
        FromEmail: "noreply@yourcompany.com",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                      | Type                                                           | Required                                                       | Description                                                    |
| -------------------------------------------------------------- | -------------------------------------------------------------- | -------------------------------------------------------------- | -------------------------------------------------------------- |
| `ctx`                                                          | [context.Context](https://pkg.go.dev/context#Context)          | :heavy_check_mark:                                             | The context to use for the request.                            |
| `request`                                                      | [components.SMTPConfig](../../models/components/smtpconfig.md) | :heavy_check_mark:                                             | The request object to use for the request.                     |
| `opts`                                                         | [][operations.Option](../../models/operations/option.md)       | :heavy_minus_sign:                                             | The options for this request.                                  |

### Response

**[*operations.CreateSMTPConfigResponse](../../models/operations/createsmtpconfigresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## Get

Retrieve the current SMTP server configuration. Password is included in the response for admin users.

### Example Usage

<!-- UsageSnippet language="go" operationID="getSMTPConfig" method="get" path="/configurationManager/smtpConfig" -->
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

    res, err := s.SMTPConfiguration.Get(ctx)
    if err != nil {
        log.Fatal(err)
    }
    if res.SMTPConfig != nil {
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

**[*operations.GetSMTPConfigResponse](../../models/operations/getsmtpconfigresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |