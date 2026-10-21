# From


## Supported Types

### 

```go
from := components.CreateFromInteger(int64{/* values here */})
```

### 

```go
from := components.CreateFromStr(string{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch from.Type {
	case components.FromTypeInteger:
		// from.Integer is populated
	case components.FromTypeStr:
		// from.Str is populated
}
```
