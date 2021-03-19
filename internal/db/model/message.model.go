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