package service

import (
	"github.com/lishimeng/app-starter"
	persistence "github.com/lishimeng/go-orm"
	"github.com/lishimeng/owl/internal/db/model"
	"github.com/lishimeng/owl/internal/db/repo"
	"github.com/lishimeng/owl/internal/messager/msg"
)

func CreateMailMessage(sender model.MailSenderInfo, template model.MailTemplateInfo, templateParams string,
	subject, receiver, cc string) (m model.MessageInfo, err error) {
	err = app.GetOrm().Transaction(func(ctx persistence.OrmContext) (e error) {
		// create message
		m, e = repo.CreateMessage(subject, msg.Email)
		if e != nil {
			return
		}
		// create mail
		_, _ = repo.CreateMailMessage(m, sender, template, templateParams, subject, receiver, cc)
		return
	})
	return
}

func UpdateMailTemplate(id, status int, body, description string) (m model.MailTemplateInfo, err error) {
	err = app.GetOrm().Transaction(func(ctx persistence.OrmContext) (e error) {
		m, e = repo.GetMailTemplateById(id)
		if e != nil {
			return
		}
		var cols []string
		if status > repo.ConditionIgnore {
			m.Status = status
			cols = append(cols, "Status")
		}
		if len(body) > 0 {
			m.Body = body
			cols = append(cols, "Body")
		}
		if len(description) > 0 {
			m.Description = description
			cols = append(cols, "Description")
		}

		m, e = repo.UpdateMailTemplate(m, cols...)

		return
	})
	return
}
