package model

// MailMessageInfo 邮件渠道详情；主题见 MessageInfo.Subject
type MailMessageInfo struct {
	TemplateChannelDetail
	Attachments string `gorm:"column:attachments"` // JSON：MailAttachmentRef 列表（id 关联 mail_attachment 表）
}

func (MailMessageInfo) TableName() string {
	return "mail_message_info"
}
