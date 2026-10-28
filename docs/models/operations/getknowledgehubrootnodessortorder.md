# GetKnowledgeHubRootNodesSortOrder

Sort direction. Omitted → default `desc`. Read case-insensitively;
any other value is answered `400` ("Invalid sort_order: ...").


## Example Usage

```go
import (
	"github.com/pipeshub-ai/pipeshub-sdk-go/v2/models/operations"
)

value := operations.GetKnowledgeHubRootNodesSortOrderAsc
```


## Values

| Name                                    | Value                                   |
| --------------------------------------- | --------------------------------------- |
| `GetKnowledgeHubRootNodesSortOrderAsc`  | asc                                     |
| `GetKnowledgeHubRootNodesSortOrderDesc` | desc                                    |