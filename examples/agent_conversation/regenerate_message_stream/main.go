// Create an agent conversation (streaming), then regenerate the last bot
// response (streaming).
//
// Regenerate constraints (server-enforced):
//   - Only the last message in the conversation can be regenerated.
//   - The target message must be of type bot_response.
//   - The request body is mandatory and must carry chatMode.
//
// Usage (from examples/):
//
//	go run ./agent_conversation/regenerate_message_stream <path-to-.env>
package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	pipeshub "github.com/pipeshub-ai/pipeshub-sdk-go"
	"github.com/pipeshub-ai/pipeshub-sdk-go/models/components"

	"enterprise_search/agui"
	"enterprise_search/auth"
)

// First user message when creating the conversation.
const firstMessage = "Who moved the cheese?"

func main() {
	if len(os.Args) < 2 {
		log.Fatal("usage: go run ./agent_conversation/regenerate_message_stream <path-to-.env>")
	}
	if err := godotenv.Load(os.Args[1]); err != nil {
		log.Fatalf("load .env: %v", err)
	}

	agentKey := os.Getenv("PIPESHUB_AGENT_KEY")
	if agentKey == "" {
		log.Fatal("PIPESHUB_AGENT_KEY is required")
	}

	sdk, err := auth.NewClient(
		os.Getenv("PIPESHUB_TEST_USER_EMAIL"),
		os.Getenv("PIPESHUB_TEST_USER_PASSWORD"),
	)
	if err != nil {
		log.Fatal(err)
	}

	ctx := context.Background()
	filters := envFilters()

	convID, botMessageID, originalAnswer, err := createAgentConversation(ctx, sdk, agentKey, firstMessage, filters)
	if err != nil {
		log.Fatalf("create conversation: %v", err)
	}
	log.Printf("conversation id: %s", convID)
	log.Printf("last bot message id: %s", botMessageID)
	fmt.Printf("\nOriginal bot response (%d chars):\n%s\n", len(originalAnswer), originalAnswer)

	regenerated, err := regenerateAgentConversationMessage(ctx, sdk, agentKey, convID, botMessageID, filters)
	if err != nil {
		log.Fatalf("regenerate message: %v", err)
	}
	fmt.Printf("\nRegenerated bot response (%d chars):\n%s\n", len(regenerated), regenerated)
}

// envFilters builds knowledge filters from KB_ID and CONNECTOR_ID.
//
// It returns nil when neither is set, so the agent falls back to its stored
// knowledge config. An empty-but-non-nil Filters would instead be sent as
// {"apps":[],"kb":[]}, which means "no knowledge sources at all".
func envFilters() *components.Filters {
	kbID, connectorID := os.Getenv("KB_ID"), os.Getenv("CONNECTOR_ID")
	if kbID == "" && connectorID == "" {
		return nil
	}
	filters := &components.Filters{}
	if kbID != "" {
		filters.Kb = []string{kbID}
	}
	if connectorID != "" {
		filters.Apps = []string{connectorID}
	}
	return filters
}

func createAgentConversation(ctx context.Context, sdk *pipeshub.Pipeshub, agentKey, query string, filters *components.Filters) (convID, botMessageID, answer string, err error) {
	res, err := sdk.Agents.StreamAgentConversation(ctx, agentKey, components.AgentStreamCreateConversationRequest{
		Query:   query,
		Filters: filters,
		// Scoped agent conversations accept only "quick".
		ChatMode: components.AgentStreamCreateConversationRequestChatModeQuick,
	})
	if err != nil {
		return "", "", "", fmt.Errorf("stream agent conversation: %w", err)
	}
	if res == nil || res.AgentStreamSSEEvent == nil {
		return "", "", "", fmt.Errorf("no SSE stream returned")
	}
	stream := res.AgentStreamSSEEvent
	defer stream.Close()

	fmt.Printf("You: %s\n\nBot: ", query)

	c := agui.Collector{Echo: os.Stdout}
	for stream.Next() {
		ev := stream.Value()
		if ev == nil || ev.Event == nil || ev.Data == nil {
			continue
		}
		done, err := c.Handle(string(*ev.Event), *ev.Data)
		if err != nil {
			return "", "", "", err
		}
		if done {
			break
		}
	}
	if err := stream.Err(); err != nil {
		return "", "", "", fmt.Errorf("stream: %w", err)
	}
	if !c.Done {
		return "", "", "", fmt.Errorf("stream ended without RUN_FINISHED")
	}
	fmt.Println()

	if c.ConversationID == "" {
		return "", "", "", fmt.Errorf("stream did not report a conversation id")
	}
	msg, ok := c.Result.Conversation.LastBotMessage()
	if !ok {
		return "", "", "", fmt.Errorf("no bot response in RUN_FINISHED result")
	}
	if msg.ID == "" {
		return "", "", "", fmt.Errorf("bot response has no message id to regenerate")
	}
	return c.ConversationID, msg.ID, msg.Content, nil
}

func regenerateAgentConversationMessage(ctx context.Context, sdk *pipeshub.Pipeshub, agentKey, convID, messageID string, filters *components.Filters) (string, error) {
	res, err := sdk.Agents.RegenerateAgentConversationMessage(ctx, agentKey, convID, messageID, components.AgentRegenerateRequest{
		Filters:  filters,
		ChatMode: components.AgentRegenerateRequestChatModeQuick,
	})
	if err != nil {
		return "", fmt.Errorf("regenerate agent conversation message: %w", err)
	}
	if res == nil || res.AgentRegenerateSSEEvent == nil {
		return "", fmt.Errorf("no SSE stream returned")
	}
	stream := res.AgentRegenerateSSEEvent
	defer stream.Close()

	fmt.Printf("\nRegenerating message %s ...\n\nBot: ", messageID)

	c := agui.Collector{Echo: os.Stdout}
	for stream.Next() {
		ev := stream.Value()
		if ev == nil || ev.Event == nil || ev.Data == nil {
			continue
		}
		done, err := c.Handle(string(*ev.Event), *ev.Data)
		if err != nil {
			return "", err
		}
		if done {
			break
		}
	}
	if err := stream.Err(); err != nil {
		return "", fmt.Errorf("stream: %w", err)
	}
	if !c.Done {
		return "", fmt.Errorf("stream ended without RUN_FINISHED")
	}
	fmt.Println()

	// Answer falls back to the streamed tokens when the regenerated turn is not
	// yet reflected in the persisted conversation.
	return c.Answer(), nil
}
