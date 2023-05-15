package model

import (
	"github.com/lishimeng/app-starter"
	"time"
)

type MessageHeader struct {
	app.Pk
	app.TableChangeInfo
	// 消息ID(外键)
	MessageId int `orm:"column(message_id)"`
}

// 消息主表
type MessageInfo struct {
	app.Pk
	app.TableChangeInfo
	Category     int       `orm:"column(category)"`
	Subject      string    `orm:"column(subject)"`
	Priority     int       `orm:"column(priority);null"`
	NextSendTime time.Time `orm:"column(next_send_time);null"`
}

const (
	MessageInit        = 1  // 新建,初始化
	MessageSending     = 2  // 投送中
	MessageSendSuccess = 3  // 投送成功
	MessageSendFailed  = 4  // 投送失败
	MessageCancelled   = -1 // 取消
	MessageSendExpired = -9 // 投送失败
)

const (
	MessagePriorityLow    = 1 // low
	MessagePriorityNormal = 2 // normal
	MessagePriorityHigh   = 3 // high
)
