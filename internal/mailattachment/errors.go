package mailattachment

import "errors"

var (
	errDisallowedMIME   = errors.New("mailattachment: disallowed mime type")
	errFileTooLarge     = errors.New("mailattachment: file exceeds max size")
	errTotalTooLarge    = errors.New("mailattachment: total attachment size exceeds limit")
	errTooManyFiles     = errors.New("mailattachment: too many attachments")
	errNotFound         = errors.New("mailattachment: attachment not found")
	errExpired          = errors.New("mailattachment: attachment expired")
	errOrgMismatch      = errors.New("mailattachment: attachment org mismatch")
	errEmptyFile        = errors.New("mailattachment: empty file")
	errDuplicateID      = errors.New("mailattachment: duplicate attachment id in request")
)
