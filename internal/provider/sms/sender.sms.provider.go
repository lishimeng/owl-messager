package sms

import "github.com/lishimeng/owl/internal/db/model"

// Request SMS发送请求
type Request struct {
	Template  string
	Sign      string
	Params    map[string]interface{}
	Receivers []string
}

// Response 服务器回复
type Response struct {
	RequestId string      // 本次请求的唯一标识，由服务器分配。用来追溯历史
	Payload   interface{} // 个性化的服务器返回信息
}

// Provider 发短信工具
type Provider interface {
	Send(req Request) (resp Response, err error)
}

type ProviderManager struct {
	factory   interface{} // 创建器
	Providers map[model.SmsVendor]Provider
}

func (pm *ProviderManager) Get(vendor model.SmsVendor) (p Provider) {
	p, ok := pm.Providers[vendor]
	if !ok {
		// TODO create
	}
	return
}

func (pm *ProviderManager) Add(vendor model.SmsVendor, p Provider) {
	pm.Providers[vendor] = p
	return
}

func New() *ProviderManager {
	pm := &ProviderManager{}
	return pm
}
