package mailattachment

import (
	"fmt"
	"strings"
	"testing"
	"time"

	"github.com/lishimeng/owl-messager/internal/db/model"
)

type memAttachmentStore struct {
	rows map[string]model.MailAttachment
}

func (s *memAttachmentStore) key(tenantCode string, id string) string {
	return fmt.Sprintf("%s:%s", tenantCode, id)
}

func (s *memAttachmentStore) Create(row *model.MailAttachment) error {
	if s.rows == nil {
		s.rows = make(map[string]model.MailAttachment)
	}
	s.rows[s.key(row.TenantCode, row.AttachmentId)] = *row
	return nil
}

func (s *memAttachmentStore) Get(tenantCode string, attachmentId string) (model.MailAttachment, error) {
	row, ok := s.rows[s.key(tenantCode, attachmentId)]
	if !ok {
		return model.MailAttachment{}, errNotFound
	}
	return row, nil
}

func (s *memAttachmentStore) UpdateBound(tenantCode string, attachmentId string, bound bool, expiresAt time.Time) error {
	row, err := s.Get(tenantCode, attachmentId)
	if err != nil {
		return err
	}
	row.Bound = bound
	row.ExpiresAt = expiresAt
	s.rows[s.key(tenantCode, attachmentId)] = row
	return nil
}

func (s *memAttachmentStore) Delete(tenantCode string, attachmentId string) error {
	delete(s.rows, s.key(tenantCode, attachmentId))
	return nil
}

func (s *memAttachmentStore) ListExpired(before time.Time) ([]model.MailAttachment, error) {
	var out []model.MailAttachment
	for _, row := range s.rows {
		if row.ExpiresAt.Before(before) {
			out = append(out, row)
		}
	}
	return out, nil
}

func TestValidateMIME(t *testing.T) {
	if err := validateMIME("a.pdf", "application/pdf"); err != nil {
		t.Fatal(err)
	}
	if err := validateMIME("a.exe", "application/octet-stream"); err == nil {
		t.Fatal("expected disallowed mime")
	}
}

func TestGenAttachmentIDIsHex(t *testing.T) {
	id := genAttachmentID("a.pdf", "default", time.Now())
	if len(id) != 32 {
		t.Fatalf("expected md5 hex length 32, got %d", len(id))
	}
	for _, c := range id {
		if (c < '0' || c > '9') && (c < 'a' || c > 'f') {
			t.Fatalf("expected hex id, got %q", id)
		}
	}
}

func TestManagerSaveResolve(t *testing.T) {
	dir := t.TempDir()
	store := &memAttachmentStore{}
	m := newManager(Config{
		Dir:          dir,
		MaxFileSize:  1024,
		MaxTotalSize: 2048,
		MaxCount:     3,
		StagingTTL:   time.Hour,
	}, store)

	ref, err := m.Save("tenant-a", "test.txt", "text/plain", strings.NewReader("hello"))
	if err != nil {
		t.Fatal(err)
	}
	if ref.ID == "" {
		t.Fatal("empty id")
	}

	refs, err := m.Resolve("tenant-a", []string{ref.ID})
	if err != nil || len(refs) != 1 {
		t.Fatalf("resolve: %v %v", refs, err)
	}
	_, err = m.Resolve("tenant-b", []string{ref.ID})
	if err != errNotFound {
		t.Fatalf("expected not found for other tenant, got %v", err)
	}
}

func TestDataPathIncludesDate(t *testing.T) {
	m := newManager(Config{Dir: "/data"}, &memAttachmentStore{})
	p := m.dataPath("2026-06-29", "acme", "abc")
	if !strings.Contains(p, "2026-06-29") || !strings.Contains(p, "acme") || !strings.Contains(p, "abc") {
		t.Fatalf("unexpected path: %s", p)
	}
}
