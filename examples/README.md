# PipesHub Go SDK examples

38 runnable Go programs that authenticate against a PipesHub instance and exercise semantic search, conversations, the agent loop, agent CRUD, and knowledge base management.

The Go module is named `enterprise_search` (which is why imports read `enterprise_search/auth`), but it lives in this `examples/` directory. Each program is its own package in its own directory, so `go build ./...` and `go vet ./...` cover the whole tree.

## Prerequisites

- Go 1.25.10+ (matching the SDK's `go` directive)
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

| Variable | Needed by |
| --- | --- |
| `PIPESHUB_BASE_URL` | all — the examples append `/api/v1` themselves |
| `PIPESHUB_BEARER_AUTH` | optional; a raw JWT access token. When set, it is used directly and the email/password login is skipped |
| `PIPESHUB_TEST_USER_EMAIL`, `PIPESHUB_TEST_USER_PASSWORD` | all, unless `PIPESHUB_BEARER_AUTH` is set |
| `PIPESHUB_AGENT_KEY` | all `agent_conversation/*` except `create_and_add_message` |
| `KB_ID`, `CONNECTOR_ID` | optional knowledge filters for `agent_conversation/*` |
| `PIPESHUB_KB_NAME`, `PIPESHUB_CONNECTOR_NAME` | optional lookup-name overrides (default `SDK-test`, `ABC News RSS` / `abc news`) |
| `PIPESHUB_AGENT_MODEL_KEY` | optional; used when an example creates an agent |
| `PIPESHUB_UPLOAD_PATH` | optional; file for `KB/*/upload_records` (falls back to a built-in sample) |

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

```bash
go run ./st_examples/search .env
go run ./conversation/internal_search .env
go run ./agent/create .env          # prints an agent key
go run ./agent/delete .env <agent-key>
```

Edit the `query` string in any `main.go` to change the question.

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
├── auth/login.go                       shared NewClient (InitAuth → Authenticate)
├── KB/
│   ├── kb/                             create, get, list, update, delete, upload_records
│   ├── folder/                         create, update, delete, upload_records
│   └── records/                        update, move, delete, reindex, stream_record_buffer
├── agent/{create,list_and_get,update,delete}/
├── agent_conversation/                 9 programs, agent-scoped conversations
├── conversation/
│   ├── internal_search/                longhand AG-UI reference
│   ├── agent/                          chatMode=agent, tools, capabilities
│   ├── agui_trace/                     every frame, by category
│   ├── knowledgebase/  connector/      Filters.Kb / Filters.Apps
│   └── web_search/
├── semantic_search/{knowledgebase,connector}/
└── st_examples/{chat,search}/
```

## Notes

- `go.mod` carries `replace github.com/pipeshub-ai/pipeshub-sdk-go => ../`, so the examples always build against the SDK source in this repo. Drop that line to build against a published version.
- Every `KB/*` program creates a fresh knowledge base named `Internal documents`, so it never touches your data. Repeated runs leave extra copies behind; delete them in the UI when you are done.
- One `main()` per directory — that is what keeps `go build ./...` and `go vet ./...` working across the tree.
- Troubleshooting: `401` means bad credentials; `500` on a conversation means no AI model provider is configured; an empty answer means nothing is indexed in the knowledge sources you scoped to.
