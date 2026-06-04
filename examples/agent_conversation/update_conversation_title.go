// Create an agent conversation (streaming), then update its title.
//
// Uses github.com/pipeshub-ai/pipeshub-sdk-go.
//
// Usage (from examples/):
//
//	go run ./agent_conversation/update_conversation_title.go <path-to-.env>
package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	pipeshub "github.com/pipeshub-ai/pipeshub-sdk-go"
	"github.com/pipeshub-ai/pipeshub-sdk-go/models/components"

	"enterprise_search/auth"
)

const (
	// Stable key for the agent that owns the conversation.
	agentKey = "ddff45f7-e534-4726-92e8-5e8e6338ad41"

	// Knowledge-base / record-group id (Filters.Kb).
	kbID = "45d5aa5b-2b2c-408d-bcd3-ce4de6dfcd5b"

	// Connector instance id (Filters.Apps).
	connectorID = "270d4bac-234a-4c0d-963f-84f152cd21f0"

	// First user message when creating the conversation.
	firstMessage = "Who moved the cheese?"

	// Title applied via UpdateAgentConversationTitle.
	newTitle = "SDK example: updated title"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatal("usage: go run ./agent_conversation/update_conversation_title.go <path-to-.env>")
	}
	if err := godotenv.Load(os.Args[1]); err != nil {
		log.Fatalf("load .env: %v", err)
	}

	sdk, err := auth.NewClient(
		os.Getenv("PIPESHUB_TEST_USER_EMAIL"),
		os.Getenv("PIPESHUB_TEST_USER_PASSWORD"),
	)
	if err != nil {
		log.Fatal(err)
	}

	ctx := context.Background()
	filters := &components.Filters{
		Kb:   []string{kbID},
		Apps: []string{connectorID},
	}

	convID, oldTitle, err := createAgentConversation(ctx, sdk, firstMessage, filters)
	if err != nil {
		log.Fatalf("create conversation: %v", err)
	}
	log.Printf("conversation id: %s", convID)

	updatedTitle, err := updateAgentConversationTitle(ctx, sdk, convID, newTitle)
	if err != nil {
		log.Fatalf("update title: %v", err)
	}

	log.Printf("old title: %q", oldTitle)
	log.Printf("new title: %q", updatedTitle)
}

func createAgentConversation(ctx context.Context, sdk *pipeshub.Pipeshub, query string, filters *components.Filters) (convID, oldTitle string, err error) {
	chatMode := components.AgentStreamCreateConversationRequestChatModeAuto

	res, err := sdk.Agents.StreamAgentConversation(ctx, agentKey, components.AgentStreamCreateConversationRequest{
		Query:    query,
		Filters:  filters,
		ChatMode: chatMode.ToPointer(),
	})
	if err != nil {
		return "", "", fmt.Errorf("stream agent conversation: %w", err)
	}
	stream := res.AgentStreamSSEEvent
	defer stream.Close()

	fmt.Printf("You: %s\n\nBot: ", query)

	for stream.Next() {
		ev := stream.Value()
		if ev == nil || ev.Event == nil || ev.Data == nil {
			continue
		}
		switch *ev.Event {
		case components.AgentStreamSSEEventEventComplete:
			answer, id, title, err := decodeConversationComplete(*ev.Data)
			if err != nil {
				return "", "", err
			}
			fmt.Println(answer)
			if id == "" {
				return "", "", fmt.Errorf("complete event missing conversation id")
			}
			if title == "" {
				title = query
			}
			return id, title, nil
		case components.AgentStreamSSEEventEventError:
			return "", "", fmt.Errorf("stream error: %s", *ev.Data)
		}
	}
	if err := stream.Err(); err != nil {
		return "", "", fmt.Errorf("stream: %w", err)
	}
	return "", "", fmt.Errorf("stream ended without complete event")
}

func updateAgentConversationTitle(ctx context.Context, sdk *pipeshub.Pipeshub, convID, title string) (string, error) {
	res, err := sdk.Agents.UpdateAgentConversationTitle(ctx, agentKey, convID, components.ConversationTitleUpdateRequest{
		Title: title,
	})
	if err != nil {
		return "", fmt.Errorf("update agent conversation title: %w", err)
	}
	resp := res.AgentConversationTitleUpdateResponse
	conv := resp.GetConversation()
	updated := (&conv).GetTitle()
	if updated == nil {
		return "", fmt.Errorf("response missing conversation title")
	}
	return *updated, nil
}

func decodeConversationComplete(data string) (answer, convID, title string, err error) {
	var payload struct {
		Conversation struct {
			ID       string  `json:"_id"`
			Title    *string `json:"title"`
			Messages []struct {
				MessageType string `json:"messageType"`
				Content     string `json:"content"`
			} `json:"messages"`
		} `json:"conversation"`
	}
	if err := json.Unmarshal([]byte(data), &payload); err != nil {
		return "", "", "", fmt.Errorf("decode complete: %w", err)
	}
	convID = payload.Conversation.ID
	if payload.Conversation.Title != nil {
		title = *payload.Conversation.Title
	}
	for _, m := range payload.Conversation.Messages {
		if title == "" && m.MessageType == "user_query" {
			title = m.Content
		}
		if m.MessageType == "bot_response" {
			answer = m.Content
		}
	}
	if answer == "" {
		return "", convID, title, fmt.Errorf("no bot_response in complete event")
	}
	return answer, convID, title, nil
}
