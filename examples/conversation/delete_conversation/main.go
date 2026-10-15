// Create a conversation, then delete it with DeleteConversationByID.
//
// Usage (from examples/):
//
//	go run ./conversation/delete_conversation <path-to-.env>
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

	"enterprise_search/auth"
)

const query = "Reply with a one-sentence greeting."

func main() {
	if len(os.Args) < 2 {
		log.Fatal("usage: go run ./conversation/delete_conversation <path-to-.env>")
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
	fmt.Printf("Created conversation: %s\n", convID)

	res, err := sdk.Conversations.DeleteConversationByID(ctx, convID)
	if err != nil {
		log.Fatalf("delete conversation: %v", err)
	}
	if res.Object == nil {
		log.Fatal("delete conversation: empty response")
	}
	b := res.Object

	status := ""
	if b.Status != nil {
		status = string(*b.Status)
	}
	deletedAt := "-"
	if b.DeletedAt != nil {
		deletedAt = b.DeletedAt.Format(time.RFC1123)
	}
	citations := int64(0)
	if b.CitationsDeleted != nil {
		citations = *b.CitationsDeleted
	}
	fmt.Printf("Deleted conversation: %s\n", ptr(b.ID))
	fmt.Printf("  status:            %s\n", status)
	fmt.Printf("  deleted at:        %s\n", deletedAt)
	fmt.Printf("  citations removed: %d\n", citations)
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
