package sms

import "errors"

// 又拍云 SMS — 尚未实现，不在 providers/init.go 注册。

const (
	apiHost = "https://sms-api.upyun.com/api/messages"
)

type TemplateResp struct {
}

type TplHandler interface {
	Get(id string) (TemplateResp, error)
	Add(params map[string]interface{}) (TemplateResp, error)
	Del(id string) (TemplateResp, error)
	Edit(params map[string]interface{}) (TemplateResp, error)
}

type Sender interface {
	SetAppId(appId string)

	GetTplHandler() TplHandler

	Send()

	Remain() int
}

type sender struct {
	appId      string
	tplHandler TplHandler
}

type templateHandler struct {
}

func (t *templateHandler) Get(id string) (r TemplateResp, err error) {
	return r, errors.New("upyun sms: not implemented")
}

func (t *templateHandler) Add(params map[string]interface{}) (r TemplateResp, err error) {
	return r, errors.New("upyun sms: not implemented")
}

func (t *templateHandler) Del(id string) (r TemplateResp, err error) {
	return r, errors.New("upyun sms: not implemented")
}

func (t *templateHandler) Edit(params map[string]interface{}) (r TemplateResp, err error) {
	return r, errors.New("upyun sms: not implemented")
}

func (s *sender) SetAppId(appId string) {
	s.appId = appId
}

func (s *sender) Send() {
	panic("upyun sms: not implemented")
}
