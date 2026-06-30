package process

import (
	"context"

	"github.com/lishimeng/owl-messager/internal/etc"
	"github.com/lishimeng/owl-messager/internal/mailattachment"
	"github.com/lishimeng/owl-messager/internal/messager/task"
)

func BeforeStarted(ctx context.Context) (err error) {
	mailattachment.Init(etc.MailAttachmentConfig())
	task.InitMonitor(ctx)
	err = messageSendProcess(ctx)
	return
}
