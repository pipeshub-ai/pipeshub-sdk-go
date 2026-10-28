# Config

Value of any other auth field the connector type declares, keyed by the field's `name`


## Supported Types

### 

```go
config := components.CreateConfigStr(string{/* values here */})
```

### 

```go
config := components.CreateConfigBoolean(bool{/* values here */})
```

### 

```go
config := components.CreateConfigNumber(float64{/* values here */})
```

### 

```go
config := components.CreateConfigArrayOfStr([]string{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch config.Type {
	case components.ConfigTypeStr:
		// config.Str is populated
	case components.ConfigTypeBoolean:
		// config.Boolean is populated
	case components.ConfigTypeNumber:
		// config.Number is populated
	case components.ConfigTypeArrayOfStr:
		// config.ArrayOfStr is populated
}
```
