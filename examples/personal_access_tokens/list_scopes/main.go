// List the scopes a personal access token may be granted, with
// ListPersonalAccessTokenScopes.
//
// Personal access token routes require an interactive user session: they
// reject PATs and OAuth tokens. So this example always logs in with
// PIPESHUB_TEST_USER_EMAIL and PIPESHUB_TEST_USER_PASSWORD, even when
// OAuth client credentials are set.
//
// Usage (from examples/):
//
//	go run ./personal_access_tokens/list_scopes <path-to-.env>
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
	if len(os.Args) < 2 {
		log.Fatal("usage: go run ./personal_access_tokens/list_scopes <path-to-.env>")
	}
	if err := godotenv.Load(os.Args[1]); err != nil {
		log.Fatalf("load .env: %v", err)
	}

	sdk, err := auth.NewSessionClient(
		os.Getenv("PIPESHUB_TEST_USER_EMAIL"),
		os.Getenv("PIPESHUB_TEST_USER_PASSWORD"),
	)
	if err != nil {
		log.Fatal(err)
	}

	res, err := sdk.PersonalAccessTokens.ListPersonalAccessTokenScopes(context.Background())
	if err != nil {
		log.Fatalf("list personal access token scopes: %v", err)
	}
	if res.PatScopesListResponse == nil {
		log.Fatal("list personal access token scopes: empty response")
	}

	scopes := res.PatScopesListResponse.Scopes
	fmt.Printf("%d scope(s) available to personal access tokens:\n", len(scopes))
	for _, s := range scopes {
		fmt.Printf("  %-28s %-14s %s\n", s.Name, s.Category, s.Description)
	}
}
