package repo

import (
	"github.com/lishimeng/app-starter"
	"github.com/lishimeng/app-starter/persistence"
	"github.com/lishimeng/owl-messager/internal/db/model"
)

func GetApnsByMessageId(msgId int) (m model.ApnsMessageInfo, err error) {
	err = app.GetOrm().Context.QueryTable(new(model.ApnsMessageInfo)).Filter("MessageId", msgId).One(&m)
	return
}

func CreateApnsMessage(ctx persistence.TxContext, message model.MessageInfo, sender model.MessageSenderInfo,
	mode int, bundleId string, params string,
	subject, receiver string) (m model.ApnsMessageInfo, err error) {

	m.MessageId = message.Id
	m.ApnsMode = model.ApnsMode(mode)
	m.BundleId = bundleId
	m.Params = params
	m.Sender = sender.Id
	m.Subject = subject
	m.Receivers = receiver

	m.Status = model.MessageInit

	_, err = ctx.Context.Insert(&m)
	return
}
