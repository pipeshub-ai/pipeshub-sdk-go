// Create and archive a conversation with a distinctive title, find it with
// SearchArchivedConversations, then delete it.
//
// The search covers archived assistant and agent conversations alike.
//
// Usage (from examples/):
//
//	go run ./conversation/search_archived_conversations <path-to-.env>
package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	pipeshub "github.com/pipeshub-ai/pipeshub-sdk-go"
	"github.com/pipeshub-ai/pipeshub-sdk-go/models/components"
	"github.com/pipeshub-ai/pipeshub-sdk-go/models/operations"

	"enterprise_search/auth"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatal("usage: go run ./conversation/search_archived_conversations <path-to-.env>")
	}
	if err := godotenv.Load(os.Args[1]); err != nil {
		log.Fatalf("load .env: %v", err)
	}

	sdk, err := auth.NewOAuthClient()
	if err != nil {
		log.Fatal(err)
	}

	ctx := context.Background()

	// A unique word to search for. The title is set explicitly so the search
	// does not depend on how the server titles a new conversation.
	needle := fmt.Sprintf("goexample%d", time.Now().Unix())

	convID, err := createConversation(ctx, sdk, "Which aviation stories are in the news?")
	if err != nil {
		log.Fatal(err)
	}
	if _, err := sdk.Conversations.UpdateConversationTitle(ctx, convID, operations.UpdateConversationTitleRequestBody{
		Title: "Aviation news " + needle,
	}); err != nil {
		log.Fatalf("update conversation title: %v", err)
	}
	if _, err := sdk.Conversations.ArchiveConversation(ctx, convID); err != nil {
		log.Fatalf("archive conversation: %v", err)
	}
	fmt.Printf("Created and archived conversation: %s\n", convID)

	res, err := sdk.Conversations.SearchArchivedConversations(ctx, needle, pipeshub.Pointer(int64(1)), pipeshub.Pointer(int64(10)))
	if err != nil {
		log.Fatalf("search archived conversations: %v", err)
	}
	if res.Object == nil {
		log.Fatal("search archived conversations: empty response")
	}

	fmt.Printf("Search %q matched %d conversation(s):\n", needle, len(res.Object.Conversations))
	for _, c := range res.Object.Conversations {
		source := "assistant"
		if c.Source != nil {
			source = string(*c.Source)
		}
		fmt.Printf("  %s  [%s]  %s\n", ptr(c.ID), source, ptr(c.Title))
	}

	if _, err := sdk.Conversations.DeleteConversationByID(ctx, convID); err != nil {
		log.Fatalf("delete conversation: %v", err)
	}
	fmt.Println("Deleted the conversation.")
}

// createConversation opens a conversation with the non-streaming
// CreateConversation call and returns its id (the "_id" field).
func createConversation(ctx context.Context, sdk *pipeshub.Pipeshub, query string) (string, error) {
	res, err := sdk.Conversations.CreateConversation(ctx, components.CreateConversationRequest{
		Query:    query,
		ChatMode: pipeshub.Pointer(components.CreateConversationRequestChatModeInternalSearch),
	})
	if err != nil {
		return "", fmt.Errorf("create conversation: %w", err)
	}
	if res.CreateConversationResponse == nil {
		return "", fmt.Errorf("create conversation: empty response")
	}
	id := res.CreateConversationResponse.Conversation.GetID()
	if id == nil || *id == "" {
		return "", fmt.Errorf("create conversation: response has no _id")
	}
	return *id, nil
}

func ptr(s *string) string {
	if s == nil {
		return ""
	}
	return *s
}
