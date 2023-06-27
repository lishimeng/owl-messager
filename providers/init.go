package providers

import (
	"encoding/json"
	"github.com/lishimeng/owl-messager/internal/db/model"
	"github.com/lishimeng/owl-messager/internal/messager"
	"github.com/lishimeng/owl-messager/internal/provider"
	"github.com/lishimeng/owl-messager/providers/mail"
	"github.com/lishimeng/owl-messager/providers/sms"
)

func registerMailProviders() {
	provider.RegisterMailProvider(model.MailVendorSmtp, func(config string) (messager.MailProvider, error) {
		return mail.NewSmtp(config)
	})

	provider.RegisterMailProvider(model.MailVendorMicrosoft, func(config string) (messager.MailProvider, error) {
		return mail.NewMicrosoft(config)
	})

	provider.RegisterMailProvider(model.MailVendorTencent, func(config string) (messager.MailProvider, error) {
		return mail.NewTencent(config)
	})
}

func registerSmsProviderBuilders() {

	provider.RegisterSmsProvider(model.SmsVendorAli, func(config string) (p messager.SmsProvider, err error) {
		var aliSmsConf model.AliSmsConfig
		h := sms.AliProvider{}
		err = json.Unmarshal([]byte(config), &aliSmsConf)
		if err != nil {
			return
		}
		err = h.Init(aliSmsConf)
		if err != nil {
			return
		}
		p = &h
		return
	})

	provider.RegisterSmsProvider(model.SmsVendorTencent, func(config string) (p messager.SmsProvider, err error) {
		var tencentSmsConf model.TencentSmsConfig
		err = json.Unmarshal([]byte(config), &tencentSmsConf)
		if err != nil {
			return
		}
		h, err := sms.NewTencent(tencentSmsConf)
		if err != nil {
			return
		}
		p = h
		return
	})
}

func init() {
	registerMailProviders()
	registerSmsProviderBuilders()
}
