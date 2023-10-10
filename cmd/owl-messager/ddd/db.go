package ddd

import "github.com/lishimeng/owl-messager/internal/db/model"

func Tables() (t []interface{}) {
	t = append(t,
		new(model.Tenant),
		new(model.OpenClient),
		new(model.MessageInfo),
		new(model.MessageTask),
		new(model.MessageRunningTask),
		new(model.MailMessageInfo),
		new(model.SmsMessageInfo),
		new(model.ApnsMessageInfo),
		new(model.MessageTemplate),
		new(model.MessageSenderInfo),
	)
	return
}
