# ToggleConnectorConflict

A Local FS desktop refusal, or the connector is locked


## Supported Types

### LocalFsDesktopRefusalError

```go
toggleConnectorConflict := apierrors.CreateToggleConnectorConflictLocalFsDesktopRefusalError(apierrors.LocalFsDesktopRefusalError{/* values here */})
```

### ErrorResponse

```go
toggleConnectorConflict := apierrors.CreateToggleConnectorConflictErrorResponse(apierrors.ErrorResponse{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch toggleConnectorConflict.Type {
	case apierrors.ToggleConnectorConflictTypeLocalFsDesktopRefusalError:
		// toggleConnectorConflict.LocalFsDesktopRefusalError is populated
	case apierrors.ToggleConnectorConflictTypeErrorResponse:
		// toggleConnectorConflict.ErrorResponse is populated
	default:
		// Unknown type - use toggleConnectorConflict.GetUnknownRaw() for raw JSON
}
```
