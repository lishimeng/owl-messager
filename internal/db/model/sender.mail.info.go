package model

// 邮件发送账号
type MailSenderInfo struct {
	Pk
	Code   string `orm:"column(code);unique"`
	Name   string `orm:"column(name)"`
	Host   string `orm:"column(host)"`
	Port   int    `orm:"column(port)"`
	Email  string `orm:"column(email)"`
	Alias  string `orm:"column(alias);null"`
	Passwd string `orm:"column(password)"`

	TableChangeInfo
}
