# AgentCreateWebSearchUnion

Web-search attachment for an agent. Accepts a provider string, an object
with at least a `provider` field, or `null`.



## Supported Types

### 

```go
agentCreateWebSearchUnion := components.CreateAgentCreateWebSearchUnionStr(string{/* values here */})
```

### AgentCreateWebSearch

```go
agentCreateWebSearchUnion := components.CreateAgentCreateWebSearchUnionAgentCreateWebSearch(components.AgentCreateWebSearch{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch agentCreateWebSearchUnion.Type {
	case components.AgentCreateWebSearchUnionTypeStr:
		// agentCreateWebSearchUnion.Str is populated
	case components.AgentCreateWebSearchUnionTypeAgentCreateWebSearch:
		// agentCreateWebSearchUnion.AgentCreateWebSearch is populated
}
```
