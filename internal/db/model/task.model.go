package model

import "github.com/lishimeng/app-starter"

// MessageTask 消息投送任务
type MessageTask struct {
	app.Pk
	MessageId         int `gorm:"column:message_id"`
	MessageInstanceId int `gorm:"column:message_instance_id"`
	app.TableChangeInfo
}

// MessageRunningTask 当前在运行的task
type MessageRunningTask struct {
	app.Pk
	TaskId int `gorm:"column:task_id"`
	app.TableInfo
}

const (
	MessageTaskInit        = iota
	MessageTaskSending
	MessageTaskSendSuccess
	MessageTaskSendFailed
	MessageTaskCancelled   = -1
	MessageTaskSendExpired = -9
)
