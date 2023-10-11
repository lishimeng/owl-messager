package repo

import (
	"errors"
	"github.com/lishimeng/app-starter"
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
