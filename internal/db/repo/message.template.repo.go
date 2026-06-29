package repo

import (
	"github.com/lishimeng/app-starter"
	"github.com/lishimeng/app-starter/persistence"
	"github.com/lishimeng/owl-messager/internal/db/model"
	"github.com/lishimeng/owl-messager/pkg/msg"
)

func GetTemplateByCode(code string, category msg.MessageCategory) (s model.MessageTemplate, err error) {
	err = app.GetOrm().Context.QueryTable(new(model.MessageTemplate)).
		Filter("Code", code).
		Filter("Category", category).
		One(&s)
	return
}

func GetMessageTemplates(
	org int,
	category msg.MessageCategory,
	provider msg.MessageProvider,
) (templates []model.MessageTemplate, err error) {
	_, err = app.GetOrm().Context.QueryTable(new(model.MessageTemplate)).
		Filter("Org", org).
		Filter("Category", category).
		Filter("Provider", provider).
		Filter("Status", model.SenderEnable).
		OrderBy("-Default").
		Limit(10).All(&templates)
	if err != nil {
		return
	}
	return
}

func GetMessageTemplateById(id int) (tpl model.MessageTemplate, err error) {
	err = app.GetOrm().Context.QueryTable(new(model.MessageTemplate)).
		Filter("Id", id).
		//Filter("Org", org).
		Filter("Status", model.SenderEnable).
		One(&tpl)
	if err != nil {
		return
	}

	return
}

func GetMessageTemplateByCode(code string, org int) (tpl model.MessageTemplate, err error) {
	err = app.GetOrm().Context.QueryTable(new(model.MessageTemplate)).
		Filter("Code", code).
		Filter("Org", org).
		Filter("Status", model.SenderEnable).
		One(&tpl)
	if err != nil {
		return
	}

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
	_, err = app.GetOrm().Context.Insert(&m)

	return
}

func UpdateMessageTemplate(status int, code, name, body, params, description string, provider string) (m model.MessageTemplate, err error) {

	err = app.GetOrm().Transaction(func(ctx persistence.TxContext) (e error) {
		e = ctx.Context.QueryTable(new(model.MessageTemplate)).Filter("Code", code).One(&m)
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

		_, err = ctx.Context.Update(&m, cols...)

		return
	})
	return
}
