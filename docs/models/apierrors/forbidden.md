# Forbidden

Forbidden. Returned either when the OAuth token is missing the required
`kb:read` scope or when the authenticated user does not have access to the record.



## Supported Types

### ErrorResponse

```go
forbidden := apierrors.CreateForbiddenErrorResponse(apierrors.ErrorResponse{/* values here */})
```

### StreamRecordErrorResponse

```go
forbidden := apierrors.CreateForbiddenStreamRecordErrorResponse(apierrors.StreamRecordErrorResponse{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch forbidden.Type {
	case apierrors.ForbiddenTypeErrorResponse:
		// forbidden.ErrorResponse is populated
	case apierrors.ForbiddenTypeStreamRecordErrorResponse:
		// forbidden.StreamRecordErrorResponse is populated
	default:
		// Unknown type - use forbidden.GetUnknownRaw() for raw JSON
}
```
