package mail

import (
	"crypto/tls"
	"errors"
	"github.com/go-gomail/gomail"
	"github.com/lishimeng/go-log"
)

type MetaInfo struct {
	Server   MetaServer
	Sender   MetaSender
	Receiver MetaReceiver
}

type MetaServer struct {
	// 邮箱服务器地址，如腾讯企业邮箱为smtp.exmail.qq.com
	Host string
	// 邮箱服务器端口，如腾讯企业邮箱为465
	Port int
}

type MetaSender struct {
	// 发件人邮箱地址
	Email string
	// 发件人名称
	Name string
	// 发件人邮箱密码（注意，这里是明文形式），TODO：如果设置成密文？
	Passwd     string
	EmailAlias string
}

type MetaReceiver struct {
	// 接收者邮件 不能为空
	To []string
	// 抄送者邮件 可以为空
	Cc []string
}

type Sender interface {
	Send(metas MetaInfo, subject string, body string) error
}

type gomailSender struct {
}

func New() (s Sender) {

	s = &gomailSender{}
	return
}

func (s *gomailSender) Send(metas MetaInfo, subject string, body string) (err error) {

	log.Debug("mail body:%s", body)
	err = s.chkParam(metas)
	if err != nil {
		return err
	}

	m := gomail.NewMessage()
	// 收件人
	m.SetHeader("To", metas.Receiver.To...)

	//抄送列表
	if len(metas.Receiver.Cc) > 0 {
		m.SetHeader("Cc", metas.Receiver.Cc...)
	}

	// 发件人
	// 发件人子邮箱
	emailAlias := metas.Sender.Email
	if len(metas.Sender.EmailAlias) > 0 {
		emailAlias = metas.Sender.EmailAlias
	}
	// 第三个参数为发件人别名，如"李大锤"，可以为空（此时则为邮箱名称）
	var senderName = ""
	if len(metas.Sender.Name) > 0 {
		senderName = metas.Sender.Name
	}
	m.SetAddressHeader("From", emailAlias, senderName)

	// -----------------------------------
	// 主题
	m.SetHeader("Subject", subject)
	// 正文
	m.SetBody("text/html", body)

	log.Info("[gomail]info before-Send: message: %+v, metas: %+v", m, metas)

	d := gomail.NewDialer(metas.Server.Host, metas.Server.Port, metas.Sender.Email, metas.Sender.Passwd)
	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}
	// 发送
	err = d.DialAndSend(m)
	return
}

func (s gomailSender) chkParam(metas MetaInfo) (err error) {

	if len(metas.Server.Host) == 0 {
		err = errors.New("Server.Host nil")
		return err
	}

	if metas.Server.Port == 0 {
		err = errors.New("Server.Port nil")
		return err
	}

	if len(metas.Sender.Email) == 0 {
		err = errors.New("Sender.Email nil")
		return err
	}

	if len(metas.Sender.Passwd) == 0 {
		err = errors.New("Sender.Passwd nil")
		return err
	}
	if len(metas.Receiver.To) == 0 {
		err = errors.New("Receiver.To nil")
		return err
	}

	return
}
