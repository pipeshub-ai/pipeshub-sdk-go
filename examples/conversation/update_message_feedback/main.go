// Create a conversation, then submit feedback on the bot's answer with
// UpdateMessageFeedback, and delete the conversation.
//
// Usage (from examples/):
//
//	go run ./conversation/update_message_feedback <path-to-.env>
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

const query = "Summarize one news story about public safety."

func main() {
	if len(os.Args) < 2 {
		log.Fatal("usage: go run ./conversation/update_message_feedback <path-to-.env>")
	}
	if err := godotenv.Load(os.Args[1]); err != nil {
		log.Fatalf("load .env: %v", err)
	}

	sdk, err := auth.NewOAuthClient()
	if err != nil {
		log.Fatal(err)
	}

	ctx := context.Background()

	conv, err := createConversation(ctx, sdk, query)
	if err != nil {
		log.Fatal(err)
	}
	convID := *conv.ID

	// Feedback goes on a bot_response message, not on the user's query.
	messageID := lastBotMessageID(conv)
	if messageID == "" {
		log.Fatal("conversation has no bot_response message to rate")
	}
	fmt.Printf("Conversation: %s\nBot message:  %s\n\n", convID, messageID)

	res, err := sdk.Conversations.UpdateMessageFeedback(ctx, convID, messageID, components.MessageFeedbackSubmitRequest{
		IsHelpful: pipeshub.Pointer(true),
		Categories: []components.MessageFeedbackSubmitRequestCategory{
			components.MessageFeedbackSubmitRequestCategoryWellExplained,
			components.MessageFeedbackSubmitRequestCategoryHelpfulCitations,
		},
		Comments: &components.MessageFeedbackSubmitRequestComments{
			Positive: pipeshub.Pointer("Clear summary with useful citations."),
		},
	})
	if err != nil {
		log.Fatalf("update message feedback: %v", err)
	}
	if res.MessageFeedbackUpdateResponse == nil {
		log.Fatal("update message feedback: empty response")
	}
	fb := res.MessageFeedbackUpdateResponse.Feedback

	helpful := "unset"
	if fb.IsHelpful != nil {
		helpful = fmt.Sprint(*fb.IsHelpful)
	}
	fmt.Printf("Feedback recorded on message %s\n", res.MessageFeedbackUpdateResponse.MessageID)
	fmt.Printf("  helpful:    %s\n", helpful)
	fmt.Printf("  categories: %v\n", fb.Categories)
	fmt.Printf("  provider:   %s\n", fb.FeedbackProvider)
	fmt.Printf("  at:         %s\n", time.UnixMilli(fb.Timestamp).Format(time.RFC1123))

	if _, err := sdk.Conversations.DeleteConversationByID(ctx, convID); err != nil {
		log.Fatalf("delete conversation: %v", err)
	}
	fmt.Println("\nDeleted the conversation.")
}

// createConversation opens a conversation with the non-streaming
// CreateConversation call and returns it, with its messages.
func createConversation(ctx context.Context, sdk *pipeshub.Pipeshub, query string) (components.Conversation, error) {
	res, err := sdk.Conversations.CreateConversation(ctx, components.CreateConversationRequest{
		Query:    query,
		ChatMode: pipeshub.Pointer(components.CreateConversationRequestChatModeInternalSearch),
	})
	if err != nil {
		return components.Conversation{}, fmt.Errorf("create conversation: %w", err)
	}
	if res.CreateConversationResponse == nil {
		return components.Conversation{}, fmt.Errorf("create conversation: empty response")
	}
	conv := res.CreateConversationResponse.Conversation
	if conv.ID == nil || *conv.ID == "" {
		return components.Conversation{}, fmt.Errorf("create conversation: response has no _id")
	}
	return conv, nil
}

// lastBotMessageID returns the "_id" of the newest bot_response message.
func lastBotMessageID(conv components.Conversation) string {
	msgs := conv.GetMessages()
	for i := len(msgs) - 1; i >= 0; i-- {
		if t := msgs[i].GetMessageType(); t != nil && *t == components.MessageMessageTypeBotResponse && msgs[i].ID != nil {
			return *msgs[i].ID
		}
	}
	return ""
}
