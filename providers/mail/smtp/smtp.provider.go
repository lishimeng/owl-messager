package smtp

import (
	"crypto/tls"
	"github.com/go-gomail/gomail"
	"github.com/lishimeng/go-log"
	"github.com/lishimeng/owl-messager/internal/db/model"
)

type MailSmtpProvider struct {
	Config model.SmtpConfig
}

func (s MailSmtpProvider) Send(subject string, body string, receivers ...string) (err error) {
	log.Debug("mail body:%s", body)

	m := gomail.NewMessage()
	// 收件人
	m.SetHeader("To", receivers...)

	// 第三个参数为发件人别名，如"李大锤"，可以为空(此时则为邮箱名称)
	m.SetAddressHeader("From", s.Config.SenderEmail, s.Config.SenderAlias)

	// -----------------------------------
	// 主题
	m.SetHeader("Subject", subject)
	// 正文
	m.SetBody("text/html", body)

	d := gomail.NewDialer(s.Config.Host, s.Config.Port, s.Config.AuthUser, s.Config.AuthPass)
	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}
	// 发送
	err = d.DialAndSend(m)
	return
}
