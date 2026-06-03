// Create an agent with DuckDuckGo web search, start a streaming conversation, then append a follow-up (streaming).
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
	"time"

	"github.com/joho/godotenv"
	pipeshub "github.com/pipeshub-ai/pipeshub-sdk-go"
	"github.com/pipeshub-ai/pipeshub-sdk-go/models/components"
	"github.com/pipeshub-ai/pipeshub-sdk-go/optionalnullable"

	"enterprise_search/auth"
)

const (
	kbID            = "45d5aa5b-2b2c-408d-bcd3-ce4de6dfcd5b"
	connectorID     = "270d4bac-234a-4c0d-963f-84f152cd21f0"
	firstMessage    = "Who moved the cheese?"
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

	webSearch, err := resolveDuckDuckGoWebSearch(ctx, sdk)
	if err != nil {
		log.Fatalf("resolve DuckDuckGo web search: %v", err)
	}
	fmt.Printf("resolved web search provider: %s\n", webSearch.Provider)

	agentKey, err := createAgentWithWebSearch(ctx, sdk, webSearch)
	if err != nil {
		log.Fatalf("create agent: %v", err)
	}
	log.Printf("agent key: %s", agentKey)

	filters := &components.Filters{
		Kb:   []string{kbID},
		Apps: []string{connectorID},
	}

	convID, err := createAgentConversation(ctx, sdk, agentKey, firstMessage, filters)
	if err != nil {
		log.Fatalf("create conversation: %v", err)
	}
	log.Printf("conversation id: %s", convID)

	if err := addAgentConversationMessage(ctx, sdk, agentKey, convID, followUpMessage, filters); err != nil {
		log.Fatalf("add message: %v", err)
	}
}

func resolveDuckDuckGoWebSearch(ctx context.Context, sdk *pipeshub.Pipeshub) (components.AgentCreateWebSearch, error) {
	res, err := sdk.WebSearch.GetWebSearchProviders(ctx)
	if err != nil {
		return components.AgentCreateWebSearch{}, fmt.Errorf("get web search providers: %w", err)
	}
	if res == nil || res.WebSearchProvidersResponse == nil {
		return components.AgentCreateWebSearch{}, fmt.Errorf("get web search providers: empty response")
	}

	for _, item := range res.WebSearchProvidersResponse.GetProviders() {
		if item.Provider != components.WebSearchProviderTypeDuckduckgo {
			continue
		}
		ws := components.AgentCreateWebSearch{
			Provider: string(components.WebSearchProviderTypeDuckduckgo),
		}
		if key := item.GetProviderKey(); key != "" {
			ws.ProviderKey = pipeshub.Pointer(key)
		}
		return ws, nil
	}

	return components.AgentCreateWebSearch{
		Provider: string(components.WebSearchProviderTypeDuckduckgo),
	}, nil
}

func createAgentWithWebSearch(ctx context.Context, sdk *pipeshub.Pipeshub, webSearch components.AgentCreateWebSearch) (string, error) {
	modelKey, err := firstReasoningModelKey(ctx, sdk)
	if err != nil {
		return "", err
	}

	isReasoning := true
	webSearchUnion := components.CreateAgentCreateWebSearchUnionAgentCreateWebSearch(webSearch)

	res, err := sdk.Agents.CreateAgent(ctx, components.AgentCreateRequest{
		Name: fmt.Sprintf("SDK example %d", time.Now().Unix()),
		Models: []components.AgentCreateModelEntryUnion{
			components.CreateAgentCreateModelEntryUnionAgentCreateModelEntry(components.AgentCreateModelEntry{
				ModelKey:    modelKey,
				IsReasoning: &isReasoning,
			}),
		},
		WebSearch: optionalnullable.From(&webSearchUnion),
	})
	if err != nil {
		return "", fmt.Errorf("create agent: %w", err)
	}
	if res == nil || res.AgentCreateResponse == nil {
		return "", fmt.Errorf("create agent: empty response")
	}

	agent := res.AgentCreateResponse.GetAgent()
	key := agent.GetKey()
	if key == "" {
		return "", fmt.Errorf("create agent: response missing agent key")
	}
	return key, nil
}

func firstReasoningModelKey(ctx context.Context, sdk *pipeshub.Pipeshub) (string, error) {
	if key := os.Getenv("PIPESHUB_AGENT_MODEL_KEY"); key != "" {
		return key, nil
	}

	res, err := sdk.AIModelsProviders.GetAvailableModelsByType(ctx, components.ModelTypeLlm)
	if err != nil {
		return "", fmt.Errorf("list LLM models: %w", err)
	}
	body := res.GetObject()
	if body == nil {
		return "", fmt.Errorf("list LLM models: empty response")
	}

	for _, m := range body.GetModels() {
		if m.GetIsReasoning() && m.GetModelKey() != "" {
			return m.GetModelKey(), nil
		}
	}
	models := body.GetModels()
	if len(models) > 0 && models[0].GetModelKey() != "" {
		return models[0].GetModelKey(), nil
	}

	return "", fmt.Errorf("no LLM model configured; set PIPESHUB_AGENT_MODEL_KEY in .env")
}

func createAgentConversation(ctx context.Context, sdk *pipeshub.Pipeshub, agentKey, query string, filters *components.Filters) (string, error) {
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

func addAgentConversationMessage(ctx context.Context, sdk *pipeshub.Pipeshub, agentKey, convID, query string, filters *components.Filters) error {
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
