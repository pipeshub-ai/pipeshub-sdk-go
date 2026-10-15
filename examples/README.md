# PipesHub Go SDK examples

69 runnable Go programs that authenticate against a PipesHub instance and exercise semantic search, conversations, the agent loop, agent CRUD, personal access tokens, and knowledge base management.

The Go module is named `enterprise_search` (which is why imports read `enterprise_search/auth`), but it lives in this `examples/` directory. Each program is its own package in its own directory, so `go build ./...` and `go vet ./...` cover the whole tree.

## Prerequisites

- Go 1.26.8+ (matching the SDK's `go` directive)
- A running PipesHub instance (default: `http://localhost:3000`)
- A user account on that instance
- At least one knowledge base or connector with indexed content
- An AI model provider configured on the instance — conversation calls return `500 — Failed to get AI response` without one

## Setup

```bash
cp .env.example .env    # then fill it in
go mod tidy
```

Every program takes the `.env` path as its first argument.

### Authentication

The programs added in SDK 1.7.1 (`KB/kb/get_upload_limits`, `KB/records/get_record_by_id`, `KB/records/reindex_record_group`, `KB/trash/*`, and the non-streaming `conversation/*` programs) authenticate with **OAuth2 client credentials** through `auth.NewOAuthClient`. Create an OAuth app with the `client_credentials` grant and the scopes the examples use, then set `PIPESHUB_CLIENT_ID` and `PIPESHUB_CLIENT_SECRET`. The SDK fetches the access token from `PIPESHUB_TOKEN_URL` or, when that is unset, `${PIPESHUB_BASE_URL}/api/v1/oauth2/token`, and fetches a new one when it expires. The SDK's built-in default token URL is relative, so the helper always passes an absolute one.

```go
// What auth.NewOAuthClient passes to pipeshub.WithSecurity
components.Security{Oauth2: &components.SchemeOauth2{
	ClientID:     os.Getenv("PIPESHUB_CLIENT_ID"),
	ClientSecret: os.Getenv("PIPESHUB_CLIENT_SECRET"),
	TokenURL:     "http://localhost:3000/api/v1/oauth2/token",
}}
```

The `personal_access_tokens/*` programs use `auth.NewSessionClient`, which logs in with `PIPESHUB_TEST_USER_EMAIL` and `PIPESHUB_TEST_USER_PASSWORD`: the personal access token routes require an interactive user session and reject OAuth tokens with `403`.

The older programs use `auth.NewClient`: `PIPESHUB_BEARER_AUTH` when it is set, otherwise the same email and password login.


| Variable | Needed by |
| --- | --- |
| `PIPESHUB_BASE_URL` | all — the examples append `/api/v1` themselves |
| `PIPESHUB_CLIENT_ID`, `PIPESHUB_CLIENT_SECRET` | the 1.7.1 programs (see [Authentication](#authentication)); OAuth2 client credentials |
| `PIPESHUB_TOKEN_URL` | optional; OAuth2 token endpoint (default `${PIPESHUB_BASE_URL}/api/v1/oauth2/token`) |
| `PIPESHUB_TEST_USER_EMAIL`, `PIPESHUB_TEST_USER_PASSWORD` | `personal_access_tokens/*`; the older programs unless `PIPESHUB_BEARER_AUTH` is set |
| `PIPESHUB_BEARER_AUTH` | optional, older programs only; a JWT access token used instead of the email and password login |
| `PIPESHUB_AGENT_KEY` | all `agent_conversation/*` except `create_and_add_message` |
| `KB_ID`, `CONNECTOR_ID` | optional knowledge filters for `agent_conversation/*` |
| `PIPESHUB_KB_NAME`, `PIPESHUB_CONNECTOR_NAME` | optional lookup-name overrides (default `SDK-test`, `ABC News RSS` / `abc news`). `KB/records/reindex_record_group` falls back to the first connector |
| `PIPESHUB_AGENT_MODEL_KEY` | optional; used when an example creates an agent |
| `PIPESHUB_UPLOAD_PATH` | optional; file for `KB/*/upload_records`, `KB/records/get_record_by_id` and `KB/trash/*` (falls back to a built-in sample) |
| `PIPESHUB_GOOGLE_SERVICE_ACCOUNT_FILE`, `PIPESHUB_GOOGLE_ADMIN_EMAIL`, `PIPESHUB_ATLASSIAN_*`, `PIPESHUB_S3_*` | the Gmail, Drive, Confluence, Jira and S3 programs; see `.env.example`. Gmail, Drive, Confluence, Jira and Slack keep the source's permissions, so search only finds their records for PipesHub users whose email has access in the source; the examples report that instead of failing. |
| `PIPESHUB_AZURE_FILES_CONNECTION_STRING` | `connectors/azure_files`; a storage account connection string. `PIPESHUB_AZURE_FILES_SHARE` and `PIPESHUB_AZURE_FILES_SEARCH_QUERY` are optional |
| `PIPESHUB_SLACK_BOT_TOKEN` | `connectors/slack_workspace`; a Slack bot token (`xoxb-…`). `PIPESHUB_SLACK_CHANNEL_ID` and `PIPESHUB_SLACK_SEARCH_QUERY` are optional |
| `PIPESHUB_RSS_FEED_URL` | optional; the feed `connectors/rss` syncs (default Hacker News, which currently refuses the RSS connector with HTTP 419, so its connection test fails) |

## Run

Suggested order — each step only needs what the ones before it produced.

| Program | Demonstrates | Extra env | Needs AI backend |
| --- | --- | --- | --- |
| `./st_examples/search` | `SemanticSearch.Search` across every KB | — | no |
| `./conversation/internal_search` | `StreamChat`, AG-UI decoded **longhand** | — | yes |
| `./st_examples/chat` | `StreamChat` + `AddMessageStream` (two turns) | — | yes |
| `./conversation/web_search` | `chatMode: web_search` | — | yes |
| `./conversation/agent` | `chatMode: agent`, `Tools`, `AgentCapabilities` | — | yes |
| `./conversation/agui_trace` | every AG-UI frame, by category | — | yes |
| `./agent/create` | `CreateAgent` — prints an agent key | — | no |
| `./agent/list_and_get` | `ListAgents`, `GetAgent` | — | no |
| `./agent/update` | `UpdateAgent` | *(agent key as arg 2)* | no |
| `./agent/delete` | `DeleteAgent` | *(agent key as arg 2)* | no |
| `./KB/kb/create_knowledge_base` | `CreateKnowledgeBase` | — | no |
| `./KB/kb/list_knowledge_bases` | `ListKnowledgeBases` | — | no |
| `./KB/kb/get_knowledge_base` | `GetKnowledgeBase` | — | no |
| `./KB/kb/update_knowledge_base` | `UpdateKnowledgeBase` | — | no |
| `./KB/kb/delete_knowledge_base` | `DeleteKnowledgeBase` | — | no |
| `./KB/kb/upload_records` | `UploadRecords` SSE into a KB root | `PIPESHUB_UPLOAD_PATH` (optional) | no |
| `./KB/folder/create_folder` | `CreateFolder` | — | no |
| `./KB/folder/update_folder` | `UpdateFolder` | — | no |
| `./KB/folder/delete_folder` | `DeleteFolder` | — | no |
| `./KB/folder/upload_records` | `UploadRecords` into a folder | `PIPESHUB_UPLOAD_PATH` (optional) | no |
| `./KB/records/update_record` | `UpdateRecord` | — | no |
| `./KB/records/move_record` | `MoveRecord` between folders | — | no |
| `./KB/records/delete_record` | `DeleteRecord` | — | no |
| `./KB/records/reindex_record` | `ReindexRecord` | — | no |
| `./KB/records/stream_record_buffer` | `StreamRecordBuffer` | — | no |
| `./agent_conversation/create_and_add_message` | creates its own agent, then streams two turns | — | yes |
| `./agent_conversation/regenerate_message_stream` | `RegenerateAgentConversationMessage` | `PIPESHUB_AGENT_KEY` | yes |
| `./agent_conversation/add_message_feedback` | `UpdateAgentConversationMessageFeedback` | `PIPESHUB_AGENT_KEY` | yes |
| `./agent_conversation/archive_unarchive` | archive / unarchive | `PIPESHUB_AGENT_KEY` | yes |
| `./agent_conversation/update_conversation_title` | `UpdateAgentConversationTitle` | `PIPESHUB_AGENT_KEY` | yes |
| `./agent_conversation/get_conversation_by_id` | `GetAgentConversationByID` | `PIPESHUB_AGENT_KEY` | yes |
| `./agent_conversation/get_all_conversations` | paginated `ListAgentConversations` | `PIPESHUB_AGENT_KEY` | no |
| `./agent_conversation/list_all_archived_conversations` | `ListAgentConversationArchives` | `PIPESHUB_AGENT_KEY` | no |
| `./agent_conversation/list_archives_grouped` | `ListAgentArchivedConversationsGrouped` | `PIPESHUB_AGENT_KEY` | yes |
| `./conversation/knowledgebase` | `StreamChat` scoped by `Filters.Kb` | a KB named `SDK-test` | yes |
| `./conversation/connector` | `StreamChat` scoped by `Filters.Apps` | a connector named `ABC News RSS` | yes |
| `./semantic_search/knowledgebase` | `Search` scoped by `Filters.Kb` | a KB named `SDK-test` | no |
| `./semantic_search/connector` | `Search` scoped by `Filters.Apps` | a connector named `abc news` | no |
| `./conversation/create_conversation` | `CreateConversation` (non-streaming) | — | yes |
| `./conversation/add_message` | `AddMessage` (non-streaming follow-up) | — | yes |
| `./conversation/get_all_conversations` | paginated `GetAllConversations` | — | yes |
| `./conversation/get_conversation_by_id` | `GetConversationByID` | — | yes |
| `./conversation/update_conversation_title` | `UpdateConversationTitle` | — | yes |
| `./conversation/archive_unarchive` | `ArchiveConversation`, `UnarchiveConversation` | — | yes |
| `./conversation/get_archived_conversations` | `GetArchivedConversations` | — | yes |
| `./conversation/search_archived_conversations` | `SearchArchivedConversations` | — | yes |
| `./conversation/update_message_feedback` | `UpdateMessageFeedback` | — | yes |
| `./conversation/delete_conversation` | `DeleteConversationByID` | — | yes |
| `./conversation/cancel_conversation_stream` | `StreamChat` with a `runId`, then `CancelConversationStream` | — | yes |
| `./KB/kb/get_upload_limits` | `GetUploadLimits` | — | no |
| `./KB/records/get_record_by_id` | `GetRecordByID` | `PIPESHUB_UPLOAD_PATH` (optional) | no |
| `./KB/records/reindex_record_group` | knowledge-hub lookup, then `ReindexRecordGroup` on a connector record group | a connector, `PIPESHUB_CONNECTOR_NAME` (optional) | no |
| `./KB/trash/list_trash` | `ListTrash` | trash enabled (see below) | no |
| `./KB/trash/restore_record` | `RestoreRecord` | trash enabled (see below) | no |
| `./KB/trash/restore_records` | `RestoreRecords` (bulk) | trash enabled (see below) | no |
| `./personal_access_tokens/list_scopes` | `ListPersonalAccessTokenScopes` | email + password | no |
| `./personal_access_tokens/create_token` | `CreatePersonalAccessToken` (then revokes it) | email + password | no |
| `./personal_access_tokens/list_tokens` | `ListPersonalAccessTokens` | email + password | no |
| `./personal_access_tokens/revoke_token` | `RevokePersonalAccessToken` | email + password | no |
| `./personal_access_tokens/admin_list_tokens` | `AdminListPersonalAccessTokens` | email + password, org admin | no |
| `./personal_access_tokens/admin_revoke_token` | `AdminRevokePersonalAccessToken` | email + password, org admin | no |
| `./connectors/gmail_workspace` | Gmail Workspace: a Google service account with domain-wide delegation (`PIPESHUB_GOOGLE_SERVICE_ACCOUNT_FILE`, `PIPESHUB_GOOGLE_ADMIN_EMAIL`), mail received in the last 7 days; syncs, waits until indexing finishes, searches it, then turns the sync off and deletes it | see `.env.example` | no |
| `./connectors/drive_workspace` | Drive Workspace: the same Google service account, files modified in the last 30 days, one drive when `PIPESHUB_GOOGLE_DRIVE_ID` is set or the connector lists any; syncs, waits until indexing finishes, searches it, then turns the sync off and deletes it | see `.env.example` | no |
| `./connectors/confluence` | Confluence: an Atlassian API token (`PIPESHUB_ATLASSIAN_BASE_URL`, `PIPESHUB_ATLASSIAN_EMAIL`, `PIPESHUB_ATLASSIAN_API_TOKEN`), one space (`PIPESHUB_CONFLUENCE_SPACE_KEY`, or the first); syncs, waits until indexing finishes, searches it, then turns the sync off and deletes it | see `.env.example` | no |
| `./connectors/jira` | Jira: the same Atlassian API token, one project (`PIPESHUB_JIRA_PROJECT_KEY`, or the first) and issues modified in the last 30 days; syncs, waits until indexing finishes, searches it, then turns the sync off and deletes it | see `.env.example` | no |
| `./connectors/s3` | S3: an access key (`PIPESHUB_S3_ACCESS_KEY`, `PIPESHUB_S3_SECRET_KEY`), one bucket (`PIPESHUB_S3_BUCKET`, or the first); syncs, waits until indexing finishes, searches it, then turns the sync off and deletes it | see `.env.example` | no |
| `./connectors/azure_files` | `CreateConnectorInstance` (Azure Files, connection string), `GetFilterFieldOptions` for shares, `UpdateConnectorFiltersSyncConfig` (one share), `ToggleConnector`, wait for indexing, scoped `Search`, then sync off and `DeleteConnectorInstance` | `PIPESHUB_AZURE_FILES_CONNECTION_STRING`; `PIPESHUB_AZURE_FILES_SHARE`, `PIPESHUB_AZURE_FILES_SEARCH_QUERY` (optional) | no |
| `./connectors/slack_workspace` | `CreateConnectorInstance` (Slack Workspace, bot token), `GetFilterFieldOptions` for channels, `UpdateConnectorFiltersSyncConfig` (one channel, last 7 days), `ToggleConnector`, wait for indexing, scoped `Search`, then sync off and `DeleteConnectorInstance`. Slack records keep Slack's permissions, so search only finds them for PipesHub users whose email matches a member of the channel; the example says so instead of failing when the caller sees none. | `PIPESHUB_SLACK_BOT_TOKEN`; `PIPESHUB_SLACK_CHANNEL_ID`, `PIPESHUB_SLACK_SEARCH_QUERY` (optional) | no |
| `./connectors/rss` | registry lookup, `CreateConnectorInstance` (RSS), `ToggleConnector`, `GetConnectorStats` until indexed, scoped `Search`, then sync off and `DeleteConnectorInstance` | `PIPESHUB_RSS_FEED_URL` (optional) | no |

```bash
go run ./st_examples/search .env
go run ./conversation/internal_search .env
go run ./agent/create .env          # prints an agent key
go run ./agent/delete .env <agent-key>
```

Edit the `query` string in any `main.go` to change the question.

The programs added for conversations, the KB trash and personal access tokens are self-contained: each creates the conversations, knowledge bases, agents or tokens it needs and deletes or revokes them before it exits. The exception is `KB/records/reindex_record_group`, which queues a reindex of an existing connector record group.

A few need something on the instance first:

- **`KB/trash/*`** need the workspace Labs setting **"Move Deleted Records to the Trash"** (`ENABLE_SOFT_DELETE`). With it off, a deleted record is gone for good, `ListTrash` stays empty and restoring fails.
- **`KB/records/reindex_record_group`** needs a connector record group. A knowledge base, or a folder in one, is not a record group here and returns `404`; use `KB/records/reindex_record` for KB content. The record group is found through the knowledge-hub calls, which the SDK marks deprecated but which are still the only way to browse a connector's record groups.
- **`personal_access_tokens/*`** log in with email and password (see [Authentication](#authentication)) and print token ids and metadata but never the secret. `expiryDays` accepts only `30`, `90`, `365` or `"never"`. The `admin_*` programs need an org admin.

## Understanding the stream (AG-UI)

Conversations stream over Server-Sent Events using the **AG-UI** protocol. The older `connected` / `answer_chunk` / `complete` / `error` vocabulary no longer exists.

| Legacy event | AG-UI replacement |
| --- | --- |
| `connected` | `CUSTOM` with `name: "conversation_created"`; `value.conversationId` on create routes |
| `answer_chunk` | `TEXT_MESSAGE_CONTENT`, token text in `delta` |
| `complete` | `RUN_FINISHED`, the old payload now nested under `result` |
| `error` | `RUN_ERROR` with `message` and `code` |

There are 21 event names in total, including `STEP_STARTED`/`STEP_FINISHED`, `TOOL_CALL_START`/`ARGS`/`END`/`RESULT`, the `REASONING_*` family, `STATE_DELTA`, `STATE_SNAPSHOT`, and `HEARTBEAT`. **Clients must ignore names they do not recognise** — new ones get added.

Two traps worth knowing before you write your own client:

1. **Child runs carry `parentRunId`.** In agent mode each sub-agent emits its own `RUN_STARTED`/`RUN_FINISHED`/`RUN_ERROR`, forwarded verbatim. A child `RUN_FINISHED` arrives *before* the terminal one and has no `result`, so treating the first one as terminal ends the stream with an empty answer. A child `RUN_ERROR` is one sub-agent failing, not the stream.
2. **`TEXT_MESSAGE_CONTENT` is raw text.** Citation references are rewritten as the answer is finalized, so the authoritative reply is the persisted message in `RUN_FINISHED.result.conversation`. (The citation-resolved running text rides on `STATE_DELTA` at path `/normalizedAnswer`.) Stream the tokens for the typing effect; trust the persisted message for the answer.

Also note `delta` is a **string** on `TEXT_MESSAGE_CONTENT` but a **JSON-Patch array** on `STATE_DELTA` — one flat envelope struct cannot decode both.

The shared decoder in [`agui/agui.go`](agui/agui.go) handles all of this; it is the Go counterpart of `integration-tests/helper/agui_sse.py` in the `pipeshub-ai` repo. [`conversation/internal_search/main.go`](conversation/internal_search/main.go) deliberately spells the same logic out longhand against the generated event constants, as a protocol reference. [`conversation/agui_trace/main.go`](conversation/agui_trace/main.go) prints every frame so you can watch it live.

## chatMode

`chatMode` is **required and non-pointer** on every streaming request, with two disjoint vocabularies:

| Route | Accepted values |
| --- | --- |
| `Conversations.StreamChat`, `Conversations.AddMessageStream` | `agent`, `internal_search`, `web_search` |
| `Agents.StreamAgentConversation`, `…Message`, `…Regenerate` | `quick` only |

```go
ChatMode: components.ConversationStreamRequestChatModeAgent
ChatMode: components.AgentStreamCreateConversationRequestChatModeQuick
```

On a follow-up, use the same mode the conversation was opened with. There is also a `Protocol` field whose only legal value is `agui`; omit it.

## Filters, tools, and capabilities

`Filters` is nullable and the distinction matters:

| Value | Sent as | Meaning |
| --- | --- | --- |
| `nil` | *(omitted)* | agent uses its stored knowledge config |
| `&Filters{Kb: []string{id}}` | `{"kb":["…"]}` | scoped to those sources |
| `&Filters{Apps: []string{}, Kb: []string{}}` | `{"apps":[],"kb":[]}` | **no** knowledge sources at all |
| `&Filters{}` | `{}` | ambiguous — avoid |

So pass `nil`, never an empty-but-non-nil `Filters`, when you mean "no opinion".

`Tools []string` is an allow-list of fully-qualified action names (`jira.create_issue`). Omit it to offer every configured tool; an empty non-nil slice offers none. `AgentCapabilities` toggles `InternalSearch` / `WebSearch` / `DeepSearch`, each treated as enabled when its field is absent. Both apply only under `chatMode: agent`.

## Project layout

```
examples/
├── go.mod  .env.example
├── agui/agui.go                        shared AG-UI SSE decoder
├── auth/login.go                       shared NewOAuthClient, NewSessionClient, NewClient
├── connectorsync/                      shared connector steps: options, filters, sync, wait, search, clean up
├── KB/
│   ├── kb/                             create, get, list, update, delete, upload_records,
│   │                                   get_upload_limits
│   ├── folder/                         create, update, delete, upload_records
│   ├── records/                        update, move, delete, reindex, stream_record_buffer,
│   │                                   get_record_by_id, reindex_record_group
│   └── trash/                          list_trash, restore_record, restore_records
├── agent/{create,list_and_get,update,delete}/
├── agent_conversation/                 9 programs, agent-scoped conversations
├── conversation/
│   ├── internal_search/                longhand AG-UI reference
│   ├── agent/                          chatMode=agent, tools, capabilities
│   ├── agui_trace/                     every frame, by category
│   ├── knowledgebase/  connector/      Filters.Kb / Filters.Apps
│   ├── web_search/
│   └── create_conversation/ …          11 non-streaming / management programs
├── personal_access_tokens/             list_scopes, create, list, revoke, admin list / revoke
├── semantic_search/{knowledgebase,connector}/
└── st_examples/{chat,search}/
```

## Notes

- `go.mod` requires `github.com/pipeshub-ai/pipeshub-sdk-go v1.8.0`. `go mod tidy` downloads it.
- Every `KB/*` program creates a fresh knowledge base named `Internal documents`, so it never touches your data. Repeated runs leave extra copies behind; delete them in the UI when you are done.
- One `main()` per directory — that is what keeps `go build ./...` and `go vet ./...` working across the tree.
- Troubleshooting: `401` means bad credentials; `500` on a conversation means no AI model provider is configured; an empty answer means nothing is indexed in the knowledge sources you scoped to.
