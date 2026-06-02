// Create an agent conversation (streaming), then regenerate the last bot response (streaming).
//
// Uses github.com/pipeshub-ai/pipeshub-sdk-go.
//
// Regenerate constraints (server-enforced):
//   - Only the last message in the conversation can be regenerated.
//   - The target message must be of type bot_response.
//
// Usage (from examples/):
//
//	go run ./agent_conversation/regenerate_message_stream.go <path-to-.env>
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

	"enterprise_search/auth"
)

const (
	// Stable key for the agent that owns the conversation.
	agentKey = "02a7d998-d21b-4015-aaf7-5cda765c1012"

	// Knowledge-base / record-group id (Filters.Kb).
	kbID = "8747da12-4724-4a95-ac92-827b88d79647"

	// Connector instance id (Filters.Apps).
	connectorID = "aeab9ddc-fb9b-47c8-ad98-bd4744e19555"

	// First user message when creating the conversation.
	firstMessage = "What are some latest tech news?"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatal("usage: go run ./agent_conversation/regenerate_message_stream.go <path-to-.env>")
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

	convID, botMessageID, originalAnswer, err := createAgentConversation(ctx, sdk, firstMessage, filters)
	if err != nil {
		log.Fatalf("create conversation: %v", err)
	}
	log.Printf("conversation id: %s", convID)
	log.Printf("last bot message id: %s", botMessageID)
	fmt.Printf("\nOriginal bot response (%d chars):\n%s\n", len(originalAnswer), summarize(originalAnswer, 400))

	regenerated, err := regenerateAgentConversationMessage(ctx, sdk, convID, botMessageID, filters)
	if err != nil {
		log.Fatalf("regenerate message: %v", err)
	}
	fmt.Printf("\nRegenerated bot response (%d chars):\n%s\n", len(regenerated), summarize(regenerated, 400))
}

func createAgentConversation(ctx context.Context, sdk *pipeshub.Pipeshub, query string, filters *components.Filters) (convID, botMessageID, answer string, err error) {
	chatMode := components.AgentStreamCreateConversationRequestChatModeAuto

	res, err := sdk.Agents.StreamAgentConversation(ctx, agentKey, components.AgentStreamCreateConversationRequest{
		Query:    query,
		Filters:  filters,
		ChatMode: chatMode.ToPointer(),
	})
	if err != nil {
		return "", "", "", fmt.Errorf("stream agent conversation: %w", err)
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
			answer, convID, botMessageID, err = decodeConversationComplete(*ev.Data)
			if err != nil {
				return "", "", "", err
			}
			fmt.Println(answer)
			if convID == "" {
				return "", "", "", fmt.Errorf("complete event missing conversation id")
			}
			if botMessageID == "" {
				return "", "", "", fmt.Errorf("complete event missing bot message id")
			}
			return convID, botMessageID, answer, nil
		case components.AgentStreamSSEEventEventError:
			return "", "", "", fmt.Errorf("stream error: %s", *ev.Data)
		}
	}
	if err := stream.Err(); err != nil {
		return "", "", "", fmt.Errorf("stream: %w", err)
	}
	return "", "", "", fmt.Errorf("stream ended without complete event")
}

func regenerateAgentConversationMessage(ctx context.Context, sdk *pipeshub.Pipeshub, convID, messageID string, filters *components.Filters) (string, error) {
	res, err := sdk.Agents.RegenerateAgentConversationMessage(ctx, agentKey, convID, messageID, &components.RegenerateRequest{
		Filters: filters,
	})
	if err != nil {
		return "", fmt.Errorf("regenerate agent conversation message: %w", err)
	}
	stream := res.AgentRegenerateSSEEvent
	defer stream.Close()

	fmt.Printf("\nRegenerating message %s ...\n\nBot: ", messageID)

	var accumulated string
	for stream.Next() {
		ev := stream.Value()
		if ev == nil || ev.Event == nil || ev.Data == nil {
			continue
		}
		switch *ev.Event {
		case components.AgentRegenerateSSEEventEventAnswerChunk:
			var chunk struct {
				Accumulated string `json:"accumulated"`
			}
			if err := json.Unmarshal([]byte(*ev.Data), &chunk); err == nil && chunk.Accumulated != "" {
				accumulated = chunk.Accumulated
			}
		case components.AgentRegenerateSSEEventEventComplete:
			answer, _, _, err := decodeConversationComplete(*ev.Data)
			if err != nil {
				return "", err
			}
			if answer == "" && accumulated != "" {
				answer = accumulated
			}
			fmt.Println(answer)
			return answer, nil
		case components.AgentRegenerateSSEEventEventError:
			return "", fmt.Errorf("stream error: %s", *ev.Data)
		}
	}
	if err := stream.Err(); err != nil {
		return "", fmt.Errorf("stream: %w", err)
	}
	return "", fmt.Errorf("stream ended without complete event")
}

func decodeConversationComplete(data string) (answer, convID, botMessageID string, err error) {
	var payload struct {
		Conversation struct {
			ID       string `json:"_id"`
			Messages []struct {
				ID          string `json:"_id"`
				MessageType string `json:"messageType"`
				Content     string `json:"content"`
			} `json:"messages"`
		} `json:"conversation"`
	}
	if err := json.Unmarshal([]byte(data), &payload); err != nil {
		return "", "", "", fmt.Errorf("decode complete: %w", err)
	}
	convID = payload.Conversation.ID
	for i := len(payload.Conversation.Messages) - 1; i >= 0; i-- {
		m := payload.Conversation.Messages[i]
		if m.MessageType == "bot_response" {
			return m.Content, convID, m.ID, nil
		}
	}
	return "", convID, "", fmt.Errorf("no bot_response in complete event")
}

func summarize(s string, max int) string {
	s = strings.TrimSpace(s)
	if max <= 0 || len(s) <= max {
		return s
	}
	return s[:max] + "..."
}
