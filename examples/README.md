# PipesHub Go SDK examples

85 runnable Go programs that authenticate against a PipesHub instance and exercise semantic search, conversations, the agent loop, agent CRUD, connectors and knowledge base management.

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

The programs added in SDK 1.7.1 (`KB/kb/get_upload_limits`, `KB/records/get_record_by_id`, `KB/records/reindex_record_group`, `KB/trash/*`, the non-streaming `conversation/*` programs) and `connectors/*` authenticate with **OAuth2 client credentials**, each building the client inline. Create an OAuth app with the `client_credentials` grant and the scopes the examples use, then set `PIPESHUB_CLIENT_ID` and `PIPESHUB_CLIENT_SECRET`. The SDK fetches the access token from `PIPESHUB_TOKEN_URL` or, when that is unset, `${PIPESHUB_BASE_URL}/api/v1/oauth2/token`, and fetches a new one when it expires. The SDK's built-in default token URL is relative, so the programs always pass an absolute one.

```go
// What those programs pass to pipeshub.WithSecurity
components.Security{Oauth2: &components.SchemeOauth2{
	ClientID:     os.Getenv("PIPESHUB_CLIENT_ID"),
	ClientSecret: os.Getenv("PIPESHUB_CLIENT_SECRET"),
	TokenURL:     "http://localhost:3000/api/v1/oauth2/token",
}}
```

The older programs use `auth.NewClient`: `PIPESHUB_BEARER_AUTH` when it is set, otherwise the same email and password login.


