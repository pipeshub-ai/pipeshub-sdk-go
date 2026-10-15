# ConnectorAuthType

How a connector instance signs in to its source. Each connector type supports a subset; see `supportedAuthTypes` in the registry.

## Example Usage

```go
import (
	"github.com/pipeshub-ai/pipeshub-sdk-go/models/components"
)

value := components.ConnectorAuthTypeOauth

// Open enum: custom values can be created with a direct type cast
custom := components.ConnectorAuthType("custom_value")
```


## Values

| Name                                 | Value                                |
| ------------------------------------ | ------------------------------------ |
| `ConnectorAuthTypeOauth`             | OAUTH                                |
| `ConnectorAuthTypeOauthAdminConsent` | OAUTH_ADMIN_CONSENT                  |
| `ConnectorAuthTypeAPIToken`          | API_TOKEN                            |
| `ConnectorAuthTypeBearerToken`       | BEARER_TOKEN                         |
| `ConnectorAuthTypeBasicAuth`         | BASIC_AUTH                           |
| `ConnectorAuthTypeAccessKey`         | ACCESS_KEY                           |
| `ConnectorAuthTypeAccountKey`        | ACCOUNT_KEY                          |
| `ConnectorAuthTypeConnectionString`  | CONNECTION_STRING                    |
| `ConnectorAuthTypeCustom`            | CUSTOM                               |
| `ConnectorAuthTypeNone`              | NONE                                 |