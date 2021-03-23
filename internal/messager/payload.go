package messager

type Payload struct {
	MessageCategory int         `json:"messageCategory,omitempty"` // 消息类型
	Sender          string      `json:"sender"`                    // 发送者ID,必须在系统中注册过
	Payload         interface{} `json:"payload"`
}

// 消息
type Message struct {
	MessageId       int `json:"id"`
	MessageCategory int `json:"messageCategory,omitempty"` // 消息类型
}
