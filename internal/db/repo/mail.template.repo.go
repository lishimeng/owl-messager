package repo

import (
	"github.com/lishimeng/app-starter"
	"github.com/lishimeng/owl/internal/db/model"
	"time"
)

// 查询邮件Template
func GetMailTemplateByCode(code string) (s model.MailTemplateInfo, err error) {
	err = app.GetOrm().Context.QueryTable(new(model.MailTemplateInfo)).Filter("Code", code).One(&s)
	return
}

// 查询邮件Template
func GetMailTemplateById(id int) (s model.MailTemplateInfo, err error) {
	s.Id = id
	err = app.GetOrm().Context.Read(&s)
	return
}

// 查询邮件Template列表
func GetMailTemplateList() (s []model.MailTemplateInfo, err error) {
	_, err = app.GetOrm().Context.QueryTable(new(model.MailSenderInfo)).All(&s)
	return
}

func DeleteMailTemplate(id int) (err error) {
	var t model.MailTemplateInfo
	t.Id = id
	_, err = app.GetOrm().Context.Delete(&t)
	return
}

func CreateMailTemplate(code, body, description string) (m model.MailTemplateInfo, err error) {
	tci := model.TableChangeInfo{
		Status:     10,
		CreateTime: time.Now(),
		UpdateTime: time.Now(),
	}
	m = model.MailTemplateInfo{
		Code:            code,
		Body:            body,
		TableChangeInfo: tci,
	}
	if len(description) > 0 {
		m.Description = description
	}
	_, err = app.GetOrm().Context.Insert(&m)

	return
}
