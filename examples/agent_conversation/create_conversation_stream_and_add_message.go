// Create an agent conversation (streaming), then append a follow-up message (streaming).
//
// Uses github.com/pipeshub-ai/pipeshub-sdk-go.
//
// Usage (from examples/):
//
//	go run ./agent_conversation/create_conversation_stream_and_add_message.go <path-to-.env>
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

	// Follow-up user message appended to the same conversation.
	followUpMessage = "Can you give me more details on that?"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatal("usage: go run ./agent_conversation/create_conversation_stream_and_add_message.go <path-to-.env>")
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

	convID, err := createAgentConversation(ctx, sdk, firstMessage, filters)
	if err != nil {
		log.Fatalf("create conversation: %v", err)
	}
	log.Printf("conversation id: %s", convID)

	if err := addAgentConversationMessage(ctx, sdk, convID, followUpMessage, filters); err != nil {
		log.Fatalf("add message: %v", err)
	}
}

func createAgentConversation(ctx context.Context, sdk *pipeshub.Pipeshub, query string, filters *components.Filters) (string, error) {
	chatMode := components.AgentStreamCreateConversationRequestChatModeAuto

	res, err := sdk.Agents.StreamAgentConversation(ctx, agentKey, components.AgentStreamCreateConversationRequest{
		Query:    query,
		Filters:  filters,
		ChatMode: chatMode.ToPointer(),
	})
	if err != nil {
		return "", fmt.Errorf("stream agent conversation: %w", err)
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
			answer, convID, err := decodeAgentComplete(*ev.Data)
			if err != nil {
				return "", err
			}
			fmt.Println(answer)
			if convID == "" {
				return "", fmt.Errorf("complete event missing conversation id")
			}
			return convID, nil
		case components.AgentStreamSSEEventEventError:
			return "", fmt.Errorf("stream error: %s", *ev.Data)
		}
	}
	if err := stream.Err(); err != nil {
		return "", fmt.Errorf("stream: %w", err)
	}
	return "", fmt.Errorf("stream ended without complete event")
}

func addAgentConversationMessage(ctx context.Context, sdk *pipeshub.Pipeshub, convID, query string, filters *components.Filters) error {
	chatMode := components.AgentAddMessageStreamRequestChatModeAuto

	res, err := sdk.Agents.StreamAgentConversationMessage(ctx, agentKey, convID, components.AgentAddMessageStreamRequest{
		Query:    query,
		Filters:  filters,
		ChatMode: chatMode.ToPointer(),
	})
	if err != nil {
		return fmt.Errorf("stream agent conversation message: %w", err)
	}
	stream := res.AgentMessageStreamSSEEvent
	defer stream.Close()

	fmt.Printf("\nYou: %s\n\nBot: ", query)

	for stream.Next() {
		ev := stream.Value()
		if ev == nil || ev.Event == nil || ev.Data == nil {
			continue
		}
		switch *ev.Event {
		case components.AgentMessageStreamSSEEventEventComplete:
			answer, _, err := decodeAgentComplete(*ev.Data)
			if err != nil {
				return err
			}
			fmt.Println(answer)
			return nil
		case components.AgentMessageStreamSSEEventEventError:
			return fmt.Errorf("stream error: %s", *ev.Data)
		}
	}
	if err := stream.Err(); err != nil {
		return fmt.Errorf("stream: %w", err)
	}
	return fmt.Errorf("stream ended without complete event")
}

func decodeAgentComplete(data string) (answer, convID string, err error) {
	var payload struct {
		Conversation struct {
			ID       string `json:"_id"`
			Messages []struct {
				MessageType string `json:"messageType"`
				Content     string `json:"content"`
			} `json:"messages"`
		} `json:"conversation"`
	}
	if err := json.Unmarshal([]byte(data), &payload); err != nil {
		return "", "", fmt.Errorf("decode complete: %w", err)
	}
	for i := len(payload.Conversation.Messages) - 1; i >= 0; i-- {
		m := payload.Conversation.Messages[i]
		if m.MessageType == "bot_response" {
			return m.Content, payload.Conversation.ID, nil
		}
	}
	return "", "", fmt.Errorf("no bot response in complete event")
}
