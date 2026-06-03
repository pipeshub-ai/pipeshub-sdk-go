# AgentCreateWebSearchUnion

Accepted web-search attachment for `POST /agents/create`.
The gateway accepts either a provider string or an object with at least
a `provider` field.



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
