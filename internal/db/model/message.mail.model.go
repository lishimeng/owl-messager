package model

// 邮件
type MailMessageInfo struct {
	MessageHeader

	Template int  // mail template

	Params string // json params

	Sender int // sender's Id

	Receivers string // receiver list. comma split

	Cc string // CC

	// 主题
	Subject string
	// 正文
	Body string
}