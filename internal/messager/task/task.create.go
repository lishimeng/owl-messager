package task

import (
	"context"
	"github.com/lishimeng/go-log"
	"github.com/lishimeng/owl/internal/db/model"
	"github.com/lishimeng/owl/internal/db/repo"
	"github.com/lishimeng/owl/internal/db/service"
	"github.com/lishimeng/owl/internal/messager/sender"
	"time"
)

type MessageTask interface {
	Run()
	HandleMessage(message model.MessageInfo) (err error)
}

type messageTask struct {
	running  bool
	ctx      context.Context
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
				time.Sleep(time.Second * 10)
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
		e := t.HandleMessage(message)
		if e != nil {
			log.Info("handle message failed")
			log.Info(e)
		}
	}
}

func (t *messageTask) HandleMessage(message model.MessageInfo) (err error) {
	instanceId, err := service.GetMessageInstanceId(message)
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
	AfterHandleTask(task, err)
	return
}

func AfterHandleTask(task model.MessageTask, err error) {
	if err != nil {
		err = service.OnTaskHandleFail(task)
	} else {
		err = service.OnTaskHandleSuccess(task)
	}
}
