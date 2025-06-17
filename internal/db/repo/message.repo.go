package repo

import (
	"github.com/lishimeng/app-starter"
	"github.com/lishimeng/app-starter/persistence"
	"github.com/lishimeng/go-log"
	"github.com/lishimeng/owl-messager/internal/db/model"
	"github.com/lishimeng/owl-messager/pkg/msg"
	"time"
)

func GetMessageById(id int) (m model.MessageInfo, err error) {
	log.Debug("get message from db: %d", id)
	m.Id = id
	err = app.GetOrm().Context.Read(&m)
	return
}

// 查询需要发送的消息
func GetMessageToSend(size int) (messages []model.MessageInfo, err error) {
	_, err = app.GetOrm().Context.
		QueryTable(new(model.MessageInfo)).
		Filter("Status", model.MessageInit).
		OrderBy("UpdateTime").Limit(size).All(&messages)
	return
}

func UpdateMessageStatus(ctx persistence.TxContext, id int, status int) (m model.MessageInfo, err error) {
	m.Id = id
	m.Status = status
	m.UpdateTime = time.Now()
	_, err = ctx.Context.Update(&m, "Status")
	return
}

func UpdateMessagePriority(id int, priority int) (m model.MessageInfo, err error) {
	m.Id = id
	m.Priority = priority
	m.UpdateTime = time.Now()
	_, err = app.GetOrm().Context.Update(&m, "Priority")
	return
}

func CreateMessage(ctx persistence.TxContext, org int, subject string, category msg.MessageCategory) (m model.MessageInfo, err error) {
	log.Debug("create message %s[category:%s]", subject, category)
	m.Org = org
	m.Subject = subject
	m.Priority = model.MessagePriorityNormal
	m.Category = category
	m.Status = model.MessageInit

	_, err = ctx.Context.Insert(&m)
	return
}
