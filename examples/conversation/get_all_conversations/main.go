// Create a conversation, list your conversations newest first with
// GetAllConversations, then delete the conversation again.
//
// Usage (from examples/):
//
//	go run ./conversation/get_all_conversations <path-to-.env>
package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	pipeshub "github.com/pipeshub-ai/pipeshub-sdk-go"
	"github.com/pipeshub-ai/pipeshub-sdk-go/models/components"
	"github.com/pipeshub-ai/pipeshub-sdk-go/models/operations"

	"enterprise_search/auth"
)

const query = "Name one recent headline from the news feed."

func main() {
	if len(os.Args) < 2 {
		log.Fatal("usage: go run ./conversation/get_all_conversations <path-to-.env>")
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
	fmt.Printf("Created conversation: %s\n\n", convID)

	res, err := sdk.Conversations.GetAllConversations(ctx, operations.GetAllConversationsRequest{
		Source:    pipeshub.Pointer(operations.QueryParamSourceOwned),
		Page:      pipeshub.Pointer(int64(1)),
		Limit:     pipeshub.Pointer(int64(10)),
		SortBy:    pipeshub.Pointer(operations.GetAllConversationsSortByEnumLastActivityAt),
		SortOrder: pipeshub.Pointer(operations.GetAllConversationsSortOrderEnumDesc),
	})
	if err != nil {
		log.Fatalf("get all conversations: %v", err)
	}
	if res.Object == nil {
		log.Fatal("get all conversations: empty response")
	}

	p := res.Object.Pagination
	fmt.Printf("Page %d of %d (%d conversations in total)\n", deref(p.Page), deref(p.TotalPages), deref(p.TotalCount))
	for _, c := range res.Object.Conversations {
		marker := " "
		if c.ID != nil && *c.ID == convID {
			marker = "*"
		}
		status := ""
		if c.Status != nil {
			status = string(*c.Status)
		}
		fmt.Printf("%s %s  %-10s  %s\n", marker, ptr(c.ID), status, ptr(c.Title))
	}
	fmt.Println("\n* = the conversation this example created")

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

func deref(n *int64) int64 {
	if n == nil {
		return 0
	}
	return *n
}

func ptr(s *string) string {
	if s == nil {
		return ""
	}
	return *s
}
