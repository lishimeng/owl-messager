package provider

import (
	"errors"
	"github.com/lishimeng/owl-messager/internal/messager"
	"github.com/lishimeng/owl-messager/pkg/msg"
)

type MailFactory struct {
}

var DefaultMailFactory *MailFactory

func init() {
	DefaultMailFactory = &MailFactory{}
}

func (f *MailFactory) Create(vendor msg.MessageProvider, config string) (s messager.MailProvider, err error) {

	b, ok := providerBuilders[vendor]
	if !ok {
		err = errors.New("unknown mail vendor")
		return
	}
	s, err = b(config)
	return
}

var providerBuilders = map[msg.MessageProvider]func(config string) (messager.MailProvider, error){}

func RegisterMailProvider(vendor msg.MessageProvider, h func(config string) (messager.MailProvider, error)) {
	if h == nil {
		return
	}
	providerBuilders[vendor] = h
}
