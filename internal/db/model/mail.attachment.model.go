package model

import (
	"time"

	"github.com/lishimeng/app-starter"
)

// MailAttachment 邮件附件元数据（文件内容在本地 staging 目录）
type MailAttachment struct {
	app.TenantPk
	AttachmentId string    `gorm:"column:attachment_id;uniqueIndex;size:32"`
	FileName     string    `gorm:"column:file_name"`
	Mime         string    `gorm:"column:mime"`
	Size         int64     `gorm:"column:size"`
	StorageDate  string    `gorm:"column:storage_date;index;size:10"` // YYYY-MM-DD，与磁盘路径一致
	ExpiresAt    time.Time `gorm:"column:expires_at;index"`
	Bound        bool      `gorm:"column:bound"`
	app.TableChangeInfo
}
