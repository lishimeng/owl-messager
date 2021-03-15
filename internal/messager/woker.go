package messager

import (
	"fmt"
	"github.com/lishimeng/owl/internal/messager/msg"
)

func sendMessage(payload Payload) {
	switch payload.MessageCategory {
	case msg.Email:
		sendEmailMessage(payload.Payload)
	case msg.Sms:
		sendSmsMessage(payload.Payload)
	default:
		fmt.Println("unknown message category")
	}
}

// 发邮件
func sendEmailMessage(payload interface{}) {

}

// 发短信
func sendSmsMessage(payload interface{}) {

}