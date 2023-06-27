package provider

import (
	"errors"
	"github.com/lishimeng/app-starter/factory"
	"github.com/lishimeng/go-log"
	"github.com/lishimeng/owl-messager/internal/db/model"
	"github.com/lishimeng/owl-messager/internal/messager"
)

type SmsFactory struct {
}

func init() {
	factory.Add(&SmsFactory{})
}

func GetFactory() (f *SmsFactory) {
	_ = factory.Get(f)
	return
}

func (f *SmsFactory) Create(vendor model.SmsVendor, config string) (p messager.SmsProvider, err error) {

	log.Info("create sms provider")
	b, ok := smsProviderBuilders[vendor]
	if !ok {
		err = errors.New("unknown mail vendor")
		return
	}
	p, err = b(config)

	return
}

var smsProviderBuilders = map[model.SmsVendor]func(config string) (messager.SmsProvider, error){}

func RegisterSmsProvider(vendor model.SmsVendor, h func(config string) (messager.SmsProvider, error)) {
	if h == nil {
		return
	}
	smsProviderBuilders[vendor] = h
}
