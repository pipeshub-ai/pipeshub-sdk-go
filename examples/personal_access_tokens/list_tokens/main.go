// Mint a personal access token, list your tokens with
// ListPersonalAccessTokens, then revoke the one this example created.
//
// Personal access token routes require an interactive user session: they
// reject PATs and OAuth tokens. So this example always logs in with
// PIPESHUB_TEST_USER_EMAIL and PIPESHUB_TEST_USER_PASSWORD, even when
// OAuth client credentials are set.
//
// Usage (from examples/):
//
//	go run ./personal_access_tokens/list_tokens <path-to-.env>
package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/joho/godotenv"
	pipeshub "github.com/pipeshub-ai/pipeshub-sdk-go"
	"github.com/pipeshub-ai/pipeshub-sdk-go/models/components"

	"enterprise_search/auth"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatal("usage: go run ./personal_access_tokens/list_tokens <path-to-.env>")
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

	ctx := context.Background()

	token, err := createToken(ctx, sdk)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Created token: %s\n\n", token.ID)

	res, err := sdk.PersonalAccessTokens.ListPersonalAccessTokens(ctx)
	if err != nil {
		log.Fatalf("list personal access tokens: %v", err)
	}
	if res.ListPatResponse == nil {
		log.Fatal("list personal access tokens: empty response")
	}

	// Listing never returns token secrets, only their metadata.
	tokens := res.ListPatResponse.Tokens
	fmt.Printf("You have %d active token(s):\n", len(tokens))
	for _, t := range tokens {
		marker := " "
		if t.ID == token.ID {
			marker = "*"
		}
		lastUsed := "never"
		if t.LastUsedAt != nil {
			lastUsed = t.LastUsedAt.Format(time.DateTime)
		}
		fmt.Printf("%s %s  %-28s expires %s  last used %s\n", marker, t.ID, t.Name, t.ExpiresAt.Format(time.DateOnly), lastUsed)
		fmt.Printf("      scopes: %s\n", strings.Join(t.Scopes, ", "))
	}
	fmt.Println("\n* = the token this example created")

	revokeToken(ctx, sdk, token.ID)
}

// createToken mints a 30-day token with a small scope list. The secret in the
// response is shown only once; this example never prints it.
func createToken(ctx context.Context, sdk *pipeshub.Pipeshub) (components.PatWithSecret, error) {
	res, err := sdk.PersonalAccessTokens.CreatePersonalAccessToken(ctx, components.CreatePatRequest{
		Name:   fmt.Sprintf("Go SDK example %d", time.Now().Unix()),
		Scopes: []string{"kb:read", "semantic:read"},
		// 30, 90 or 365 days, or ExpiryDaysNeverNever for no expiry.
		ExpiryDays: pipeshub.Pointer(components.CreateExpiryDaysExpiryDaysEnum(components.ExpiryDaysEnumThirty)),
	})
	if err != nil {
		return components.PatWithSecret{}, fmt.Errorf("create personal access token: %w", err)
	}
	if res.CreatePatResponse == nil {
		return components.PatWithSecret{}, fmt.Errorf("create personal access token: empty response")
	}
	return res.CreatePatResponse.Token, nil
}

// revokeToken revokes a token the example created.
func revokeToken(ctx context.Context, sdk *pipeshub.Pipeshub, tokenID string) {
	if _, err := sdk.PersonalAccessTokens.RevokePersonalAccessToken(ctx, tokenID, &components.RevokePatRequest{
		Reason: pipeshub.Pointer("Go SDK example cleanup"),
	}); err != nil {
		log.Fatalf("revoke personal access token: %v", err)
	}
	fmt.Printf("\nRevoked token %s.\n", tokenID)
}
