package setup

import (
	"context"
	"github.com/lishimeng/owl/internal/messager/task"
)

func BeforeStarted(ctx context.Context) (err error) {

	task.InitMonitor(ctx)
	return
}
