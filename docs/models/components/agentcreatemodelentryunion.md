# AgentCreateModelEntryUnion

Accepted model entry for `POST /agents/create`.
The gateway accepts either a non-empty string model entry or an object entry
with a required `modelKey`.

The `models` array must include at least one object entry with `isReasoning: true`.
String-only entries are schema-valid but are rejected at the gateway with HTTP 400.



## Supported Types

### 

```go
agentCreateModelEntryUnion := components.CreateAgentCreateModelEntryUnionStr(string{/* values here */})
```

### AgentCreateModelEntry

```go
agentCreateModelEntryUnion := components.CreateAgentCreateModelEntryUnionAgentCreateModelEntry(components.AgentCreateModelEntry{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch agentCreateModelEntryUnion.Type {
	case components.AgentCreateModelEntryUnionTypeStr:
		// agentCreateModelEntryUnion.Str is populated
	case components.AgentCreateModelEntryUnionTypeAgentCreateModelEntry:
		// agentCreateModelEntryUnion.AgentCreateModelEntry is populated
}
```
