// Register a Gmail Workspace connector with a service account, sync the last
// 7 days of mail and index it, then search it.
//
// Reads the connector's auth fields from the registry, creates a team "Gmail
// Workspace" connector with a Google service account that has domain-wide
// delegation (the key file at PIPESHUB_GOOGLE_SERVICE_ACCOUNT_FILE) and the
// Workspace admin email in PIPESHUB_GOOGLE_ADMIN_EMAIL, limits its sync to
// mail received in the last 7 days, turns the sync on, waits until the sync
// and indexing have finished, runs a semantic search scoped to the
// connector, and finally turns the sync off and deletes the connector.
//
// Gmail records keep Google's permissions, so search only finds a mailbox's
// messages for the PipesHub user with that email. The example reports that
// instead of failing when the caller sees none.
//
// Usage (from examples/):
//
//	go run ./connectors/gmail_workspace <path-to-.env>
package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"sort"
	"time"

	"github.com/joho/godotenv"
	pipeshub "github.com/pipeshub-ai/pipeshub-sdk-go"
	"github.com/pipeshub-ai/pipeshub-sdk-go/models/components"

	"enterprise_search/auth"
	"enterprise_search/connectorsync"
)

const (
	connectorType = "Gmail Workspace"
	authType      = components.ConnectorAuthTypeCustom
)

func main() {
	if len(os.Args) < 2 {
		log.Fatal("usage: go run ./connectors/gmail_workspace <path-to-.env>")
	}
	if err := godotenv.Load(os.Args[1]); err != nil {
		log.Fatalf("load .env: %v", err)
	}
	googleAdminEmail := mustEnv("PIPESHUB_GOOGLE_ADMIN_EMAIL")
	googleServiceAccountJSON, err := os.ReadFile(mustEnv("PIPESHUB_GOOGLE_SERVICE_ACCOUNT_FILE"))
	if err != nil {
		log.Fatalf("read PIPESHUB_GOOGLE_SERVICE_ACCOUNT_FILE: %v", err)
	}
	query := os.Getenv("PIPESHUB_GMAIL_SEARCH_QUERY")
	if query == "" {
		query = "meeting"
	}

	sdk, err := auth.NewOAuthClient()
	if err != nil {
		log.Fatal(err)
	}
	ctx := context.Background()

	schema, err := sdk.ConnectorRegistry.GetConnectorSchema(ctx, connectorType)
	if err != nil {
		log.Fatalf("get connector schema: %v", err)
	}
	var fields []string
	for _, f := range schema.Object.Schema.GetAuth().GetSchemas()[string(authType)].Fields {
		if f.GetName() != nil {
			fields = append(fields, *f.GetName())
		}
	}
	log.Printf("%s auth fields: %v", connectorType, fields)

	// Auth fields are keyed by the registry schema's field names.
	connectorID, err := createConnector(ctx, sdk, map[string]any{"adminEmail": googleAdminEmail, "serviceAccountJson": string(googleServiceAccountJSON)})
	if err != nil {
		log.Fatalf("create connector: %v", err)
	}
	log.Printf("created connector %s", connectorID)
	// log.Fatal skips deferred calls, so failures below return an error to here instead.
	runErr := syncAndSearch(ctx, sdk, connectorID, query)
	if err := connectorsync.CleanUp(ctx, sdk, connectorID); err != nil {
		log.Printf("clean up: %v", err)
	}
	if runErr != nil {
		log.Fatal(runErr)
	}
}

func createConnector(ctx context.Context, sdk *pipeshub.Pipeshub, authFields map[string]any) (string, error) {
	at := authType
	strategy := components.ConnectorSyncConfigSelectedStrategyManual
	res, err := sdk.ConnectorInstances.CreateConnectorInstance(ctx, components.CreateConnectorRequest{
		ConnectorType: connectorType,
		InstanceName:  "Gmail Workspace example",
		Scope:         components.ConnectorScopeTeam,
		AuthType:      &at,
		Config: &components.CreateConnectorRequestConfig{
			Auth: &components.ConnectorAuthConfig{AdditionalProperties: authFields},
			Sync: &components.ConnectorSyncConfig{SelectedStrategy: &strategy},
		},
	})
	if err != nil {
		return "", err
	}
	return res.Object.Connector.ConnectorID, nil
}

func syncAndSearch(ctx context.Context, sdk *pipeshub.Pipeshub, connectorID, query string) error {
	filters := map[string]components.ConnectorFilterValue{
		"received_date": {Operator: "last_7_days", Value: nil, Type: components.ConnectorFilterValueTypeDatetime},
	}
	// Filter values are keyed by the schema's filter names; each has an operator and a type.
	if err := connectorsync.SaveSyncFilters(ctx, sdk, connectorID, filters); err != nil {
		return err
	}
	log.Printf("sync filters: %v", keys(filters))

	// Toggling flips the sync state; turning it on tests the credentials and starts a full sync.
	if err := connectorsync.ToggleSync(ctx, sdk, connectorID); err != nil {
		return fmt.Errorf("turn sync on: %w", err)
	}
	log.Print("sync on")

	stats, err := connectorsync.WaitUntilIndexed(ctx, sdk, connectorID, 30*time.Minute)
	if err != nil {
		return err
	}
	log.Printf("indexed %d of %d records", connectorsync.Count(stats.IndexingStatus.Completed), stats.Total)
	return connectorsync.Search(ctx, sdk, connectorID, query)
}

func mustEnv(name string) string {
	v := os.Getenv(name)
	if v == "" {
		log.Fatalf("%s is required", name)
	}
	return v
}

func keys(m map[string]components.ConnectorFilterValue) []string {
	var out []string
	for k := range m {
		out = append(out, k)
	}
	sort.Strings(out)
	return out
}
