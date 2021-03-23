package model

type MailTemplateInfo struct {
	Pk
	Code     string `orm:"column(code);unique"`
	Body     string `orm:"column(body)"`
	Category int    `orm:"column(category)"`
	TableChangeInfo
}

const (
	MailTemplateEnable  = 1 // enable
	MailTemplateDisable = 0 // disable
)

const (
	MailTemplateCategoryText = 1 // text
	MailTemplateCategoryHtml = 2 // html
)
