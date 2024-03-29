package model

import "github.com/lishimeng/app-starter"

// MessageTask 消息投送任务
type MessageTask struct {
	app.Pk
	MessageId         int `orm:"column(message_id)"`          // message
	MessageInstanceId int `orm:"column(message_instance_id)"` // sms id/mail id
	app.TableChangeInfo
}

// MessageRunningTask 当前在运行的task,投送前添加进来,投送完成后删除,超时没完成job清理
type MessageRunningTask struct {
	app.Pk
	TaskId int `orm:"column(task_id)"`
	app.TableInfo
}

const (
	MessageTaskInit        = iota // 新建,初始化
	MessageTaskSending            // 投送中
	MessageTaskSendSuccess        // 投送成功
	MessageTaskSendFailed         // 投送失败
	MessageTaskCancelled   = -1   // 取消
	MessageTaskSendExpired = -9   // 投送失败
)
