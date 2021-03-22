package repo

import (
	"github.com/lishimeng/app-starter"
	"github.com/lishimeng/go-log"
	persistence "github.com/lishimeng/go-orm"
	"github.com/lishimeng/owl/internal/db/model"
)

func GetMessageById(id int) (m model.MessageInfo,err error) {
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

func GetMessageByStatus(status int) (messages []model.MessageInfo, err error) {
	return
}

func UpdateMessageStatus(id int, status int) (m model.MessageInfo,err error) {
	err = app.GetOrm().Transaction(func(ctx persistence.OrmContext) (e error) {
		m.Id = id
		e = ctx.Context.Read(&m)
		if e != nil {
			return
		}
		m.Status = status
		_, e = ctx.Context.Update(m, "Status")
		if e != nil {
			return
		}
		return
	})
	return
}

func CreateMessage(subject string, category int) (m model.MessageInfo, err error) {
	log.Debug("create message %s[category:%d]", subject, category)
	m.Subject = subject
	m.Category = category
	m.Status = model.MessageInit
	_, err = app.GetOrm().Context.Insert(&m)
	return
}