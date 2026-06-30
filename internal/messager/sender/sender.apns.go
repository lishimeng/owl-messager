package sender

import (
	"context"
	"crypto/tls"
	"encoding/json"
	"errors"

	"github.com/lishimeng/go-log"
	"github.com/lishimeng/owl-messager/internal/certificate"
	"github.com/lishimeng/owl-messager/internal/db/model"
	"github.com/lishimeng/owl-messager/internal/db/repo"
	"github.com/lishimeng/owl-messager/pkg/msg"
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
	log.Info("send apns:%d", p.Id)

	si, err := repo.GetMessageSenderById(p.SenderId)
	if err != nil {
		log.Info("apns sender not found:%d", p.SenderId)
		return
	}

	tlsCert, err := apnsTLSFromConfig(string(si.Config))
	if err != nil {
		return err
	}

	var payload map[string]interface{}
	if len(p.Params) > 0 {
		if err = json.Unmarshal([]byte(p.Params), &payload); err != nil {
			return err
		}
	}
	if payload == nil {
		payload = map[string]interface{}{}
	}

	msgInfo, err := repo.GetMessageById(p.MessageId)
	if err != nil {
		return err
	}
	if msgInfo.Subject != "" {
		if _, ok := payload["aps"]; !ok {
			pl := apns.NewPayload().Alert(msgInfo.Subject)
			b, _ := pl.MarshalJSON()
			_ = json.Unmarshal(b, &payload)
		}
	}

	deviceToken := msg.SplitReceivers(p.Receivers)
	if len(deviceToken) == 0 {
		return errors.New("apns: no device token")
	}

	topic := p.BundleId
	if topic == "" {
		return errors.New("apns: bundle id required")
	}

	production := p.ApnsMode == model.ProductionMode
	return m.provider.Send(production, tlsCert, deviceToken[0], topic, payload)
}

func apnsTLSFromConfig(config string) (tls.Certificate, error) {
	var cfg msg.ApnsConfig
	if err := json.Unmarshal([]byte(config), &cfg); err != nil {
		return tls.Certificate{}, err
	}
	if cfg.Certificate == "" || cfg.CertificateKey == "" {
		return tls.Certificate{}, errors.New("apns: certificate config incomplete")
	}

	crtHandler := certificate.PemHandler{Pem: cfg.Certificate}
	crt, err := crtHandler.ParseCrt()
	if err != nil {
		return tls.Certificate{}, err
	}
	keyHandler := certificate.PemHandler{Pem: cfg.CertificateKey}
	key, err := keyHandler.ParseKey()
	if err != nil {
		return tls.Certificate{}, err
	}
	return tls.Certificate{
		Certificate: [][]byte{crt.Raw},
		PrivateKey:  key,
		Leaf:        crt,
	}, nil
}
