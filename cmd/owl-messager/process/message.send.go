package process

import (
	"context"
	"github.com/lishimeng/app-starter/factory"
	"github.com/lishimeng/owl-messager/internal/etc"
	"github.com/lishimeng/owl-messager/internal/messager/sender"
	"github.com/lishimeng/owl-messager/internal/messager/task"
)

//var taskExecutor sender.TaskExecutor

func messageSendProcess(ctx context.Context) (err error) {
	taskExecutor, err := sender.New(ctx)
	if err != nil {
		return
	}

	var messageTask task.MessageTask
	var opts []task.Option

	var strategy = task.Strategy(etc.Config.Sender.Strategy)
	switch strategy {
	case task.MemQueue:
		opts = append(opts, task.WithQueue(etc.Config.Sender.Buff))
	case task.Db:
		opts = append(opts, task.WithDb(10)) // etc
	default:
		opts = append(opts, task.WithQueue(etc.Config.Sender.Buff)) // 默认使用queue
	}

	messageTask, err = task.New(ctx, taskExecutor, opts...)
	if err != nil {
		return
	}

	go messageTask.Run()

	factory.Add(&messageTask)

	return
}
