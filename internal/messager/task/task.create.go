package task

import (
	"context"
	"github.com/lishimeng/go-log"
	"github.com/lishimeng/owl-messager/internal/db/model"
	"github.com/lishimeng/owl-messager/internal/db/repo"
	"github.com/lishimeng/owl-messager/internal/db/service"
	"github.com/lishimeng/owl-messager/internal/messager/sender"
	"math"
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
			messages, err := t.getMessages(1)
			if err != nil {
				log.Info("get message failed")
				log.Info(err)
			} else {
				t.handleMessages(messages...)
			}
			if len(messages) == 0 {
				time.Sleep(time.Second * 10)
			} else {
				time.Sleep(time.Second * time.Duration(math.Round(10)))
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
	GetMonitor().Pub(MonitorData{
		TaskId:            task.Id,
		MessageId:         message.Id,
		MessageCategory:   message.Category,
		Subject:           message.Subject,
		MessageInstanceId: task.MessageInstanceId,
		Status:            task.Status,
	})
	// 提交给发送器
	err = t.executor.Execute(task)
	AfterHandleTask(task, message, err)
	return
}

func AfterHandleTask(task model.MessageTask, message model.MessageInfo, err error) {
	pubData := MonitorData{
		TaskId:            task.Id,
		MessageId:         message.Id,
		MessageCategory:   message.Category,
		Subject:           message.Subject,
		MessageInstanceId: task.MessageInstanceId,
	}
	if err != nil {
		err = service.OnTaskHandleFail(task)
		pubData.Status = model.MessageTaskSendFailed
	} else {
		err = service.OnTaskHandleSuccess(task)
		pubData.Status = model.MessageTaskSendSuccess
	}

	GetMonitor().Pub(pubData)
}
