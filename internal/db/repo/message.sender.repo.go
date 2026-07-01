package repo

import (
	"errors"

	"github.com/lishimeng/app-starter/persistence"
	"github.com/lishimeng/owl-messager/internal/db/model"
	"github.com/lishimeng/owl-messager/pkg/msg"
)

func GetDefMessageSender(tenantCode string, category msg.MessageCategory, provider msg.MessageProvider) (s model.MessageSenderInfo, err error) {
	provider = msg.NormalizeProvider(provider)
	var senders []model.MessageSenderInfo
	err = orm().Model(&model.MessageSenderInfo{}).
		Equal("tenant_code", tenantCode).
		Equal("category", category).
		Equal("message_provider", provider).
		Equal("status", model.SenderEnable).
		Order("-Default").
		Limit(10).
		Find(&senders)
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
	tenantCode string,
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
	m.TenantCode = tenantCode
	m.Status = model.SenderEnable
	err = create(&m)
	return m, err
}

func GetMessageSenderByCode(code string) (s model.MessageSenderInfo, err error) {
	err = orm().Model(&model.MessageSenderInfo{}).Equal("code", code).First(&s)
	return
}

func GetMessageSenderById(id int) (s model.MessageSenderInfo, err error) {
	err = orm().Model(&model.MessageSenderInfo{}).Equal("id", id).First(&s)
	if err != nil {
		return
	}
	err = s.Config.Decode()
	return
}

func UpdateMessageSender(code string, config msg.SenderConfig) (s model.MessageSenderInfo, err error) {
	err = orm().Transaction(func(ctx persistence.TxContext) (e error) {
		e = ctx.Model(&model.MessageSenderInfo{}).Equal("code", code).First(&s)
		if e != nil {
			return
		}
		if len(config) > 0 {
			s.Config = config
			e = s.Config.Encode()
			if e != nil {
				return
			}
			e = updateSelect(ctx, &s, "Config")
		}
		return
	})
	return
}
