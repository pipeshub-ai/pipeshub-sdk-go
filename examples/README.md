# Enterprise Search Quickstart

Four minimal Go programs that authenticate against a PipesHub instance and run either a **semantic search** or an **AI conversation** — each scoped either by knowledge base or by connector.

## Prerequisites

- Go 1.22+
- A running PipesHub instance (default: `http://localhost:3000`)
- A user account on that instance
- At least one knowledge base and/or connector with indexed content
- The PipesHub Go SDK (`github.com/pipeshub-ai/pipeshub-sdk-go`) — pulled in automatically by `go mod tidy`

## Setup

Create a `.env` file in this directory:

```dotenv
PIPESHUB_BASE_URL=http://localhost:3000
PIPESHUB_TEST_USER_EMAIL=you@example.com
PIPESHUB_TEST_USER_PASSWORD=your-password

KB_ID=<knowledge-base-uuid>
CONNECTOR_ID=<connector-instance-uuid>
```

Install dependencies:

```bash
go mod tidy
```

## Run

From inside `enterprise_search/`:

```bash
# Semantic search
go run ./semantic_search/connector   .env

# AI conversation
go run ./conversation/knowledgebase   .env
go run ./conversation/connector       .env
go run ./conversation/web_search      .env
go run ./conversation/internal_search .env
```

The path argument is the location of your `.env` file.

- **Semantic search** prints the top hits as `Record / ID / Chunk` blocks.
- **Conversation** streams the AI's answer to stdout as it's generated, using `Conversations.StreamChat` (SSE). The non-streaming `Conversations.CreateConversation` SDK method is not used here because the corresponding non-streaming route on the AI backend is currently unavailable; only the `/chat/stream` route is implemented.

Edit the `Query` string in any `main.go` to change the question.

## Project layout

```
enterprise_search/
├── go.mod
├── .env
├── auth/login.go                       shared NewClient (InitAuth → Authenticate)
├── semantic_search/
│   ├── knowledgebase/main.go           SemanticSearch.Search, KB-filtered
│   └── connector/main.go               SemanticSearch.Search, connector-filtered
└── conversation/
    ├── knowledgebase/main.go           Conversations.StreamChat, KB-filtered
    ├── connector/main.go               Conversations.StreamChat, connector-filtered
    ├── web_search/main.go              Conversations.StreamChat, chatMode=web_search
    └── internal_search/main.go         Conversations.StreamChat, chatMode=internal_search (KB + connector filters)
```

All four programs import `enterprise_search/auth` to get an authenticated `*pipeshub.Pipeshub`, then differ only in which SDK call and which `Filters` they use:

- `Filters.Kb: []string{...}` — restrict to one or more knowledge bases
- `Filters.Apps: []string{...}` — restrict to specific connector instances (instance UUIDs)

## Notes

- The SDK module (`github.com/pipeshub-ai/pipeshub-sdk-go`) is fetched from the public module proxy via `go mod tidy`. To pin a specific version, run `go get github.com/pipeshub-ai/pipeshub-sdk-go@<version>`.
- Each leaf subdirectory is its own Go package with its own `main()` — Go requires one `main` per package, so each runnable demo lives in its own directory, but they all share the parent module and the `auth/` package.
- Conversation calls require an AI model provider to be configured on your PipesHub instance; otherwise the server returns `500 — Failed to get AI response`.

