package hooks

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"strings"
	"sync"
	"testing"

	"github.com/pipeshub-ai/pipeshub-sdk-go/v2/internal/config"
)

// runHook sends one SDK request to a fake instance whose /health returns the given body,
// and returns the paths it was asked for and the warnings logged.
func runHook(t *testing.T, server *httptest.Server, path string) []string {
	t.Helper()
	var warnings []string
	warnf = func(format string, args ...any) { warnings = append(warnings, fmt.Sprintf(format, args...)) }
	t.Cleanup(func() {
		warnf = log.Printf
		checked = sync.Map{}
	})

	req, err := http.NewRequest(http.MethodGet, server.URL+path, nil)
	if err != nil {
		t.Fatal(err)
	}
	ctx := BeforeRequestContext{HookContext{SDKConfiguration: config.SDKConfiguration{Client: server.Client(), SDKVersion: "2.0.0"}}}
	got, err := (&VersionCheckHook{}).BeforeRequest(ctx, req)
	if err != nil || got != req {
		t.Fatalf("BeforeRequest changed the request: %v", err)
	}
	return warnings
}

func fakeInstance(t *testing.T, status int, body string, paths *[]string) *httptest.Server {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		*paths = append(*paths, r.URL.Path)
		w.WriteHeader(status)
		_, _ = w.Write([]byte(body))
	}))
	t.Cleanup(server.Close)
	return server
}

func TestMismatchWarningIsEmpty(t *testing.T) {
	for _, body := range []string{
		`{"status":"healthy","build":{"version":"1.0.0","commitId":"abc"}}`,
		`{"status":"healthy","build":{"version":"1.0.7","commitId":"abc"}}`,
		`{"status":"healthy","build":{"version":null,"commitId":null}}`,
		`{"status":"healthy","build":{"version":"dev","commitId":null}}`,
		`{"not":"pipeshub"}`,
		`["not","an","object"]`,
	} {
		if got := mismatchWarning("2.0.0", "https://pipeshub.example", []byte(body)); got != "" {
			t.Errorf("%s: got %q, want no warning", body, got)
		}
	}
}

func TestMismatchWarningOnAnotherMinorRelease(t *testing.T) {
	got := mismatchWarning("2.0.0", "https://pipeshub.example", []byte(`{"status":"healthy","build":{"version":"1.1.0","commitId":"abc"}}`))
	want := "pipeshub-sdk-go 2.0.0 is built for PipesHub 1.0.0, but https://pipeshub.example runs PipesHub 1.1.0. Some calls may not work as expected."
	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

func TestMismatchWarningWithoutBuildInfo(t *testing.T) {
	got := mismatchWarning("2.0.0", "https://pipeshub.example", []byte(`{"status":"healthy"}`))
	if !strings.Contains(got, "a PipesHub release older than 1.0.0") {
		t.Errorf("got %q", got)
	}
}

func TestHookWarnsOncePerInstance(t *testing.T) {
	var paths []string
	server := fakeInstance(t, http.StatusOK, `{"status":"healthy","build":{"version":"0.10.0","commitId":"abc"}}`, &paths)

	first := runHook(t, server, "/api/v1/knowledgeBase")
	second := runHook(t, server, "/api/v1/knowledgeBase")

	if len(paths) != 1 || paths[0] != "/api/v1/health" {
		t.Fatalf("health calls: %v", paths)
	}
	if len(first) != 1 || !strings.Contains(first[0], "runs PipesHub 0.10.0") || len(second) != 0 {
		t.Fatalf("warnings: first %v, second %v", first, second)
	}
}

func TestHookIsQuietWhenHealthFails(t *testing.T) {
	var paths []string
	server := fakeInstance(t, http.StatusInternalServerError, `{"error":"boom"}`, &paths)
	if warnings := runHook(t, server, "/api/v1/knowledgeBase"); len(warnings) != 0 || len(paths) != 1 {
		t.Fatalf("warnings %v, paths %v", warnings, paths)
	}
}

func TestHookSkipsRequestsOutsideTheAPI(t *testing.T) {
	var paths []string
	server := fakeInstance(t, http.StatusOK, `{"status":"healthy"}`, &paths)
	if warnings := runHook(t, server, "/mcp"); len(warnings) != 0 || len(paths) != 0 {
		t.Fatalf("warnings %v, paths %v", warnings, paths)
	}
}
