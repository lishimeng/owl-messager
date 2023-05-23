package providers

import (
	"github.com/lishimeng/owl/internal/db/model"
	"github.com/lishimeng/owl/internal/messager"
	"github.com/lishimeng/owl/internal/provider"
	"github.com/lishimeng/owl/providers/mail"
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

func init() {
	registerMailProviders()
}
