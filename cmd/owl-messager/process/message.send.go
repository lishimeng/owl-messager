package process

import (
	"context"

	"github.com/lishimeng/go-log"
	"github.com/lishimeng/owl-messager/internal/db/repo"
	"github.com/lishimeng/owl-messager/internal/messager/sender"
	"github.com/lishimeng/owl-messager/internal/messager/task"
	"github.com/lishimeng/x/container"
)

func messageSendProcess(ctx context.Context) (err error) {
	taskExecutor, err := sender.New(ctx)
	if err != nil {
		return
	}

	settings := repo.LoadTaskChannelSettings()
	opts := task.Configure(settings.Channel, settings.ScanInterval)
	if task.UseMemQueue() {
		log.Info("message task channel: memqueue")
	} else {
		log.Info("message task channel: db (scanInterval=%ds)", settings.ScanInterval)
	}

	messageTask, err := task.New(ctx, taskExecutor, opts...)
	if err != nil {
		return
	}

	go messageTask.Run()

	container.Add(&messageTask)

	return
}
