package model

// 实体（模板、发送方等）启用状态
const (
	EntityDisabled = 0
	EntityEnabled  = 1
)

const (
	TemplateDisable = EntityDisabled
	TemplateEnable  = EntityEnabled
)

const (
	SenderDisable = EntityDisabled
	SenderEnable  = EntityEnabled
)

const (
	DefaultSenderDisable = EntityDisabled
	DefaultSenderEnable  = EntityEnabled
)

// 消息主表状态（仅 MessageInfo 使用）
const (
	MessageInit        = 1
	MessageSending     = 2
	MessageSendSuccess = 3
	MessageSendFailed  = 4
	MessageCancelled   = -1
	MessageSendExpired = -9
)

// 消息任务状态
const (
	MessageTaskInit        = iota
	MessageTaskSending
	MessageTaskSendSuccess
	MessageTaskSendFailed
	MessageTaskCancelled   = -1
	MessageTaskSendExpired = -9
)
