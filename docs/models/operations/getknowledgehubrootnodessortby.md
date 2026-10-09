# GetKnowledgeHubRootNodesSortBy

Field to sort results by. Omitted → default `updatedAt`.
Unknown value → silently falls back to `name`.


## Example Usage

```go
import (
	"github.com/pipeshub-ai/pipeshub-sdk-go/models/operations"
)

value := operations.GetKnowledgeHubRootNodesSortByName
```


## Values

| Name                                      | Value                                     |
| ----------------------------------------- | ----------------------------------------- |
| `GetKnowledgeHubRootNodesSortByName`      | name                                      |
| `GetKnowledgeHubRootNodesSortByCreatedAt` | createdAt                                 |
| `GetKnowledgeHubRootNodesSortByUpdatedAt` | updatedAt                                 |
| `GetKnowledgeHubRootNodesSortBySize`      | size                                      |
| `GetKnowledgeHubRootNodesSortByType`      | type                                      |