package etc

import (
	"time"

	"github.com/lishimeng/owl-messager/internal/mailattachment"
)

func MailAttachmentConfig() mailattachment.Config {
	c := mailattachment.Config{
		Dir:          Config.MailAttachment.Dir,
		MaxFileSize:  Config.MailAttachment.MaxFileSize,
		MaxTotalSize: Config.MailAttachment.MaxTotalSize,
		MaxCount:     Config.MailAttachment.MaxCount,
	}
	if d, err := time.ParseDuration(Config.MailAttachment.StagingTTL); err == nil && d > 0 {
		c.StagingTTL = d
	}
	if d, err := time.ParseDuration(Config.MailAttachment.RetainAfterSend); err == nil {
		c.RetainAfterSend = d
	}
	return c
}
