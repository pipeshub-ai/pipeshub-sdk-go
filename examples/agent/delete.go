// Delete an agent by key.
//
// Uses github.com/pipeshub-ai/pipeshub-sdk-go.
//
// Usage (from examples/):
//
//	go run ./agent/delete.go <path-to-.env> <agent-key>
package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"

	"enterprise_search/auth"
)

func main() {
	if len(os.Args) < 3 {
		log.Fatal("usage: go run ./agent/delete.go <path-to-.env> <agent-key>")
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

	res, err := sdk.Agents.DeleteAgent(context.Background(), agentKey)
	if err != nil {
		log.Fatalf("delete agent: %v", err)
	}

	body := res.GetAgentDeleteResponse()
	if body == nil {
		fmt.Printf("deleted agent: %s\n", agentKey)
		return
	}

	fmt.Printf("deleted agent: %s\nstatus: %s\nmessage: %s\n", agentKey, body.GetStatus(), body.GetMessage())
}
