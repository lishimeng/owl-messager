package service

import (
	"github.com/lishimeng/app-starter"
	persistence "github.com/lishimeng/go-orm"
	"github.com/lishimeng/owl-messager/internal/db/model"
	"github.com/lishimeng/owl-messager/internal/db/repo"
	"github.com/lishimeng/owl-messager/internal/messager/msg"
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
		_, _ = repo.CreateMailMessage(ctx, m, template, templateParams, subject, receiver)
		return
	})
	return
}

func CreateCloudMailMessage(sender *model.MailSenderInfo, templateId string, templateParams string,
	subject, receiver, cc string) (m model.MessageInfo, err error) {
	err = app.GetOrm().Transaction(func(ctx persistence.TxContext) (e error) {
		// create message
		m, e = repo.CreateMessage(ctx, subject, msg.Email)
		if e != nil {
			return
		}
		// create mail
		_, _ = repo.CreateCloudMailMessage(ctx, m, templateId, templateParams, subject, receiver)
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

func UpdateMailTemplateByCode(status int, code, name, body, description string) (m model.MailTemplateInfo, err error) {
	m, err = repo.GetMailTemplateByCode(code)
	if err != nil {
		return
	}
	err = app.GetOrm().Transaction(func(ctx persistence.TxContext) (e error) {
		var cols []string
		if status > repo.ConditionIgnore {
			m.Status = status
			cols = append(cols, "Status")
		}
		if len(name) > 0 {
			m.Name = name
			cols = append(cols, "Name")
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

// SetDefaultMailSender 设置默认发送账号
func SetDefaultMailSender(id int, org int) (err error) {

	senders, err := repo.GetMailSenders(org)
	if err != nil {
		return
	}
	for _, s := range senders {
		if s.Id == id {
			s.Default = model.DefaultSenderEnable
		} else {
			s.Default = model.DefaultSenderDisable
		}
	}
	return
}

func CreateMsi(code, vendor, config string, defaultSender int) (m model.MailSenderInfo, err error) {
	if err != nil {
		return
	}
	m, err = repo.CreateMailSenderInfo(code, vendor, config, defaultSender)
	return
}
func UpdateMsi(code, vendor, config string, defaultSender int) (m model.MailSenderInfo, err error) {
	m, err = repo.GetMailSenderByCode(code)
	if err != nil {
		return
	}
	var cols []string
	m.Default = defaultSender
	cols = append(cols, "Default")
	m.Config = config
	cols = append(cols, "Config")
	m, err = repo.UpdateMailSenderInfo(m, cols...)
	return
}
