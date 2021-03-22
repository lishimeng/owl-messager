package model

type MailTemplateInfo struct {
	Pk
	Code string `orm:"column(code)"`
	Body string `orm:"column(body)"`
	TableChangeInfo
}
