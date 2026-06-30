package model

import (
	"time"

	"github.com/lishimeng/app-starter"
	"github.com/lishimeng/owl-messager/pkg/msg"
)

type MessageHeader struct {
	app.TenantPk

	MessageId int `gorm:"column:message_id"`
	app.TableChangeInfo
}

// MessageInfo 消息主表
type MessageInfo struct {
	app.TenantPk
	app.TableChangeInfo
	Category     msg.MessageCategory `gorm:"column:category"`
	Subject      string              `gorm:"column:subject"`
	Priority     int                 `gorm:"column:priority"`
	NextSendTime time.Time           `gorm:"column:next_send_time"`
}

const (
	MessageInit        = 1
	MessageSending     = 2
	MessageSendSuccess = 3
	MessageSendFailed  = 4
	MessageCancelled   = -1
	MessageSendExpired = -9
)

const (
	MessagePriorityLow    = 1
	MessagePriorityNormal = 2
	MessagePriorityHigh   = 3
)
