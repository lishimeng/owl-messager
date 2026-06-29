package smtp

import (
	"strings"
	"testing"
)

func TestBuildMessage(t *testing.T) {
	msg := string(buildMessage(
		"Owl <noreply@example.com>",
		"hello",
		"<p>hi</p>",
		[]string{"a@example.com", "b@example.com"},
	))

	if !strings.Contains(msg, "From: Owl <noreply@example.com>") {
		t.Fatalf("missing From header: %s", msg)
	}
	if !strings.Contains(msg, "To: a@example.com, b@example.com") {
		t.Fatalf("missing To header: %s", msg)
	}
	if !strings.Contains(msg, "Subject: hello") {
		t.Fatalf("missing Subject header: %s", msg)
	}
	if !strings.Contains(msg, "Content-Type: text/html; charset=UTF-8") {
		t.Fatalf("missing Content-Type header: %s", msg)
	}
	if !strings.Contains(msg, "<p>hi</p>") {
		t.Fatalf("missing body: %s", msg)
	}
}

func TestEncodeFrom(t *testing.T) {
	if encodeFrom("a@b.com", "") != "a@b.com" {
		t.Fatal("bare email expected")
	}
	if encodeFrom("a@b.com", "Alias") != "Alias <a@b.com>" {
		t.Fatal("formatted from expected")
	}
}
