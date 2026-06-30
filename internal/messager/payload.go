package messager

import "github.com/lishimeng/owl-messager/pkg/msg"

type Payload struct {
	MessageCategory msg.MessageCategory `json:"messageCategory,omitempty"`
	Sender          string              `json:"sender"`
	Payload         interface{}         `json:"payload"`
}

// Message 消息
type Message struct {
	MessageId       int                 `json:"id"`
	MessageCategory msg.MessageCategory `json:"messageCategory,omitempty"`
}
