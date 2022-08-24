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
	default:
		err = errors.New("unknown mail vendor")
	}
	return
}
