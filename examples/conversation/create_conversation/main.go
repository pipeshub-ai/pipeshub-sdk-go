// Create a conversation with the non-streaming CreateConversation call, print
// the persisted answer, then delete the conversation.
//
// CreateConversation returns only once the whole answer is ready, which can
// take a minute or more; auth.NewClient uses a 5-minute HTTP timeout.
//
// Usage (from examples/):
//
//	go run ./conversation/create_conversation <path-to-.env>
package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/pipeshub-ai/pipeshub-sdk-go/models/components"

	"enterprise_search/auth"
)

const query = "In one short paragraph, what kinds of documents are in my workspace?"

func main() {
	if len(os.Args) < 2 {
		log.Fatal("usage: go run ./conversation/create_conversation <path-to-.env>")
	}
	if err := godotenv.Load(os.Args[1]); err != nil {
		log.Fatalf("load .env: %v", err)
	}

	sdk, err := auth.NewOAuthClient()
	if err != nil {
		log.Fatal(err)
	}

	ctx := context.Background()

	chatMode := components.CreateConversationRequestChatModeInternalSearch
	res, err := sdk.Conversations.CreateConversation(ctx, components.CreateConversationRequest{
		Query:    query,
		ChatMode: &chatMode,
	})
	if err != nil {
		log.Fatalf("create conversation: %v", err)
	}
	if res.CreateConversationResponse == nil {
		log.Fatal("create conversation: empty response")
	}
	conv := res.CreateConversationResponse.Conversation

	// The conversation id is the Mongo "_id" field, exposed as ID.
	convID := deref(conv.GetID())
	fmt.Printf("Conversation: %s\n", convID)
	fmt.Printf("Title:        %s\n", deref(conv.GetTitle()))
	if s := conv.GetStatus(); s != nil {
		fmt.Printf("Status:       %s\n", *s)
	}
	fmt.Printf("\nYou: %s\n\nBot: %s\n", query, lastBotAnswer(conv))

	// Clean up so repeated runs don't pile up conversations.
	if _, err := sdk.Conversations.DeleteConversationByID(ctx, convID); err != nil {
		log.Fatalf("delete conversation: %v", err)
	}
	fmt.Println("\nDeleted the conversation.")
}

// lastBotAnswer returns the content of the newest bot_response message.
func lastBotAnswer(conv components.Conversation) string {
	msgs := conv.GetMessages()
	for i := len(msgs) - 1; i >= 0; i-- {
		if t := msgs[i].GetMessageType(); t != nil && *t == components.MessageMessageTypeBotResponse {
			return deref(msgs[i].GetContent())
		}
	}
	return "(no bot response)"
}

func deref(s *string) string {
	if s == nil {
		return ""
	}
	return *s
}
