package provider

import (
	"errors"
	"github.com/lishimeng/owl-messager/internal/db/model"
	"github.com/lishimeng/owl-messager/internal/messager"
)

type MailFactory struct {
}

var DefaultMailFactory *MailFactory

func init() {
	DefaultMailFactory = &MailFactory{}
}

func (f *MailFactory) Create(vendor model.MailVendor, config string) (s messager.MailProvider, err error) {

	b, ok := providerBuilders[vendor]
	if !ok {
		err = errors.New("unknown mail vendor")
		return
	}
	s, err = b(config)
	return
}

var providerBuilders = map[model.MailVendor]func(config string) (messager.MailProvider, error){}

func RegisterMailProvider(vendor model.MailVendor, h func(config string) (messager.MailProvider, error)) {
	if h == nil {
		return
	}
	providerBuilders[vendor] = h
}
