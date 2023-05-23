package provider

import (
	"errors"
	"github.com/lishimeng/owl/internal/db/model"
	"github.com/lishimeng/owl/internal/messager"
	"github.com/lishimeng/owl/internal/provider/mail"
)

type MailFactory struct {
}

var DefaultMailFactory *MailFactory

func init() {
	DefaultMailFactory = &MailFactory{}
}

func (f *MailFactory) Create(vendor model.MailVendor, config string) (s messager.MailProvider, err error) {

	switch vendor {
	case model.MailVendorSmtp:
		s, err = mail.NewSmtp(config)
	case model.MailVendorMicrosoft:
		s, err = mail.NewMicrosoft(config)
	case model.MailVendorTencent:
		s, err = mail.NewTencent(config)
	default:
		err = errors.New("unknown mail vendor")
	}
	return
}

var providerBuilders = map[model.MailVendor]func(config string) (messager.MailProvider, error){}

func RegisterMailProvider(vendor model.MailVendor, h func(config string) (messager.MailProvider, error)) {
	if h == nil {
		return
	}
	providerBuilders[vendor] = h
}
