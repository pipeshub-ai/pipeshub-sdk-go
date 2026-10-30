// Warn once per PipesHub instance when its release differs (major.minor) from the
// one this SDK version is built for. The server reports its release on GET /api/v1/health.

package hooks

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
	"sync"
	"time"
)

// SupportedPipesHubVersion is the PipesHub release this SDK version is built for.
// Bump it with each SDK release.
const SupportedPipesHubVersion = "1.0.0"

const (
	apiPrefix     = "/api/v1"
	healthTimeout = 2 * time.Second
)

var (
	checked sync.Map
	warnf   = log.Printf // swapped in tests to capture the warning
)

type VersionCheckHook struct{}

var _ beforeRequestHook = (*VersionCheckHook)(nil)

func majorMinor(version string) (string, bool) {
	parts := strings.Split(strings.TrimPrefix(strings.TrimSpace(version), "v"), ".")
	if len(parts) < 2 {
		return "", false
	}
	major, err1 := strconv.Atoi(parts[0])
	minor, err2 := strconv.Atoi(parts[1])
	if err1 != nil || err2 != nil {
		return "", false
	}
	return fmt.Sprintf("%d.%d", major, minor), true
}

// mismatchWarning returns the warning for a /health response body, or "" when there is nothing to warn about.
func mismatchWarning(sdkVersion, instance string, body []byte) string {
	var health map[string]json.RawMessage
	if err := json.Unmarshal(body, &health); err != nil {
		return ""
	}
	if _, ok := health["status"]; !ok {
		return ""
	}
	var running string
	var build struct {
		Version *string `json:"version"`
	}
	if raw, ok := health["build"]; !ok || string(raw) == "null" {
		// PipesHub reports `build` from 1.0.0 on.
		running = "a PipesHub release older than 1.0.0"
	} else if json.Unmarshal(raw, &build) != nil {
		return ""
	} else {
		if build.Version == nil {
			return "" // a local or dev build reports no version
		}
		server, ok := majorMinor(*build.Version)
		supported, _ := majorMinor(SupportedPipesHubVersion)
		if !ok || server == supported {
			return ""
		}
		running = "PipesHub " + *build.Version
	}
	return fmt.Sprintf("pipeshub-sdk-go %s is built for PipesHub %s, but %s runs %s. Some calls may not work as expected.",
		sdkVersion, SupportedPipesHubVersion, instance, running)
}

func checkVersion(client HTTPClient, sdkVersion, instance string) {
	// Not the caller's context: the check outlives the request that started it.
	ctx, cancel := context.WithTimeout(context.Background(), healthTimeout)
	defer cancel()
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, instance+apiPrefix+"/health", nil)
	if err != nil {
		return
	}
	res, err := client.Do(req)
	if err != nil {
		return // the check must never affect the caller's requests
	}
	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		return
	}
	var body json.RawMessage
	if err := json.NewDecoder(res.Body).Decode(&body); err != nil {
		return
	}
	if warning := mismatchWarning(sdkVersion, instance, body); warning != "" {
		warnf("%s", warning)
	}
}

func (h *VersionCheckHook) BeforeRequest(hookCtx BeforeRequestContext, req *http.Request) (*http.Request, error) {
	url := req.URL.String()
	index := strings.Index(url, apiPrefix+"/")
	client := hookCtx.SDKConfiguration.Client
	if index == -1 || client == nil {
		return req, nil
	}
	instance := url[:index]
	if _, seen := checked.LoadOrStore(instance, true); seen {
		return req, nil
	}
	// Inline, not a goroutine: Go kills goroutines when main returns, so a short program would never see the warning.
	// Costs one /health round trip (at most healthTimeout) on the first request per instance.
	checkVersion(client, hookCtx.SDKConfiguration.SDKVersion, instance)
	return req, nil
}
