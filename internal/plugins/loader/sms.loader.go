package loader

import (
	"github.com/lishimeng/go-log"
	"github.com/lishimeng/owl/internal/db/model"
	"github.com/lishimeng/owl/internal/plugins/container"
	"github.com/lishimeng/owl/internal/provider/sms"
)

type SmsLoader interface {
	Load(id string) sms.Provider
	Unload(id string)
}

func New() (sms SmsLoader) {
	h := smsLoaderImpl{}
	sms = &h
	return
}

type smsLoaderImpl struct {
}

func (s *smsLoaderImpl) Load(id string) (p sms.Provider) {
	var manager, err = container.Get(new(sms.ProviderManager))
	if err != nil {
		log.Info(err)
		return
	}
	p = manager.Get(model.SmsVendor(id))
	return
}

func (s *smsLoaderImpl) Unload(id string) {
	var _, err = container.Get(new(sms.ProviderManager))
	if err != nil {
		log.Info(err)
		return
	}
	// TODO
}
