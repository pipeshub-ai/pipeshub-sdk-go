# Saml

## Overview

### Available Operations

* [SignIn](#signin) - Initiate SAML sign-in flow
* [SignInCallback](#signincallback) - SAML sign-in callback

## SignIn

Initiate SAML Single Sign-On authentication by redirecting to the Identity Provider (IDP).
<br><br>
<b>Usage:</b><br>
1. Call <code>/userAccount/initAuth</code> to get a session token<br>
2. If <code>samlSso</code> is in the allowed methods, redirect the user to this endpoint<br>
3. User authenticates with their IDP<br>
4. IDP redirects back to <code>/saml/signIn/callback</code> with SAML response<br>
5. Callback completes authentication and returns tokens
<br><br>
<b>Note:</b> This is a browser redirect endpoint, not a typical API call.
The user's browser should be redirected to this URL.
<br><br>
<b>Prerequisites:</b><br>
- Organization must have SAML SSO configured via <code>/saml/updateAppConfig</code><br>
- User must belong to an organization with SAML enabled


### Example Usage

<!-- UsageSnippet language="go" operationID="signInViaSAML" method="get" path="/saml/signIn" -->
```go
package main

import(
	"context"
	pipeshub "github.com/pipeshub-ai/pipeshub-sdk-go"
	"log"
)

func main() {
    ctx := context.Background()

    s := pipeshub.New()

    res, err := s.Saml.SignIn(ctx, "Daphney.Koss@hotmail.com", nil)
    if err != nil {
        log.Fatal(err)
    }
    if res != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                        | Type                                                                                                                             | Required                                                                                                                         | Description                                                                                                                      |
| -------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                                                            | :heavy_check_mark:                                                                                                               | The context to use for the request.                                                                                              |
| `email`                                                                                                                          | *string*                                                                                                                         | :heavy_check_mark:                                                                                                               | User's email address                                                                                                             |
| `sessionToken`                                                                                                                   | **string*                                                                                                                        | :heavy_minus_sign:                                                                                                               | Session token from `/userAccount/initAuth`. Optional but recommended<br/>for maintaining authentication state across the SAML flow.<br/> |
| `opts`                                                                                                                           | [][operations.Option](../../models/operations/option.md)                                                                         | :heavy_minus_sign:                                                                                                               | The options for this request.                                                                                                    |

### Response

**[*operations.SignInViaSAMLResponse](../../models/operations/signinviasamlresponse.md), error**

### Errors

| Error Type         | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| apierrors.APIError | 4XX, 5XX           | \*/\*              |

## SignInCallback

Handle the SAML Identity Provider callback after user authentication. This endpoint receives the SAML assertion from the IdP.


### Example Usage

<!-- UsageSnippet language="go" operationID="samlSignInCallback" method="post" path="/saml/signIn/callback" -->
```go
package main

import(
	"context"
	pipeshub "github.com/pipeshub-ai/pipeshub-sdk-go"
	"log"
)

func main() {
    ctx := context.Background()

    s := pipeshub.New()

    res, err := s.Saml.SignInCallback(ctx, nil)
    if err != nil {
        log.Fatal(err)
    }
    if res != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                    | Type                                                                                         | Required                                                                                     | Description                                                                                  |
| -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- |
| `ctx`                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                        | :heavy_check_mark:                                                                           | The context to use for the request.                                                          |
| `request`                                                                                    | [operations.SamlSignInCallbackRequest](../../models/operations/samlsignincallbackrequest.md) | :heavy_check_mark:                                                                           | The request object to use for the request.                                                   |
| `opts`                                                                                       | [][operations.Option](../../models/operations/option.md)                                     | :heavy_minus_sign:                                                                           | The options for this request.                                                                |

### Response

**[*operations.SamlSignInCallbackResponse](../../models/operations/samlsignincallbackresponse.md), error**

### Errors

| Error Type                  | Status Code                 | Content Type                |
| --------------------------- | --------------------------- | --------------------------- |
| apierrors.BadRequestError   | 400                         | application/json            |
| apierrors.UnauthorizedError | 401                         | application/json            |
| apierrors.APIError          | 4XX, 5XX                    | \*/\*                       |