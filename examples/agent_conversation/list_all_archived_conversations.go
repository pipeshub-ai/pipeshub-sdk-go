// List all archived conversations for one agent using the paginated
// (non-grouped) archives API.
//
// Uses github.com/pipeshub-ai/pipeshub-sdk-go.
//
// Usage (from examples/):
//
//	go run ./agent_conversation/list_all_archived_conversations.go <path-to-.env>
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

	"enterprise_search/auth"
)

const (
	agentKey = "02a7d998-d21b-4015-aaf7-5cda765c1012"

	pageLimit = int64(20)
)

func main() {
	if len(os.Args) < 2 {
		log.Fatal("usage: go run ./agent_conversation/list_all_archived_conversations.go <path-to-.env>")
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
	sortBy := operations.ListAgentConversationArchivesSortByLastActivityAt
	sortOrder := operations.ListAgentConversationArchivesSortOrderDesc

	var (
		page       int64 = 1
		totalCount int64
		printed    int
	)

	fmt.Printf("Archived conversations for agent %s (newest first):\n", agentKey)

	for {
		res, err := sdk.Agents.ListAgentConversationArchives(ctx, operations.ListAgentConversationArchivesRequest{
			AgentKey:  agentKey,
			Page:      pipeshub.Int64(page),
			Limit:     pipeshub.Int64(pageLimit),
			SortBy:    sortBy.ToPointer(),
			SortOrder: sortOrder.ToPointer(),
		})
		if err != nil {
			log.Fatalf("list archived conversations (page %d): %v", page, err)
		}
		if res == nil || res.AgentArchivedConversationListResponse == nil {
			log.Fatalf("no archive list response returned (page %d)", page)
		}

		body := res.AgentArchivedConversationListResponse
		if page == 1 {
			totalCount = body.GetPagination().TotalCount
		}

		for _, conv := range body.GetConversations() {
			printArchivedConversation(conv)
			printed++
		}

		pagination := body.GetPagination()
		if !pagination.HasNextPage || page >= pagination.TotalPages {
			break
		}
		page++
	}

	if printed == 0 {
		fmt.Println("  (no archived conversations for this agent)")
		return
	}

	fmt.Printf("\nListed %d archived conversation(s) (total reported: %d).\n", printed, totalCount)
}

func printArchivedConversation(conv components.AgentConversationListItem) {
	id := ""
	if v := conv.GetID(); v != nil {
		id = *v
	}

	title := "(untitled)"
	if t := conv.GetTitle(); t != nil && *t != "" {
		title = *t
	}

	archived := "—"
	if at := conv.GetArchivedAt(); at != nil {
		archived = at.Format("2006-01-02 15:04:05 MST")
	}

	fmt.Printf("  - %q — %s — archived %s\n", title, id, archived)
}
