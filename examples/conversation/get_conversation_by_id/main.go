// Create a conversation, fetch it back with GetConversationByID and print its
// messages, then delete it.
//
// Usage (from examples/):
//
//	go run ./conversation/get_conversation_by_id <path-to-.env>
package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/joho/godotenv"
	pipeshub "github.com/pipeshub-ai/pipeshub-sdk-go"
	"github.com/pipeshub-ai/pipeshub-sdk-go/models/components"
	"github.com/pipeshub-ai/pipeshub-sdk-go/models/operations"

	"enterprise_search/auth"
)

const query = "What is the most recent news story about?"

func main() {
	if len(os.Args) < 2 {
		log.Fatal("usage: go run ./conversation/get_conversation_by_id <path-to-.env>")
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

	// Messages are paginated; sort oldest first to read them in order.
	res, err := sdk.Conversations.GetConversationByID(ctx, operations.GetConversationByIDRequest{
		ConversationID: convID,
		Limit:          pipeshub.Pointer(int64(20)),
		SortOrder:      pipeshub.Pointer(operations.GetConversationByIDSortOrderEnumAsc),
	})
	if err != nil {
		log.Fatalf("get conversation: %v", err)
	}
	if res.Object == nil || res.Object.Conversation == nil {
		log.Fatal("get conversation: empty response")
	}
	conv := res.Object.Conversation

	// Unlike the create response, this view names the id "id", not "_id".
	fmt.Printf("Conversation: %s\n", ptr(conv.ID))
	fmt.Printf("Title:        %s\n", ptr(conv.Title))
	if conv.Status != nil {
		fmt.Printf("Status:       %s\n", *conv.Status)
	}
	fmt.Printf("Messages:     %d\n\n", len(conv.Messages))
	for _, m := range conv.Messages {
		kind := ""
		if m.MessageType != nil {
			kind = string(*m.MessageType)
		}
		fmt.Printf("[%s] %s\n  %s\n\n", kind, ptr(m.ID), oneLine(ptr(m.Content), 200))
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

// oneLine flattens s and truncates it to at most n runes.
func oneLine(s string, n int) string {
	s = strings.Join(strings.Fields(s), " ")
	if r := []rune(s); len(r) > n {
		return string(r[:n]) + "…"
	}
	return s
}

func ptr(s *string) string {
	if s == nil {
		return ""
	}
	return *s
}
