// Create a conversation with a first question (non-streaming).
// Needs PIPESHUB_BASE_URL, PIPESHUB_CLIENT_ID, PIPESHUB_CLIENT_SECRET.
//
// Usage (from examples/): go run ./conversation/create_conversation <path-to-.env>
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
		log.Fatal("usage: go run ./conversation/create_conversation <path-to-.env>")
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

	res, err := client.Conversations.CreateConversation(context.Background(), components.CreateConversationRequest{
		Query:    "What topics do the latest news articles cover?",
		ChatMode: pipeshub.Pointer(components.CreateConversationRequestChatModeInternalSearch),
	})
	if err != nil {
		log.Fatal(err)
	}
	out, _ := json.MarshalIndent(res.CreateConversationResponse, "", "  ")
	fmt.Println(string(out))
}
