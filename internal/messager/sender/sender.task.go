package sender

import (
	"context"
	"github.com/lishimeng/go-log"
	"github.com/lishimeng/owl-messager/internal/db/model"
	"github.com/lishimeng/owl-messager/internal/db/repo"
	"github.com/lishimeng/owl-messager/internal/messager/msg"
)

type TaskExecutor interface {
	Execute(task model.MessageTask) (err error)
}

type taskExecutor struct {
	mailSenders Mail
	smsSenders  Sms
	apnsSender  Apns

	ctx context.Context
}

func New(ctx context.Context) (t TaskExecutor, err error) {
	log.Info("start task engine")
	mail, err := NewMailSender(ctx)
	if err != nil {
		return
	}
	apns, err := NewApnsSender(ctx)
	if err != nil {
		return
	}
	sms, err := NewSmsSender(ctx)
	if err != nil {
		return
	}
	if err != nil {
		return
	}
	t = &taskExecutor{
		mailSenders: mail,
		smsSenders:  sms,
		apnsSender:  apns,
		ctx:         ctx,
	}
	return
}

func (c *taskExecutor) Execute(task model.MessageTask) (err error) {

	log.Info("task engine handle task: %d", task.Id)
	mi, err := repo.GetMessageById(task.MessageId)
	if err != nil {
		_ = log.Error("unknown task:%d", task.Id)
		_ = log.Error(err)
		return
	}

	category := mi.Category
	switch mi.Category {
	case msg.Email:
		log.Debug("mail task")
		var m model.MailMessageInfo
		m, err = repo.GetMailByMessageId(mi.Id)
		if err != nil {
			_ = log.Error("no mail refer to message:%d", mi.Id)
			return
		}
		err = c.mailSenders.Send(m)
	case msg.Sms:
		log.Debug("sms task")
		var m model.SmsMessageInfo
		m, err = repo.GetSmsByMessageId(mi.Id)
		if err != nil {
			_ = log.Error("no sms refer to message:%d", mi.Id)
			return
		}
		err = c.smsSenders.Send(m)
	default:
		log.Info("unknown category:%d[task:%d]\n", category, task.Id)
	}
	return
}
