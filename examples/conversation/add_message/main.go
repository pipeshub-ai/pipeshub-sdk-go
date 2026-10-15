// Create a conversation, then ask a follow-up with the non-streaming
// AddMessage call, print the answer, and delete the conversation.
//
// Usage (from examples/):
//
//	go run ./conversation/add_message <path-to-.env>
package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	pipeshub "github.com/pipeshub-ai/pipeshub-sdk-go"
	"github.com/pipeshub-ai/pipeshub-sdk-go/models/components"

	"enterprise_search/auth"
)

const (
	firstQuery = "What topics do the latest news articles cover?"
	followUp   = "Pick one of those topics and summarize it in two sentences."
)

func main() {
	if len(os.Args) < 2 {
		log.Fatal("usage: go run ./conversation/add_message <path-to-.env>")
	}
	if err := godotenv.Load(os.Args[1]); err != nil {
		log.Fatalf("load .env: %v", err)
	}

	sdk, err := auth.NewOAuthClient()
	if err != nil {
		log.Fatal(err)
	}

	ctx := context.Background()

	convID, err := createConversation(ctx, sdk, firstQuery)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Created conversation: %s\n", convID)

	// The follow-up should use the same chatMode the conversation was opened with.
	res, err := sdk.Conversations.AddMessage(ctx, convID, components.AddMessageRequest{
		Query:    followUp,
		ChatMode: pipeshub.Pointer(components.AddMessageRequestChatModeInternalSearch),
	})
	if err != nil {
		log.Fatalf("add message: %v", err)
	}
	if res.AddMessageResponse == nil {
		log.Fatal("add message: empty response")
	}
	body := res.AddMessageResponse

	fmt.Printf("\nYou: %s\n\nBot: %s\n", followUp, lastBotAnswer(body.Conversation))
	fmt.Printf("\nMessages in conversation: %d, records used: %d\n", len(body.Conversation.GetMessages()), body.RecordsUsed)

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

// lastBotAnswer returns the content of the newest bot_response message.
func lastBotAnswer(conv components.Conversation) string {
	msgs := conv.GetMessages()
	for i := len(msgs) - 1; i >= 0; i-- {
		if t := msgs[i].GetMessageType(); t != nil && *t == components.MessageMessageTypeBotResponse {
			if c := msgs[i].GetContent(); c != nil {
				return *c
			}
		}
	}
	return "(no bot response)"
}
