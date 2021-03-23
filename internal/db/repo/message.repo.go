package repo

import (
	"github.com/lishimeng/app-starter"
	"github.com/lishimeng/go-log"
	"github.com/lishimeng/owl/internal/db/model"
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

func GetMessageByStatus(status int) (messages []model.MessageInfo, err error) {
	return
}

func UpdateMessageStatus(id int, status int) (m model.MessageInfo, err error) {
	m.Id = id
	m.Status = status
	m.UpdateTime = time.Now()
	_, err = app.GetOrm().Context.Update(&m, "Status")
	return
}

func UpdateMessagePriority(id int, priority int) (m model.MessageInfo, err error) {
	m.Id = id
	m.Priority = priority
	m.UpdateTime = time.Now()
	_, err = app.GetOrm().Context.Update(&m, "Priority")
	return
}

func CreateMessage(subject string, category int) (m model.MessageInfo, err error) {
	log.Debug("create message %s[category:%d]", subject, category)
	m.Subject = subject
	m.Priority = model.MessagePriorityNormal
	m.Category = category
	var tci = model.TableChangeInfo{
		Status:     model.MessageInit,
		CreateTime: time.Now(),
		UpdateTime: time.Now(),
	}
	m.TableChangeInfo = tci
	_, err = app.GetOrm().Context.Insert(&m)
	return
}
