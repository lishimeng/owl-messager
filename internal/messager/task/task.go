package task

import (
	"context"
	"fmt"
	"github.com/lishimeng/owl/internal/db/model"
	"github.com/lishimeng/owl/internal/db/repo"
	"github.com/lishimeng/owl/internal/messager/sender"
	"time"
)

type MessageTask interface {
	Run()
}

type messageTask struct {
	running bool
	ctx context.Context
}

func New(ctx context.Context) (t MessageTask, err error) {
	t = &messageTask{ctx: ctx}
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
				fmt.Println("get message failed")
				fmt.Println(err)
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
		t.handleMessage(message)
	}
}

func (t *messageTask) handleMessage(message model.MessageInfo) {
	task, err := repo.AddMessageTask(message.Id)
	if err != nil {
		fmt.Println(err)
		return
	}
	// 提交给发送器
	err = sender.Container.ExecuteTask(task)
	if err != nil {
		fmt.Println(err)
		return
	}
}

