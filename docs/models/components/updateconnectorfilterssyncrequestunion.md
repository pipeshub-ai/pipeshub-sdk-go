# UpdateConnectorFiltersSyncRequestUnion

At least one of `sync` and `filters`. `filters.sync` and `filters.indexing` replace the saved ones.


## Supported Types

### UpdateConnectorFiltersSyncRequest1

```go
updateConnectorFiltersSyncRequestUnion := components.CreateUpdateConnectorFiltersSyncRequestUnionUpdateConnectorFiltersSyncRequest1(components.UpdateConnectorFiltersSyncRequest1{/* values here */})
```

### UpdateConnectorFiltersSyncRequest2

```go
updateConnectorFiltersSyncRequestUnion := components.CreateUpdateConnectorFiltersSyncRequestUnionUpdateConnectorFiltersSyncRequest2(components.UpdateConnectorFiltersSyncRequest2{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch updateConnectorFiltersSyncRequestUnion.Type {
	case components.UpdateConnectorFiltersSyncRequestUnionTypeUpdateConnectorFiltersSyncRequest1:
		// updateConnectorFiltersSyncRequestUnion.UpdateConnectorFiltersSyncRequest1 is populated
	case components.UpdateConnectorFiltersSyncRequestUnionTypeUpdateConnectorFiltersSyncRequest2:
		// updateConnectorFiltersSyncRequestUnion.UpdateConnectorFiltersSyncRequest2 is populated
}
```
