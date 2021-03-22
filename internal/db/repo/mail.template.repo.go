package repo

import (
	"github.com/lishimeng/app-starter"
	"github.com/lishimeng/owl/internal/db/model"
)

// 查询邮件Template
func GetMailTemplateByCode(code string) (s model.MailTemplateInfo,err error) {
	err = app.GetOrm().Context.QueryTable(new(model.MailTemplateInfo)).Filter("Code", code).One(&s)
	return
}

// 查询邮件Template
func GetMailTemplateById(id int) (s model.MailTemplateInfo,err error) {
	s.Id = id
	err = app.GetOrm().Context.Read(&s)
	return
}

// 查询邮件Template列表
func GetMailTemplateList() (s []model.MailTemplateInfo,err error) {
	_, err = app.GetOrm().Context.QueryTable(new(model.MailSenderInfo)).All(&s)
	return
}
