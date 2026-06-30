package mailattachment

import (
	"encoding/json"
	"fmt"
	"io"
	"mime"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"time"

	"github.com/iris-contrib/go.uuid"
	"github.com/lishimeng/owl-messager/internal/messager"
	"github.com/lishimeng/owl-messager/pkg/msg"
)

type storedMeta struct {
	msg.MailAttachmentRef
	Org       int       `json:"org"`
	ExpiresAt time.Time `json:"expiresAt"`
	Bound     bool      `json:"bound"`
}

type Manager struct {
	cfg Config
	mu  sync.Mutex
}

var defaultMgr *Manager

func Init(cfg Config) {
	defaultMgr = &Manager{cfg: cfg.withDefaults()}
	_ = os.MkdirAll(defaultMgr.cfg.Dir, 0o750)
}

func Default() *Manager {
	if defaultMgr == nil {
		Init(Config{})
	}
	return defaultMgr
}

func (m *Manager) Config() Config {
	return m.cfg
}

func (m *Manager) Save(org int, filename string, contentType string, r io.Reader) (msg.MailAttachmentRef, error) {
	if org <= 0 {
		return msg.MailAttachmentRef{}, errOrgMismatch
	}
	name := filepath.Base(strings.TrimSpace(filename))
	if name == "" || name == "." {
		return msg.MailAttachmentRef{}, errEmptyFile
	}

	id := strings.ReplaceAll(uuid.Must(uuid.NewV4()).String(), "-", "")
	dir := m.itemDir(org, id)
	if err := os.MkdirAll(dir, 0o750); err != nil {
		return msg.MailAttachmentRef{}, err
	}

	dataPath := m.dataPath(org, id)
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

	ref := msg.MailAttachmentRef{
		ID:   id,
		Name: name,
		Mime: strings.Split(detected, ";")[0],
		Size: n,
	}
	meta := storedMeta{
		MailAttachmentRef: ref,
		Org:               org,
		ExpiresAt:         time.Now().Add(m.cfg.StagingTTL),
	}
	if err = m.writeMeta(org, id, meta); err != nil {
		_ = os.RemoveAll(dir)
		return msg.MailAttachmentRef{}, err
	}
	return ref, nil
}

func (m *Manager) Resolve(org int, ids []string) ([]msg.MailAttachmentRef, error) {
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

		meta, err := m.readMeta(org, id)
		if err != nil {
			return nil, err
		}
		if meta.Org != org {
			return nil, errOrgMismatch
		}
		if time.Now().After(meta.ExpiresAt) {
			return nil, errExpired
		}
		total += meta.Size
		if total > m.cfg.MaxTotalSize {
			return nil, errTotalTooLarge
		}
		refs = append(refs, meta.MailAttachmentRef)
	}
	return refs, nil
}

func (m *Manager) MarkBound(org int, ids []string) error {
	for _, id := range ids {
		id = strings.TrimSpace(id)
		if id == "" {
			continue
		}
		meta, err := m.readMeta(org, id)
		if err != nil {
			return err
		}
		if meta.Org != org {
			return errOrgMismatch
		}
		meta.Bound = true
		meta.ExpiresAt = time.Now().Add(m.cfg.StagingTTL)
		if err = m.writeMeta(org, id, meta); err != nil {
			return err
		}
	}
	return nil
}

func (m *Manager) LoadForSend(org int, attachmentsJSON string) ([]messager.MailAttachment, error) {
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
		meta, err := m.readMeta(org, ref.ID)
		if err != nil {
			return nil, err
		}
		if meta.Org != org {
			return nil, errOrgMismatch
		}
		if time.Now().After(meta.ExpiresAt) {
			return nil, errExpired
		}
		total += meta.Size
		if total > m.cfg.MaxTotalSize {
			return nil, errTotalTooLarge
		}
		data, err := os.ReadFile(m.dataPath(org, ref.ID))
		if err != nil {
			return nil, err
		}
		if int64(len(data)) != meta.Size {
			return nil, fmt.Errorf("mailattachment: size mismatch for %s", ref.ID)
		}
		out = append(out, messager.MailAttachment{
			Filename:    meta.Name,
			ContentType: meta.Mime,
			Data:        data,
		})
	}
	return out, nil
}

func (m *Manager) ScheduleRemove(org int, attachmentsJSON string) {
	refs, err := msg.ParseMailAttachmentRefs(attachmentsJSON)
	if err != nil || len(refs) == 0 {
		return
	}
	delay := m.cfg.RetainAfterSend
	if delay <= 0 {
		for _, ref := range refs {
			_ = m.remove(org, ref.ID)
		}
		return
	}
	ids := append([]msg.MailAttachmentRef(nil), refs...)
	go func() {
		time.Sleep(delay)
		for _, ref := range ids {
			_ = m.remove(org, ref.ID)
		}
	}()
}

func (m *Manager) CleanupExpired() {
	entries, err := os.ReadDir(m.cfg.Dir)
	if err != nil {
		return
	}
	now := time.Now()
	for _, orgEntry := range entries {
		if !orgEntry.IsDir() {
			continue
		}
		orgDir := filepath.Join(m.cfg.Dir, orgEntry.Name())
		ids, err := os.ReadDir(orgDir)
		if err != nil {
			continue
		}
		for _, idEntry := range ids {
			if !idEntry.IsDir() {
				continue
			}
			id := idEntry.Name()
			org := parseOrgDir(orgEntry.Name())
			meta, err := m.readMeta(org, id)
			if err != nil {
				_ = os.RemoveAll(filepath.Join(orgDir, id))
				continue
			}
			if now.After(meta.ExpiresAt) {
				_ = m.remove(org, id)
			}
		}
	}
}

func (m *Manager) itemDir(org int, id string) string {
	return filepath.Join(m.cfg.Dir, fmt.Sprintf("%d", org), id)
}

func (m *Manager) dataPath(org int, id string) string {
	return filepath.Join(m.itemDir(org, id), "data")
}

func (m *Manager) metaPath(org int, id string) string {
	return filepath.Join(m.itemDir(org, id), "meta.json")
}

func (m *Manager) readMeta(org int, id string) (storedMeta, error) {
	b, err := os.ReadFile(m.metaPath(org, id))
	if err != nil {
		if os.IsNotExist(err) {
			return storedMeta{}, errNotFound
		}
		return storedMeta{}, err
	}
	var meta storedMeta
	if err = json.Unmarshal(b, &meta); err != nil {
		return storedMeta{}, err
	}
	return meta, nil
}

func (m *Manager) writeMeta(org int, id string, meta storedMeta) error {
	b, err := json.Marshal(meta)
	if err != nil {
		return err
	}
	return os.WriteFile(m.metaPath(org, id), b, 0o640)
}

func (m *Manager) remove(org int, id string) error {
	return os.RemoveAll(m.itemDir(org, id))
}

func parseOrgDir(name string) int {
	var org int
	_, _ = fmt.Sscanf(name, "%d", &org)
	return org
}

func mimeTypeByName(name string) string {
	ext := filepath.Ext(name)
	if ext == "" {
		return ""
	}
	return mime.TypeByExtension(ext)
}
