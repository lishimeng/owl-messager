package service

import (
	"github.com/lishimeng/app-starter"
	persistence "github.com/lishimeng/go-orm"
	"github.com/lishimeng/owl-messager/internal/db/model"
	"github.com/lishimeng/owl-messager/internal/db/repo"
	"github.com/lishimeng/owl-messager/internal/messager/msg"
)

func CreateApnsMessage(sender model.ApnsSenderInfo,
	mode int, bundleId string, params string,
	subject string, receiver string) (m model.MessageInfo, err error) {

	err = app.GetOrm().Transaction(func(ctx persistence.TxContext) (e error) {
		// create message
		m, e = repo.CreateMessage(ctx, subject, msg.Apns)
		if e != nil {
			return
		}
		// create mail
		_, _ = repo.CreateApnsMessage(ctx, m, sender, mode, bundleId, params, subject, receiver)
		return
	})
	return
}
