// Create and archive a conversation, list archived conversations with
// GetArchivedConversations, then delete the conversation.
//
// Usage (from examples/):
//
//	go run ./conversation/get_archived_conversations <path-to-.env>
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

const query = "Are there any business stories in the news?"

func main() {
	if len(os.Args) < 2 {
		log.Fatal("usage: go run ./conversation/get_archived_conversations <path-to-.env>")
	}
	if err := godotenv.Load(os.Args[1]); err != nil {
		log.Fatalf("load .env: %v", err)
	}

	sdk, err := auth.NewOAuthClient()
	if err != nil {
		log.Fatal(err)
	}

	ctx := context.Background()

	convID, err := createConversation(ctx, sdk, query)
	if err != nil {
		log.Fatal(err)
	}
	if _, err := sdk.Conversations.ArchiveConversation(ctx, convID); err != nil {
		log.Fatalf("archive conversation: %v", err)
	}
	fmt.Printf("Created and archived conversation: %s\n\n", convID)

	res, err := sdk.Conversations.GetArchivedConversations(ctx, operations.GetArchivedConversationsRequest{
		Page:      pipeshub.Pointer(int64(1)),
		Limit:     pipeshub.Pointer(int64(10)),
		SortBy:    pipeshub.Pointer(operations.GetArchivedConversationsSortByEnumLastActivityAt),
		SortOrder: pipeshub.Pointer(operations.GetArchivedConversationsSortOrderEnumDesc),
	})
	if err != nil {
		log.Fatalf("get archived conversations: %v", err)
	}
	if res.Object == nil {
		log.Fatal("get archived conversations: empty response")
	}

	if s := res.Object.Summary; s != nil && s.TotalArchived != nil {
		fmt.Printf("Archived conversations: %d\n", *s.TotalArchived)
	}
	for _, c := range res.Object.Conversations {
		marker := " "
		if c.ID != nil && *c.ID == convID {
			marker = "*"
		}
		fmt.Printf("%s %s  archived %s  %s\n", marker, ptr(c.ID), when(c.ArchivedAt), ptr(c.Title))
	}
	fmt.Println("\n* = the conversation this example archived")

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

func when(t *time.Time) string {
	if t == nil {
		return "-"
	}
	return t.Format(time.DateTime)
}

func ptr(s *string) string {
	if s == nil {
		return ""
	}
	return *s
}
