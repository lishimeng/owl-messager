package mailattachment

import (
	"mime"
	"path/filepath"
	"strings"
)

var allowedMIME = map[string]struct{}{
	"application/pdf": {},
	"image/png":       {},
	"image/jpeg":      {},
	"image/gif":       {},
	"text/plain":      {},
	"text/csv":        {},
	"application/zip": {},
	"application/vnd.openxmlformats-officedocument.wordprocessingml.document": {},
	"application/vnd.openxmlformats-officedocument.spreadsheetml.sheet":       {},
}

func validateMIME(name, detected string) error {
	mimeType := detected
	if mimeType == "" || mimeType == "application/octet-stream" {
		mimeType = mime.TypeByExtension(filepath.Ext(name))
	}
	if mimeType == "" {
		return errDisallowedMIME
	}
	base := strings.Split(mimeType, ";")[0]
	if _, ok := allowedMIME[strings.TrimSpace(base)]; !ok {
		return errDisallowedMIME
	}
	return nil
}
