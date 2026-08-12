# ExpiryDays

Token lifetime. Defaults to `30` if omitted — a token minted
without an explicit choice shouldn't default to the longest
lifetime. `"never"` is stored as a ~100-year expiry (the
underlying schema field is required and TTL-indexed, so there's
no literal null option).



## Supported Types

### ExpiryDaysEnum

```go
expiryDays := components.CreateExpiryDaysExpiryDaysEnum(components.ExpiryDaysEnum{/* values here */})
```

### ExpiryDaysNever

```go
expiryDays := components.CreateExpiryDaysExpiryDaysNever(components.ExpiryDaysNever{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch expiryDays.Type {
	case components.ExpiryDaysTypeExpiryDaysEnum:
		// expiryDays.ExpiryDaysEnum is populated
	case components.ExpiryDaysTypeExpiryDaysNever:
		// expiryDays.ExpiryDaysNever is populated
}
```
