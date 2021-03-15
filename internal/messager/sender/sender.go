package sender

import (
	"fmt"
	"github.com/lishimeng/owl/internal/db/model"
	"github.com/lishimeng/owl/internal/db/repo"
	"github.com/lishimeng/owl/internal/messager/msg"
)

type FactoryContainer struct {
	mailSenders map[string]Mail
	smsSenders map[string]Mail
}

var Container FactoryContainer

func (c *FactoryContainer) getMail() (m Mail, err error) {
	return
}

func (c *FactoryContainer) getSms() (m Mail, err error) {
	return
}

func (c *FactoryContainer) ExecuteTask(task model.MessageTask) (err error) {

	mi, err := repo.GetMessageById(task.MessageId)
	if err != nil {
		fmt.Println(err)
		return
	}

	category := mi.Category
	switch mi.Category {
	case msg.Email:
		fmt.Println("mail task")

	case msg.Sms:
		fmt.Println("sms  task")
	default:
		fmt.Printf("unknown category:%d\n", category)
	}
	return
}
