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
		new(model.MailTemplateInfo),
		new(model.MailSenderInfo),
		new(model.SmsMessageInfo),
		new(model.SmsSenderInfo),
		new(model.SmsTemplateInfo),
		new(model.ApnsMessageInfo),
		new(model.ApnsSenderInfo),
	)
	return
}
