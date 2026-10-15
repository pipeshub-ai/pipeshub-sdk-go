// Package connectorsync holds the steps every connector example shares: pick a
// filter value from a connector's live options, save sync filters, turn the
// sync on, wait until the sync and indexing have finished, search what was
// indexed, and clean the connector up again.
package connectorsync

import (
	"context"
	"errors"
	"fmt"
	"log"
	"time"

	pipeshub "github.com/pipeshub-ai/pipeshub-sdk-go"
	"github.com/pipeshub-ai/pipeshub-sdk-go/models/apierrors"
	"github.com/pipeshub-ai/pipeshub-sdk-go/models/components"
	"github.com/pipeshub-ai/pipeshub-sdk-go/models/operations"
)

// Option is one value a connector offers for a filter, such as a channel or a bucket.
type Option struct {
	ID    string
	Label string
}

// ListOptions returns the first page of values a connector offers for filterKey.
func ListOptions(ctx context.Context, sdk *pipeshub.Pipeshub, connectorID, filterKey string) ([]Option, error) {
	limit := int64(10)
	res, err := sdk.ConnectorFilters.GetFilterFieldOptions(ctx, operations.GetFilterFieldOptionsRequest{
		ConnectorID: connectorID,
		FilterKey:   filterKey,
		Limit:       &limit,
	})
	if err != nil {
		return nil, fmt.Errorf("list %s options: %w", filterKey, err)
	}
	var opts []Option
	for _, o := range res.Object.Options {
		opts = append(opts, Option{ID: value(o.ID), Label: value(o.Label)})
	}
	return opts, nil
}

// Labels returns the options' display names, for printing.
func Labels(opts []Option) []string {
	var out []string
	for _, o := range opts {
		out = append(out, o.Label)
	}
	return out
}

// Choose returns preferred when it is set, otherwise the first option's id.
// The options are paged, so an explicit value needs no lookup.
func Choose(preferred string, opts []Option) (string, bool) {
	if preferred != "" {
		return preferred, true
	}
	if len(opts) == 0 {
		return "", false
	}
	return opts[0].ID, true
}

// SaveSyncFilters stores sync filter values, keyed by the registry schema's filter names.
func SaveSyncFilters(ctx context.Context, sdk *pipeshub.Pipeshub, connectorID string, values map[string]components.ConnectorFilterValue) error {
	_, err := sdk.ConnectorConfiguration.UpdateConnectorFiltersSyncConfig(ctx, connectorID, components.UpdateConnectorFiltersSyncRequest{
		Filters: &components.ConnectorFiltersConfig{Sync: &components.ConnectorFiltersConfigSync{Values: values}},
	})
	if err != nil {
		return fmt.Errorf("save filters: %w", err)
	}
	return nil
}

// ToggleSync flips the connector's sync state. Turning it on tests the
// credentials and starts a full sync.
func ToggleSync(ctx context.Context, sdk *pipeshub.Pipeshub, connectorID string) error {
	_, err := sdk.ConnectorControl.ToggleConnector(ctx, connectorID, components.ConnectorToggleRequest{
		Type: components.ConnectorToggleRequestTypeSync,
	})
	return err
}

// WaitUntilIndexed polls until the sync has finished and no record is still
// waiting to be indexed. A sync can legitimately find nothing, so finishing is
// judged from the connector's status rather than from the record count.
func WaitUntilIndexed(ctx context.Context, sdk *pipeshub.Pipeshub, connectorID string, timeout time.Duration) (*components.Stats, error) {
	start := time.Now()
	seenSync := false
	for {
		inst, err := sdk.ConnectorInstances.GetConnectorInstance(ctx, connectorID)
		if err != nil {
			return nil, fmt.Errorf("get connector: %w", err)
		}
		status, _ := inst.Object.Connector.GetStatus().GetOrZero()
		syncing := status == "SYNCING" || status == "FULL_SYNCING"

		res, err := sdk.ConnectorInstances.GetConnectorStats(ctx, connectorID)
		if err != nil {
			return nil, fmt.Errorf("get connector stats: %w", err)
		}
		stats := res.ConnectorStats.Data.Stats
		c := stats.IndexingStatus
		pending := Count(c.NotStarted) + Count(c.Queued) + Count(c.InProgress)
		seenSync = seenSync || syncing || stats.Total > 0
		log.Printf("status: %s  records: %d  indexed: %d  pending: %d  failed: %d",
			status, stats.Total, Count(c.Completed), pending, Count(c.Failed))
		// Give the sync a minute to show up before trusting an idle status.
		if !syncing && pending == 0 && (seenSync || time.Since(start) > time.Minute) {
			return &stats, nil
		}
		if time.Since(start) > timeout {
			return nil, fmt.Errorf("sync or indexing still running after %s", timeout)
		}
		time.Sleep(15 * time.Second)
	}
}

// Search runs a semantic search scoped to the connector and prints each
// matching record once. Hits are chunks, so one record can match many times.
//
// Records from connectors that keep the source's permissions are only visible
// to PipesHub users the source lets see them; search answers 404 when the
// caller can see none, which is reported rather than treated as a failure.
func Search(ctx context.Context, sdk *pipeshub.Pipeshub, connectorID, query string) error {
	limit := int64(3)
	res, err := sdk.SemanticSearch.Search(ctx, components.SemanticSearchRequest{
		Query:   query,
		Filters: &components.Filters{Apps: []string{connectorID}},
		Limit:   &limit,
	})
	var apiErr *apierrors.APIError
	if errors.As(err, &apiErr) && apiErr.StatusCode == 404 {
		log.Print("search: nothing you can see yet. This connector keeps the source's permissions, " +
			"so only PipesHub users whose email has access in the source find its records")
		return nil
	}
	if err != nil {
		return fmt.Errorf("search: %w", err)
	}
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

// CleanUp turns the sync off if it is on, then deletes the connector.
func CleanUp(ctx context.Context, sdk *pipeshub.Pipeshub, connectorID string) error {
	res, err := sdk.ConnectorInstances.GetConnectorInstance(ctx, connectorID)
	if err != nil {
		return fmt.Errorf("get connector: %w", err)
	}
	if active := res.Object.Connector.GetIsActive(); active != nil && *active {
		if err := ToggleSync(ctx, sdk, connectorID); err != nil {
			return fmt.Errorf("turn sync off: %w", err)
		}
	}
	if _, err := sdk.ConnectorInstances.DeleteConnectorInstance(ctx, connectorID); err != nil {
		return fmt.Errorf("delete connector: %w", err)
	}
	log.Print("sync off, connector deleted")
	return nil
}

// Count dereferences an optional count, treating a missing one as zero.
func Count(p *int64) int64 {
	if p == nil {
		return 0
	}
	return *p
}

func value(p *string) string {
	if p == nil {
		return ""
	}
	return *p
}
