package app

import (
	"context"
	"github.com/lishimeng/owl/internal/messager/sender"
	"github.com/lishimeng/owl/internal/messager/task"
)

var taskExecutor sender.TaskExecutor

func SetupSender(ctx context.Context) (err error) {
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

func SetupClearExpireTask(ctx context.Context) (err error) {
	go task.RunClearExpired(ctx)
	return
}
