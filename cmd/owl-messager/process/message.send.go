package process

import (
	"context"
	"github.com/lishimeng/owl-messager/internal/messager/sender"
	"github.com/lishimeng/owl-messager/internal/messager/task"
)

var taskExecutor sender.TaskExecutor

func messageSendProcess(ctx context.Context) (err error) {
	taskExecutor, err = sender.New(ctx)
	if err != nil {
		return
	}

	messageTask, err := task.New(ctx, taskExecutor)
	if err != nil {
		return
	}

	go messageTask.Run()

	return
}
