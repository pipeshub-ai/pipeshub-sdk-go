// Create an agent conversation (streaming), then fetch it by id.
//
// Usage (from examples/):
//
//	go run ./agent_conversation/get_conversation_by_id <path-to-.env>
package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	pipeshub "github.com/pipeshub-ai/pipeshub-sdk-go"
	"github.com/pipeshub-ai/pipeshub-sdk-go/models/components"
	"github.com/pipeshub-ai/pipeshub-sdk-go/models/operations"

	"enterprise_search/agui"
	"enterprise_search/auth"
)

// First user message when creating the conversation.
const firstMessage = "Who moved the cheese?"

func main() {
	if len(os.Args) < 2 {
		log.Fatal("usage: go run ./agent_conversation/get_conversation_by_id <path-to-.env>")
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

	convID, err := createAgentConversation(ctx, sdk, agentKey, firstMessage, envFilters())
	if err != nil {
		log.Fatalf("create conversation: %v", err)
	}
	log.Printf("created conversation id: %s", convID)

	if err := printAgentConversationByID(ctx, sdk, agentKey, convID); err != nil {
		log.Fatalf("get conversation: %v", err)
	}
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

func createAgentConversation(ctx context.Context, sdk *pipeshub.Pipeshub, agentKey, query string, filters *components.Filters) (string, error) {
	res, err := sdk.Agents.StreamAgentConversation(ctx, agentKey, components.AgentStreamCreateConversationRequest{
		Query:   query,
		Filters: filters,
		// Scoped agent conversations accept only "quick".
		ChatMode: components.AgentStreamCreateConversationRequestChatModeQuick,
	})
	if err != nil {
		return "", fmt.Errorf("stream agent conversation: %w", err)
	}
	if res == nil || res.AgentStreamSSEEvent == nil {
		return "", fmt.Errorf("no SSE stream returned")
	}
	stream := res.AgentStreamSSEEvent
	defer stream.Close()

	// The conversation id arrives on the CUSTOM/conversation_created frame,
	// well before the answer, and is confirmed again by RUN_FINISHED.
	var c agui.Collector
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
	if c.ConversationID == "" {
		return "", fmt.Errorf("stream did not report a conversation id")
	}
	return c.ConversationID, nil
}

func printAgentConversationByID(ctx context.Context, sdk *pipeshub.Pipeshub, agentKey, conversationID string) error {
	res, err := sdk.Agents.GetAgentConversationByID(ctx, operations.GetAgentConversationByIDRequest{
		AgentKey:       agentKey,
		ConversationID: conversationID,
	})
	if err != nil {
		return err
	}
	conv := res.AgentConversationDetailResponse.GetConversation()
	fmt.Printf("\n--- conversation by id: %s ---\n", conv.GetID())
	if title := conv.GetTitle(); title != nil {
		fmt.Printf("title: %q\n", *title)
	}
	fmt.Printf("messages: %d\n", len(conv.GetMessages()))
	for i, m := range conv.GetMessages() {
		msgType := ""
		if t := m.GetMessageType(); t != nil {
			msgType = string(*t)
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
