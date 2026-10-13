# Changelog

## 1.7.0

Compatible with PipesHub [v0.9.1](https://github.com/pipeshub-ai/pipeshub-ai/releases/tag/v0.9.1).

> **1.7.0 is a minor version, but it contains breaking changes.** Read this list before you upgrade.

### Breaking changes and how to fix them

#### Go `1.26.8` or newer is required

`go.mod` now declares `go 1.26.8` (it was `go 1.22`). Upgrade Go, or keep the default `GOTOOLCHAIN=auto` so that Go downloads the toolchain.

```sh
go version            # must print go1.26.8 or newer
go get github.com/pipeshub-ai/pipeshub-sdk-go@v1.7.0
```

#### Renaming a conversation no longer returns `messages`

`UpdateConversationTitle` returns only the updated conversation fields. Read the messages with a second call.

```go
// ❌ 1.6.0
// res, err := s.Conversations.UpdateConversationTitle(ctx, conversationID, operations.UpdateConversationTitleRequestBody{Title: title})
// messages := res.Object.Conversation.Messages

// ✅ 1.7.0
if _, err := s.Conversations.UpdateConversationTitle(ctx, conversationID, operations.UpdateConversationTitleRequestBody{Title: title}); err != nil {
	return err
}
res, err := s.Conversations.GetConversationByID(ctx, operations.GetConversationByIDRequest{ConversationID: conversationID})
if err != nil {
	return err
}
var messages []operations.Message
if res.Object != nil && res.Object.Conversation != nil {
	messages = res.Object.Conversation.Messages
}
```

#### Deleting an agent conversation: `Conversation` can be empty

`Conversation` is now a `*components.ConversationUnion`, which can be `nil`, a stored conversation, or an empty object.

```go
res, err := s.Agents.DeleteAgentConversationByID(ctx, agentKey, conversationID)
if err != nil {
	return err
}

// ❌ 1.6.0
// fmt.Println(*res.AgentConversationDeleteResponse.Conversation.ID)

// ✅ 1.7.0
if c := res.AgentConversationDeleteResponse.Conversation; c != nil && c.StoredAgentConversation != nil && c.StoredAgentConversation.ID != nil {
	fmt.Println(*c.StoredAgentConversation.ID)
}
```

#### Some type names changed

Use the new name. The data inside is the same. Enum constants take the new prefix too, for example `operations.IsArchivedTrue` becomes `operations.ListAgentConversationsIsArchivedTrue`.

| Old name | New name |
|---|---|
| `components.Meta` | `components.MessageFeedbackUpdateResponseMeta` |
| `components.ModelInfo` | `components.ConversationListItemModelInfo` |
| `operations.IsArchived` | `operations.ListAgentConversationsIsArchived` |
| `operations.GetConversationByIDMessage` | `operations.Message` |
| `operations.GetConversationByIDMessageMessageType` | `operations.MessageMessageType` |
| `operations.GetConversationByIDContentFormat` | `operations.ContentFormat` |
| `operations.GetConversationByIDReferenceDatum` | `operations.ReferenceDatum` |
| `operations.GetConversationByIDAppliedFilters` | `operations.AppliedFilters` |

The `operations.UpdateConversationTitle*` message types are removed, because renaming a conversation no longer returns messages.

```go
// ❌ 1.6.0
// var info *components.ModelInfo
// archived := operations.IsArchivedTrue

// ✅ 1.7.0
var info *components.ConversationListItemModelInfo
archived := operations.ListAgentConversationsIsArchivedTrue
```

#### Two errors have a new error type

A `500` from `KnowledgeBase.UploadRecords` is now returned as `*apierrors.ErrorResponse`. A `410` from `UserAccount.Authenticate` is now the default `*apierrors.APIError`, not `*apierrors.ErrorResponse`. Check for both types.

```go
var errResp *apierrors.ErrorResponse
var apiErr *apierrors.APIError
switch {
case errors.As(err, &errResp):
	// error with a JSON body
case errors.As(err, &apiErr) && apiErr.StatusCode == 410:
	// the 410 from Authenticate
}
```

#### `Success` is always `true`

On the delete-record, reindex-record, reindex-record-group and knowledge-hub-nodes responses, `Success` is no longer a public field. A `200` already means success, so you can drop the check, or use the getter.

```go
// ❌ 1.6.0
// if res.DeleteRecordResponseSchema.Success { ... }

// ✅ 1.7.0
if res.DeleteRecordResponseSchema.GetSuccess() { ... }
```

#### A citation can have two shapes

A message citation is now a `components.CitationUnion`: a `CitationReference` or a `PopulatedCitationReference`. Read the citation through its wrapper.

```go
// ❌ 1.6.0
// for _, c := range msg.Citations { fmt.Println(*c.CitationID) }

// ✅ 1.7.0
for _, c := range msg.Citations {
	if c.CitationReference != nil && c.CitationReference.CitationID != nil {
		fmt.Println(*c.CitationReference.CitationID)
	}
}
```

### Behavior notes (not breaking)

- **A user token on 22 endpoints.** OAuth apps, personal access tokens, user info, org auth config and reset password now send only the user token. If you set up only an app token, add a user token:
  ```go
  s := pipeshub.New(pipeshub.WithSecurity(components.Security{BearerAuth: pipeshub.Pointer(userToken)}))
  ```
- **Stream event `data` stays a string.** Code that reads `*ev.Data` keeps working.