| Variable | Needed by |
| --- | --- |
| `PIPESHUB_BASE_URL` | all — the examples append `/api/v1` themselves |
| `PIPESHUB_CLIENT_ID`, `PIPESHUB_CLIENT_SECRET` | the 1.7.1 programs (see [Authentication](#authentication)); OAuth2 client credentials |
| `PIPESHUB_TOKEN_URL` | optional; OAuth2 token endpoint (default `${PIPESHUB_BASE_URL}/api/v1/oauth2/token`) |
| `PIPESHUB_TEST_USER_EMAIL`, `PIPESHUB_TEST_USER_PASSWORD` | the older programs unless `PIPESHUB_BEARER_AUTH` is set |
| `PIPESHUB_BEARER_AUTH` | optional, older programs only; a JWT access token used instead of the email and password login |
| `PIPESHUB_AGENT_KEY` | all `agent_conversation/*` except `create_and_add_message` |
| `KB_ID`, `CONNECTOR_ID` | optional knowledge filters for `agent_conversation/*` |
| `PIPESHUB_KB_NAME`, `PIPESHUB_CONNECTOR_NAME` | optional lookup-name overrides (default `SDK-test`, `ABC News RSS` / `abc news`). `KB/records/reindex_record_group` falls back to the first connector |
| `PIPESHUB_AGENT_MODEL_KEY` | optional; used when an example creates an agent |
| `PIPESHUB_UPLOAD_PATH` | optional; file for `KB/*/upload_records`, `KB/records/get_record_by_id` and `KB/trash/*` (falls back to a built-in sample) |
| `PIPESHUB_CONVERSATION_ID`, `PIPESHUB_MESSAGE_ID`, `PIPESHUB_RUN_ID` | the single-operation `conversation/*` programs |
| `PIPESHUB_KB_ID`, `PIPESHUB_RECORD_ID`, `PIPESHUB_RECORD_GROUP_ID` | the single-operation `KB/*` programs |
| `PIPESHUB_CONNECTOR_ID`, `PIPESHUB_CONNECTOR_TYPE` (e.g. `RSS`), `PIPESHUB_FILTER_KEY` | the `connectors/*` programs |
| `PIPESHUB_OAUTH_CONFIG_ID`, `PIPESHUB_OAUTH_CODE`, `PIPESHUB_OAUTH_STATE` | the OAuth programs in `connectors/*` |

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
| `./conversation/add_message` | `AddMessage` (non-streaming follow-up) | `PIPESHUB_CONVERSATION_ID` | yes |
| `./conversation/get_all_conversations` | paginated `GetAllConversations` | — | yes |
| `./conversation/get_conversation_by_id` | `GetConversationByID` | `PIPESHUB_CONVERSATION_ID` | yes |
| `./conversation/update_conversation_title` | `UpdateConversationTitle` | `PIPESHUB_CONVERSATION_ID` | yes |
| `./conversation/get_archived_conversations` | `GetArchivedConversations` | — | yes |
| `./conversation/search_archived_conversations` | `SearchArchivedConversations` | — | yes |
| `./conversation/update_message_feedback` | `UpdateMessageFeedback` | `PIPESHUB_CONVERSATION_ID`, `PIPESHUB_MESSAGE_ID` | yes |
| `./conversation/delete_conversation` | `DeleteConversationByID` | `PIPESHUB_CONVERSATION_ID` | yes |
| `./conversation/cancel_conversation_stream` | `StreamChat` with a `runId`, then `CancelConversationStream` | `PIPESHUB_CONVERSATION_ID`, `PIPESHUB_RUN_ID` | yes |
| `./KB/kb/get_upload_limits` | `GetUploadLimits` | — | no |
| `./KB/records/get_record_by_id` | `GetRecordByID` | `PIPESHUB_RECORD_ID` | no |
| `./KB/records/reindex_record_group` | `ReindexRecordGroup` on a connector record group | `PIPESHUB_RECORD_GROUP_ID` | no |
| `./KB/trash/list_trash` | `ListTrash` | `PIPESHUB_KB_ID`; trash enabled (see below) | no |
| `./KB/trash/restore_record` | `RestoreRecord` | `PIPESHUB_RECORD_ID`; trash enabled (see below) | no |
| `./KB/trash/restore_records` | `RestoreRecords` | `PIPESHUB_RECORD_ID`; trash enabled (see below) | no |
| `./conversation/archive_conversation` | Archive a conversation. | `PIPESHUB_CONVERSATION_ID` | no |
| `./conversation/unarchive_conversation` | Unarchive a conversation. | `PIPESHUB_CONVERSATION_ID` | no |
| `./connectors/create_connector_instance` | Create an RSS connector instance (team scope, manual sync) for one feed. | — | no |
| `./connectors/create_oauth_config` | Create an OAuth app config for a connector type. | `PIPESHUB_CONNECTOR_TYPE` | no |
| `./connectors/delete_connector_instance` | Delete a connector instance and its records. | `PIPESHUB_CONNECTOR_ID` | no |
| `./connectors/delete_oauth_config` | Delete an OAuth app config. | `PIPESHUB_CONNECTOR_TYPE`, `PIPESHUB_OAUTH_CONFIG_ID` | no |
| `./connectors/get_connector_config` | Read a connector instance's auth, sync and filter config. | `PIPESHUB_CONNECTOR_ID` | no |
| `./connectors/get_connector_egress_ips` | List the IP addresses PipesHub connects from (for allow-lists). | — | no |
| `./connectors/get_connector_filters` | Read a connector instance's saved filters. | `PIPESHUB_CONNECTOR_ID` | no |
| `./connectors/get_connector_instance` | Fetch one connector instance. | `PIPESHUB_CONNECTOR_ID` | no |
| `./connectors/get_connector_registry` | List the connector types available on the instance. | — | no |
| `./connectors/get_connector_schema` | Get a connector type's config schema (auth, sync and filter fields). | `PIPESHUB_CONNECTOR_TYPE` | no |
| `./connectors/get_connector_stats` | Get indexing statistics for a connector instance. | `PIPESHUB_CONNECTOR_ID` | no |
| `./connectors/get_filter_field_options` | List the values you can pick for one filter field (a channel, space, folder...). | `PIPESHUB_CONNECTOR_ID`, `PIPESHUB_FILTER_KEY` | no |
| `./connectors/get_oauth_authorization_url` | Get the provider consent URL for an OAuth connector instance. | `PIPESHUB_CONNECTOR_ID` | no |
| `./connectors/get_oauth_config` | Fetch one OAuth app config (secrets are masked). | `PIPESHUB_CONNECTOR_TYPE`, `PIPESHUB_OAUTH_CONFIG_ID` | no |
| `./connectors/get_oauth_connector_type` | Get the OAuth config fields for one connector type. | `PIPESHUB_CONNECTOR_TYPE` | no |
| `./connectors/get_oauth_registry` | List the connector types that support OAuth app configs. | — | no |
| `./connectors/handle_oauth_callback` | Finish the OAuth flow with the code and state the provider redirected back with. | `PIPESHUB_OAUTH_CODE`, `PIPESHUB_OAUTH_STATE` | no |
| `./connectors/list_connector_instances` | List the team connector instances. | — | no |
| `./connectors/list_oauth_configs` | List all OAuth app configs. | — | no |
| `./connectors/list_oauth_configs_by_type` | List the OAuth app configs of one connector type. | `PIPESHUB_CONNECTOR_TYPE` | no |
| `./connectors/resync_connector` | Start a full resync of a connector instance. | `PIPESHUB_CONNECTOR_ID`, `PIPESHUB_CONNECTOR_TYPE` | no |
| `./connectors/save_connector_filters` | Save a connector instance's filters (keys come from its schema). | `PIPESHUB_CONNECTOR_ID` | no |
| `./connectors/stop_connector_sync` | Stop a running sync on a connector instance. | `PIPESHUB_CONNECTOR_ID` | no |
| `./connectors/test_connector_connection` | Test a connector's connection with the auth config of an existing instance. | `PIPESHUB_CONNECTOR_TYPE`, `PIPESHUB_CONNECTOR_ID` | no |
| `./connectors/toggle_connector` | Flip a connector's sync on or off (turning it on starts a sync). | `PIPESHUB_CONNECTOR_ID` | no |
| `./connectors/update_connector_auth_config` | Point a connector at an OAuth app config. | `PIPESHUB_CONNECTOR_ID`, `PIPESHUB_OAUTH_CONFIG_ID` | no |
| `./connectors/update_connector_config` | Change a connector's sync strategy; auth and filters stay as they are. | `PIPESHUB_CONNECTOR_ID` | no |
| `./connectors/update_connector_filters_sync_config` | Change a connector's sync strategy (filters can go in the same call). | `PIPESHUB_CONNECTOR_ID` | no |
| `./connectors/update_oauth_config` | Rename an OAuth app config. | `PIPESHUB_CONNECTOR_TYPE`, `PIPESHUB_OAUTH_CONFIG_ID` | no |

```bash
go run ./st_examples/search .env
go run ./conversation/internal_search .env
go run ./agent/create .env          # prints an agent key
go run ./agent/delete .env <agent-key>
```

Edit the `query` string in any `main.go` to change the question.

The programs added in 1.7.1 and for connectors are quick-start snippets: each makes one call, takes the ids it needs from environment variables (see `.env.example`) and prints the response. Create the conversation, record or connector first and put its id in your `.env`.

A few need something on the instance first:

- **`KB/trash/*`** need the workspace Labs setting **"Move Deleted Records to the Trash"** (`ENABLE_SOFT_DELETE`). With it off, a deleted record is gone for good, `ListTrash` stays empty and restoring fails.
- **`KB/records/reindex_record_group`** needs a connector record group. A knowledge base, or a folder in one, is not a record group here and returns `404`; use `KB/records/reindex_record` for KB content. The record group is found through the knowledge-hub calls, which the SDK marks deprecated but which are still the only way to browse a connector's record groups.

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
├── auth/login.go                       shared NewOAuthClient, NewSessionClient, NewClient (older programs)
├── connectors/                         one program per connector / OAuth app config operation
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
│   └── create_conversation/ …          one program per non-streaming / management operation
├── semantic_search/{knowledgebase,connector}/
└── st_examples/{chat,search}/
```

## Notes

- `go.mod` requires `github.com/pipeshub-ai/pipeshub-sdk-go v1.8.0`. `go mod tidy` downloads it.
- Every `KB/*` program creates a fresh knowledge base named `Internal documents`, so it never touches your data. Repeated runs leave extra copies behind; delete them in the UI when you are done.
- One `main()` per directory — that is what keeps `go build ./...` and `go vet ./...` working across the tree.
- Troubleshooting: `401` means bad credentials; `500` on a conversation means no AI model provider is configured; an empty answer means nothing is indexed in the knowledge sources you scoped to.
