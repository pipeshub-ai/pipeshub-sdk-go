# To


## Supported Types

### 

```go
to := components.CreateToInteger(int64{/* values here */})
```

### 

```go
to := components.CreateToStr(string{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch to.Type {
	case components.ToTypeInteger:
		// to.Integer is populated
	case components.ToTypeStr:
		// to.Str is populated
}
```
