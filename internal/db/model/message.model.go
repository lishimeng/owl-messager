package model

import "time"

type MessageHeader struct {
	Pk
	TableChangeInfo
	// 消息ID(外键)
	MessageId int
}

// 消息主表
type MessageInfo struct {
	Pk
	TableChangeInfo
	Category int
	Subject string
	NextSendTime time.Time
}

const (
	MessageInit = 1 // 新建,初始化
	MessageSending = 2 // 投送中
	MessageSendSuccess = 3 // 投送成功
	MessageSendFailed = 4 // 投送失败
	MessageCancelled = -1 // 取消
	MessageSendExpired = -9 // 投送失败
)