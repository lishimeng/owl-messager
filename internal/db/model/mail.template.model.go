package model

type MailTemplateInfo struct {
	Pk
	Code string `orm:"column(code);unique"`
	Body string `orm:"column(body)"`
	TableChangeInfo
}

const (
	MailTemplateEnable  = 1 // enable
	MailTemplateDisable = 0 // disable
)
