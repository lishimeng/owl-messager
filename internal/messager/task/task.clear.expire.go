package task

import (
	"context"
	"github.com/lishimeng/go-log"
	"github.com/lishimeng/owl-messager/internal/db/repo"
	"github.com/lishimeng/owl-messager/internal/db/service"
	"time"
)

/////////////////
// clear expired task
// set message status: expired
////////////////

func RunClearExpired(ctx context.Context) {
	var duration = time.Minute * 10
	var timer = time.NewTimer(duration)
	defer func() {
		timer.Stop()
	}()
	for {
		select {
		case <-ctx.Done():
			return
		case <-timer.C:
			clearExpiredTaskOnce(ctx)
			timer.Reset(duration)
		}
	}
}

func clearExpiredTaskOnce(ctx context.Context) {
	// TODO job
	for {
		select {
		case <-ctx.Done():
			return
		default:
			size, err := clearExpiredTask()
			if err != nil {
				log.Info("get expired tasks failed")
				return
			}
			if size <= 0 {
				return
			}
		}
	}
}

func clearExpiredTask() (size int, err error) {
	var now = time.Now().Add(-time.Minute * 12)
	tasks, err := repo.GetExpiredTasks(10, now)
	if err != nil {
		log.Info(err)
		return
	}
	size = len(tasks)
	if size == 0 {
		return
	}
	for _, item := range tasks {
		err = service.HandleExpiredTask(item)
		if err != nil {
			return
		}
	}
	return
}
