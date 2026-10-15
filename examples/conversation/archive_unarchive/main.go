// Create a conversation, archive it, unarchive it, then delete it.
//
// Usage (from examples/):
//
//	go run ./conversation/archive_unarchive <path-to-.env>
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

const query = "Which sports stories are in the news?"

func main() {
	if len(os.Args) < 2 {
		log.Fatal("usage: go run ./conversation/archive_unarchive <path-to-.env>")
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

	archived, err := sdk.Conversations.ArchiveConversation(ctx, convID)
	if err != nil {
		log.Fatalf("archive conversation: %v", err)
	}
	if b := archived.Object; b != nil {
		fmt.Printf("Archived:   status=%s at=%s\n", str(b.Status), when(b.ArchivedAt))
	}

	unarchived, err := sdk.Conversations.UnarchiveConversation(ctx, convID)
	if err != nil {
		log.Fatalf("unarchive conversation: %v", err)
	}
	if b := unarchived.Object; b != nil {
		fmt.Printf("Unarchived: status=%s at=%s\n", str(b.Status), when(b.UnarchivedAt))
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

// str prints an optional string-typed enum.
func str[T ~string](v *T) string {
	if v == nil {
		return ""
	}
	return string(*v)
}

func when(t *time.Time) string {
	if t == nil {
		return "-"
	}
	return t.Format(time.RFC1123)
}
