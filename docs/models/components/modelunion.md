# ModelUnion


## Supported Types

### 

```go
modelUnion := components.CreateModelUnionStr(string{/* values here */})
```

### Model

```go
modelUnion := components.CreateModelUnionModel(components.Model{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch modelUnion.Type {
	case components.ModelUnionTypeStr:
		// modelUnion.Str is populated
	case components.ModelUnionTypeModel:
		// modelUnion.Model is populated
	default:
		// Unknown type - use modelUnion.GetUnknownRaw() for raw JSON
}
```
