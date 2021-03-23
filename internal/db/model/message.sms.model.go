package model

// 邮件
type SmsMessageInfo struct {
	MessageHeader
	// 主题
	Subject string `orm:"column(subject)"`
	// 正文
	Body string `orm:"column(body)"`
}
