// Find a connector's record group through the knowledge hub and queue it for
// reindexing with ReindexRecordGroup.
//
// The record group must belong to a connector. A knowledge base, or a folder
// inside one, is not a record group here and returns 404; use ReindexRecord
// for KB content instead.
//
// The connector is the one named PIPESHUB_CONNECTOR_NAME, or the first
// connector on the instance when that variable is empty. The knowledge-hub
// calls are marked deprecated in the SDK but are still the way to browse a
// connector's record groups.
//
// Usage (from examples/):
//
//	go run ./KB/records/reindex_record_group <path-to-.env>
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

func main() {
	if len(os.Args) < 2 {
		log.Fatal("usage: go run ./KB/records/reindex_record_group <path-to-.env>")
	}
	if err := godotenv.Load(os.Args[1]); err != nil {
		log.Fatalf("load .env: %v", err)
	}

	sdk, err := auth.NewOAuthClient()
	if err != nil {
		log.Fatal(err)
	}

	ctx := context.Background()

	app, err := findConnectorApp(ctx, sdk, os.Getenv("PIPESHUB_CONNECTOR_NAME"))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Connector:    %s (%s)\n", app.Name, app.ID)

	group, err := firstRecordGroup(ctx, sdk, app.ID)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Record group: %s (%s)\n\n", group.Name, group.ID)

	res, err := sdk.KnowledgeBase.ReindexRecordGroup(ctx, group.ID, &components.ReindexRecordGroupRequestBody{
		Depth: pipeshub.Pointer(int64(0)),
		Force: pipeshub.Pointer(false),
		// Only requeue records whose indexing failed; omit to reindex everything.
		StatusFilters: []components.IndexingStatusFilter{components.IndexingStatusFilterFailed},
	})
	if err != nil {
		log.Fatalf("reindex record group: %v", err)
	}
	if res.ReIndexRecordGroupResponseSchema == nil {
		log.Fatal("reindex record group: empty response")
	}
	body := res.ReIndexRecordGroupResponseSchema

	connector, _ := body.Connector.GetOrZero()
	fmt.Println(body.Message)
	fmt.Printf("  connector:       %s\n", connector)
	fmt.Printf("  depth:           %d\n", body.Depth)
	fmt.Printf("  event published: %v\n", body.EventPublished)
}

// findConnectorApp returns the knowledge-hub root node for the connector named
// name, or the first connector when name is empty.
func findConnectorApp(ctx context.Context, sdk *pipeshub.Pipeshub, name string) (components.KnowledgeHubNode, error) {
	res, err := sdk.KnowledgeHub.GetKnowledgeHubRootNodes(ctx, operations.GetKnowledgeHubRootNodesRequest{})
	if err != nil {
		return components.KnowledgeHubNode{}, fmt.Errorf("get knowledge hub root nodes: %w", err)
	}
	if res.KnowledgeHubNodesResponse == nil {
		return components.KnowledgeHubNode{}, fmt.Errorf("get knowledge hub root nodes: empty response")
	}
	for _, n := range res.KnowledgeHubNodesResponse.GetItems() {
		if n.NodeType != components.NodeTypeApp || n.Origin != components.KnowledgeHubNodeOriginConnector {
			continue
		}
		if name == "" || n.Name == name {
			return n, nil
		}
	}
	if name == "" {
		return components.KnowledgeHubNode{}, fmt.Errorf("no connector found on this instance")
	}
	return components.KnowledgeHubNode{}, fmt.Errorf("connector %q not found", name)
}

// firstRecordGroup returns the first record-group child of a connector app.
func firstRecordGroup(ctx context.Context, sdk *pipeshub.Pipeshub, appID string) (components.KnowledgeHubNode, error) {
	res, err := sdk.KnowledgeHub.GetKnowledgeHubChildNodes(ctx, operations.GetKnowledgeHubChildNodesRequest{
		ParentType:     operations.ParentTypeApp,
		ParentID:       appID,
		OnlyContainers: pipeshub.Pointer(true),
	})
	if err != nil {
		return components.KnowledgeHubNode{}, fmt.Errorf("get knowledge hub child nodes: %w", err)
	}
	if res.KnowledgeHubNodesResponse == nil {
		return components.KnowledgeHubNode{}, fmt.Errorf("get knowledge hub child nodes: empty response")
	}
	for _, n := range res.KnowledgeHubNodesResponse.GetItems() {
		if n.NodeType == components.NodeTypeRecordGroup {
			return n, nil
		}
	}
	return components.KnowledgeHubNode{}, fmt.Errorf("connector has no record groups")
}
