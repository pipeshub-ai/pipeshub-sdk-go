// Create an agent conversation (streaming), then archive and unarchive it.
//
// Uses github.com/pipeshub-ai/pipeshub-sdk-go.
//
// Usage (from examples/):
//
//	go run ./agent_conversation/archive_unarchive.go <path-to-.env>
package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"

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
)

func main() {
	if len(os.Args) < 2 {
		log.Fatal("usage: go run ./agent_conversation/archive_unarchive.go <path-to-.env>")
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

	convID, title, err := createAgentConversation(ctx, sdk, firstMessage, filters)
	if err != nil {
		log.Fatalf("create conversation: %v", err)
	}
	if title == "" {
		title = firstMessage
	}

	fmt.Printf("Created conversation: %s\n", convID)
	fmt.Printf("Title: %q\n", title)

	if err := archiveAgentConversation(ctx, sdk, convID); err != nil {
		log.Fatalf("archive conversation: %v", err)
	}
	if err := unarchiveAgentConversation(ctx, sdk, convID); err != nil {
		log.Fatalf("unarchive conversation: %v", err)
	}
}

func createAgentConversation(ctx context.Context, sdk *pipeshub.Pipeshub, query string, filters *components.Filters) (convID, title string, err error) {
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

	for stream.Next() {
		ev := stream.Value()
		if ev == nil || ev.Event == nil || ev.Data == nil {
			continue
		}
		switch *ev.Event {
		case components.AgentStreamSSEEventEventComplete:
			id, t, err := decodeConversationComplete(*ev.Data)
			if err != nil {
				return "", "", err
			}
			if id == "" {
				return "", "", fmt.Errorf("complete event missing conversation id")
			}
			return id, t, nil
		case components.AgentStreamSSEEventEventError:
			return "", "", fmt.Errorf("stream error: %s", *ev.Data)
		}
	}
	if err := stream.Err(); err != nil {
		return "", "", fmt.Errorf("stream: %w", err)
	}
	return "", "", fmt.Errorf("stream ended without complete event")
}

func archiveAgentConversation(ctx context.Context, sdk *pipeshub.Pipeshub, convID string) error {
	res, err := sdk.Agents.ArchiveAgentConversation(ctx, agentKey, convID)
	if err != nil {
		return fmt.Errorf("archive agent conversation: %w", err)
	}
	body := res.AgentConversationArchiveResponse
	at := body.GetArchivedAt()
	if at.IsZero() {
		fmt.Println("Archived (by you): conversation is now in archives")
		return nil
	}
	fmt.Printf("Archived (by you at %s): conversation is now in archives\n", at.Format(time.RFC1123))
	return nil
}

func unarchiveAgentConversation(ctx context.Context, sdk *pipeshub.Pipeshub, convID string) error {
	res, err := sdk.Agents.UnarchiveAgentConversation(ctx, agentKey, convID)
	if err != nil {
		return fmt.Errorf("unarchive agent conversation: %w", err)
	}
	body := res.AgentConversationUnarchiveResponse
	at := body.GetUnarchivedAt()
	if at.IsZero() {
		fmt.Println("Unarchived: conversation is back in your active list")
		return nil
	}
	fmt.Printf("Unarchived (at %s): conversation is back in your active list\n", at.Format(time.RFC1123))
	return nil
}

func decodeConversationComplete(data string) (convID, title string, err error) {
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
		return "", "", fmt.Errorf("decode complete: %w", err)
	}
	convID = payload.Conversation.ID
	if payload.Conversation.Title != nil {
		title = *payload.Conversation.Title
	}
	for _, m := range payload.Conversation.Messages {
		if title == "" && m.MessageType == "user_query" {
			title = m.Content
		}
	}
	return convID, title, nil
}
