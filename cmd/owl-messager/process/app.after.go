package process

import (
	"context"
	"time"

	"github.com/lishimeng/owl-messager/internal/mailattachment"
)

func AfterStarted(ctx context.Context) (err error) {
	go runAttachmentCleanup(ctx)
	err = messageSendProcess(ctx)
	return
}

func runAttachmentCleanup(ctx context.Context) {
	ticker := time.NewTicker(time.Hour)
	defer ticker.Stop()
	for {
		select {
		case <-ctx.Done():
			return
		case <-ticker.C:
			mailattachment.Default().CleanupExpired()
		}
	}
}
