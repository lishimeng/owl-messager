package model

// 消息投送任务
type MessageTask struct {
	Pk
	MessageId int
	TableChangeInfo
}

// 当前在运行的task,投送前添加进来,投送完成后删除,超时没完成job清理
type MessageRunningTask struct {
	Pk
	TaskId int
	TableInfo
}

const (
	MessageStatusInit = iota // 新建,初始化
	MessageSending // 投送中
	MessageSendSuccess // 投送成功
	MessageSendFailed // 投送失败
	MessageCancelled = -1 // 取消
)