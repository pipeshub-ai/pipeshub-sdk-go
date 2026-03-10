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

	res, err := s.UserAccount.InitAuth(ctx, components.InitAuthRequest{
		Email: "user@example.com",
	})
	if err != nil {
		log.Fatal(err)
	}
	if res.InitAuthResponse != nil {
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

	res, err := s.UserAccount.InitAuth(ctx, components.InitAuthRequest{
		Email: "user@example.com",
	})
	if err != nil {
		log.Fatal(err)
	}
	if res.InitAuthResponse != nil {
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
	"github.com/pipeshub-ai/pipeshub-sdk-go/models/components"
	"github.com/pipeshub-ai/pipeshub-sdk-go/models/operations"
	"log"
	"os"
)

func main() {
	ctx := context.Background()

	s := pipeshub.New()

	res, err := s.UserAccount.ResetPasswordWithToken(ctx, components.TokenPasswordResetRequest{
		Password: "H9GEHoL829GXj06",
	}, operations.ResetPasswordWithTokenSecurity{
		ScopedToken: os.Getenv("PIPESHUB_SCOPED_TOKEN"),
	})
	if err != nil {
		log.Fatal(err)
	}
	if res.PasswordResetResponse != nil {
		// handle response
	}
}

```
<!-- End Authentication [security] -->

<!-- Start Available Resources and Operations [operations] -->
## Available Resources and Operations

<details open>
<summary>Available methods</summary>

### [AgentConversations](docs/sdks/agentconversations/README.md)

* [List](docs/sdks/agentconversations/README.md#list) - List agent conversations
* [CreateConversation](docs/sdks/agentconversations/README.md#createconversation) - Create agent conversation
* [Stream](docs/sdks/agentconversations/README.md#stream) - Create agent conversation with streaming
* [Get](docs/sdks/agentconversations/README.md#get) - Get agent conversation
* [Delete](docs/sdks/agentconversations/README.md#delete) - Delete agent conversation
* [AddMessage](docs/sdks/agentconversations/README.md#addmessage) - Add message to agent conversation
* [StreamMessage](docs/sdks/agentconversations/README.md#streammessage) - Add message with streaming
* [RegenerateAnswer](docs/sdks/agentconversations/README.md#regenerateanswer) - Regenerate agent response

### [Agents](docs/sdks/agents/README.md)

* [List](docs/sdks/agents/README.md#list) - List agents
* [Create](docs/sdks/agents/README.md#create) - Create agent
* [ListTools](docs/sdks/agents/README.md#listtools) - List available tools
* [Get](docs/sdks/agents/README.md#get) - Get agent
* [Update](docs/sdks/agents/README.md#update) - Update agent
* [Delete](docs/sdks/agents/README.md#delete) - Delete agent
* [GetPermissions](docs/sdks/agents/README.md#getpermissions) - Get agent permissions
* [UpdatePermissions](docs/sdks/agents/README.md#updatepermissions) - Update agent permissions
* [Share](docs/sdks/agents/README.md#share) - Share agent
* [Unshare](docs/sdks/agents/README.md#unshare) - Unshare an agent

### [AgentTemplates](docs/sdks/agenttemplates/README.md)

* [List](docs/sdks/agenttemplates/README.md#list) - List agent templates
* [Create](docs/sdks/agenttemplates/README.md#create) - Create agent template
* [Get](docs/sdks/agenttemplates/README.md#get) - Get agent template
* [Update](docs/sdks/agenttemplates/README.md#update) - Update agent template
* [Delete](docs/sdks/agenttemplates/README.md#delete) - Delete agent template

### [AiModelsProviders](docs/sdks/aimodelsproviders/README.md)

* [GetByType](docs/sdks/aimodelsproviders/README.md#getbytype) - Get models by type
* [GetAvailableModels](docs/sdks/aimodelsproviders/README.md#getavailablemodels) - Get available models for selection
* [Add](docs/sdks/aimodelsproviders/README.md#add) - Add new AI model provider
* [Update](docs/sdks/aimodelsproviders/README.md#update) - Update AI model provider
* [Delete](docs/sdks/aimodelsproviders/README.md#delete) - Delete AI model provider
* [SetDefault](docs/sdks/aimodelsproviders/README.md#setdefault) - Set default AI model

### [AuthConfig](docs/sdks/authconfig/README.md)

* [SetGoogle](docs/sdks/authconfig/README.md#setgoogle) - Configure Google authentication

### [AuthConfigurations](docs/sdks/authconfigurations/README.md)

* [SetOAuth](docs/sdks/authconfigurations/README.md#setoauth) - Configure generic OAuth provider

### [AuthenticationConfiguration](docs/sdks/authenticationconfiguration/README.md)

* [SetAzureAdAuthConfig](docs/sdks/authenticationconfiguration/README.md#setazureadauthconfig) - Configure Azure AD authentication
* [GetAzureAd](docs/sdks/authenticationconfiguration/README.md#getazuread) - Get Azure AD configuration
* [SetMicrosoftAuth](docs/sdks/authenticationconfiguration/README.md#setmicrosoftauth) - Configure Microsoft authentication
* [GetMicrosoft](docs/sdks/authenticationconfiguration/README.md#getmicrosoft) - Get Microsoft authentication configuration
* [GetGoogleAuthConfig](docs/sdks/authenticationconfiguration/README.md#getgoogleauthconfig) - Get Google authentication configuration
* [SetSsoAuthConfig](docs/sdks/authenticationconfiguration/README.md#setssoauthconfig) - Configure SAML SSO authentication
* [GetSSO](docs/sdks/authenticationconfiguration/README.md#getsso) - Get SAML SSO configuration
* [GetGenericOAuth](docs/sdks/authenticationconfiguration/README.md#getgenericoauth) - Get generic OAuth configuration

### [ConfigurationManager](docs/sdks/configurationmanager/README.md)

* [GetSlackBotConfigs](docs/sdks/configurationmanager/README.md#getslackbotconfigs) - Get Slack bot configurations
* [CreateSlackBotConfig](docs/sdks/configurationmanager/README.md#createslackbotconfig) - Create Slack bot configuration
* [UpdateSlackBotConfig](docs/sdks/configurationmanager/README.md#updateslackbotconfig) - Update Slack bot configuration
* [DeleteSlackBotConfig](docs/sdks/configurationmanager/README.md#deleteslackbotconfig) - Delete Slack bot configuration
* [SetMetricsCollectionPushInterval](docs/sdks/configurationmanager/README.md#setmetricscollectionpushinterval) - Set metrics push interval
* [SetMetricsServerURL](docs/sdks/configurationmanager/README.md#setmetricsserverurl) - Set metrics remote server URL
* [GetAIModels](docs/sdks/configurationmanager/README.md#getaimodels) - Get AI models configuration
* [CreateAIModelsConfig](docs/sdks/configurationmanager/README.md#createaimodelsconfig) - Create AI models configuration
* [GetAIModelsProviders](docs/sdks/configurationmanager/README.md#getaimodelsproviders) - Get AI model providers

### [Connector](docs/sdks/connector/README.md)

* [ReindexRecord](docs/sdks/connector/README.md#reindexrecord) - Reindex single record
* [ReindexGroup](docs/sdks/connector/README.md#reindexgroup) - Reindex record group
* [Resync](docs/sdks/connector/README.md#resync) - Resync connector

### [ConnectorConfiguration](docs/sdks/connectorconfiguration/README.md)

* [Get](docs/sdks/connectorconfiguration/README.md#get) - Get connector configuration
* [Update](docs/sdks/connectorconfiguration/README.md#update) - Update connector configuration
* [UpdateAuth](docs/sdks/connectorconfiguration/README.md#updateauth) - Update authentication configuration
* [UpdateFiltersSync](docs/sdks/connectorconfiguration/README.md#updatefilterssync) - Update filters and sync configuration

### [ConnectorControl](docs/sdks/connectorcontrol/README.md)

* [Toggle](docs/sdks/connectorcontrol/README.md#toggle) - Toggle connector sync or agent

### [ConnectorFilters](docs/sdks/connectorfilters/README.md)

* [Get](docs/sdks/connectorfilters/README.md#get) - Get filter options
* [Save](docs/sdks/connectorfilters/README.md#save) - Save filter selections
* [GetFilterOptions](docs/sdks/connectorfilters/README.md#getfilteroptions) - Get dynamic filter options

### [ConnectorInstances](docs/sdks/connectorinstances/README.md)

* [List](docs/sdks/connectorinstances/README.md#list) - List connector instances
* [Create](docs/sdks/connectorinstances/README.md#create) - Create connector instance
* [ListActive](docs/sdks/connectorinstances/README.md#listactive) - List active connector instances
* [ListInactive](docs/sdks/connectorinstances/README.md#listinactive) - List inactive connector instances
* [ListConfigured](docs/sdks/connectorinstances/README.md#listconfigured) - List configured connector instances
* [ListActiveAgents](docs/sdks/connectorinstances/README.md#listactiveagents) - List active agent connectors
* [Get](docs/sdks/connectorinstances/README.md#get) - Get connector instance
* [Delete](docs/sdks/connectorinstances/README.md#delete) - Delete connector instance
* [UpdateName](docs/sdks/connectorinstances/README.md#updatename) - Update connector instance name

### [ConnectorOAuth](docs/sdks/connectoroauth/README.md)

* [Authorize](docs/sdks/connectoroauth/README.md#authorize) - Get OAuth authorization URL
* [HandleCallback](docs/sdks/connectoroauth/README.md#handlecallback) - OAuth callback handler
* [~~ExchangeLegacyToken~~](docs/sdks/connectoroauth/README.md#exchangelegacytoken) - Exchange Google authorization code for tokens :warning: **Deprecated**

### [ConnectorRegistry](docs/sdks/connectorregistry/README.md)

* [List](docs/sdks/connectorregistry/README.md#list) - List available connector types
* [GetSchemaForType](docs/sdks/connectorregistry/README.md#getschemafortype) - Get connector configuration schema

### [Connectors](docs/sdks/connectors/README.md)

* [GetStats](docs/sdks/connectors/README.md#getstats) - Get connector statistics

### [Conversations](docs/sdks/conversations/README.md)

* [Create](docs/sdks/conversations/README.md#create) - Create a new AI conversation
* [Stream](docs/sdks/conversations/README.md#stream) - Create conversation with streaming response
* [List](docs/sdks/conversations/README.md#list) - List all conversations
* [ListArchives](docs/sdks/conversations/README.md#listarchives) - List archived conversations
* [Get](docs/sdks/conversations/README.md#get) - Get conversation by ID
* [Delete](docs/sdks/conversations/README.md#delete) - Delete conversation
* [AddMessage](docs/sdks/conversations/README.md#addmessage) - Add message to conversation
* [AddMessageStream](docs/sdks/conversations/README.md#addmessagestream) - Add message with streaming response
* [Share](docs/sdks/conversations/README.md#share) - Share conversation with users
* [UpdateTitle](docs/sdks/conversations/README.md#updatetitle) - Update conversation title
* [Archive](docs/sdks/conversations/README.md#archive) - Archive conversation
* [Unarchive](docs/sdks/conversations/README.md#unarchive) - Unarchive conversation
* [Regenerate](docs/sdks/conversations/README.md#regenerate) - Regenerate AI response
* [SubmitFeedback](docs/sdks/conversations/README.md#submitfeedback) - Submit feedback on AI response
* [Unshare](docs/sdks/conversations/README.md#unshare) - Unshare a conversation

### [CrawlingJobs](docs/sdks/crawlingjobs/README.md)

* [Schedule](docs/sdks/crawlingjobs/README.md#schedule) - Schedule a crawling job
* [GetStatus](docs/sdks/crawlingjobs/README.md#getstatus) - Get crawling job status
* [Remove](docs/sdks/crawlingjobs/README.md#remove) - Remove a crawling job
* [ListAllStatuses](docs/sdks/crawlingjobs/README.md#listallstatuses) - Get all crawling job statuses
* [RemoveAll](docs/sdks/crawlingjobs/README.md#removeall) - Remove all crawling jobs
* [Pause](docs/sdks/crawlingjobs/README.md#pause) - Pause a crawling job
* [Resume](docs/sdks/crawlingjobs/README.md#resume) - Resume a crawling job
* [GetQueueStats](docs/sdks/crawlingjobs/README.md#getqueuestats) - Get queue statistics

### [DocumentManagement](docs/sdks/documentmanagement/README.md)

* [Download](docs/sdks/documentmanagement/README.md#download) - Download document

### [Folders](docs/sdks/folders/README.md)

* [CreateRoot](docs/sdks/folders/README.md#createroot) - Create root folder
* [GetContents](docs/sdks/folders/README.md#getcontents) - Get folder contents
* [Update](docs/sdks/folders/README.md#update) - Update folder
* [Delete](docs/sdks/folders/README.md#delete) - Delete folder
* [GetChildren](docs/sdks/folders/README.md#getchildren) - Get folder children (alias for folder contents)
* [Create](docs/sdks/folders/README.md#create) - Create subfolder

### [KnowledgeBases](docs/sdks/knowledgebases/README.md)

* [Create](docs/sdks/knowledgebases/README.md#create) - Create a new knowledge base
* [List](docs/sdks/knowledgebases/README.md#list) - List all knowledge bases
* [Get](docs/sdks/knowledgebases/README.md#get) - Get knowledge base by ID
* [Update](docs/sdks/knowledgebases/README.md#update) - Update knowledge base
* [Delete](docs/sdks/knowledgebases/README.md#delete) - Delete knowledge base
* [ReindexFailedRecords](docs/sdks/knowledgebases/README.md#reindexfailedrecords) - Reindex failed records for connector
* [MoveRecord](docs/sdks/knowledgebases/README.md#moverecord) - Move record to another location
* [GetRootNodes](docs/sdks/knowledgebases/README.md#getrootnodes) - Get knowledge hub root nodes
* [GetChildNodes](docs/sdks/knowledgebases/README.md#getchildnodes) - Get knowledge hub child nodes

### [Mcp](docs/sdks/mcp/README.md)

* [HandleRequest](docs/sdks/mcp/README.md#handlerequest) - Handle MCP JSON-RPC request
* [Stream](docs/sdks/mcp/README.md#stream) - MCP SSE streaming endpoint

### [MetricsCollection](docs/sdks/metricscollection/README.md)

* [Get](docs/sdks/metricscollection/README.md#get) - Get metrics collection configuration
* [Toggle](docs/sdks/metricscollection/README.md#toggle) - Enable or disable metrics collection

### [Oauth](docs/sdks/oauth/README.md)

* [ExchangeCode](docs/sdks/oauth/README.md#exchangecode) - Exchange OAuth authorization code for tokens

### [OauthApps](docs/sdks/oauthapps/README.md)

* [List](docs/sdks/oauthapps/README.md#list) - List OAuth apps
* [Create](docs/sdks/oauthapps/README.md#create) - Create OAuth app
* [ListScopes](docs/sdks/oauthapps/README.md#listscopes) - List available scopes
* [Get](docs/sdks/oauthapps/README.md#get) - Get OAuth app details
* [Update](docs/sdks/oauthapps/README.md#update) - Update OAuth app
* [Delete](docs/sdks/oauthapps/README.md#delete) - Delete OAuth app
* [RegenerateSecret](docs/sdks/oauthapps/README.md#regeneratesecret) - Regenerate client secret
* [Suspend](docs/sdks/oauthapps/README.md#suspend) - Suspend OAuth app
* [Activate](docs/sdks/oauthapps/README.md#activate) - Activate suspended OAuth app
* [ListTokens](docs/sdks/oauthapps/README.md#listtokens) - List app tokens
* [RevokeAllTokens](docs/sdks/oauthapps/README.md#revokealltokens) - Revoke all app tokens

### [OauthConfigs](docs/sdks/oauthconfigs/README.md)

* [ListAll](docs/sdks/oauthconfigs/README.md#listall) - List OAuth configurations
* [UpdateByID](docs/sdks/oauthconfigs/README.md#updatebyid) - Update OAuth configuration

### [OauthConfiguration](docs/sdks/oauthconfiguration/README.md)

* [ListToolsetConfigs](docs/sdks/oauthconfiguration/README.md#listtoolsetconfigs) - List OAuth configs by toolset type
* [DeleteConfig](docs/sdks/oauthconfiguration/README.md#deleteconfig) - Delete OAuth config
* [List](docs/sdks/oauthconfiguration/README.md#list) - List OAuth-capable connector types
* [GetConnectorTypeDetails](docs/sdks/oauthconfiguration/README.md#getconnectortypedetails) - Get OAuth connector type details
* [ListConfigsByType](docs/sdks/oauthconfiguration/README.md#listconfigsbytype) - List OAuth configs for connector type
* [Create](docs/sdks/oauthconfiguration/README.md#create) - Create OAuth configuration
* [Get](docs/sdks/oauthconfiguration/README.md#get) - Get OAuth configuration

### [OauthConfigurations](docs/sdks/oauthconfigurations/README.md)

* [Update](docs/sdks/oauthconfigurations/README.md#update) - Update OAuth config
* [Delete](docs/sdks/oauthconfigurations/README.md#delete) - Delete OAuth configuration

### [OauthProvider](docs/sdks/oauthprovider/README.md)

* [Authorize](docs/sdks/oauthprovider/README.md#authorize) - Initiate OAuth authorization flow
* [AuthorizeConsent](docs/sdks/oauthprovider/README.md#authorizeconsent) - Submit authorization consent
* [ExchangeToken](docs/sdks/oauthprovider/README.md#exchangetoken) - Exchange authorization code for tokens
* [RevokeToken](docs/sdks/oauthprovider/README.md#revoketoken) - Revoke an access or refresh token
* [IntrospectToken](docs/sdks/oauthprovider/README.md#introspecttoken) - Introspect a token

### [OpenIDConnect](docs/sdks/openidconnect/README.md)

* [UserInfo](docs/sdks/openidconnect/README.md#userinfo) - Get authenticated user information
* [OauthAuthorizationServerMetadata](docs/sdks/openidconnect/README.md#oauthauthorizationservermetadata) - OAuth 2.0 Authorization Server Metadata
* [Jwks](docs/sdks/openidconnect/README.md#jwks) - JSON Web Key Set
* [GetProtectedResourceMetadata](docs/sdks/openidconnect/README.md#getprotectedresourcemetadata) - OAuth Protected Resource Metadata
* [GetConfiguration](docs/sdks/openidconnect/README.md#getconfiguration) - OpenID Connect Discovery

### [OrganizationAuthConfig](docs/sdks/organizationauthconfig/README.md)

* [GetAuthMethods](docs/sdks/organizationauthconfig/README.md#getauthmethods) - Get organization authentication methods
* [UpdateAuthMethod](docs/sdks/organizationauthconfig/README.md#updateauthmethod) - Update organization authentication methods
* [SetUp](docs/sdks/organizationauthconfig/README.md#setup) - Set up auth configuration

### [Organizations](docs/sdks/organizations/README.md)

* [CheckExists](docs/sdks/organizations/README.md#checkexists) - Check if organization exists
* [Create](docs/sdks/organizations/README.md#create) - Create organization
* [GetCurrent](docs/sdks/organizations/README.md#getcurrent) - Get current organization
* [Update](docs/sdks/organizations/README.md#update) - Update organization
* [Delete](docs/sdks/organizations/README.md#delete) - Delete organization
* [UploadLogo](docs/sdks/organizations/README.md#uploadlogo) - Upload organization logo
* [GetLogo](docs/sdks/organizations/README.md#getlogo) - Get organization logo
* [DeleteLogo](docs/sdks/organizations/README.md#deletelogo) - Delete organization logo
* [GetOnboardingStatus](docs/sdks/organizations/README.md#getonboardingstatus) - Get onboarding status
* [UpdateOnboardingStatus](docs/sdks/organizations/README.md#updateonboardingstatus) - Update onboarding status

### [Permissions](docs/sdks/permissions/README.md)

* [GrantKBAccess](docs/sdks/permissions/README.md#grantkbaccess) - Grant permissions
* [List](docs/sdks/permissions/README.md#list) - List permissions
* [Update](docs/sdks/permissions/README.md#update) - Update permissions
* [DeleteFromKB](docs/sdks/permissions/README.md#deletefromkb) - Remove permissions

### [PlatformSettings](docs/sdks/platformsettings/README.md)

* [Update](docs/sdks/platformsettings/README.md#update) - Update platform settings
* [Get](docs/sdks/platformsettings/README.md#get) - Get platform settings
* [ListFeatureFlags](docs/sdks/platformsettings/README.md#listfeatureflags) - Get available feature flags
* [SetCustomSystemPrompt](docs/sdks/platformsettings/README.md#setcustomsystemprompt) - Update custom system prompt
* [GetCustomSystemPrompt](docs/sdks/platformsettings/README.md#getcustomsystemprompt) - Get custom system prompt

### [PublicUrls](docs/sdks/publicurls1/README.md)

* [SetFrontend](docs/sdks/publicurls1/README.md#setfrontend) - Set frontend public URL
* [Set](docs/sdks/publicurls1/README.md#set) - Set connector public URL
* [GetConnector](docs/sdks/publicurls1/README.md#getconnector) - Get connector public URL

### [PublicURLs](docs/sdks/publicurls2/README.md)

* [GetFrontend](docs/sdks/publicurls2/README.md#getfrontend) - Get frontend public URL

### [Records](docs/sdks/records/README.md)

* [GetAll](docs/sdks/records/README.md#getall) - Get all records across knowledge bases
* [GetByKB](docs/sdks/records/README.md#getbykb) - Get records for a knowledge base
* [GetChildren](docs/sdks/records/README.md#getchildren) - Get KB children (alias for records)
* [GetByID](docs/sdks/records/README.md#getbyid) - Get record by ID
* [Update](docs/sdks/records/README.md#update) - Update record
* [Delete](docs/sdks/records/README.md#delete) - Delete record
* [StreamContent](docs/sdks/records/README.md#streamcontent) - Stream record content

### [Saml](docs/sdks/saml/README.md)

* [SignIn](docs/sdks/saml/README.md#signin) - Initiate SAML sign-in flow
* [SignInCallback](docs/sdks/saml/README.md#signincallback) - SAML sign-in callback

### [SemanticSearch](docs/sdks/semanticsearch/README.md)

* [Execute](docs/sdks/semanticsearch/README.md#execute) - Perform semantic search
* [GetHistory](docs/sdks/semanticsearch/README.md#gethistory) - Get search history
* [DeleteAllHistory](docs/sdks/semanticsearch/README.md#deleteallhistory) - Clear all search history
* [GetByID](docs/sdks/semanticsearch/README.md#getbyid) - Get search by ID
* [Delete](docs/sdks/semanticsearch/README.md#delete) - Delete search by ID
* [Share](docs/sdks/semanticsearch/README.md#share) - Share a search
* [Unshare](docs/sdks/semanticsearch/README.md#unshare) - Unshare a search
* [Archive](docs/sdks/semanticsearch/README.md#archive) - Archive a search
* [Unarchive](docs/sdks/semanticsearch/README.md#unarchive) - Unarchive a search

### [SmtpConfiguration](docs/sdks/smtpconfiguration/README.md)

* [CreateOrUpdate](docs/sdks/smtpconfiguration/README.md#createorupdate) - Create or update SMTP configuration
* [Get](docs/sdks/smtpconfiguration/README.md#get) - Get SMTP configuration

### [StorageConfiguration](docs/sdks/storageconfiguration/README.md)

* [Get](docs/sdks/storageconfiguration/README.md#get) - Get current storage configuration

### [Teams](docs/sdks/teams/README.md)

* [Create](docs/sdks/teams/README.md#create) - Create a team
* [List](docs/sdks/teams/README.md#list) - List teams
* [GetByID](docs/sdks/teams/README.md#getbyid) - Get team by ID
* [Update](docs/sdks/teams/README.md#update) - Update team
* [Delete](docs/sdks/teams/README.md#delete) - Delete team
* [ListUserTeams](docs/sdks/teams/README.md#listuserteams) - Get current user's teams
* [GetUsers](docs/sdks/teams/README.md#getusers) - Get users in team
* [AddUsers](docs/sdks/teams/README.md#addusers) - Add users to team
* [Remove](docs/sdks/teams/README.md#remove) - Remove user from team
* [UpdateUserPermissions](docs/sdks/teams/README.md#updateuserpermissions) - Update team users permissions
* [GetUserCreated](docs/sdks/teams/README.md#getusercreated) - Get user created teams

### [ToolsetConfiguration](docs/sdks/toolsetconfiguration/README.md)

* [GetConfig](docs/sdks/toolsetconfiguration/README.md#getconfig) - Get toolset configuration
* [Update](docs/sdks/toolsetconfiguration/README.md#update) - Update toolset configuration
* [Delete](docs/sdks/toolsetconfiguration/README.md#delete) - Delete toolset configuration

### [~~ToolsetConfigurations~~](docs/sdks/toolsetconfigurations/README.md)

* [~~Save~~](docs/sdks/toolsetconfigurations/README.md#save) - Save toolset configuration :warning: **Deprecated**

### [ToolsetInstances](docs/sdks/toolsetinstances/README.md)

* [ListConfigured](docs/sdks/toolsetinstances/README.md#listconfigured) - List configured toolsets
* [CheckStatus](docs/sdks/toolsetinstances/README.md#checkstatus) - Check toolset status
* [ListMine](docs/sdks/toolsetinstances/README.md#listmine) - List my toolsets with auth status
* [List](docs/sdks/toolsetinstances/README.md#list) - List toolset instances
* [Create](docs/sdks/toolsetinstances/README.md#create) - Create toolset instance
* [Get](docs/sdks/toolsetinstances/README.md#get) - Get toolset instance
* [Update](docs/sdks/toolsetinstances/README.md#update) - Update toolset instance
* [Delete](docs/sdks/toolsetinstances/README.md#delete) - Delete toolset instance
* [Authenticate](docs/sdks/toolsetinstances/README.md#authenticate) - Authenticate toolset instance
* [DeleteCredentials](docs/sdks/toolsetinstances/README.md#deletecredentials) - Remove toolset credentials
* [Reauthenticate](docs/sdks/toolsetinstances/README.md#reauthenticate) - Mark instance for reauthentication
* [GetStatus](docs/sdks/toolsetinstances/README.md#getstatus) - Get instance authentication status

### [ToolsetOAuth](docs/sdks/toolsetoauth/README.md)

* [Authorize](docs/sdks/toolsetoauth/README.md#authorize) - Get OAuth authorization URL
* [Callback](docs/sdks/toolsetoauth/README.md#callback) - Handle OAuth callback
* [GetAuthURL](docs/sdks/toolsetoauth/README.md#getauthurl) - Get OAuth authorization URL for instance

### [ToolsetRegistry](docs/sdks/toolsetregistry/README.md)

* [List](docs/sdks/toolsetregistry/README.md#list) - List available toolsets
* [GetSchema](docs/sdks/toolsetregistry/README.md#getschema) - Get toolset schema

### [Toolsets](docs/sdks/toolsets/README.md)

* [CreateInstance](docs/sdks/toolsets/README.md#createinstance) - Create toolset instance
* [ReauthenticateByID](docs/sdks/toolsets/README.md#reauthenticatebyid) - Reauthenticate toolset

### [Upload](docs/sdks/upload/README.md)

* [Files](docs/sdks/upload/README.md#files) - Upload files to knowledge base
* [ToFolder](docs/sdks/upload/README.md#tofolder) - Upload files to folder
* [GetLimits](docs/sdks/upload/README.md#getlimits) - Get upload limits

### [UserAccount](docs/sdks/useraccount/README.md)

* [InitAuth](docs/sdks/useraccount/README.md#initauth) - Initialize authentication session
* [Authenticate](docs/sdks/useraccount/README.md#authenticate) - Authenticate user with credentials
* [GenerateLoginOtp](docs/sdks/useraccount/README.md#generateloginotp) - Generate and send OTP for login
* [RequestPasswordReset](docs/sdks/useraccount/README.md#requestpasswordreset) - Request password reset email
* [ResetPasswordWithToken](docs/sdks/useraccount/README.md#resetpasswordwithtoken) - Reset password with email token
* [RefreshToken](docs/sdks/useraccount/README.md#refreshtoken) - Refresh access token
* [Logout](docs/sdks/useraccount/README.md#logout) - Logout current session
* [ResetPassword](docs/sdks/useraccount/README.md#resetpassword) - Reset password

### [UserGroups](docs/sdks/usergroups/README.md)

* [Create](docs/sdks/usergroups/README.md#create) - Create user group
* [GetAll](docs/sdks/usergroups/README.md#getall) - Get all user groups
* [GetByID](docs/sdks/usergroups/README.md#getbyid) - Get user group by ID
* [Update](docs/sdks/usergroups/README.md#update) - Update user group
* [Delete](docs/sdks/usergroups/README.md#delete) - Delete user group
* [AddUsers](docs/sdks/usergroups/README.md#addusers) - Add users to group
* [RemoveUsers](docs/sdks/usergroups/README.md#removeusers) - Remove users from group
* [GetUserGroups](docs/sdks/usergroups/README.md#getusergroups) - Get groups for a user
* [GetInGroup](docs/sdks/usergroups/README.md#getingroup) - Get users in group
* [GetStatistics](docs/sdks/usergroups/README.md#getstatistics) - Get group statistics

### [Users](docs/sdks/users/README.md)

* [GetAll](docs/sdks/users/README.md#getall) - Get all users
* [Create](docs/sdks/users/README.md#create) - Create a new user
* [Get](docs/sdks/users/README.md#get) - Get user by ID
* [Update](docs/sdks/users/README.md#update) - Update user
* [Delete](docs/sdks/users/README.md#delete) - Delete user
* [GetEmail](docs/sdks/users/README.md#getemail) - Get user email by ID
* [UpdateEmail](docs/sdks/users/README.md#updateemail) - Update user email
* [UploadDisplayPicture](docs/sdks/users/README.md#uploaddisplaypicture) - Upload display picture
* [GetDisplayPicture](docs/sdks/users/README.md#getdisplaypicture) - Get display picture
* [RemoveDisplayPicture](docs/sdks/users/README.md#removedisplaypicture) - Remove display picture
* [BulkInvite](docs/sdks/users/README.md#bulkinvite) - Bulk invite users
* [ResendInvite](docs/sdks/users/README.md#resendinvite) - Resend user invite
* [ListGraph](docs/sdks/users/README.md#listgraph) - List users (paginated with graph data)
* [Unblock](docs/sdks/users/README.md#unblock) - Unblock a user in organization
* [GetAllWithGroups](docs/sdks/users/README.md#getallwithgroups) - Get all users with groups
* [GetByIds](docs/sdks/users/README.md#getbyids) - Get users by IDs
* [UpdateFullName](docs/sdks/users/README.md#updatefullname) - Update user full name
* [UpdateFirstName](docs/sdks/users/README.md#updatefirstname) - Update user first name
* [UpdateLastName](docs/sdks/users/README.md#updatelastname) - Update user last name
* [UpdateDesignation](docs/sdks/users/README.md#updatedesignation) - Update user designation
* [CheckAdmin](docs/sdks/users/README.md#checkadmin) - Check if user is admin
* [ListTeams](docs/sdks/users/README.md#listteams) - Get user teams

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

	res, err := s.Conversations.Stream(ctx, components.CreateConversationRequest{
		Query: "What are the key findings from our Q4 financial report?",
		RecordIds: []string{
			"507f1f77bcf86cd799439011",
			"507f1f77bcf86cd799439012",
		},
		ModelKey:  pipeshub.Pointer("gpt-4-turbo"),
		ModelName: pipeshub.Pointer("GPT-4 Turbo"),
		ChatMode:  pipeshub.Pointer("balanced"),
	})
	if err != nil {
		log.Fatal(err)
	}
	if res.SSEEvent != nil {
		defer res.SSEEvent.Close()

		for res.SSEEvent.Next() {
			event := res.SSEEvent.Value()
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

	res, err := s.UserAccount.InitAuth(ctx, components.InitAuthRequest{
		Email: "user@example.com",
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
	if res.InitAuthResponse != nil {
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

	res, err := s.UserAccount.InitAuth(ctx, components.InitAuthRequest{
		Email: "user@example.com",
	})
	if err != nil {
		log.Fatal(err)
	}
	if res.InitAuthResponse != nil {
		// handle response
	}
}

```
<!-- End Retries [retries] -->

<!-- Start Error Handling [errors] -->
## Error Handling

Handling errors in this SDK should largely match your expectations. All operations return a response object or an error, they will never return both.

By Default, an API error will return `apierrors.APIError`. When custom error responses are specified for an operation, the SDK may also return their associated error. You can refer to respective *Errors* tables in SDK docs for more details on possible error types for each operation.

For example, the `InitAuth` function may return the following errors:

| Error Type          | Status Code   | Content Type     |
| ------------------- | ------------- | ---------------- |
| apierrors.AuthError | 400, 403, 404 | application/json |
| apierrors.APIError  | 4XX, 5XX      | \*/\*            |

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

	res, err := s.UserAccount.InitAuth(ctx, components.InitAuthRequest{
		Email: "user@example.com",
	})
	if err != nil {

		var e *apierrors.AuthError
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

	res, err := s.UserAccount.InitAuth(ctx, components.InitAuthRequest{
		Email: "user@example.com",
	})
	if err != nil {
		log.Fatal(err)
	}
	if res.InitAuthResponse != nil {
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

	res, err := s.UserAccount.InitAuth(ctx, components.InitAuthRequest{
		Email: "user@example.com",
	})
	if err != nil {
		log.Fatal(err)
	}
	if res.InitAuthResponse != nil {
		// handle response
	}
}

```

### Override Server URL Per-Operation

The server URL can also be overridden on a per-operation basis, provided a server list was specified for the operation. For example:
```go
package main

import (
	"context"
	pipeshub "github.com/pipeshub-ai/pipeshub-sdk-go"
	"github.com/pipeshub-ai/pipeshub-sdk-go/models/operations"
	"log"
)

func main() {
	ctx := context.Background()

	s := pipeshub.New()

	res, err := s.OpenIDConnect.OauthAuthorizationServerMetadata(ctx, operations.WithServerURL(""))
	if err != nil {
		log.Fatal(err)
	}
	if res.OpenIDConfiguration != nil {
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
