# PipesHub Go SDK

[![Go Reference](https://pkg.go.dev/badge/github.com/pipeshub-ai/pipeshub-sdk-go/v2.svg)](https://pkg.go.dev/github.com/pipeshub-ai/pipeshub-sdk-go/v2)
[![License](https://img.shields.io/github/license/pipeshub-ai/pipeshub-sdk-go.svg)](https://github.com/pipeshub-ai/pipeshub-sdk-go/blob/main/LICENSE)

This is the official Go SDK for [PipesHub](https://pipeshub.com). Search your company's knowledge and get answers with citations, run agents, and manage knowledge bases and connectors, all from your Go code.

PipesHub connects workplace apps such as Google Drive, Slack, Confluence and Jira and indexes their content. The SDK has typed requests and responses for every operation, and every method takes a `context.Context`.

## Documentation

- [API reference](https://docs.pipeshub.com/developer/api-reference): the PipesHub HTTP API.
- [Go package reference](https://pkg.go.dev/github.com/pipeshub-ai/pipeshub-sdk-go/v2)
- [Available methods](#available-resources-and-operations): every SDK method, linked to its parameters, response and errors.
- [Examples](examples/): runnable programs for search, chat, agents, knowledge bases and connectors.

The client signs in with a [personal access token](https://docs.pipeshub.com/developer/personal-access-tokens) and connects to PipesHub Cloud (`https://app.pipeshub.com`) by default. For a self-hosted instance, pass `pipeshub.WithServerURL("https://<your-instance>/api/v1")` (see [Server Selection](#server-selection)).

<!-- No Summary [summary] -->

<!-- Start Table of Contents [toc] -->
## Table of Contents
<!-- $toc-max-depth=2 -->
* [PipesHub Go SDK](#pipeshub-go-sdk)
  * [Documentation](#documentation)
  * [SDK Installation](#sdk-installation)
  * [SDK Example Usage](#sdk-example-usage)
  * [Authentication](#authentication)
  * [Server-sent event streaming](#server-sent-event-streaming)
  * [Retries](#retries)
  * [Error Handling](#error-handling)
  * [Server Selection](#server-selection)
  * [Custom HTTP Client](#custom-http-client)
  * [Special Types](#special-types)
  * [Available Resources and Operations](#available-resources-and-operations)
  * [Support](#support)
  * [Contributing](#contributing)
  * [License](#license)

<!-- End Table of Contents [toc] -->

<!-- Start SDK Installation [installation] -->
## SDK Installation

To add the SDK as a dependency to your project:
```bash
go get github.com/pipeshub-ai/pipeshub-sdk-go/v2
```
<!-- End SDK Installation [installation] -->

<!-- Start SDK Example Usage [usage] -->
## SDK Example Usage

### Retrieve passages with citations

```go
package main

import (
	"context"
	pipeshub "github.com/pipeshub-ai/pipeshub-sdk-go/v2"
	"github.com/pipeshub-ai/pipeshub-sdk-go/v2/models/components"
	"log"
	"os"
)

func main() {
	ctx := context.Background()

	s := pipeshub.New(
		pipeshub.WithSecurity(components.Security{
			BearerAuth: pipeshub.Pointer(os.Getenv("PIPESHUB_BEARER_AUTH")),
		}),
	)

	res, err := s.SemanticSearch.Retrieve(ctx, components.RetrieveRequest{
		Query: "how do we roll back a canary?",
	})
	if err != nil {
		log.Fatal(err)
	}
	if res.RetrieveResponse != nil {
		// handle response
	}
}

```
<!-- End SDK Example Usage [usage] -->

<!-- Start Authentication [security] -->
## Authentication

### Per-Client Security Schemes

This SDK supports the following security schemes globally:

| Name         | Type   | Scheme       | Environment Variable   |
| ------------ | ------ | ------------ | ---------------------- |
| `BearerAuth` | http   | HTTP Bearer  | `PIPESHUB_BEARER_AUTH` |
| `Oauth2`     | oauth2 | OAuth2 token | `PIPESHUB_OAUTH2`      |

You can set the security parameters through the `WithSecurity` option when initializing the SDK client instance. The selected scheme will be used by default to authenticate with the API for all operations that support it. For example:
```go
package main

import (
	"context"
	pipeshub "github.com/pipeshub-ai/pipeshub-sdk-go/v2"
	"github.com/pipeshub-ai/pipeshub-sdk-go/v2/models/components"
	"log"
	"os"
)

func main() {
	ctx := context.Background()

	s := pipeshub.New(
		pipeshub.WithSecurity(components.Security{
			BearerAuth: pipeshub.Pointer(os.Getenv("PIPESHUB_BEARER_AUTH")),
		}),
	)

	res, err := s.OAuthProvider.OauthToken(ctx, components.OAuthTokenRequest{
		GrantType: components.GrantTypeClientCredentials,
	})
	if err != nil {
		log.Fatal(err)
	}
	if res.OAuthTokenResponse != nil {
		// handle response
	}
}

```

### Per-Operation Security Schemes

Some operations in this SDK require the security scheme to be specified at the request level. For example:
```go
package main

import (
	"context"
	pipeshub "github.com/pipeshub-ai/pipeshub-sdk-go/v2"
	"github.com/pipeshub-ai/pipeshub-sdk-go/v2/models/operations"
	"log"
	"os"
)

func main() {
	ctx := context.Background()

	s := pipeshub.New()

	res, err := s.UserAccount.RefreshToken(ctx, operations.RefreshTokenSecurity{
		ScopedToken: os.Getenv("PIPESHUB_SCOPED_TOKEN"),
	})
	if err != nil {
		log.Fatal(err)
	}
	if res.RefreshTokenResponse != nil {
		// handle response
	}
}

```
<!-- End Authentication [security] -->

<!-- Start Server-sent event streaming [eventstream] -->
## Server-sent event streaming

[Server-sent events][mdn-sse] are used to stream content from certain
operations. These operations will expose the stream as an iterable that
can be consumed using a simple `for` loop. The loop will
terminate when the server no longer has any events to send and closes the
underlying connection.

```go
package main

import (
	"context"
	pipeshub "github.com/pipeshub-ai/pipeshub-sdk-go/v2"
	"github.com/pipeshub-ai/pipeshub-sdk-go/v2/models/components"
	"github.com/pipeshub-ai/pipeshub-sdk-go/v2/models/operations"
	"log"
	"os"
)

func main() {
	ctx := context.Background()

	s := pipeshub.New(
		pipeshub.WithSecurity(components.Security{
			BearerAuth: pipeshub.Pointer(os.Getenv("PIPESHUB_BEARER_AUTH")),
		}),
	)

	res, err := s.KnowledgeBase.UploadRecords(ctx, "<id>", operations.UploadRecordsRequestBody{
		Files: []operations.UploadRecordsFile{},
		FilesMetadata: []operations.FilesMetadatum{
			operations.FilesMetadatum{
				FilePath:     "/docs/report.pdf",
				LastModified: 2024.0,
			},
		},
	}, nil)
	if err != nil {
		log.Fatal(err)
	}
	if res.UploadStreamSSEEvent != nil {
		defer res.UploadStreamSSEEvent.Close()

		for res.UploadStreamSSEEvent.Next() {
			event := res.UploadStreamSSEEvent.Value()
			log.Print(event)
			// Handle the event
		}
	}
}

```

[mdn-sse]: https://developer.mozilla.org/en-US/docs/Web/API/Server-sent_events/Using_server-sent_events
<!-- End Server-sent event streaming [eventstream] -->

<!-- Start Retries [retries] -->
## Retries

Some of the endpoints in this SDK support retries. If you use the SDK without any configuration, it will fall back to the default retry strategy provided by the API. However, the default retry strategy can be overridden on a per-operation basis, or across the entire SDK.

To change the default retry strategy for a single API call, simply provide a `retry.Config` object to the call by using the `WithRetries` option:
```go
package main

import (
	"context"
	pipeshub "github.com/pipeshub-ai/pipeshub-sdk-go/v2"
	"github.com/pipeshub-ai/pipeshub-sdk-go/v2/models/components"
	"github.com/pipeshub-ai/pipeshub-sdk-go/v2/retry"
	"log"
	"models/operations"
)

func main() {
	ctx := context.Background()

	s := pipeshub.New()

	res, err := s.OAuthProvider.OauthToken(ctx, components.OAuthTokenRequest{
		GrantType: components.GrantTypeClientCredentials,
	}, operations.WithRetries(
		retry.Config{
			Strategy: "backoff",
			Backoff: &retry.BackoffStrategy{
				InitialInterval: 1,
				MaxInterval:     50,
				Exponent:        1.1,
				MaxElapsedTime:  100,
			},
			RetryConnectionErrors: false,
		}))
	if err != nil {
		log.Fatal(err)
	}
	if res.OAuthTokenResponse != nil {
		// handle response
	}
}

```

If you'd like to override the default retry strategy for all operations that support retries, you can use the `WithRetryConfig` option at SDK initialization:
```go
package main

import (
	"context"
	pipeshub "github.com/pipeshub-ai/pipeshub-sdk-go/v2"
	"github.com/pipeshub-ai/pipeshub-sdk-go/v2/models/components"
	"github.com/pipeshub-ai/pipeshub-sdk-go/v2/retry"
	"log"
)

func main() {
	ctx := context.Background()

	s := pipeshub.New(
		pipeshub.WithRetryConfig(
			retry.Config{
				Strategy: "backoff",
				Backoff: &retry.BackoffStrategy{
					InitialInterval: 1,
					MaxInterval:     50,
					Exponent:        1.1,
					MaxElapsedTime:  100,
				},
				RetryConnectionErrors: false,
			}),
	)

	res, err := s.OAuthProvider.OauthToken(ctx, components.OAuthTokenRequest{
		GrantType: components.GrantTypeClientCredentials,
	})
	if err != nil {
		log.Fatal(err)
	}
	if res.OAuthTokenResponse != nil {
		// handle response
	}
}

```
<!-- End Retries [retries] -->

<!-- Start Error Handling [errors] -->
## Error Handling

Handling errors in this SDK should largely match your expectations. All operations return a response object or an error, they will never return both.

By Default, an API error will return `apierrors.APIError`. When custom error responses are specified for an operation, the SDK may also return their associated error. You can refer to respective *Errors* tables in SDK docs for more details on possible error types for each operation.

For example, the `OauthToken` function may return the following errors:

| Error Type                                    | Status Code | Content Type     |
| --------------------------------------------- | ----------- | ---------------- |
| apierrors.BadRequest                          | 400         | application/json |
| apierrors.OAuthErrorResponse                  | 401         | application/json |
| apierrors.OAuthClientManagementRateLimitError | 429         | application/json |
| apierrors.ErrorResponse                       | 500         | application/json |
| apierrors.APIError                            | 4XX, 5XX    | \*/\*            |

### Example

```go
package main

import (
	"context"
	"errors"
	pipeshub "github.com/pipeshub-ai/pipeshub-sdk-go/v2"
	"github.com/pipeshub-ai/pipeshub-sdk-go/v2/models/apierrors"
	"github.com/pipeshub-ai/pipeshub-sdk-go/v2/models/components"
	"log"
)

func main() {
	ctx := context.Background()

	s := pipeshub.New()

	res, err := s.OAuthProvider.OauthToken(ctx, components.OAuthTokenRequest{
		GrantType: components.GrantTypeClientCredentials,
	})
	if err != nil {

		var e *apierrors.BadRequest
		if errors.As(err, &e) {
			// handle error
			log.Fatal(e.Error())
		}

		var e *apierrors.OAuthErrorResponse
		if errors.As(err, &e) {
			// handle error
			log.Fatal(e.Error())
		}

		var e *apierrors.OAuthClientManagementRateLimitError
		if errors.As(err, &e) {
			// handle error
			log.Fatal(e.Error())
		}

		var e *apierrors.ErrorResponse
		if errors.As(err, &e) {
			// handle error
			log.Fatal(e.Error())
		}

		var e *apierrors.APIError
		if errors.As(err, &e) {
			// handle error
			log.Fatal(e.Error())
		}
	}
}

```
<!-- End Error Handling [errors] -->

<!-- Start Server Selection [server] -->
## Server Selection

### Select Server by Index

You can override the default server globally using the `WithServerIndex(serverIndex int)` option when initializing the SDK client instance. The selected server will then be used as the default on the operations that use it. This table lists the indexes associated with the available servers:

| #   | Server                          | Variables      | Description                                       |
| --- | ------------------------------- | -------------- | ------------------------------------------------- |
| 0   | `https://{instance_url}/api/v1` | `instance_url` | Base API URL                                      |
| 1   | `https://{instance_url}`        | `instance_url` | Root URL (used for MCP endpoints mounted at /mcp) |

If the selected server has variables, you may override its default values using the associated option(s):

| Variable       | Option                                | Default              | Description                                                                                                                                                  |
| -------------- | ------------------------------------- | -------------------- | ------------------------------------------------------------------------------------------------------------------------------------------------------------ |
| `instance_url` | `WithInstanceURL(instanceURL string)` | `"app.pipeshub.com"` | Host name, with an optional port, of a PipesHub instance served over HTTPS. For a plain-HTTP instance, such as a local one, set the full server URL instead. |

#### Example

```go
package main

import (
	"context"
	pipeshub "github.com/pipeshub-ai/pipeshub-sdk-go/v2"
	"github.com/pipeshub-ai/pipeshub-sdk-go/v2/models/components"
	"log"
)

func main() {
	ctx := context.Background()

	s := pipeshub.New(
		pipeshub.WithServerIndex(0),
		pipeshub.WithInstanceURL("app.pipeshub.com"),
	)

	res, err := s.OAuthProvider.OauthToken(ctx, components.OAuthTokenRequest{
		GrantType: components.GrantTypeClientCredentials,
	})
	if err != nil {
		log.Fatal(err)
	}
	if res.OAuthTokenResponse != nil {
		// handle response
	}
}

```

### Override Server URL Per-Client

The default server can also be overridden globally using the `WithServerURL(serverURL string)` option when initializing the SDK client instance. For example:
```go
package main

import (
	"context"
	pipeshub "github.com/pipeshub-ai/pipeshub-sdk-go/v2"
	"github.com/pipeshub-ai/pipeshub-sdk-go/v2/models/components"
	"log"
)

func main() {
	ctx := context.Background()

	s := pipeshub.New(
		pipeshub.WithServerURL("https://app.pipeshub.com"),
	)

	res, err := s.OAuthProvider.OauthToken(ctx, components.OAuthTokenRequest{
		GrantType: components.GrantTypeClientCredentials,
	})
	if err != nil {
		log.Fatal(err)
	}
	if res.OAuthTokenResponse != nil {
		// handle response
	}
}

```
<!-- End Server Selection [server] -->

<!-- Start Custom HTTP Client [http-client] -->
## Custom HTTP Client

The Go SDK makes API calls that wrap an internal HTTP client. The requirements for the HTTP client are very simple. It must match this interface:

```go
type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}
```

The built-in `net/http` client satisfies this interface and a default client based on the built-in is provided by default. To replace this default with a client of your own, you can implement this interface yourself or provide your own client configured as desired. Here's a simple example, which adds a client with a 30 second timeout.

```go
import (
	"net/http"
	"time"

	"github.com/pipeshub-ai/pipeshub-sdk-go/v2"
)

var (
	httpClient = &http.Client{Timeout: 30 * time.Second}
	sdkClient  = pipeshub.New(pipeshub.WithClient(httpClient))
)
```

This can be a convenient way to configure timeouts, cookies, proxies, custom headers, and other low-level configuration.
<!-- End Custom HTTP Client [http-client] -->

<!-- Start Special Types [types] -->
## Special Types

This SDK defines the following custom types to assist with marshalling and unmarshalling data.

### Date

`types.Date` is a wrapper around time.Time that allows for JSON marshaling a date string formatted as "2006-01-02".

#### Usage

```go
d1 := types.NewDate(time.Now()) // returns *types.Date

d2 := types.DateFromTime(time.Now()) // returns types.Date

d3, err := types.NewDateFromString("2019-01-01") // returns *types.Date, error

d4, err := types.DateFromString("2019-01-01") // returns types.Date, error

d5 := types.MustNewDateFromString("2019-01-01") // returns *types.Date and panics on error

d6 := types.MustDateFromString("2019-01-01") // returns types.Date and panics on error
```
<!-- End Special Types [types] -->

<!-- Start Available Resources and Operations [operations] -->
## Available Resources and Operations

<details open>
<summary>Available methods</summary>

### [Agents](docs/sdks/agents/README.md)

* [ListAgents](docs/sdks/agents/README.md#listagents) - List agents
* [CreateAgent](docs/sdks/agents/README.md#createagent) - Create agent
* [GetAgent](docs/sdks/agents/README.md#getagent) - Get agent
* [UpdateAgent](docs/sdks/agents/README.md#updateagent) - Update agent
* [DeleteAgent](docs/sdks/agents/README.md#deleteagent) - Delete agent
* [ListAgentArchivedConversationsGrouped](docs/sdks/agents/README.md#listagentarchivedconversationsgrouped) - List archived agent conversations grouped by agent
* [ListAgentConversationArchives](docs/sdks/agents/README.md#listagentconversationarchives) - List archived conversations for an agent
* [UploadAgentConversationChatAttachments](docs/sdks/agents/README.md#uploadagentconversationchatattachments) - Upload agent chat attachments
* [DeleteAgentConversationChatAttachment](docs/sdks/agents/README.md#deleteagentconversationchatattachment) - Delete an agent chat attachment
* [StreamAgentConversation](docs/sdks/agents/README.md#streamagentconversation) - Create agent conversation with streaming response
* [StreamAgentConversationMessage](docs/sdks/agents/README.md#streamagentconversationmessage) - Add message to agent conversation with streaming response
* [RegenerateAgentConversationMessage](docs/sdks/agents/README.md#regenerateagentconversationmessage) - Regenerate agent conversation message
* [ForkAgentConversation](docs/sdks/agents/README.md#forkagentconversation) - Fork agent conversation from a message
* [CancelAgentConversationStream](docs/sdks/agents/README.md#cancelagentconversationstream) - Cancel an in-flight agent chat stream
* [UpdateAgentConversationMessageFeedback](docs/sdks/agents/README.md#updateagentconversationmessagefeedback) - Submit feedback for an agent message
* [AddAgentConversationMessage](docs/sdks/agents/README.md#addagentconversationmessage) - Add message to agent conversation (non-streaming)
* [ArchiveAgentConversation](docs/sdks/agents/README.md#archiveagentconversation) - Archive an agent conversation
* [UnarchiveAgentConversation](docs/sdks/agents/README.md#unarchiveagentconversation) - Unarchive an agent conversation
* [UpdateAgentConversationTitle](docs/sdks/agents/README.md#updateagentconversationtitle) - Update agent conversation title
* [DeleteAgentConversationByID](docs/sdks/agents/README.md#deleteagentconversationbyid) - Delete an agent conversation
* [GetAgentConversationByID](docs/sdks/agents/README.md#getagentconversationbyid) - Get agent conversation by ID
* [ListAgentConversations](docs/sdks/agents/README.md#listagentconversations) - List agent conversations
* [CreateAgentConversation](docs/sdks/agents/README.md#createagentconversation) - Create agent conversation (non-streaming)

### [AIModelsProviders](docs/sdks/aimodelsproviders/README.md)

* [GetAvailableModelsByType](docs/sdks/aimodelsproviders/README.md#getavailablemodelsbytype) - Get available models by type

### [Connector](docs/sdks/connector/README.md)

* [GetRecordContent](docs/sdks/connector/README.md#getrecordcontent) - Get a record's full parsed content and metadata
* [NavigateKnowledgeGraph](docs/sdks/connector/README.md#navigateknowledgegraph) - Browse the knowledge graph from a node
* [LookupRecordByIdentifier](docs/sdks/connector/README.md#lookuprecordbyidentifier) - Resolve a URL, issue key or external ID to a Record ID

### [ConnectorConfiguration](docs/sdks/connectorconfiguration/README.md)

* [GetConnectorConfig](docs/sdks/connectorconfiguration/README.md#getconnectorconfig) - Get connector configuration
* [UpdateConnectorConfig](docs/sdks/connectorconfiguration/README.md#updateconnectorconfig) - Update connector configuration
* [UpdateConnectorAuthConfig](docs/sdks/connectorconfiguration/README.md#updateconnectorauthconfig) - Update authentication configuration
* [UpdateConnectorFiltersSyncConfig](docs/sdks/connectorconfiguration/README.md#updateconnectorfilterssyncconfig) - Update filters and sync configuration

### [ConnectorControl](docs/sdks/connectorcontrol/README.md)

* [ToggleConnector](docs/sdks/connectorcontrol/README.md#toggleconnector) - Toggle connector sync or agent

### [ConnectorFilters](docs/sdks/connectorfilters/README.md)

* [GetConnectorFilters](docs/sdks/connectorfilters/README.md#getconnectorfilters) - Get filter options
* [~~SaveConnectorFilters~~](docs/sdks/connectorfilters/README.md#saveconnectorfilters) - Save filter selections (deprecated) :warning: **Deprecated**
* [GetFilterFieldOptions](docs/sdks/connectorfilters/README.md#getfilterfieldoptions) - Get dynamic filter options

### [ConnectorInstances](docs/sdks/connectorinstances/README.md)

* [ListConnectorInstances](docs/sdks/connectorinstances/README.md#listconnectorinstances) - List connector instances
* [CreateConnectorInstance](docs/sdks/connectorinstances/README.md#createconnectorinstance) - Create connector instance
* [GetConnectorInstance](docs/sdks/connectorinstances/README.md#getconnectorinstance) - Get connector instance
* [DeleteConnectorInstance](docs/sdks/connectorinstances/README.md#deleteconnectorinstance) - Delete connector instance
* [GetConnectorStats](docs/sdks/connectorinstances/README.md#getconnectorstats) - Get connector statistics
* [ResyncConnector](docs/sdks/connectorinstances/README.md#resyncconnector) - Resync connector
* [StopConnectorSync](docs/sdks/connectorinstances/README.md#stopconnectorsync) - Stop the running sync for a connector

### [ConnectorOAuth](docs/sdks/connectoroauth/README.md)

* [GetOAuthAuthorizationURL](docs/sdks/connectoroauth/README.md#getoauthauthorizationurl) - Get OAuth authorization URL
* [HandleOAuthCallback](docs/sdks/connectoroauth/README.md#handleoauthcallback) - OAuth callback handler

### [ConnectorRegistry](docs/sdks/connectorregistry/README.md)

* [GetConnectorRegistry](docs/sdks/connectorregistry/README.md#getconnectorregistry) - List available connector types
* [GetConnectorSchema](docs/sdks/connectorregistry/README.md#getconnectorschema) - Get connector configuration schema
* [TestConnectorConnection](docs/sdks/connectorregistry/README.md#testconnectorconnection) - Test connector credentials before saving
* [GetConnectorEgressIps](docs/sdks/connectorregistry/README.md#getconnectoregressips) - Get the IP addresses connectors connect from

### [Conversations](docs/sdks/conversations/README.md)

* [CreateConversation](docs/sdks/conversations/README.md#createconversation) - Create conversation (non-streaming)
* [StreamChat](docs/sdks/conversations/README.md#streamchat) - Create conversation with streaming response
* [GetAllConversations](docs/sdks/conversations/README.md#getallconversations) - List all conversations
* [GetArchivedConversations](docs/sdks/conversations/README.md#getarchivedconversations) - List archived conversations
* [SearchArchivedConversations](docs/sdks/conversations/README.md#searcharchivedconversations) - Search archived conversations
* [GetConversationByID](docs/sdks/conversations/README.md#getconversationbyid) - Get conversation by ID
* [DeleteConversationByID](docs/sdks/conversations/README.md#deleteconversationbyid) - Delete conversation
* [AddMessage](docs/sdks/conversations/README.md#addmessage) - Add message (non-streaming)
* [AddMessageStream](docs/sdks/conversations/README.md#addmessagestream) - Add message to a conversation with streaming response
* [UpdateConversationTitle](docs/sdks/conversations/README.md#updateconversationtitle) - Update conversation title
* [ArchiveConversation](docs/sdks/conversations/README.md#archiveconversation) - Archive conversation
* [UnarchiveConversation](docs/sdks/conversations/README.md#unarchiveconversation) - Unarchive conversation
* [PinConversation](docs/sdks/conversations/README.md#pinconversation) - Pin conversation
* [UnpinConversation](docs/sdks/conversations/README.md#unpinconversation) - Unpin conversation
* [RegenerateAnswer](docs/sdks/conversations/README.md#regenerateanswer) - Regenerate AI response
* [ForkConversation](docs/sdks/conversations/README.md#forkconversation) - Fork conversation from a message
* [CancelConversationStream](docs/sdks/conversations/README.md#cancelconversationstream) - Cancel an in-flight chat stream
* [UpdateMessageFeedback](docs/sdks/conversations/README.md#updatemessagefeedback) - Submit feedback on AI response

### [KnowledgeBase](docs/sdks/knowledgebase/README.md)

* [CreateKnowledgeBase](docs/sdks/knowledgebase/README.md#createknowledgebase) - Create a new knowledge base
* [ListKnowledgeBases](docs/sdks/knowledgebase/README.md#listknowledgebases) - List all knowledge bases
* [GetKnowledgeBase](docs/sdks/knowledgebase/README.md#getknowledgebase) - Get knowledge base by ID
* [UpdateKnowledgeBase](docs/sdks/knowledgebase/README.md#updateknowledgebase) - Update knowledge base
* [DeleteKnowledgeBase](docs/sdks/knowledgebase/README.md#deleteknowledgebase) - Delete knowledge base
* [GetRecordByID](docs/sdks/knowledgebase/README.md#getrecordbyid) - Get record by ID
* [UpdateRecordMultipart](docs/sdks/knowledgebase/README.md#updaterecordmultipart) - Update record
* [UpdateRecord](docs/sdks/knowledgebase/README.md#updaterecord) - Update record
* [UpdateRecordForm](docs/sdks/knowledgebase/README.md#updaterecordform) - Update record
* [DeleteRecord](docs/sdks/knowledgebase/README.md#deleterecord) - Delete record
* [RestoreRecord](docs/sdks/knowledgebase/README.md#restorerecord) - Restore a deleted record
* [RestoreRecords](docs/sdks/knowledgebase/README.md#restorerecords) - Restore several deleted records
* [ListTrash](docs/sdks/knowledgebase/README.md#listtrash) - List a collection's recently deleted items
* [StreamRecordBuffer](docs/sdks/knowledgebase/README.md#streamrecordbuffer) - Stream record content
* [CreateFolder](docs/sdks/knowledgebase/README.md#createfolder) - Create folder
* [CreateSubfolder](docs/sdks/knowledgebase/README.md#createsubfolder) - Create subfolder
* [UpdateFolder](docs/sdks/knowledgebase/README.md#updatefolder) - Update folder
* [DeleteFolder](docs/sdks/knowledgebase/README.md#deletefolder) - Delete folder
* [UploadRecords](docs/sdks/knowledgebase/README.md#uploadrecords) - Upload files to knowledge base or folder
* [GetUploadLimits](docs/sdks/knowledgebase/README.md#getuploadlimits) - Get knowledge base upload limits
* [ReindexRecord](docs/sdks/knowledgebase/README.md#reindexrecord) - Reindex single record
* [ReindexRecordGroup](docs/sdks/knowledgebase/README.md#reindexrecordgroup) - Reindex record group
* [MoveRecord](docs/sdks/knowledgebase/README.md#moverecord) - Move record to another location
* [~~GetKnowledgeHubRootNodes~~](docs/sdks/knowledgebase/README.md#getknowledgehubrootnodes) - Get knowledge hub root nodes :warning: **Deprecated**
* [~~GetKnowledgeHubChildNodes~~](docs/sdks/knowledgebase/README.md#getknowledgehubchildnodes) - Get knowledge hub child nodes :warning: **Deprecated**

### [~~KnowledgeHub~~](docs/sdks/knowledgehub/README.md)

* [~~GetKnowledgeHubRootNodes~~](docs/sdks/knowledgehub/README.md#getknowledgehubrootnodes) - Get knowledge hub root nodes :warning: **Deprecated**
* [~~GetKnowledgeHubChildNodes~~](docs/sdks/knowledgehub/README.md#getknowledgehubchildnodes) - Get knowledge hub child nodes :warning: **Deprecated**

### [OAuthApps](docs/sdks/oauthapps/README.md)

* [ListOAuthApps](docs/sdks/oauthapps/README.md#listoauthapps) - List OAuth apps
* [CreateOAuthApp](docs/sdks/oauthapps/README.md#createoauthapp) - Create OAuth app
* [ListOAuthScopes](docs/sdks/oauthapps/README.md#listoauthscopes) - List available scopes
* [GetOAuthApp](docs/sdks/oauthapps/README.md#getoauthapp) - Get OAuth app details
* [UpdateOAuthApp](docs/sdks/oauthapps/README.md#updateoauthapp) - Update OAuth app
* [DeleteOAuthApp](docs/sdks/oauthapps/README.md#deleteoauthapp) - Delete OAuth app
* [RegenerateOAuthAppSecret](docs/sdks/oauthapps/README.md#regenerateoauthappsecret) - Regenerate client secret
* [SuspendOAuthApp](docs/sdks/oauthapps/README.md#suspendoauthapp) - Suspend OAuth app
* [ActivateOAuthApp](docs/sdks/oauthapps/README.md#activateoauthapp) - Activate suspended OAuth app
* [ListOAuthAppTokens](docs/sdks/oauthapps/README.md#listoauthapptokens) - List app tokens
* [RevokeAllOAuthAppTokens](docs/sdks/oauthapps/README.md#revokealloauthapptokens) - Revoke all app tokens

### [OAuthConfiguration](docs/sdks/oauthconfiguration/README.md)

* [GetOAuthRegistry](docs/sdks/oauthconfiguration/README.md#getoauthregistry) - List OAuth-capable connector types
* [GetOAuthConnectorType](docs/sdks/oauthconfiguration/README.md#getoauthconnectortype) - Get OAuth connector type details
* [ListOAuthConfigs](docs/sdks/oauthconfiguration/README.md#listoauthconfigs) - List OAuth configurations
* [ListOAuthConfigsByType](docs/sdks/oauthconfiguration/README.md#listoauthconfigsbytype) - List OAuth configs for connector type
* [CreateOAuthConfig](docs/sdks/oauthconfiguration/README.md#createoauthconfig) - Create OAuth configuration
* [GetOAuthConfig](docs/sdks/oauthconfiguration/README.md#getoauthconfig) - Get OAuth configuration
* [UpdateOAuthConfig](docs/sdks/oauthconfiguration/README.md#updateoauthconfig) - Update OAuth configuration
* [DeleteOAuthConfig](docs/sdks/oauthconfiguration/README.md#deleteoauthconfig) - Delete OAuth configuration

### [OAuthProvider](docs/sdks/oauthprovider/README.md)

* [OauthToken](docs/sdks/oauthprovider/README.md#oauthtoken) - Exchange authorization code for tokens
* [OauthRevoke](docs/sdks/oauthprovider/README.md#oauthrevoke) - Revoke an access or refresh token
* [OauthIntrospect](docs/sdks/oauthprovider/README.md#oauthintrospect) - Introspect a token

### [OpenIDConnect](docs/sdks/openidconnect/README.md)

* [OauthUserInfo](docs/sdks/openidconnect/README.md#oauthuserinfo) - Get authenticated user information

### [OrganizationAuthConfig](docs/sdks/organizationauthconfig/README.md)

* [GetAuthMethods](docs/sdks/organizationauthconfig/README.md#getauthmethods) - Get organization authentication methods
* [UpdateAuthMethod](docs/sdks/organizationauthconfig/README.md#updateauthmethod) - Update organization authentication methods
* [SetUpAuthConfig](docs/sdks/organizationauthconfig/README.md#setupauthconfig) - Set up auth configuration

### [Organizations](docs/sdks/organizations/README.md)

* [GetCurrentOrganization](docs/sdks/organizations/README.md#getcurrentorganization) - Get current organization

### [PersonalAccessTokens](docs/sdks/personalaccesstokens/README.md)

* [ListPersonalAccessTokens](docs/sdks/personalaccesstokens/README.md#listpersonalaccesstokens) - List your own personal access tokens
* [CreatePersonalAccessToken](docs/sdks/personalaccesstokens/README.md#createpersonalaccesstoken) - Create a personal access token
* [ListPersonalAccessTokenScopes](docs/sdks/personalaccesstokens/README.md#listpersonalaccesstokenscopes) - List scopes available for a new personal access token
* [RevokePersonalAccessToken](docs/sdks/personalaccesstokens/README.md#revokepersonalaccesstoken) - Revoke one of your own personal access tokens
* [AdminListPersonalAccessTokens](docs/sdks/personalaccesstokens/README.md#adminlistpersonalaccesstokens) - Admin: list every active personal access token in the org
* [AdminRevokePersonalAccessToken](docs/sdks/personalaccesstokens/README.md#adminrevokepersonalaccesstoken) - Admin: revoke any user's personal access token by id

### [SemanticSearch](docs/sdks/semanticsearch/README.md)

* [Retrieve](docs/sdks/semanticsearch/README.md#retrieve) - Retrieve chunks with citations
* [Search](docs/sdks/semanticsearch/README.md#search) - Perform semantic search
* [SearchHistory](docs/sdks/semanticsearch/README.md#searchhistory) - Get search history
* [DeleteSearchHistory](docs/sdks/semanticsearch/README.md#deletesearchhistory) - Clear all search history
* [GetSearchByID](docs/sdks/semanticsearch/README.md#getsearchbyid) - Get search by ID
* [DeleteSearchByID](docs/sdks/semanticsearch/README.md#deletesearchbyid) - Delete search by ID
* [ArchiveSearch](docs/sdks/semanticsearch/README.md#archivesearch) - Archive a search
* [UnarchiveSearch](docs/sdks/semanticsearch/README.md#unarchivesearch) - Unarchive a search

### [UserAccount](docs/sdks/useraccount/README.md)

* [InitAuth](docs/sdks/useraccount/README.md#initauth) - Initialize authentication session
* [Authenticate](docs/sdks/useraccount/README.md#authenticate) - Authenticate user with credentials
* [RefreshToken](docs/sdks/useraccount/README.md#refreshtoken) - Refresh access token
* [GetPasswordStatus](docs/sdks/useraccount/README.md#getpasswordstatus) - Get password status
* [SendPasswordSetupEmail](docs/sdks/useraccount/README.md#sendpasswordsetupemail) - Email a link to set a first password
* [ResetPassword](docs/sdks/useraccount/README.md#resetpassword) - Reset password

### [WebSearch](docs/sdks/websearch/README.md)

* [GetWebSearchProviders](docs/sdks/websearch/README.md#getwebsearchproviders) - Get all web search providers

</details>
<!-- End Available Resources and Operations [operations] -->

<!-- Placeholder for Future Speakeasy SDK Sections -->

## Support

- Bugs and feature requests: [open an issue](https://github.com/pipeshub-ai/pipeshub-sdk-go/issues).
- Questions: the [PipesHub Discord](https://discord.com/invite/K5RskzJBm2).
- PipesHub itself: [pipeshub-ai/pipeshub-ai](https://github.com/pipeshub-ai/pipeshub-ai).

## Contributing

Most of this SDK is generated, so direct changes to generated files are overwritten by the next release. Please open an issue instead. See [CONTRIBUTING.md](CONTRIBUTING.md).

## License

Licensed under the [Apache License 2.0](LICENSE).
