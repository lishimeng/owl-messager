package repo

import (
	"github.com/lishimeng/app-starter/persistence"
	"github.com/lishimeng/go-log"
	"github.com/lishimeng/owl-messager/internal/db/model"
	"github.com/lishimeng/owl-messager/pkg/msg"
	"time"
)

func GetMessageById(id int) (m model.MessageInfo, err error) {
	log.Debug("get message from db: %d", id)
	err = orm().Model(&model.MessageInfo{}).Equal("id", id).First(&m)
	return
}

func GetMessageToSend(size int) (messages []model.MessageInfo, err error) {
	now := time.Now()
	err = orm().Model(&model.MessageInfo{}).
		Equal("status", model.MessageInit).
		Where("next_send_time <= ?", now).
		Order("-priority").
		Order("mtime").
		Limit(size).
		Find(&messages)
	return
}

func UpdateMessageStatus(ctx persistence.TxContext, id int, status int) (m model.MessageInfo, err error) {
	m.Id = id
	m.Status = status
	m.UpdateTime = time.Now()
	err = updateSelect(ctx, &m, "Status", "UpdateTime")
	return
}

func UpdateMessagePriority(id int, priority int) (m model.MessageInfo, err error) {
	m.Id = id
	m.Priority = priority
	m.UpdateTime = time.Now()
	err = orm().Model(&m).Select("Priority", []interface{}{"UpdateTime"}...).Updates(&m)
	return
}

func CreateMessage(ctx persistence.TxContext, org int, subject string, category msg.MessageCategory) (m model.MessageInfo, err error) {
	log.Debug("create message %s[category:%s]", subject, category)
	m.Org = org
	m.Subject = subject
	m.Priority = model.MessagePriorityNormal
	m.Category = category
	m.Status = model.MessageInit
	m.NextSendTime = time.Now()
	err = ctx.Create(&m)
	return
}
