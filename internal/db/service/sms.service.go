package service

import (
	"github.com/lishimeng/app-starter"
	persistence "github.com/lishimeng/go-orm"
	"github.com/lishimeng/owl/internal/db/model"
	"github.com/lishimeng/owl/internal/db/repo"
	"github.com/lishimeng/owl/internal/messager/msg"
)

func CreateSmsMessage(template model.SmsTemplateInfo, templateParams string,
	receiver string) (m model.MessageInfo, err error) {
	err = app.GetOrm().Transaction(func(ctx persistence.TxContext) (e error) {
		// create message
		m, e = repo.CreateMessage(ctx, template.Name, msg.Sms)
		if e != nil {
			return
		}
		// create mail
		_, _ = repo.CreateSmsMessage(ctx, m, template, templateParams, receiver)
		return
	})
	return
}

func UpdateSmsTemplate(id, status int, body, description string) (m model.MailTemplateInfo, err error) {
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

func CreateSsi(code, vendor, config string, defaultSender int) (m model.SmsSenderInfo, err error) {
	if err != nil {
		return
	}
	m, err = repo.CreateSmslSenderInfo(code, vendor, config, defaultSender)
	return
}

func UpdateSsi(code, vendor, config string, defaultSender int) (m model.SmsSenderInfo, err error) {
	m, err = repo.GetSmsSenderByCode(code)
	if err != nil {
		return
	}
	var cols []string
	m.Default = defaultSender
	cols = append(cols, "Default")
	m.Config = config
	cols = append(cols, "Config")
	m, err = repo.UpdateSmsSenderInfo(m, cols...)
	return
}
