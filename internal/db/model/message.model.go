package model

import (
	"time"

	"github.com/lishimeng/app-starter"
	"github.com/lishimeng/owl-messager/pkg/msg"
)

// MessageHeader 渠道子表公共头；状态由 MessageInfo 维护，子表仅保留创建时间。
type MessageHeader struct {
	TenantScope
	MessageId int `gorm:"column:message_id;uniqueIndex"`
	app.TableInfo
}

// MessageInfo 消息主表
type MessageInfo struct {
	TenantScope
	app.TableChangeInfo
	Category     msg.MessageCategory `gorm:"column:category"`
	Subject      string              `gorm:"column:subject"`
	Priority     int                 `gorm:"column:priority"`
	NextSendTime time.Time           `gorm:"column:next_send_time"`
}

const (
	MessagePriorityLow    = 1
	MessagePriorityNormal = 2
	MessagePriorityHigh   = 3
)
