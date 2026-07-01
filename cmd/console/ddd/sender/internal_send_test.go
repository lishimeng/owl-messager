package sender

import "testing"

func TestShouldUseInternalSend(t *testing.T) {
	if !shouldUseInternalSend("") {
		t.Fatal("empty host should use internal")
	}
	if !shouldUseInternalSend("internal") {
		t.Fatal("internal keyword should use internal")
	}
	if shouldUseInternalSend("http://localhost:81") {
		t.Fatal("explicit messager host should use SDK")
	}
}

func TestResolveMessagerHost(t *testing.T) {
	if got := resolveMessagerHost("http://localhost:81"); got != "http://localhost:81" {
		t.Fatalf("got %q", got)
	}
}

func TestNormalizeMessagerHost(t *testing.T) {
	cases := map[string]string{
		"http://127.0.0.1:81":      "http://127.0.0.1:81",
		"http://127.0.0.1:81/":     "http://127.0.0.1:81",
		"http://127.0.0.1:81/api":  "http://127.0.0.1:81",
		"http://127.0.0.1:81/api/": "http://127.0.0.1:81",
	}
	for input, want := range cases {
		if got := normalizeMessagerHost(input); got != want {
			t.Fatalf("normalizeMessagerHost(%q) = %q, want %q", input, got, want)
		}
	}
}
