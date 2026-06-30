// Stamp an x-request-id trace header onto every outbound SDK request.

package hooks

import (
	"crypto/rand"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

const requestIDHeader = "x-request-id"

type RequestIDHook struct{}

var _ beforeRequestHook = (*RequestIDHook)(nil)

// userIDFromToken decodes the JWT payload (without verifying the signature) and
// returns the `userId` claim, or "" when it cannot be found.
func userIDFromToken(req *http.Request) string {
	auth := req.Header.Get("Authorization")
	if !strings.HasPrefix(strings.ToLower(auth), "bearer ") {
		return ""
	}
	parts := strings.Split(strings.TrimSpace(auth[7:]), ".")
	if len(parts) < 2 || parts[1] == "" {
		return ""
	}
	// JWT segments are unpadded base64url; trim any stray padding to be safe.
	payload, err := base64.RawURLEncoding.DecodeString(strings.TrimRight(parts[1], "="))
	if err != nil {
		return ""
	}
	var claims map[string]any
	if err := json.Unmarshal(payload, &claims); err != nil {
		return ""
	}
	if userID, ok := claims["userId"].(string); ok {
		return userID
	}
	return ""
}

func (h *RequestIDHook) BeforeRequest(hookCtx BeforeRequestContext, req *http.Request) (*http.Request, error) {
	if req.Header.Get(requestIDHeader) != "" {
		return req, nil
	}
	buf := make([]byte, 16)
	if _, err := rand.Read(buf); err != nil {
		return req, nil
	}
	random := hex.EncodeToString(buf)
	if userID := userIDFromToken(req); userID != "" {
		req.Header.Set(requestIDHeader, fmt.Sprintf("sdk-go-%s-%s", userID, random))
	} else {
		req.Header.Set(requestIDHeader, fmt.Sprintf("sdk-go-%s", random))
	}
	return req, nil
}
