package model

type MailTemplateInfo struct {
	Pk
	Code string `orm:"column(code);unique"`
	Body string `orm:"column(body)"`
	TableChangeInfo
}