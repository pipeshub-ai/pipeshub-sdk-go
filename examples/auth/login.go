package auth

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"time"

	pipeshub "github.com/pipeshub-ai/pipeshub-sdk-go"
	"github.com/pipeshub-ai/pipeshub-sdk-go/models/components"
)

// http.Client.Timeout covers the whole request, including reading the body.
// The SDK default is 60s, which cuts a long SSE stream off part way through
// the answer, so the examples use a longer one.
func streamFriendlyClient() *http.Client {
	return &http.Client{Timeout: 5 * time.Minute}
}

func NewClient(email, password string) (*pipeshub.Pipeshub, error) {
	baseURL := os.Getenv("PIPESHUB_BASE_URL")
	if baseURL == "" {
		baseURL = "http://localhost:3000"
	}
	baseURL += "/api/v1"
	ctx := context.Background()

	if token := os.Getenv("PIPESHUB_BEARER_AUTH"); token != "" {
		return pipeshub.New(
			pipeshub.WithServerURL(baseURL),
			pipeshub.WithClient(streamFriendlyClient()),
			pipeshub.WithSecurity(components.Security{BearerAuth: &token}),
		), nil
	}

	s := pipeshub.New(
		pipeshub.WithServerURL(baseURL),
		pipeshub.WithClient(streamFriendlyClient()),
	)

	initRes, err := s.UserAccount.InitAuth(ctx, &components.InitAuthRequest{Email: &email})
	if err != nil {
		return nil, fmt.Errorf("init auth: %w", err)
	}
	sessionToken := http.Header(initRes.Headers).Get("x-session-token")

	authRes, err := s.UserAccount.Authenticate(ctx, sessionToken, components.AuthenticateRequest{
		Method: components.MethodPassword,
		Credentials: components.CreateCredentialsPasswordCredentials(
			components.PasswordCredentials{Password: password},
		),
	})
	if err != nil {
		return nil, fmt.Errorf("authenticate: %w", err)
	}
	if authRes == nil || authRes.AuthenticateResponse == nil || authRes.AuthenticateResponse.AuthenticateFinalResponse == nil {
		return nil, fmt.Errorf("authenticate: expected final response, got multi-step or empty")
	}
	accessToken := authRes.AuthenticateResponse.AuthenticateFinalResponse.AccessToken

	return pipeshub.New(
		pipeshub.WithServerURL(baseURL),
		pipeshub.WithClient(streamFriendlyClient()),
		pipeshub.WithSecurity(components.Security{BearerAuth: &accessToken}),
	), nil
}
