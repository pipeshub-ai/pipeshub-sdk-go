package auth

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"strings"
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

// instanceURL is PIPESHUB_BASE_URL without a trailing slash. A trailing slash
// would leave "//api/v1" in every request URL; the SDK only trims a slash at
// the end of the server URL.
func instanceURL() string {
	u := strings.TrimRight(os.Getenv("PIPESHUB_BASE_URL"), "/")
	if u == "" {
		u = "http://localhost:3000"
	}
	return u
}

// NewClient builds a client for the existing examples: PIPESHUB_BEARER_AUTH
// when it is set, otherwise an email/password login through NewSessionClient.
func NewClient(email, password string) (*pipeshub.Pipeshub, error) {
	if token := os.Getenv("PIPESHUB_BEARER_AUTH"); token != "" {
		return pipeshub.New(
			pipeshub.WithServerURL(instanceURL()+"/api/v1"),
			pipeshub.WithClient(streamFriendlyClient()),
			pipeshub.WithSecurity(components.Security{BearerAuth: &token}),
		), nil
	}
	return NewSessionClient(email, password)
}

// NewOAuthClient authenticates with OAuth2 client credentials from
// PIPESHUB_CLIENT_ID and PIPESHUB_CLIENT_SECRET (an OAuth app with the
// client_credentials grant). The SDK fetches the access token, and fetches a
// new one when it expires.
func NewOAuthClient() (*pipeshub.Pipeshub, error) {
	clientID, clientSecret := os.Getenv("PIPESHUB_CLIENT_ID"), os.Getenv("PIPESHUB_CLIENT_SECRET")
	if clientID == "" || clientSecret == "" {
		return nil, fmt.Errorf("PIPESHUB_CLIENT_ID and PIPESHUB_CLIENT_SECRET are required")
	}
	baseURL := instanceURL() + "/api/v1"
	// The SDK's default token URL is a relative path, so always pass an
	// absolute one.
	tokenURL := os.Getenv("PIPESHUB_TOKEN_URL")
	if tokenURL == "" {
		tokenURL = baseURL + "/oauth2/token"
	}
	return pipeshub.New(
		pipeshub.WithServerURL(baseURL),
		pipeshub.WithClient(streamFriendlyClient()),
		pipeshub.WithSecurity(components.Security{Oauth2: &components.SchemeOauth2{
			ClientID:     clientID,
			ClientSecret: clientSecret,
			TokenURL:     tokenURL,
		}}),
	), nil
}

// NewSessionClient logs in with email and password (InitAuth, then
// Authenticate) and uses the returned session JWT. Routes that require an
// interactive user session, such as personal access token management, reject
// OAuth tokens and personal access tokens, so their examples use this.
func NewSessionClient(email, password string) (*pipeshub.Pipeshub, error) {
	baseURL := instanceURL() + "/api/v1"
	ctx := context.Background()

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
