# OAuthConfigCredentialValue

A stored credential value, returned exactly as it was sent on create or update. Usually a
string, boolean, number or list of strings, but the value is not checked, so `null`, a
nested object or an array of other items comes back as well.



## Supported Types

### 

```go
oAuthConfigCredentialValue := components.CreateOAuthConfigCredentialValueStr(string{/* values here */})
```

### 

```go
oAuthConfigCredentialValue := components.CreateOAuthConfigCredentialValueBoolean(bool{/* values here */})
```

### 

```go
oAuthConfigCredentialValue := components.CreateOAuthConfigCredentialValueNumber(float64{/* values here */})
```

### 

```go
oAuthConfigCredentialValue := components.CreateOAuthConfigCredentialValueArrayOfAny([]any{/* values here */})
```

### 

```go
oAuthConfigCredentialValue := components.CreateOAuthConfigCredentialValueMapOfAny(map[string]any{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch oAuthConfigCredentialValue.Type {
	case components.OAuthConfigCredentialValueTypeStr:
		// oAuthConfigCredentialValue.Str is populated
	case components.OAuthConfigCredentialValueTypeBoolean:
		// oAuthConfigCredentialValue.Boolean is populated
	case components.OAuthConfigCredentialValueTypeNumber:
		// oAuthConfigCredentialValue.Number is populated
	case components.OAuthConfigCredentialValueTypeArrayOfAny:
		// oAuthConfigCredentialValue.ArrayOfAny is populated
	case components.OAuthConfigCredentialValueTypeMapOfAny:
		// oAuthConfigCredentialValue.MapOfAny is populated
	default:
		// Unknown type - use oAuthConfigCredentialValue.GetUnknownRaw() for raw JSON
}
```
