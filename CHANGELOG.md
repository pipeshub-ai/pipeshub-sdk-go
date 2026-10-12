# Changelog

## 1.7.0

Compatible with PipesHub [v0.9.1](https://github.com/pipeshub-ai/pipeshub-ai/releases/tag/v0.9.1). This version is generated from the PipesHub v0.9.1 API spec. The changes for all three SDKs are also in the [PipesHub changelog](https://github.com/pipeshub-ai/pipeshub-ai/blob/main/changelog/0.9.0.md#sdks--170-typescript-python-go).

> **1.7.0 is a minor version, but it contains breaking changes.** Read the list below before you upgrade.

### What's new

- **Projects:** a new `Projects` namespace. You can create, list, get, update, delete, archive, unarchive, pin and unpin projects. You can also manage project members, the project knowledge base, and a project's conversations. `Conversations.SetConversationProject`, `Conversations.SetConversationProjectVisibility`, `Projects.SetAgentConversationProject` and `Projects.SetAgentConversationProjectVisibility` move a conversation into a project and set its visibility.
- **Non-streaming chat:** `Conversations.CreateConversation`, `Conversations.AddMessage`, `Agents.CreateAgentConversation` and `Agents.AddAgentConversationMessage` return the finished answer as JSON.
- **Stop a running answer:** `Conversations.CancelConversationStream` and `Agents.CancelAgentConversationStream`.
- **Examples:** `examples/` has one folder for each example, built against the published SDK.

### Breaking changes and how to fix them

#### Go `1.26.8` or newer is required

`go.mod` now declares `go 1.26.8` (it was `1.22`). Upgrade Go, or let Go download the toolchain with the default `GOTOOLCHAIN=auto`.

#### Renaming a conversation no longer returns `messages`

Get the messages with a second call.

```go
sdk.Conversations.UpdateConversationTitle(ctx, ...)
res, _ := sdk.Conversations.GetConversationByID(ctx, operations.GetConversationByIDRequest{ConversationID: id})
```

#### Delete agent conversation: `Conversation` can be empty

Check that the conversation exists before you read it.

```go
if c := res.AgentConversationDeleteResponse.Conversation; c != nil && c.StoredAgentConversation != nil {
    fmt.Println(c.StoredAgentConversation.ID)
}
```

#### Some type names changed

Use the new name. The data inside is the same.

| Old name | New name |
|---|---|
| `Meta` | `MessageFeedbackUpdateResponseMeta` |
| `ModelInfo` | `ConversationListItemModelInfo` |
| `IsArchived` | `ListAgentConversationsIsArchived` |
| `GetConversationByIDMessage` | `Message` |
| `GetConversationByIDMessageMessageType` | `MessageMessageType` |
| `GetConversationByIDContentFormat` | `ContentFormat` |
| `GetConversationByIDReferenceDatum` | `ReferenceDatum` |
| `GetConversationByIDAppliedFilters` | `AppliedFilters` |

The `UpdateConversationTitle*` message types are removed, because renaming a conversation no longer returns messages.

#### Two errors have a new error class

A `500` from `KnowledgeBase.UploadRecords` is now an `apierrors.ErrorResponse`. A `410` from `UserAccount.Authenticate` is now the default `apierrors.APIError` instead of an `apierrors.ErrorResponse`. Check for both.

```go
var apiErr *apierrors.APIError
var errResp *apierrors.ErrorResponse
if errors.As(err, &errResp) || errors.As(err, &apiErr) { ... }
```

#### `Success` is always `true`

On the delete-record, reindex-record, reindex-record-group and knowledge-hub-nodes responses, `Success` is no longer a public field. Use the method, or drop the check: a 200 response already means success.

```go
// ❌ if res.DeleteRecordResponseSchema.Success { ... }
// ✅
if res.DeleteRecordResponseSchema.GetSuccess() { ... }
```

#### A citation can have two shapes

A message citation is now a `CitationReference` or a `PopulatedCitationReference`. Read it through its wrapper.

```go
// ❌ c.CitationID
// ✅
for _, c := range msg.Citations {
    if c.CitationReference != nil { fmt.Println(*c.CitationReference.CitationID) }
}
```

### Behavior notes (not breaking)

- **Login on 22 endpoints** (OAuth apps, access tokens, user info, org auth config, reset password): these calls send only the user token. If you set up only an app token, add a user token.
  ```go
  s := pipeshub.New(pipeshub.WithSecurity(components.Security{BearerAuth: pipeshub.Pointer(userToken)}))
  ```
- **Stream event `data` stays a string.** Code that reads `ev.Data` keeps working.
