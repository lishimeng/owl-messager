package mailattachment

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"io"
	"mime"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"time"

	"github.com/lishimeng/owl-messager/internal/db/model"
	"github.com/lishimeng/owl-messager/internal/db/repo"
	"github.com/lishimeng/owl-messager/internal/messager"
	"github.com/lishimeng/owl-messager/pkg/msg"
)

type attachmentMetaStore interface {
	Create(row *model.MailAttachment) error
	Get(tenantCode string, attachmentId string) (model.MailAttachment, error)
	UpdateBound(tenantCode string, attachmentId string, bound bool, expiresAt time.Time) error
	Delete(tenantCode string, attachmentId string) error
	ListExpired(before time.Time) ([]model.MailAttachment, error)
}

type dbAttachmentMetaStore struct{}

func (dbAttachmentMetaStore) Create(row *model.MailAttachment) error {
	return repo.CreateMailAttachment(row)
}
func (dbAttachmentMetaStore) Get(tenantCode string, attachmentId string) (model.MailAttachment, error) {
	return repo.GetMailAttachment(tenantCode, attachmentId)
}
func (dbAttachmentMetaStore) UpdateBound(tenantCode string, attachmentId string, bound bool, expiresAt time.Time) error {
	return repo.UpdateMailAttachmentBound(tenantCode, attachmentId, bound, expiresAt)
}
func (dbAttachmentMetaStore) Delete(tenantCode string, attachmentId string) error {
	return repo.DeleteMailAttachment(tenantCode, attachmentId)
}
func (dbAttachmentMetaStore) ListExpired(before time.Time) ([]model.MailAttachment, error) {
	return repo.ListExpiredMailAttachments(before)
}

type Manager struct {
	cfg   Config
	mu    sync.Mutex
	store attachmentMetaStore
}

var defaultMgr *Manager

func Init(cfg Config) {
	defaultMgr = newManager(cfg, dbAttachmentMetaStore{})
}

func newManager(cfg Config, store attachmentMetaStore) *Manager {
	m := &Manager{cfg: cfg.withDefaults(), store: store}
	_ = os.MkdirAll(m.cfg.Dir, 0o750)
	return m
}

func Default() *Manager {
	if defaultMgr == nil {
		defaultMgr = newManager(Config{}, dbAttachmentMetaStore{})
	}
	return defaultMgr
}

func (m *Manager) Config() Config {
	return m.cfg
}

func genAttachmentID(filename string, tenantCode string, at time.Time) string {
	raw := fmt.Sprintf("%s:%s:%s", filename, tenantCode, at.Format(time.RFC3339Nano))
	sum := md5.Sum([]byte(raw))
	return hex.EncodeToString(sum[:])
}

func (m *Manager) itemDir(storageDate string, tenantCode string, id string) string {
	return filepath.Join(m.cfg.Dir, storageDate, tenantCode, id)
}

func (m *Manager) dataPath(storageDate string, tenantCode string, id string) string {
	return filepath.Join(m.itemDir(storageDate, tenantCode, id), "data")
}

func (m *Manager) refFromRow(row model.MailAttachment) msg.MailAttachmentRef {
	return msg.MailAttachmentRef{
		ID:   row.AttachmentId,
		Name: row.FileName,
		Mime: row.Mime,
		Size: row.Size,
	}
}

func (m *Manager) Save(tenantCode string, filename string, contentType string, r io.Reader) (msg.MailAttachmentRef, error) {
	if tenantCode == "" {
		return msg.MailAttachmentRef{}, errOrgMismatch
	}
	name := filepath.Base(strings.TrimSpace(filename))
	if name == "" || name == "." {
		return msg.MailAttachmentRef{}, errEmptyFile
	}

	at := time.Now()
	id := genAttachmentID(name, tenantCode, at)
	storageDate := at.Format("2006-01-02")
	dir := m.itemDir(storageDate, tenantCode, id)
	if err := os.MkdirAll(dir, 0o750); err != nil {
		return msg.MailAttachmentRef{}, err
	}

	dataPath := m.dataPath(storageDate, tenantCode, id)
	f, err := os.OpenFile(dataPath, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0o640)
	if err != nil {
		return msg.MailAttachmentRef{}, err
	}
	defer f.Close()

	limited := io.LimitReader(r, m.cfg.MaxFileSize+1)
	n, err := io.Copy(f, limited)
	if err != nil {
		_ = os.RemoveAll(dir)
		return msg.MailAttachmentRef{}, err
	}
	if n == 0 {
		_ = os.RemoveAll(dir)
		return msg.MailAttachmentRef{}, errEmptyFile
	}
	if n > m.cfg.MaxFileSize {
		_ = os.RemoveAll(dir)
		return msg.MailAttachmentRef{}, errFileTooLarge
	}

	detected := contentType
	if detected == "" {
		detected = mimeTypeByName(name)
	}
	if err = validateMIME(name, detected); err != nil {
		_ = os.RemoveAll(dir)
		return msg.MailAttachmentRef{}, err
	}
	mimeBase := strings.Split(detected, ";")[0]

	row := &model.MailAttachment{
		TenantScope:  model.TenantScope{TenantCode: tenantCode},
		AttachmentId: id,
		FileName:     name,
		Mime:         mimeBase,
		Size:         n,
		StorageDate:  storageDate,
		ExpiresAt:    at.Add(m.cfg.StagingTTL),
		Bound:        false,
	}
	if err = m.store.Create(row); err != nil {
		_ = os.RemoveAll(dir)
		return msg.MailAttachmentRef{}, err
	}
	return m.refFromRow(*row), nil
}

