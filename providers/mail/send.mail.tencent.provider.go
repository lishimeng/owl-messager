package mail

import (
	"encoding/json"
	"github.com/lishimeng/owl-messager/internal/db/model"
	"github.com/lishimeng/owl-messager/internal/messager"
	"github.com/lishimeng/owl-messager/providers/mail/tencent"
)

type tencentSender struct {
	proxy *tencent.MailTencentProvider
}

func NewTencent(config string) (s messager.MailProvider, err error) {
	p := tencentSender{}
	var tencentConfig model.TencentConfig
	err = json.Unmarshal([]byte(config), &tencentConfig)
	if err != nil {
		return
	}
	p.proxy, err = tencent.New(tencentConfig)
	s = &p
	return
}

func (s *tencentSender) Send(req messager.MailRequest) (err error) {

	err = s.proxy.Send(req)
	return
}
