package auth

import (
	"context"
	"fmt"
	"net/http"
	"os"

	pipeshub "github.com/pipeshub-ai/pipeshub-sdk-go"
	"github.com/pipeshub-ai/pipeshub-sdk-go/models/components"
)

func NewClient(email, password string) (*pipeshub.Pipeshub, error) {
	baseURL := os.Getenv("PIPESHUB_BASE_URL") + "/api/v1"
	ctx := context.Background()

	s := pipeshub.New(pipeshub.WithServerURL(baseURL))

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
	if authRes.AuthenticateResponse == nil || authRes.AuthenticateResponse.AuthenticateFinalResponse == nil {
		return nil, fmt.Errorf("authenticate: expected final response, got multi-step or empty")
	}
	accessToken := authRes.AuthenticateResponse.AuthenticateFinalResponse.AccessToken

	return pipeshub.New(
		pipeshub.WithServerURL(baseURL),
		pipeshub.WithSecurity(components.Security{BearerAuth: &accessToken}),
	), nil
}
