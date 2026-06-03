# pipeshub-sdk-go
pipeshub-sdk is the official Go client library for integrating Pipeshub into your product and internal tools

<!-- Start Summary [summary] -->
## Summary

PipesHub API: Unified API documentation for PipesHub services.

PipesHub is an enterprise-grade platform providing:
- User authentication and management
- Document storage and version control
- Knowledge base management
- Enterprise search and conversational AI
- Third-party integrations via connectors
- System configuration management
- Crawling job scheduling
- Email services

## Authentication
Most endpoints require JWT Bearer token authentication. Some internal endpoints use scoped tokens for service-to-service communication.

**OAuth 2.0 Bearer tokens** from `POST /oauth2/token` use the same `Authorization: Bearer` header. For **`client_credentials`**, machine tokens may encode `userId === client_id` in the JWT; the **Node API gateway** resolves the OAuth **app creator**, sets the authenticated user accordingly, and forwards **`x-oauth-user-id`** to Python services on proxied calls. Do not send **`x-oauth-user-id`** yourself—the gateway removes untrusted values on ingress. See the **OAuth Provider** tag for full behavior.

## Base URLs
All endpoints use the `/api/v1` prefix unless otherwise noted.
<!-- End Summary [summary] -->

<!-- Start Table of Contents [toc] -->
## Table of Contents
<!-- $toc-max-depth=2 -->
* [pipeshub-sdk-go](#pipeshub-sdk-go)
  * [Authentication](#authentication)
  * [Base URLs](#base-urls)
  * [SDK Installation](#sdk-installation)
  * [SDK Example Usage](#sdk-example-usage)
  * [Authentication](#authentication-1)
  * [Available Resources and Operations](#available-resources-and-operations)
  * [Server-sent event streaming](#server-sent-event-streaming)
  * [Retries](#retries)
  * [Error Handling](#error-handling)
  * [Server Selection](#server-selection)
  * [Custom HTTP Client](#custom-http-client)

<!-- End Table of Contents [toc] -->

<!-- Start SDK Installation [installation] -->
## SDK Installation

To add the SDK as a dependency to your project:
```bash
go get github.com/pipeshub-ai/pipeshub-sdk-go
```
<!-- End SDK Installation [installation] -->

<!-- Start SDK Example Usage [usage] -->
## SDK Example Usage

### Example

```go
package main

import (
	"context"
	pipeshub "github.com/pipeshub-ai/pipeshub-sdk-go"
	"github.com/pipeshub-ai/pipeshub-sdk-go/models/components"
	"log"
)

func main() {
	ctx := context.Background()

	s := pipeshub.New()

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
	pipeshub "github.com/pipeshub-ai/pipeshub-sdk-go"
	"github.com/pipeshub-ai/pipeshub-sdk-go/models/components"
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
	pipeshub "github.com/pipeshub-ai/pipeshub-sdk-go"
	"github.com/pipeshub-ai/pipeshub-sdk-go/models/operations"
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
* [UpdateAgentConversationMessageFeedback](docs/sdks/agents/README.md#updateagentconversationmessagefeedback) - Submit feedback for an agent message
* [ArchiveAgentConversation](docs/sdks/agents/README.md#archiveagentconversation) - Archive an agent conversation
* [UnarchiveAgentConversation](docs/sdks/agents/README.md#unarchiveagentconversation) - Unarchive an agent conversation
* [UpdateAgentConversationTitle](docs/sdks/agents/README.md#updateagentconversationtitle) - Update agent conversation title
* [DeleteAgentConversationByID](docs/sdks/agents/README.md#deleteagentconversationbyid) - Delete an agent conversation
* [GetAgentConversationByID](docs/sdks/agents/README.md#getagentconversationbyid) - Get agent conversation by ID
* [ListAgentConversations](docs/sdks/agents/README.md#listagentconversations) - List agent conversations

### [AIModelsProviders](docs/sdks/aimodelsproviders/README.md)

* [GetAvailableModelsByType](docs/sdks/aimodelsproviders/README.md#getavailablemodelsbytype) - Get available models by type

### [Conversations](docs/sdks/conversations/README.md)

* [StreamChat](docs/sdks/conversations/README.md#streamchat) - Create conversation with streaming response
* [GetAllConversations](docs/sdks/conversations/README.md#getallconversations) - List all conversations
* [GetArchivedConversations](docs/sdks/conversations/README.md#getarchivedconversations) - List archived conversations
* [SearchArchivedConversations](docs/sdks/conversations/README.md#searcharchivedconversations) - Search archived conversations
* [GetConversationByID](docs/sdks/conversations/README.md#getconversationbyid) - Get conversation by ID
* [DeleteConversationByID](docs/sdks/conversations/README.md#deleteconversationbyid) - Delete conversation
* [AddMessageStream](docs/sdks/conversations/README.md#addmessagestream) - Add message to a conversation with streaming response
* [UpdateConversationTitle](docs/sdks/conversations/README.md#updateconversationtitle) - Update conversation title
* [ArchiveConversation](docs/sdks/conversations/README.md#archiveconversation) - Archive conversation
* [UnarchiveConversation](docs/sdks/conversations/README.md#unarchiveconversation) - Unarchive conversation
* [RegenerateAnswer](docs/sdks/conversations/README.md#regenerateanswer) - Regenerate AI response
* [UpdateMessageFeedback](docs/sdks/conversations/README.md#updatemessagefeedback) - Submit feedback on AI response

### [KnowledgeHub](docs/sdks/knowledgehub/README.md)

* [GetKnowledgeHubRootNodes](docs/sdks/knowledgehub/README.md#getknowledgehubrootnodes) - Get knowledge hub root nodes
* [GetKnowledgeHubChildNodes](docs/sdks/knowledgehub/README.md#getknowledgehubchildnodes) - Get knowledge hub child nodes

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

### [SemanticSearch](docs/sdks/semanticsearch/README.md)

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
* [ResetPassword](docs/sdks/useraccount/README.md#resetpassword) - Reset password

### [WebSearch](docs/sdks/websearch/README.md)

* [GetWebSearchProviders](docs/sdks/websearch/README.md#getwebsearchproviders) - Get all web search providers

</details>
<!-- End Available Resources and Operations [operations] -->

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
	pipeshub "github.com/pipeshub-ai/pipeshub-sdk-go"
	"github.com/pipeshub-ai/pipeshub-sdk-go/models/components"
	"github.com/pipeshub-ai/pipeshub-sdk-go/types"
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

	res, err := s.Conversations.StreamChat(ctx, components.CreateConversationRequest{
		Query: "What are the key findings from our Q4 financial report?",
		RecordIds: []string{
			"507f1f77bcf86cd799439011",
			"507f1f77bcf86cd799439012",
		},
		ModelKey:          pipeshub.Pointer("gpt-4-turbo"),
		ModelName:         pipeshub.Pointer("GPT-4 Turbo"),
		ModelFriendlyName: pipeshub.Pointer("GPT-4 Turbo"),
		ChatMode:          components.CreateConversationRequestChatModeWebSearch.ToPointer(),
		Timezone:          pipeshub.Pointer("America/New_York"),
		CurrentTime:       types.MustNewTimeFromString("2026-04-12T16:00:00+05:30"),
		Tools: []string{
			"jira.create_issue",
			"confluence.search_content",
		},
	})
	if err != nil {
		log.Fatal(err)
	}
	if res.AssistantStreamSSEEvent != nil {
		defer res.AssistantStreamSSEEvent.Close()

		for res.AssistantStreamSSEEvent.Next() {
			event := res.AssistantStreamSSEEvent.Value()
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
	pipeshub "github.com/pipeshub-ai/pipeshub-sdk-go"
	"github.com/pipeshub-ai/pipeshub-sdk-go/models/components"
	"github.com/pipeshub-ai/pipeshub-sdk-go/retry"
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
	pipeshub "github.com/pipeshub-ai/pipeshub-sdk-go"
	"github.com/pipeshub-ai/pipeshub-sdk-go/models/components"
	"github.com/pipeshub-ai/pipeshub-sdk-go/retry"
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
| apierrors.ErrorResponse                       | 400         | application/json |
| apierrors.OAuthErrorResponse                  | 401         | application/json |
| apierrors.OAuthClientManagementRateLimitError | 429         | application/json |
| apierrors.APIError                            | 4XX, 5XX    | \*/\*            |

### Example

```go
package main

import (
	"context"
	"errors"
	pipeshub "github.com/pipeshub-ai/pipeshub-sdk-go"
	"github.com/pipeshub-ai/pipeshub-sdk-go/models/apierrors"
	"github.com/pipeshub-ai/pipeshub-sdk-go/models/components"
	"log"
)

func main() {
	ctx := context.Background()

	s := pipeshub.New()

	res, err := s.OAuthProvider.OauthToken(ctx, components.OAuthTokenRequest{
		GrantType: components.GrantTypeClientCredentials,
	})
	if err != nil {

		var e *apierrors.ErrorResponse
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

| Variable       | Option                                | Default                      | Description     |
| -------------- | ------------------------------------- | ---------------------------- | --------------- |
| `instance_url` | `WithInstanceURL(instanceURL string)` | `"https://app.pipeshub.com"` | Base server URL |

#### Example

```go
package main

import (
	"context"
	pipeshub "github.com/pipeshub-ai/pipeshub-sdk-go"
	"github.com/pipeshub-ai/pipeshub-sdk-go/models/components"
	"log"
)

func main() {
	ctx := context.Background()

	s := pipeshub.New(
		pipeshub.WithServerIndex(0),
		pipeshub.WithInstanceURL("https://app.pipeshub.com"),
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
	pipeshub "github.com/pipeshub-ai/pipeshub-sdk-go"
	"github.com/pipeshub-ai/pipeshub-sdk-go/models/components"
	"log"
)

func main() {
	ctx := context.Background()

	s := pipeshub.New(
		pipeshub.WithServerURL("https://https://app.pipeshub.com"),
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

	"github.com/pipeshub-ai/pipeshub-sdk-go"
)

var (
	httpClient = &http.Client{Timeout: 30 * time.Second}
	sdkClient  = pipeshub.New(pipeshub.WithClient(httpClient))
)
```

This can be a convenient way to configure timeouts, cookies, proxies, custom headers, and other low-level configuration.
<!-- End Custom HTTP Client [http-client] -->

<!-- Placeholder for Future Speakeasy SDK Sections -->
