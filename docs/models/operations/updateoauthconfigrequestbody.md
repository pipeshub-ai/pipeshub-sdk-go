# UpdateOAuthConfigRequestBody

At least one of `oauthInstanceName` and `config` must be present. Other fields are ignored.


## Supported Types

### RequestBody1

```go
updateOAuthConfigRequestBody := operations.CreateUpdateOAuthConfigRequestBodyRequestBody1(operations.RequestBody1{/* values here */})
```

### RequestBody2

```go
updateOAuthConfigRequestBody := operations.CreateUpdateOAuthConfigRequestBodyRequestBody2(operations.RequestBody2{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch updateOAuthConfigRequestBody.Type {
	case operations.UpdateOAuthConfigRequestBodyTypeRequestBody1:
		// updateOAuthConfigRequestBody.RequestBody1 is populated
	case operations.UpdateOAuthConfigRequestBodyTypeRequestBody2:
		// updateOAuthConfigRequestBody.RequestBody2 is populated
}
```
