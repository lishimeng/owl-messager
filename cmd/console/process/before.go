package process

import (
	"context"

	"github.com/lishimeng/owl-messager/internal/consoleauth"
)

func BeforeStarted(ctx context.Context) error {
	consoleauth.Load()
	return nil
}
