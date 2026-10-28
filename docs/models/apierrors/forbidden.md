# Forbidden

Forbidden. Returned when the OAuth token is missing the required `kb:read` scope
(`ErrorResponse`), or as `{"error": "You do not have permission to access this record"}`
when the caller has no access to the record or the record is in the trash.



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
