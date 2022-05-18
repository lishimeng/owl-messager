package service

import (
	"github.com/lishimeng/app-starter"
	persistence "github.com/lishimeng/go-orm"
	"github.com/lishimeng/owl/internal/db/model"
	"github.com/lishimeng/owl/internal/db/repo"
	"github.com/lishimeng/owl/internal/messager/msg"
)

func CreateMailMessage(sender *model.MailSenderInfo, template model.MailTemplateInfo, templateParams string,
	subject, receiver, cc string) (m model.MessageInfo, err error) {
	err = app.GetOrm().Transaction(func(ctx persistence.TxContext) (e error) {
		// create message
		m, e = repo.CreateMessage(ctx, subject, msg.Email)
		if e != nil {
			return
		}
		// create mail
		_, _ = repo.CreateMailMessage(ctx, m, sender, template, templateParams, subject, receiver, cc)
		return
	})
	return
}

func UpdateMailTemplate(id, status int, body, description string) (m model.MailTemplateInfo, err error) {
	m, err = repo.GetMailTemplateById(id)
	if err != nil {
		return
	}
	err = app.GetOrm().Transaction(func(ctx persistence.TxContext) (e error) {
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

		m, e = repo.UpdateMailTemplate(ctx, m, cols...)

		return
	})
	return
}
