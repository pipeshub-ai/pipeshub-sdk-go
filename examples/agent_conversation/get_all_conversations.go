// List all active (non-archived) conversations for one agent.
//
// Uses github.com/pipeshub-ai/pipeshub-sdk-go.
//
// Usage (from examples/):
//
//	go run ./agent_conversation/get_all_conversations.go <path-to-.env>
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
	"github.com/pipeshub-ai/pipeshub-sdk-go/models/operations"

	"enterprise_search/auth"
)

const (
	agentKey = "02a7d998-d21b-4015-aaf7-5cda765c1012"

	pageLimit = int64(20)
)

func main() {
	if len(os.Args) < 2 {
		log.Fatal("usage: go run ./agent_conversation/get_all_conversations.go <path-to-.env>")
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
	sortBy := string(operations.SearchHistorySortByLastActivityAt)
	sortOrder := string(operations.SearchHistorySortOrderDesc)

	var (
		page            int64 = 1
		ownedTotalCount int64
		owned           []components.AgentConversationListItem
		shared          []components.AgentConversationListItem
	)

	fmt.Printf("Active conversations for agent %s (newest first):\n", agentKey)

	for {
		res, err := sdk.Agents.ListAgentConversations(ctx, operations.ListAgentConversationsRequest{
			AgentKey:  agentKey,
			Page:      pipeshub.Int64(page),
			Limit:     pipeshub.Int64(pageLimit),
			SortBy:    &sortBy,
			SortOrder: &sortOrder,
		})
		if err != nil {
			log.Fatalf("list conversations (page %d): %v", page, err)
		}
		if res == nil || res.AgentConversationListResponse == nil {
			log.Fatalf("no conversation list response returned (page %d)", page)
		}

		body := res.AgentConversationListResponse
		if page == 1 {
			ownedTotalCount = body.GetPagination().TotalCount
		}

		owned = append(owned, body.GetConversations()...)
		shared = append(shared, body.GetSharedWithMeConversations()...)

		pagination := body.GetPagination()
		if !pagination.HasNextPage || page >= pagination.TotalPages {
			break
		}
		page++
	}

	printConversationSection("Your conversations", owned)
	printConversationSection("Shared with you", shared)

	ownedCount := len(owned)
	sharedCount := len(shared)
	if ownedCount == 0 && sharedCount == 0 {
		fmt.Println("\n(no active conversations for this agent)")
		return
	}

	fmt.Printf("\nListed %d owned and %d shared conversation(s)", ownedCount, sharedCount)
	if ownedTotalCount > 0 || ownedCount > 0 {
		fmt.Printf(" (owned total reported: %d)", ownedTotalCount)
	}
	fmt.Println(".")
}

func printConversationSection(heading string, convs []components.AgentConversationListItem) {
	fmt.Printf("\n%s:\n", heading)
	if len(convs) == 0 {
		fmt.Println("  (none)")
		return
	}
	for _, conv := range convs {
		printActiveConversation(conv)
	}
}

func printActiveConversation(conv components.AgentConversationListItem) {
	id := ""
	if v := conv.GetID(); v != nil {
		id = *v
	}

	title := "(untitled)"
	if t := conv.GetTitle(); t != nil && *t != "" {
		title = *t
	}

	when := formatActivityTime(conv)

	fmt.Printf("  - %q — %s — %s\n", title, id, when)
}

func formatActivityTime(conv components.AgentConversationListItem) string {
	if ms := conv.GetLastActivityAt(); ms != nil && *ms > 0 {
		return time.UnixMilli(*ms).Format("2006-01-02 15:04:05 MST")
	}
	if ut := conv.GetUpdatedAt(); ut != nil {
		return ut.Format("2006-01-02 15:04:05 MST")
	}
	return "—"
}
