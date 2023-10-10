package service

import (
	"github.com/lishimeng/app-starter"
	"github.com/lishimeng/app-starter/persistence"
	"github.com/lishimeng/owl-messager/internal/db/model"
	"github.com/lishimeng/owl-messager/internal/db/repo"
	"github.com/lishimeng/owl-messager/pkg/msg"
)

func CreateSmsMessage(org int, template model.MessageTemplate, templateParams string,
	receiver string) (m model.MessageInfo, err error) {
	err = app.GetOrm().Transaction(func(ctx persistence.TxContext) (e error) {
		// create message
		m, e = repo.CreateMessage(ctx, org, template.Name, msg.SmsMessage)
		if e != nil {
			return
		}
		// create mail
		_, _ = repo.CreateSmsMessage(ctx, m, template, templateParams, receiver)
		return
	})
	return
}
