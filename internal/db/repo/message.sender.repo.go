package repo

import (
	"errors"
	"github.com/lishimeng/app-starter"
	"github.com/lishimeng/app-starter/persistence"
	"github.com/lishimeng/owl-messager/internal/db/model"
	"github.com/lishimeng/owl-messager/pkg/msg"
)

// GetDefMessageSender 获取默认的消息发送者
func GetDefMessageSender(org int, category msg.MessageCategory, provider msg.MessageProvider) (s model.MessageSenderInfo, err error) {

	var senders []model.MessageSenderInfo
	_, err = app.GetOrm().Context.QueryTable(new(model.MessageSenderInfo)).
		Filter("Org", org).
		Filter("Category", category).
		Filter("Provider", provider).
		Filter("Status", model.SenderEnable).
		OrderBy("-Default").
		Limit(10).All(&senders)
	if err != nil {
		return
	}
	if len(senders) <= 0 {
		err = errors.New("not found")
		return
	}
	s = senders[0]
	err = s.Config.Decode()
	return
}

func CreateMessageSender(
	org int,
	category msg.MessageCategory,
	provider msg.MessageProvider,
	isDefault int,
	code string,
	config msg.SenderConfig,
) (model.MessageSenderInfo, error) {
	m := model.MessageSenderInfo{
		Code:     code,
		Category: category,
		Provider: provider,
		Default:  isDefault,
		Config:   config,
	}
	err := m.Config.Encode()
	if err != nil {
		return m, err
	}
	m.Org = org
	m.Status = model.SenderEnable
	_, err = app.GetOrm().Context.Insert(&m)
	return m, err
}

func GetMessageSenderByCode(code string) (s model.MessageSenderInfo, err error) {
	err = app.GetOrm().Context.QueryTable(new(model.MessageSenderInfo)).
		Filter("Code", code).
		//Filter("Org", org).
		One(&s)
	return
}

func UpdateMessageSender(code string, config msg.SenderConfig) (s model.MessageSenderInfo, err error) {

	err = app.GetOrm().Transaction(func(ctx persistence.TxContext) (e error) {
		e = ctx.Context.QueryTable(new(model.MessageSenderInfo)).Filter("Code", code).One(&s)
		if e != nil {
			return
		}
		var cols []string
		if len(config) > 0 {
			s.Config = config
			err = s.Config.Encode()
			if err != nil {
				return err
			}
			cols = append(cols, "Config")
		}
		_, err = ctx.Context.Update(&s, cols...)

		return
	})
	return
}
