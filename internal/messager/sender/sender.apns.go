package sender

import (
	"context"
	"github.com/lishimeng/go-log"
	"github.com/lishimeng/owl-messager/internal/db/model"
	"github.com/lishimeng/owl-messager/providers/apns"
)

type Apns interface {
	Send(model.ApnsMessageInfo) (err error)
}

type apnsSender struct {
	ctx       context.Context
	provider  *apns.Provider
	maxWorker int
}

func NewApnsSender(ctx context.Context) (m Apns, err error) {
	m = &apnsSender{
		ctx:       ctx,
		provider:  apns.New(),
		maxWorker: 1,
	}
	return
}

func (m *apnsSender) Send(p model.ApnsMessageInfo) (err error) {
	// sender info
	log.Info("send apns:%d", p.Id)
	// 获取sender

	return
}
