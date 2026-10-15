// Create a conversation, rename it with UpdateConversationTitle, then delete it.
//
// Usage (from examples/):
//
//	go run ./conversation/update_conversation_title <path-to-.env>
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

const query = "Give me a one-line summary of today's news."

func main() {
	if len(os.Args) < 2 {
		log.Fatal("usage: go run ./conversation/update_conversation_title <path-to-.env>")
	}
	if err := godotenv.Load(os.Args[1]); err != nil {
		log.Fatalf("load .env: %v", err)
	}

	sdk, err := auth.NewOAuthClient()
	if err != nil {
		log.Fatal(err)
	}

	ctx := context.Background()

	convID, oldTitle, err := createConversation(ctx, sdk, query)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Created conversation: %s\n", convID)
	fmt.Printf("Old title: %q\n", oldTitle)

	newTitle := fmt.Sprintf("News summary (renamed by the Go SDK, %s)", time.Now().Format(time.DateOnly))
	res, err := sdk.Conversations.UpdateConversationTitle(ctx, convID, operations.UpdateConversationTitleRequestBody{
		Title: newTitle,
	})
	if err != nil {
		log.Fatalf("update conversation title: %v", err)
	}
	if res.Object == nil {
		log.Fatal("update conversation title: empty response")
	}
	fmt.Printf("New title: %q\n", ptr(res.Object.Conversation.Title))

	if _, err := sdk.Conversations.DeleteConversationByID(ctx, convID); err != nil {
		log.Fatalf("delete conversation: %v", err)
	}
	fmt.Println("Deleted the conversation.")
}

// createConversation opens a conversation with the non-streaming
// CreateConversation call and returns its id (the "_id" field) and title.
func createConversation(ctx context.Context, sdk *pipeshub.Pipeshub, query string) (id, title string, err error) {
	res, err := sdk.Conversations.CreateConversation(ctx, components.CreateConversationRequest{
		Query:    query,
		ChatMode: pipeshub.Pointer(components.CreateConversationRequestChatModeInternalSearch),
	})
	if err != nil {
		return "", "", fmt.Errorf("create conversation: %w", err)
	}
	if res.CreateConversationResponse == nil {
		return "", "", fmt.Errorf("create conversation: empty response")
	}
	conv := res.CreateConversationResponse.Conversation
	if conv.ID == nil || *conv.ID == "" {
		return "", "", fmt.Errorf("create conversation: response has no _id")
	}
	return *conv.ID, ptr(conv.Title), nil
}

func ptr(s *string) string {
	if s == nil {
		return ""
	}
	return *s
}
