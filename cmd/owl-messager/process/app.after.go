package process

import (
	"context"
)

func AfterStarted(ctx context.Context) (err error) {
	err = messageSendProcess(ctx)
	if err != nil {
		return
	}
	return
}
