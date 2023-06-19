package apns

import (
	"crypto/tls"
	"github.com/lishimeng/go-log"
)

type Provider struct {
	clientManager *ClientManager
}

func New() (p *Provider) {

	p = &Provider{clientManager: NewClientManager()}
	return
}

// Send 发送消息
// mode 模式：development、production
// certificate 证书文件
func (p *Provider) Send(productionMode bool, certificate tls.Certificate, deviceToken string, topic string, payload map[string]interface{}) (err error) {

	notification := Notification{
		DeviceToken: deviceToken,
		Topic:       topic,
		Payload:     payload,
	}

	client := p.clientManager.Get(certificate)
	if client == nil {
		return // 出错了
	}

	if productionMode {
		client = client.Production()
	} else {
		client = client.Development()
	}
	resp, err := client.Push(&notification)
	if err != nil {
		return
	}
	log.Info(resp.ApnsID) // TODO response
	return
}
