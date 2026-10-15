# ResyncConnectorConflict

Conflict. Possible reasons:<br>
<ul>
<li>A sync is already running, or the connector is locked</li>
<li><code>DESKTOP_OFFLINE</code> - Local FS: the owner device's desktop app is not connected</li>
<li><code>DESKTOP_UNCLAIMED</code> - Local FS: no device owns the connector yet</li>
</ul>
Local FS refusals carry <code>details.code</code> and use the
LocalFsDesktopRefusal body; the others use the standard error body.



## Supported Types

### LocalFsDesktopRefusalError

```go
resyncConnectorConflict := apierrors.CreateResyncConnectorConflictLocalFsDesktopRefusalError(apierrors.LocalFsDesktopRefusalError{/* values here */})
```

### ErrorResponse

```go
resyncConnectorConflict := apierrors.CreateResyncConnectorConflictErrorResponse(apierrors.ErrorResponse{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch resyncConnectorConflict.Type {
	case apierrors.ResyncConnectorConflictTypeLocalFsDesktopRefusalError:
		// resyncConnectorConflict.LocalFsDesktopRefusalError is populated
	case apierrors.ResyncConnectorConflictTypeErrorResponse:
		// resyncConnectorConflict.ErrorResponse is populated
	default:
		// Unknown type - use resyncConnectorConflict.GetUnknownRaw() for raw JSON
}
```
