// Mint a personal access token with CreatePersonalAccessToken, print its
// details (never the secret), then revoke it.
//
// Personal access token routes require an interactive user session: they
// reject PATs and OAuth tokens. So this example always logs in with
// PIPESHUB_TEST_USER_EMAIL and PIPESHUB_TEST_USER_PASSWORD, even when
// OAuth client credentials are set.
//
// Usage (from examples/):
//
//	go run ./personal_access_tokens/create_token <path-to-.env>
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
		log.Fatal("usage: go run ./personal_access_tokens/create_token <path-to-.env>")
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

	fmt.Printf("Created token: %s\n", token.ID)
	fmt.Printf("  name:    %s\n", token.Name)
	fmt.Printf("  scopes:  %s\n", strings.Join(token.Scopes, ", "))
	fmt.Printf("  created: %s\n", token.CreatedAt.Format(time.RFC1123))
	fmt.Printf("  expires: %s\n", token.ExpiresAt.Format(time.RFC1123))
	// The secret (prefixed "phpat_") is returned only here, once. Store it
	// somewhere safe; it is deliberately not printed.
	fmt.Printf("  secret:  %s… (%d characters, shown only once)\n", prefix(token.AccessToken), len(token.AccessToken))

	revokeToken(ctx, sdk, token.ID)
}

// prefix returns the non-secret "phpat_" marker of a token, if present.
func prefix(token string) string {
	if strings.HasPrefix(token, "phpat_") {
		return "phpat_"
	}
	return ""
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
