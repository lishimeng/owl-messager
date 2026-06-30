package repo

import (
	"github.com/lishimeng/app-starter/persistence"
	"github.com/lishimeng/owl-messager/internal/db/model"
	"github.com/lishimeng/owl-messager/pkg/msg"
)

func GetTemplateByCode(code string, category msg.MessageCategory) (s model.MessageTemplate, err error) {
	err = orm().Model(&model.MessageTemplate{}).
		Equal("code", code).
		Equal("message_category", category).
		First(&s)
	return
}

func GetMessageTemplates(
	org int,
	category msg.MessageCategory,
	provider msg.MessageProvider,
) (templates []model.MessageTemplate, err error) {
	err = orm().Model(&model.MessageTemplate{}).
		Equal("org", org).
		Equal("message_category", category).
		Equal("message_provider", provider).
		Equal("status", model.SenderEnable).
		Order("-Default").
		Limit(10).
		Find(&templates)
	return
}

func GetMessageTemplateById(id int) (tpl model.MessageTemplate, err error) {
	err = orm().Model(&model.MessageTemplate{}).
		Equal("id", id).
		Equal("status", model.SenderEnable).
		First(&tpl)
	return
}

func GetMessageTemplateByCode(code string, org int) (tpl model.MessageTemplate, err error) {
	err = orm().Model(&model.MessageTemplate{}).
		Equal("code", code).
		Equal("org", org).
		Equal("status", model.SenderEnable).
		First(&tpl)
	return
}

func CreateMessageTemplate(
	org int,
	code,
	name,
	body,
	cloudTemplate,
	params,
	description string,
	category msg.MessageCategory,
	provider msg.MessageProvider,
) (m model.MessageTemplate, err error) {
	m = model.MessageTemplate{
		Code:          code,
		Name:          name,
		Params:        params,
		Category:      category,
		Provider:      provider,
		Body:          body,
		CloudTemplate: cloudTemplate,
	}
	m.Org = org
	if len(description) > 0 {
		m.Description = description
	}
	m.Status = model.TemplateEnable
	err = create(&m)
	return
}

func UpdateMessageTemplate(status int, code, name, body, params, description string, provider string) (m model.MessageTemplate, err error) {
	err = orm().Transaction(func(ctx persistence.TxContext) (e error) {
		e = ctx.Model(&model.MessageTemplate{}).Equal("code", code).First(&m)
		if e != nil {
			return
		}
		var cols []string
		if status > ConditionIgnore {
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
		if len(params) > 0 {
			m.Params = params
			cols = append(cols, "Params")
		}
		if len(description) > 0 {
			m.Description = description
			cols = append(cols, "Description")
		}
		if len(provider) > 0 {
			m.Provider = msg.MessageProvider(provider)
			cols = append(cols, "Provider")
		}
		e = updateSelect(ctx, &m, cols...)
		return
	})
	return
}
