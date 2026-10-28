# StreamRecordBufferInternalServerError

Internal streaming failure, downstream conversion failure, or connector/backend
error proxied by the gateway.


A JSON body larger than 10 MB, or a form-encoded body larger than 100 KB or with more than 1000 fields, is answered `413` by the app-wide body parsers instead, also before authentication.



## Supported Types

### StreamRecordErrorResponse

```go
streamRecordBufferInternalServerError := apierrors.CreateStreamRecordBufferInternalServerErrorStreamRecordErrorResponse(apierrors.StreamRecordErrorResponse{/* values here */})
```

### ErrorResponse

```go
streamRecordBufferInternalServerError := apierrors.CreateStreamRecordBufferInternalServerErrorErrorResponse(apierrors.ErrorResponse{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch streamRecordBufferInternalServerError.Type {
	case apierrors.StreamRecordBufferInternalServerErrorTypeStreamRecordErrorResponse:
		// streamRecordBufferInternalServerError.StreamRecordErrorResponse is populated
	case apierrors.StreamRecordBufferInternalServerErrorTypeErrorResponse:
		// streamRecordBufferInternalServerError.ErrorResponse is populated
	default:
		// Unknown type - use streamRecordBufferInternalServerError.GetUnknownRaw() for raw JSON
}
```
