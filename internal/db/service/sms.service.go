package service

import (
	"github.com/lishimeng/app-starter"
	persistence "github.com/lishimeng/go-orm"
	"github.com/lishimeng/owl-messager/internal/db/model"
	"github.com/lishimeng/owl-messager/internal/db/repo"
	"github.com/lishimeng/owl-messager/internal/messager/msg"
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

func UpdateSmsTemplate(id, status int, body, description string) (m model.SmsTemplateInfo, err error) {
	m, err = repo.GetSmsTemplateById(id)
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

		m, e = repo.UpdateSmsTemplateInfo(ctx, m, cols...)

		return
	})
	return
}
func UpdateSmsTemplateByCode(status, sender int, code, name, body, templateId, signature, description, params, vendor string) (m model.SmsTemplateInfo, err error) {
	m, err = repo.GetSmsTemplateByCode(code)
	if err != nil {
		return
	}
	err = app.GetOrm().Transaction(func(ctx persistence.TxContext) (e error) {
		var cols []string
		if status > repo.ConditionIgnore {
			m.Status = status
			cols = append(cols, "Status")
		}
		if sender > 0 {
			m.Sender = sender
			cols = append(cols, "Sender")
		}
		if len(name) > 0 {
			m.Name = name
			cols = append(cols, "Name")
		}
		if len(body) > 0 {
			m.Body = body
			cols = append(cols, "Body")
		}
		if len(templateId) > 0 {
			m.CloudTemplateId = templateId
			cols = append(cols, "CloudTemplateId")
		}
		if len(signature) > 0 {
			m.Signature = signature
			cols = append(cols, "Signature")
		}
		if len(description) > 0 {
			m.Description = description
			cols = append(cols, "Description")
		}
		if len(params) > 0 {
			m.Params = params
			cols = append(cols, "Params")
		}
		if len(vendor) > 0 {
			m.Vendor = vendor
			cols = append(cols, "Vendor")
		}
		m, e = repo.UpdateSmsTemplateInfo(ctx, m, cols...)
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
