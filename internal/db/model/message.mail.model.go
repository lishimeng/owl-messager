package model

// MailMessageInfo 邮件渠道详情；主题见 MessageInfo.Subject
type MailMessageInfo struct {
	TemplateChannelDetail
	Attachments string `gorm:"column:attachments"` // JSON，预留附件列表
}
