package setup

import (
	"context"
)

func AfterStarted(ctx context.Context) (err error) {
	err = loadSmsProviders(ctx)
	if err != nil {
		return
	}

	err = messageSendProcess(ctx)
	if err != nil {
		return
	}
	return
}
