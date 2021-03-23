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
