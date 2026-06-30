package repo

import (
	"github.com/lishimeng/app-starter/persistence"
	"github.com/lishimeng/owl-messager/internal/db/model"
)

func GetApnsByMessageId(msgId int) (m model.ApnsMessageInfo, err error) {
	err = orm().Model(&model.ApnsMessageInfo{}).Equal("message_id", msgId).First(&m)
	return
}

func CreateApnsMessage(ctx persistence.TxContext, message model.MessageInfo, sender model.MessageSenderInfo,
	mode int, bundleId string, params string,
	receiver string) (m model.ApnsMessageInfo, err error) {

	initPushChannel(&m.PushChannelDetail, message, sender, mode, bundleId, params, receiver)
	err = ctx.Create(&m)
	return
}
