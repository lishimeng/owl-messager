package repo

import (
	"github.com/lishimeng/app-starter"
	"github.com/lishimeng/owl/internal/db/model"
)

// 查询邮件发送账号
func GetMailSenderByCode(code string) (s model.MailSenderInfo,err error) {
	err = app.GetOrm().Context.QueryTable(new(model.MailSenderInfo)).Filter("Code", code).One(&s)
	return
}

// 查询邮件发送账号
func GetMailSenderById(id int) (s model.MailSenderInfo,err error) {
	s.Id = id
	err = app.GetOrm().Context.Read(&s)
	return
}

// 查询邮件发送账号列表
func GetMailSenderList() (s []model.MailSenderInfo,err error) {
	return
}