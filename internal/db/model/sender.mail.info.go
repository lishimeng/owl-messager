package model


// 邮件发送账号
type MailSenderInfo struct {
	Pk
	Code string
	Host string
	Port int
	Email string
	Alias string
	Passwd string

	TableChangeInfo
}