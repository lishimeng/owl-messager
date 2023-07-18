package model

import "github.com/lishimeng/app-starter"

type MailTemplateInfo struct {
	app.TenantPk
	Code        string `orm:"column(code);unique"`
	Name        string `orm:"column(name)"`
	Body        string `orm:"column(body)"`
	Cloud       int    `orm:"column(cloud)"`
	CloudId     string `orm:"column(cloud_id)"`
	Category    int    `orm:"column(category)"`
	Description string `orm:"column(description);null"`
	Vendor      string `orm:"column(vendor);null"` // vendor
	app.TableChangeInfo
}

const (
	MailTemplateEnable  = 1 // enable
	MailTemplateDisable = 0 // disable
)

const (
	MailCloudTemplate  = 1 // enable
	MailNativeTemplate = 0 // disable
)

const (
	MailTemplateCategoryText    = 1                        // text
	MailTemplateCategoryHtml    = 2                        // html
	MailTemplateCategoryDefault = MailTemplateCategoryText // text
)

var MailTemplateStatus []int

func init() {
	MailTemplateStatus = append(MailTemplateStatus, MailTemplateEnable, MailTemplateDisable)
}
