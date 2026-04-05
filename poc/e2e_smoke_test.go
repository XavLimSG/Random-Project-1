//go:build e2e

package poc

import (
	"os"
	"testing"
)

func TestE2ESmoke(t *testing.T) {
	if os.Getenv("API_TOKEN") == "" {
		t.Skip("API_TOKEN not set")
	}
}
