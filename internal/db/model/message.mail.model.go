package model

// 邮件
type MailMessageInfo struct {
	MessageHeader
	// 主题
	Subject string
	// 正文
	Body string
}