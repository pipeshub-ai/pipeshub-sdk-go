// Rate a bot answer (a bot_response message, not the user's question).
// Needs PIPESHUB_BASE_URL, PIPESHUB_CLIENT_ID, PIPESHUB_CLIENT_SECRET, PIPESHUB_CONVERSATION_ID, PIPESHUB_MESSAGE_ID.
//
// Usage (from examples/): go run ./conversation/update_message_feedback <path-to-.env>
package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/joho/godotenv"
	pipeshub "github.com/pipeshub-ai/pipeshub-sdk-go"
	"github.com/pipeshub-ai/pipeshub-sdk-go/models/components"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatal("usage: go run ./conversation/update_message_feedback <path-to-.env>")
	}
	if err := godotenv.Load(os.Args[1]); err != nil {
		log.Fatal(err)
	}

	baseURL := strings.TrimRight(os.Getenv("PIPESHUB_BASE_URL"), "/") + "/api/v1"
	tokenURL := os.Getenv("PIPESHUB_TOKEN_URL") // optional
	if tokenURL == "" {
		tokenURL = baseURL + "/oauth2/token"
	}
	client := pipeshub.New(
		pipeshub.WithServerURL(baseURL),
		pipeshub.WithSecurity(components.Security{Oauth2: &components.SchemeOauth2{
			ClientID:     os.Getenv("PIPESHUB_CLIENT_ID"),
			ClientSecret: os.Getenv("PIPESHUB_CLIENT_SECRET"),
			TokenURL:     tokenURL,
		}}),
	)

	res, err := client.Conversations.UpdateMessageFeedback(context.Background(), os.Getenv("PIPESHUB_CONVERSATION_ID"), os.Getenv("PIPESHUB_MESSAGE_ID"), components.MessageFeedbackSubmitRequest{
		IsHelpful:  pipeshub.Pointer(true),
		Categories: []components.MessageFeedbackSubmitRequestCategory{components.MessageFeedbackSubmitRequestCategoryWellExplained},
	})
	if err != nil {
		log.Fatal(err)
	}
	out, _ := json.MarshalIndent(res.MessageFeedbackUpdateResponse, "", "  ")
	fmt.Println(string(out))
}
