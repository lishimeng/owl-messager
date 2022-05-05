package setup

import (
	"context"
	"github.com/lishimeng/go-log"
	"github.com/lishimeng/owl/internal/db/model"
	"github.com/lishimeng/owl/internal/db/repo"
	"github.com/lishimeng/owl/internal/plugins/loader"
	"github.com/lishimeng/owl/internal/provider/sms"
)

func loadSmsProviders(_ context.Context) (err error) {
	log.Info("load sms providers")

	senders, err := repo.GetSmsSenders(model.SmsSenderEnable)
	if err != nil {
		return
	}
	for _, s := range senders {
		var v = s.Vendor
		var p sms.Provider
		switch v {
		case model.SmsVendorAli:
			log.Info("load sms provider: %d[%s]", s.Id, s.Code)
			ali := sms.AliProvider{}
			err = s.UnmarshalConfig()
			if err != nil {
				return
			}
			// TODO key of config map
			appKey := s.ConfigMap["appKey"]
			appSecret := s.ConfigMap["appSecret"]
			region := s.ConfigMap["region"]
			signName := s.ConfigMap["signName"]
			err = ali.Init(appKey, appSecret, region, signName)
			if err != nil {
				return
			}
			p = &ali
		default:
			continue
		}
		loader.Load(v, p)
	}
	return
}
