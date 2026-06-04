// List all agents, then optionally fetch one agent by key.
//
// Uses github.com/pipeshub-ai/pipeshub-sdk-go.
//
// Usage (from examples/):
//
//	go run ./agent/get.go <path-to-.env> [agent-key]
//
// If no agent key is provided, the example lists all agents and then
// fetches the first one returned, if any.
package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/pipeshub-ai/pipeshub-sdk-go/models/components"
	"github.com/pipeshub-ai/pipeshub-sdk-go/models/operations"

	"enterprise_search/auth"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatal("usage: go run ./agent/get.go <path-to-.env> [agent-key]")
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

	listRes, err := sdk.Agents.ListAgents(ctx, operations.ListAgentsRequest{})
	if err != nil {
		log.Fatalf("list agents: %v", err)
	}

	listBody := listRes.GetAgentListResponse()
	if listBody == nil {
		log.Fatal("list agents: empty response")
	}

	agents := listBody.GetAgents()
	fmt.Printf("found %d agent(s)\n", len(agents))
	for _, agent := range agents {
		printAgentSummary(agent)
	}

	agentKey := ""
	if len(os.Args) >= 3 {
		agentKey = os.Args[2]
	} else if len(agents) > 0 {
		agentKey = agents[0].GetKey()
	}

	if agentKey == "" {
		return
	}

	getRes, err := sdk.Agents.GetAgent(ctx, agentKey)
	if err != nil {
		log.Fatalf("get agent: %v", err)
	}

	getBody := getRes.GetGetAgentResponse()
	if getBody == nil {
		log.Fatal("get agent: empty response")
	}

	fmt.Printf("\nagent details for %s:\n", agentKey)
	printJSON(getBody.GetAgent())
}

func printAgentSummary(agent components.AgentListItem) {
	fmt.Printf("- %s (%s) shareWithOrg=%t access=%s\n", agent.GetName(), agent.GetKey(), agent.GetShareWithOrg(), agent.GetAccessType())
}

func printJSON(v any) {
	data, err := json.MarshalIndent(v, "", "  ")
	if err != nil {
		log.Fatalf("marshal json: %v", err)
	}
	fmt.Println(string(data))
}
