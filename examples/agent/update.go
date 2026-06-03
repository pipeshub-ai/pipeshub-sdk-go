// Update an existing agent, then print the updated agent.
//
// Uses github.com/pipeshub-ai/pipeshub-sdk-go.
//
// Usage (from examples/):
//
//	go run ./agent/update.go <path-to-.env> <agent-key>
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
	if len(os.Args) < 3 {
		log.Fatal("usage: go run ./agent/update.go <path-to-.env> <agent-key>")
	}
	if err := godotenv.Load(os.Args[1]); err != nil {
		log.Fatalf("load .env: %v", err)
	}

	agentKey := os.Args[2]

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

	name := fmt.Sprintf("SDK updated agent %d", time.Now().Unix())
	description := "Agent updated by the Go SDK example."
	systemPrompt := "You are an updated SDK example agent. Answer briefly."
	startMessage := "I have been updated."
	isReasoning := true

	modelUnion := components.CreateAgentCreateModelEntryUnionAgentCreateModelEntry(components.AgentCreateModelEntry{
		ModelKey:    modelKey,
		IsReasoning: pipeshub.Pointer(isReasoning),
	})

	res, err := sdk.Agents.UpdateAgent(ctx, agentKey, components.AgentUpdateRequest{
		Name:         &name,
		Description:  &description,
		SystemPrompt: &systemPrompt,
		StartMessage: &startMessage,
		Models:       []components.AgentCreateModelEntryUnion{modelUnion},
	})
	if err != nil {
		log.Fatalf("update agent: %v", err)
	}

	updateBody := res.GetAgentUpdateResponse()
	if updateBody == nil {
		log.Fatal("update agent: empty response")
	}

	fmt.Printf("update status: %s\n%s\n\n", updateBody.GetStatus(), updateBody.GetMessage())

	detailRes, err := sdk.Agents.GetAgent(ctx, agentKey)
	if err != nil {
		log.Fatalf("get updated agent: %v", err)
	}

	detailBody := detailRes.GetGetAgentResponse()
	if detailBody == nil {
		log.Fatal("get updated agent: empty response")
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
