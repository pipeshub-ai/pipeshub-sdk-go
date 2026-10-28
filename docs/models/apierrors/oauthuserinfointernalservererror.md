# OauthUserInfoInternalServerError

The scope check failed unexpectedly (OAuth body, `error` is `server_error`).

A JSON body larger than 10 MB, or a form-encoded body larger than 100 KB or with more than 1000 fields, is answered `413` by the app-wide body parsers instead, also before authentication.



## Supported Types

### OAuthBearerTokenError

```go
oauthUserInfoInternalServerError := apierrors.CreateOauthUserInfoInternalServerErrorOAuthBearerTokenError(apierrors.OAuthBearerTokenError{/* values here */})
```

### ErrorResponse

```go
oauthUserInfoInternalServerError := apierrors.CreateOauthUserInfoInternalServerErrorErrorResponse(apierrors.ErrorResponse{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch oauthUserInfoInternalServerError.Type {
	case apierrors.OauthUserInfoInternalServerErrorTypeOAuthBearerTokenError:
		// oauthUserInfoInternalServerError.OAuthBearerTokenError is populated
	case apierrors.OauthUserInfoInternalServerErrorTypeErrorResponse:
		// oauthUserInfoInternalServerError.ErrorResponse is populated
	default:
		// Unknown type - use oauthUserInfoInternalServerError.GetUnknownRaw() for raw JSON
}
```
