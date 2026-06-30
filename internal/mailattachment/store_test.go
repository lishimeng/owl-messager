package mailattachment

import (
	"strings"
	"testing"
)

func TestValidateMIME(t *testing.T) {
	if err := validateMIME("a.pdf", "application/pdf"); err != nil {
		t.Fatal(err)
	}
	if err := validateMIME("a.exe", "application/octet-stream"); err == nil {
		t.Fatal("expected disallowed mime")
	}
}

func TestManagerSaveResolve(t *testing.T) {
	dir := t.TempDir()
	m := &Manager{cfg: Config{
		Dir:          dir,
		MaxFileSize:  1024,
		MaxTotalSize: 2048,
		MaxCount:     3,
	}.withDefaults()}

	ref, err := m.Save(1, "test.txt", "text/plain", strings.NewReader("hello"))
	if err != nil {
		t.Fatal(err)
	}
	refs, err := m.Resolve(1, []string{ref.ID})
	if err != nil || len(refs) != 1 {
		t.Fatalf("resolve: %v %v", refs, err)
	}
	_, err = m.Resolve(2, []string{ref.ID})
	if err != errNotFound {
		t.Fatalf("expected not found for other org, got %v", err)
	}
}
