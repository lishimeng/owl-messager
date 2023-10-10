package smtp

import (
	"encoding/json"
	"github.com/lishimeng/owl-messager/pkg/msg"
)

func New(config string) (p *MailSmtpProvider, err error) {

	p = &MailSmtpProvider{}
	var c msg.SmtpConfig
	err = json.Unmarshal([]byte(config), &c)
	if err != nil {
		return
	}

	p.Config = c
	return
}
