package sender

import (
	"context"
	"encoding/json"
	"github.com/lishimeng/go-log"
	"github.com/lishimeng/owl-messager/internal/certificate"
	"github.com/lishimeng/owl-messager/internal/db/model"
	"github.com/lishimeng/owl-messager/internal/db/repo"
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
	si, err := repo.GetApnsSenderById(p.Sender)
	if err != nil {
		log.Info("apns sender not exist:%d", p.Sender)
		return
	}

	cert, err := certificate.FromPemBytes([]byte(si.Cert), si.CertPassword)
	if err != nil {
		log.Info("parse certificate failed", err)
		return
	}

	var params map[string]interface{}
	err = json.Unmarshal([]byte(p.Params), &params)
	if err != nil {
		log.Info("params is not json format:%s", p.Params)
		return
	}

	productionMode := p.ApnsMode == model.ProductionMode
	err = m.provider.Send(productionMode, cert, p.Receivers, p.BundleId, params)
	return
}
