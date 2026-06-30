package model

import (
	"github.com/lishimeng/app-starter"
	"github.com/lishimeng/owl-messager/pkg/msg"
)

// MessageTask 消息投送任务
type MessageTask struct {
	app.Pk
	MessageId         int                 `gorm:"column:message_id"`
	MessageInstanceId int                 `gorm:"column:message_instance_id"`
	Category          msg.MessageCategory `gorm:"column:category"`
	app.TableChangeInfo
}

// MessageRunningTask 当前在运行的 task
type MessageRunningTask struct {
	app.Pk
	TaskId int `gorm:"column:task_id"`
	app.TableChangeInfo
}
