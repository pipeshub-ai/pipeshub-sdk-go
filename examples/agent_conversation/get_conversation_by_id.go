// Create an agent conversation (streaming), then fetch it by id.
//
// Uses github.com/pipeshub-ai/pipeshub-sdk-go.
//
// Usage (from examples/):
//
//	go run ./agent_conversation/get_conversation_by_id.go <path-to-.env>
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
	"github.com/pipeshub-ai/pipeshub-sdk-go/models/operations"

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
)

func main() {
	if len(os.Args) < 2 {
		log.Fatal("usage: go run ./agent_conversation/get_conversation_by_id.go <path-to-.env>")
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
	log.Printf("created conversation id: %s", convID)

	if err := printAgentConversationByID(ctx, sdk, convID); err != nil {
		log.Fatalf("get conversation: %v", err)
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

	var convID string
	for stream.Next() {
		ev := stream.Value()
		if ev == nil || ev.Event == nil || ev.Data == nil {
			continue
		}
		switch *ev.Event {
		case components.AgentStreamSSEEventEventConnected:
			var payload struct {
				ConversationID string `json:"conversationId"`
			}
			if err := json.Unmarshal([]byte(*ev.Data), &payload); err == nil {
				convID = payload.ConversationID
			}
		case components.AgentStreamSSEEventEventComplete:
			_, id, err := decodeAgentComplete(*ev.Data)
			if err != nil {
				return "", err
			}
			if id != "" {
				convID = id
			}
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

func printAgentConversationByID(ctx context.Context, sdk *pipeshub.Pipeshub, conversationID string) error {
	res, err := sdk.Agents.GetAgentConversationByID(ctx, operations.GetAgentConversationByIDRequest{
		AgentKey:       agentKey,
		ConversationID: conversationID,
	})
	if err != nil {
		return err
	}
	conv := res.AgentConversationDetailResponse.GetConversation()
	id := conv.GetID()
	fmt.Printf("\n--- conversation by id: %s ---\n", id)
	if conv.GetTitle() != nil {
		fmt.Printf("title: %q\n", *conv.GetTitle())
	}
	fmt.Printf("messages: %d\n", len(conv.GetMessages()))
	for i, m := range conv.GetMessages() {
		msgType := ""
		if m.GetMessageType() != nil {
			msgType = string(*m.GetMessageType())
		}
		content := ""
		if c := m.GetContent(); c != nil {
			content = *c
		}
		fmt.Printf("\n--- message %d [%s] ---\n", i+1, msgType)
		fmt.Println(content)
	}
	fmt.Println()

	return nil
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
	return "", payload.Conversation.ID, nil
}
