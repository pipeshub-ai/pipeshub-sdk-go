// Create an agent with a resolved LLM model, then print the created agent.
//
// Uses github.com/pipeshub-ai/pipeshub-sdk-go.
//
// Usage (from examples/):
//
//	go run ./agent/create.go <path-to-.env>
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

func main() {
	if len(os.Args) < 2 {
		log.Fatal("usage: go run ./agent/create.go <path-to-.env>")
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

	modelKey, err := firstReasoningModelKey(ctx, sdk)
	if err != nil {
		log.Fatalf("resolve model key: %v", err)
	}

	description := "Agent created by the Go SDK example."
	systemPrompt := "You are a concise SDK example agent."
	startMessage := "Ask me anything about your workspace."
	shareWithOrg := false
	isReasoning := true

	modelUnion := components.CreateAgentCreateModelEntryUnionAgentCreateModelEntry(components.AgentCreateModelEntry{
		ModelKey:    modelKey,
		IsReasoning: pipeshub.Pointer(isReasoning),
	})

	res, err := sdk.Agents.CreateAgent(ctx, components.AgentCreateRequest{
		Name:         fmt.Sprintf("SDK example agent %d", time.Now().Unix()),
		Description:  &description,
		SystemPrompt: &systemPrompt,
		StartMessage: &startMessage,
		Models:       []components.AgentCreateModelEntryUnion{modelUnion},
		ShareWithOrg: &shareWithOrg,
	})
	if err != nil {
		log.Fatalf("create agent: %v", err)
	}

	createBody := res.GetAgentCreateResponse()
	if createBody == nil {
		log.Fatal("create agent: empty response")
	}

	created := createBody.GetAgent()
	agentKey := created.GetKey()
	if agentKey == "" {
		log.Fatal("create agent: response missing agent key")
	}

	fmt.Printf("created agent key: %s\n\n", agentKey)

	detailRes, err := sdk.Agents.GetAgent(ctx, agentKey)
	if err != nil {
		log.Fatalf("get created agent: %v", err)
	}

	detailBody := detailRes.GetGetAgentResponse()
	if detailBody == nil {
		log.Fatal("get created agent: empty response")
	}

	printJSON(detailBody.GetAgent())
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

	for _, model := range body.GetModels() {
		if model.GetIsReasoning() && model.GetModelKey() != "" {
			return model.GetModelKey(), nil
		}
	}

	models := body.GetModels()
	if len(models) > 0 && models[0].GetModelKey() != "" {
		return models[0].GetModelKey(), nil
	}

	return "", fmt.Errorf("no LLM model configured; set PIPESHUB_AGENT_MODEL_KEY in .env")
}

func printJSON(v any) {
	data, err := json.MarshalIndent(v, "", "  ")
	if err != nil {
		log.Fatalf("marshal json: %v", err)
	}
	fmt.Println(string(data))
}