func (m *Manager) Resolve(tenantCode string, ids []string) ([]msg.MailAttachmentRef, error) {
	if len(ids) == 0 {
		return nil, nil
	}
	if len(ids) > m.cfg.MaxCount {
		return nil, errTooManyFiles
	}
	seen := make(map[string]struct{}, len(ids))
	var refs []msg.MailAttachmentRef
	var total int64
	for _, id := range ids {
		id = strings.TrimSpace(id)
		if id == "" {
			continue
		}
		if _, dup := seen[id]; dup {
			return nil, errDuplicateID
		}
		seen[id] = struct{}{}

		row, err := m.store.Get(tenantCode, id)
		if err != nil {
			return nil, errNotFound
		}
		if row.TenantCode != tenantCode {
			return nil, errOrgMismatch
		}
		if time.Now().After(row.ExpiresAt) {
			return nil, errExpired
		}
		total += row.Size
		if total > m.cfg.MaxTotalSize {
			return nil, errTotalTooLarge
		}
		refs = append(refs, m.refFromRow(row))
	}
	return refs, nil
}

func (m *Manager) MarkBound(tenantCode string, ids []string) error {
	for _, id := range ids {
		id = strings.TrimSpace(id)
		if id == "" {
			continue
		}
		row, err := m.store.Get(tenantCode, id)
		if err != nil {
			return err
		}
		if row.TenantCode != tenantCode {
			return errOrgMismatch
		}
		expiresAt := time.Now().Add(m.cfg.StagingTTL)
		if err = m.store.UpdateBound(tenantCode, id, true, expiresAt); err != nil {
			return err
		}
	}
	return nil
}

func (m *Manager) LoadForSend(tenantCode string, attachmentsJSON string) ([]messager.MailAttachment, error) {
	refs, err := msg.ParseMailAttachmentRefs(attachmentsJSON)
	if err != nil {
		return nil, err
	}
	if len(refs) == 0 {
		return nil, nil
	}
	if len(refs) > m.cfg.MaxCount {
		return nil, errTooManyFiles
	}
	var total int64
	var out []messager.MailAttachment
	for _, ref := range refs {
		row, err := m.store.Get(tenantCode, ref.ID)
		if err != nil {
			return nil, err
		}
		if row.TenantCode != tenantCode {
			return nil, errOrgMismatch
		}
		if time.Now().After(row.ExpiresAt) {
			return nil, errExpired
		}
		total += row.Size
		if total > m.cfg.MaxTotalSize {
			return nil, errTotalTooLarge
		}
		data, err := os.ReadFile(m.dataPath(row.StorageDate, tenantCode, row.AttachmentId))
		if err != nil {
			return nil, err
		}
		if int64(len(data)) != row.Size {
			return nil, fmt.Errorf("mailattachment: size mismatch for %s", ref.ID)
		}
		out = append(out, messager.MailAttachment{
			Filename:    row.FileName,
			ContentType: row.Mime,
			Data:        data,
		})
	}
	return out, nil
}

func (m *Manager) ScheduleRemove(tenantCode string, attachmentsJSON string) {
	refs, err := msg.ParseMailAttachmentRefs(attachmentsJSON)
	if err != nil || len(refs) == 0 {
		return
	}
	delay := m.cfg.RetainAfterSend
	if delay <= 0 {
		for _, ref := range refs {
			_ = m.remove(tenantCode, ref.ID)
		}
		return
	}
	ids := append([]msg.MailAttachmentRef(nil), refs...)
	go func() {
		time.Sleep(delay)
		for _, ref := range ids {
			_ = m.remove(tenantCode, ref.ID)
		}
	}()
}

func (m *Manager) CleanupExpired() {
	rows, err := m.store.ListExpired(time.Now())
	if err != nil {
		return
	}
	for _, row := range rows {
		_ = m.remove(row.TenantCode, row.AttachmentId)
	}
}

func (m *Manager) remove(tenantCode string, id string) error {
	row, err := m.store.Get(tenantCode, id)
	if err == nil {
		_ = os.RemoveAll(m.itemDir(row.StorageDate, tenantCode, id))
	}
	return m.store.Delete(tenantCode, id)
}

func mimeTypeByName(name string) string {
	ext := filepath.Ext(name)
	if ext == "" {
		return ""
	}
	return mime.TypeByExtension(ext)
}
