package task

import (
	"os"
	"testing"
)

// TestSendLive is an optional integration test; set SMTP_TEST=1 and related env vars to run.
func TestSendLive(t *testing.T) {
	if os.Getenv("SMTP_TEST") != "1" {
		t.Skip("set SMTP_TEST=1 to run live SMTP integration test")
	}
	t.Skip("configure SMTP_HOST, SMTP_PORT, SMTP_USER, SMTP_PASS, SMTP_TO env vars before enabling")
}
