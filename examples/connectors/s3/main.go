// Register an Amazon S3 connector with an access key, sync one bucket and
// index it, then search it.
//
// Reads the connector's auth fields from the registry, creates a team "S3"
// connector with an access key (PIPESHUB_S3_ACCESS_KEY,
// PIPESHUB_S3_SECRET_KEY), limits its sync to one bucket
// (PIPESHUB_S3_BUCKET, or the first bucket the key can list), turns the sync
// on, waits until the sync and indexing have finished, runs a semantic
// search scoped to the connector, and finally turns the sync off and deletes
// the connector.
//
// Usage (from examples/):
//
//	go run ./connectors/s3 <path-to-.env>
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
	connectorType = "S3"
	authType      = components.ConnectorAuthTypeAccessKey
)

func main() {
	if len(os.Args) < 2 {
		log.Fatal("usage: go run ./connectors/s3 <path-to-.env>")
	}
	if err := godotenv.Load(os.Args[1]); err != nil {
		log.Fatalf("load .env: %v", err)
	}
	s3AccessKey := mustEnv("PIPESHUB_S3_ACCESS_KEY")
	s3SecretKey := mustEnv("PIPESHUB_S3_SECRET_KEY")
	query := os.Getenv("PIPESHUB_S3_SEARCH_QUERY")
	if query == "" {
		query = "report"
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
	connectorID, err := createConnector(ctx, sdk, map[string]any{"accessKey": s3AccessKey, "secretKey": s3SecretKey})
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
		InstanceName:  "S3 example",
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
	opts, err := connectorsync.ListOptions(ctx, sdk, connectorID, "buckets")
	if err != nil {
		return err
	}
	log.Printf("buckets it can sync: %v", connectorsync.Labels(opts))
	chosen, ok := connectorsync.Choose(os.Getenv("PIPESHUB_S3_BUCKET"), opts)
	if !ok {
		return fmt.Errorf("the connector can't see any bucket")
	}
	filters := map[string]components.ConnectorFilterValue{
		"buckets": {Operator: "in", Value: []string{chosen}, Type: components.ConnectorFilterValueTypeMultiselect},
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
