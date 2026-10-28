# GetKnowledgeHubChildNodesSortOrder

Sort direction. Omitted → default `desc`. Read case-insensitively;
any other value is answered `400` ("Invalid sort_order: ...").


## Example Usage

```go
import (
	"github.com/pipeshub-ai/pipeshub-sdk-go/models/operations"
)

value := operations.GetKnowledgeHubChildNodesSortOrderAsc
```


## Values

| Name                                     | Value                                    |
| ---------------------------------------- | ---------------------------------------- |
| `GetKnowledgeHubChildNodesSortOrderAsc`  | asc                                      |
| `GetKnowledgeHubChildNodesSortOrderDesc` | desc                                     |