// Register an RSS connector, sync and index a feed, then search what it indexed.
//
// Looks RSS up in the connector registry, creates a team RSS connector for
// PIPESHUB_RSS_FEED_URL (default: Hacker News), turns its sync on, waits until
// every record is indexed, runs a semantic search scoped to the connector, and
// finally turns the sync off and deletes the connector.
//
// Usage (from examples/):
//
//	go run ./connectors/rss <path-to-.env>
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

	"enterprise_search/auth"
)

const defaultFeedURL = "https://news.ycombinator.com/rss"

func main() {
	if len(os.Args) < 2 {
		log.Fatal("usage: go run ./connectors/rss <path-to-.env>")
	}
	if err := godotenv.Load(os.Args[1]); err != nil {
		log.Fatalf("load .env: %v", err)
	}
	feedURL := os.Getenv("PIPESHUB_RSS_FEED_URL")
	if feedURL == "" {
		feedURL = defaultFeedURL
	}

	sdk, err := auth.NewOAuthClient()
	if err != nil {
		log.Fatal(err)
	}
	ctx := context.Background()

	if err := describeRSS(ctx, sdk); err != nil {
		log.Fatal(err)
	}

	connectorID, err := createRSSConnector(ctx, sdk, feedURL)
	if err != nil {
		log.Fatalf("create connector: %v", err)
	}
	log.Printf("created connector %s", connectorID)
	// log.Fatal skips deferred calls, so failures below return an error to here instead.
	runErr := syncAndSearch(ctx, sdk, connectorID)
	if err := cleanUp(ctx, sdk, connectorID); err != nil {
		log.Printf("clean up: %v", err)
	}
	if runErr != nil {
		log.Fatal(runErr)
	}
}

// describeRSS prints the RSS registry entry and the sync fields its schema asks for.
func describeRSS(ctx context.Context, sdk *pipeshub.Pipeshub) error {
	search := "RSS"
	reg, err := sdk.ConnectorRegistry.GetConnectorRegistry(ctx, nil, nil, nil, &search)
	if err != nil {
		return fmt.Errorf("get connector registry: %w", err)
	}
	for _, c := range reg.Object.Connectors {
		if c.GetType() != nil && *c.GetType() == "RSS" {
			log.Printf("registry: %s, auth types %v", *c.GetName(), c.GetSupportedAuthTypes())
		}
	}

	schema, err := sdk.ConnectorRegistry.GetConnectorSchema(ctx, "RSS")
	if err != nil {
		return fmt.Errorf("get connector schema: %w", err)
	}
	var fields []string
	if sync := schema.Object.Schema.GetSync(); sync != nil {
		for _, f := range sync.GetCustomFields() {
			if f.GetName() != nil {
				fields = append(fields, *f.GetName())
			}
		}
	}
	log.Printf("sync fields: %v", fields)
	return nil
}

func createRSSConnector(ctx context.Context, sdk *pipeshub.Pipeshub, feedURL string) (string, error) {
	authType := components.ConnectorAuthTypeNone
	strategy := components.ConnectorSyncConfigSelectedStrategyManual
	res, err := sdk.ConnectorInstances.CreateConnectorInstance(ctx, components.CreateConnectorRequest{
		ConnectorType: "RSS",
		InstanceName:  "RSS example",
		Scope:         components.ConnectorScopeTeam,
		AuthType:      &authType,
		Config: &components.CreateConnectorRequestConfig{
			Auth: &components.ConnectorAuthConfig{},
			// RSS's own sync fields come from its registry schema; they travel as additional properties.
			Sync: &components.ConnectorSyncConfig{
				SelectedStrategy: &strategy,
				AdditionalProperties: map[string]any{
					"feed_urls":             feedURL,
					"max_articles_per_feed": 10,
					"fetch_full_content":    false,
				},
			},
		},
	})
	if err != nil {
		return "", err
	}
	return res.Object.Connector.ConnectorID, nil
}

func syncAndSearch(ctx context.Context, sdk *pipeshub.Pipeshub, connectorID string) error {
	// Toggling flips the sync state; turning it on tests the feed and starts a full sync.
	if err := toggleSync(ctx, sdk, connectorID); err != nil {
		return fmt.Errorf("turn sync on: %w", err)
	}
	log.Print("sync on")

	stats, err := waitUntilIndexed(ctx, sdk, connectorID, 10*time.Minute)
	if err != nil {
		return err
	}
	log.Printf("indexed %d of %d records", value(stats.IndexingStatus.Completed), stats.Total)

	limit := int64(3)
	res, err := sdk.SemanticSearch.Search(ctx, components.SemanticSearchRequest{
		Query:   "technology news",
		Filters: &components.Filters{Apps: []string{connectorID}},
		Limit:   &limit,
	})
	if err != nil {
		return fmt.Errorf("search: %w", err)
	}
	// Hits are chunks, so one record can match several times; print each record once.
	seen := map[string]bool{}
	for _, hit := range res.SemanticSearchExecuteResponse.SearchResponse.SearchResults {
		// Metadata is optional on a hit; the generated getters are nil-safe.
		name, _ := hit.Metadata.GetRecordName().GetOrZero()
		if name != "" && !seen[name] {
			seen[name] = true
			fmt.Printf("- %s\n", name)
		}
	}
	return nil
}

// waitUntilIndexed polls the connector's stats until it has records and none is still pending.
func waitUntilIndexed(ctx context.Context, sdk *pipeshub.Pipeshub, connectorID string, timeout time.Duration) (*components.Stats, error) {
	deadline := time.Now().Add(timeout)
	for {
		res, err := sdk.ConnectorInstances.GetConnectorStats(ctx, connectorID)
		if err != nil {
			return nil, fmt.Errorf("get connector stats: %w", err)
		}
		stats := res.ConnectorStats.Data.Stats
		c := stats.IndexingStatus
		pending := value(c.NotStarted) + value(c.Queued) + value(c.InProgress)
		log.Printf("records: %d  indexed: %d  pending: %d  failed: %d", stats.Total, value(c.Completed), pending, value(c.Failed))
		if stats.Total > 0 && pending == 0 {
			return &stats, nil
		}
		if time.Now().After(deadline) {
			return nil, fmt.Errorf("records still pending after %s", timeout)
		}
		time.Sleep(10 * time.Second)
	}
}

func toggleSync(ctx context.Context, sdk *pipeshub.Pipeshub, connectorID string) error {
	_, err := sdk.ConnectorControl.ToggleConnector(ctx, connectorID, components.ConnectorToggleRequest{
		Type: components.ConnectorToggleRequestTypeSync,
	})
	return err
}

// cleanUp turns the sync off if it is on, then deletes the connector.
func cleanUp(ctx context.Context, sdk *pipeshub.Pipeshub, connectorID string) error {
	res, err := sdk.ConnectorInstances.GetConnectorInstance(ctx, connectorID)
	if err != nil {
		return fmt.Errorf("get connector: %w", err)
	}
	if active := res.Object.Connector.GetIsActive(); active != nil && *active {
		if err := toggleSync(ctx, sdk, connectorID); err != nil {
			return fmt.Errorf("turn sync off: %w", err)
		}
	}
	if _, err := sdk.ConnectorInstances.DeleteConnectorInstance(ctx, connectorID); err != nil {
		return fmt.Errorf("delete connector: %w", err)
	}
	log.Print("sync off, connector deleted")
	return nil
}

func value(p *int64) int64 {
	if p == nil {
		return 0
	}
	return *p
}
