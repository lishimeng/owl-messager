package repo

import "github.com/lishimeng/owl/internal/db/model"

func GetMessageById(id int) (m model.MessageInfo,err error) {
	return
}

// 查询需要发送的消息
func GetMessageToSend(size int) (messages []model.MessageInfo, err error) {
	messages, err = GetMessageByStatus(30) // TODO 假的
	return
}

func GetMessageByStatus(status int) (messages []model.MessageInfo, err error) {
	return
}