// Create an agent with DuckDuckGo web search, start a streaming conversation, then append a follow-up (streaming).
//
// Usage (from examples/):
//
//	go run ./agent_conversation/create_and_add_message <path-to-.env>
//
// This example creates its own agent, so it needs no PIPESHUB_AGENT_KEY.
package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	pipeshub "github.com/pipeshub-ai/pipeshub-sdk-go"
	"github.com/pipeshub-ai/pipeshub-sdk-go/models/components"
	"github.com/pipeshub-ai/pipeshub-sdk-go/optionalnullable"

	"enterprise_search/agui"
	"enterprise_search/auth"
)

const (
	firstMessage    = "Who moved the cheese?"
	followUpMessage = "Can you give me more details on that?"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatal("usage: go run ./agent_conversation/create_and_add_message <path-to-.env>")
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

	filters := envFilters()

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

	fmt.Printf("You: %s\n\nBot: ", query)

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

	if c.ConversationID == "" {
		return "", fmt.Errorf("stream did not report a conversation id")
	}
	return c.ConversationID, nil
}

func addAgentConversationMessage(ctx context.Context, sdk *pipeshub.Pipeshub, agentKey, convID, query string, filters *components.Filters) error {
	res, err := sdk.Agents.StreamAgentConversationMessage(ctx, agentKey, convID, components.AgentAddMessageStreamRequest{
		Query:   query,
		Filters: filters,
		// Required on follow-ups too, and "quick" is the only accepted value.
		ChatMode: components.AgentAddMessageStreamRequestChatModeQuick,
	})
	if err != nil {
		return fmt.Errorf("stream agent conversation message: %w", err)
	}
	if res == nil || res.AgentMessageStreamSSEEvent == nil {
		return fmt.Errorf("no SSE stream returned")
	}
	stream := res.AgentMessageStreamSSEEvent
	defer stream.Close()

	fmt.Printf("\nYou: %s\n\nBot: ", query)

	c := agui.Collector{Echo: os.Stdout}
	for stream.Next() {
		ev := stream.Value()
		if ev == nil || ev.Event == nil || ev.Data == nil {
			continue
		}
		done, err := c.Handle(string(*ev.Event), *ev.Data)
		if err != nil {
			return err
		}
		if done {
			break
		}
	}
	if err := stream.Err(); err != nil {
		return fmt.Errorf("stream: %w", err)
	}
	if !c.Done {
		return fmt.Errorf("stream ended without RUN_FINISHED")
	}
	fmt.Println()

	return nil
}
