package model

// 邮件
type MailMessageInfo struct {
	MessageHeader

	SenderCode string // sender's code

	Receivers string // receiver list. comma split

	// 主题
	Subject string
	// 正文
	Body string
}