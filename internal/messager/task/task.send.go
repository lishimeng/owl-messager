package task

import (
	"context"
	"fmt"
	"github.com/lishimeng/go-log"
	"github.com/lishimeng/owl/internal/db/model"
	"github.com/lishimeng/owl/internal/db/repo"
	"github.com/lishimeng/owl/internal/db/service"
	"github.com/lishimeng/owl/internal/messager/msg"
	"github.com/lishimeng/owl/internal/messager/sender"
	"time"
)

type MessageTask interface {
	Run()
}

type messageTask struct {
	running bool
	ctx context.Context
	executor sender.TaskExecutor
}

func New(ctx context.Context, executor sender.TaskExecutor) (t MessageTask, err error) {
	t = &messageTask{ctx: ctx, executor: executor}
	return
}

func (t *messageTask) Run() {
	if t.running {
		return
	}
	go t.loop()
}

func (t *messageTask) loop() {
	for {
		select {
		case <-t.ctx.Done():
			return
		default:
			messages, err := t.getMessages(5)
			if err != nil {
				log.Info("get message failed")
				log.Info(err)
			} else {
				t.handleMessages(messages...)
			}
			if len(messages) == 0 {
				time.Sleep(time.Second*10)
			}
		}
	}
}

func (t messageTask) getMessages(size int) (messages []model.MessageInfo, err error) {
	messages, err = repo.GetMessageToSend(size)
	return
}

func (t *messageTask) handleMessages(messages ...model.MessageInfo) {
	for _, message := range messages {
		e := t.handleMessage(message)
		if e != nil {
			log.Info("handle message failed")
			log.Info(e)
		}
	}
}

func (t *messageTask) handleMessage(message model.MessageInfo) (err error) {
	instanceId, err := getMessageInstanceId(message)
	if err != nil {
		log.Info("get message instance id failed")
		return
	}
	task, err := service.CreateMessageTask(message, instanceId)
	if err != nil {
		log.Info("create message task failed")
		log.Info(err)
		return
	}
	// 提交给发送器
	err = t.executor.Execute(task)
	if err != nil {
		// TODO message status -> fail
		// TODO task status -> fail
		// TODO delete running task
		return
	}
	return
}

func getMessageInstanceId(message model.MessageInfo) (id int, err error) {
	switch message.Category {
	case msg.Email:
		var mail model.MailMessageInfo
		mail, err = repo.GetMailByMessageId(message.Id)
		if err == nil {
			id = mail.Id
		}
	case msg.Sms:
		var sms model.SmsMessageInfo
		sms, err = repo.GetSmsByMessageId(message.Id)
		if err == nil {
			id = sms.Id
		}
	default:
		log.Info("known message category:%d[message id:%d]", message.Category, message.Id)
		err = fmt.Errorf("known message category:%d[message id:%d]", message.Category, message.Id)
	}
	return
}

