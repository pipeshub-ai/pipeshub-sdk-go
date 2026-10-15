// Mint a personal access token, then revoke it as an admin with
// AdminRevokePersonalAccessToken. Admins can revoke any user's token this way.
//
// Requires an organization admin. Personal access token routes also require
// an interactive user session: they reject PATs and OAuth tokens. So this
// example always logs in with PIPESHUB_TEST_USER_EMAIL and
// PIPESHUB_TEST_USER_PASSWORD, even when OAuth client
// credentials are set.
//
// Usage (from examples/):
//
//	go run ./personal_access_tokens/admin_revoke_token <path-to-.env>
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

func main() {
	if len(os.Args) < 2 {
		log.Fatal("usage: go run ./personal_access_tokens/admin_revoke_token <path-to-.env>")
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
	fmt.Printf("Created token: %s (%s)\n", token.ID, token.Name)

	res, err := sdk.PersonalAccessTokens.AdminRevokePersonalAccessToken(ctx, token.ID, &components.RevokePatRequest{
		Reason: pipeshub.Pointer("Revoked by an administrator (Go SDK example)"),
	})
	if err != nil {
		log.Fatalf("admin revoke personal access token: %v", err)
	}
	if res.RevokePatResponse == nil {
		log.Fatal("admin revoke personal access token: empty response")
	}
	fmt.Printf("Revoked token: %s (%s)\n", token.ID, res.RevokePatResponse.Message)
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
