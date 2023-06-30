package process

import (
	"context"
	"github.com/lishimeng/owl-messager/internal/messager/task"
)

func BeforeStarted(ctx context.Context) (err error) {

	task.InitMonitor(ctx)
	return
}
