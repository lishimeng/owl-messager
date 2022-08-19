package mail

import (
	"errors"
	"github.com/lishimeng/owl/internal/db/model"
	"github.com/lishimeng/owl/internal/provider/mail/email"
)

type Factory struct {
}

func (f *Factory) Create(vendor model.MailVendor, config string) (s email.Sender, err error) {

	switch vendor {
	case model.MailVendorSmtp:
		s, err = NewSmtp(config)
	case model.MailVendorMicrosoft:
		s, err = NewMicrosoft(config)
	default:
		err = errors.New("unknown mail vendor")
	}
	return
}
