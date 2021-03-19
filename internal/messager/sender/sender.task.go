package sender

import (
	"context"
	"fmt"
	"github.com/lishimeng/owl/internal/db/model"
	"github.com/lishimeng/owl/internal/db/repo"
	"github.com/lishimeng/owl/internal/messager/msg"
)

type TaskExecutor interface {
	Execute(task model.MessageTask) (err error)
}

type taskExecutor struct {
	mailSenders Mail
	smsSenders Sms

	ctx context.Context
}

func New(ctx context.Context) (t TaskExecutor, err error) {
	mail, err := NewMailSender(ctx)
	if err != nil {
		return
	}
	t = &taskExecutor{
		mailSenders: mail,
		smsSenders:  nil,
		ctx:         ctx,
	}
	return
}

func (c *taskExecutor) Execute(task model.MessageTask) (err error) {

	mi, err := repo.GetMessageById(task.MessageId)
	if err != nil {
		fmt.Println(err)
		return
	}

	category := mi.Category
	switch mi.Category {
	case msg.Email:
		fmt.Println("mail task")
		var m model.MailMessageInfo
		m, err = repo.GetMailByMessageId(mi.Id)
		if err != nil {
			return
		}
		err = c.mailSenders.Send(m)
	case msg.Sms:
		fmt.Println("sms task")
		var m model.SmsMessageInfo
		m, err = repo.GetSmsByMessageId(mi.Id)
		if err != nil {
			return
		}
		err = c.smsSenders.Send(m)
	default:
		fmt.Printf("unknown category:%d\n", category)
	}
	return
}
