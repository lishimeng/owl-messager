package model

// 邮件发送账号
type MailSenderInfo struct {
	Pk
	Code   string `orm:"column(code)"`
	Host   string `orm:"column(host)"`
	Port   int    `orm:"column(port)"`
	Email  string `orm:"column(email)"`
	Alias  string `orm:"column(alias)"`
	Passwd string `orm:"column(password)"`

	TableChangeInfo
}
