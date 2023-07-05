package sms

// 又拍云SMS

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
	// TODO
	return
}

func (t *templateHandler) Add(params map[string]interface{}) (r TemplateResp, err error) {
	// TODO
	return
}

func (t *templateHandler) Del(id string) (r TemplateResp, err error) {
	// TODO
	return
}

func (t *templateHandler) Edit(params map[string]interface{}) (r TemplateResp, err error) {
	// TODO
	return
}

func (s *sender) SetAppId(appId string) {
	s.appId = appId
}

func (s *sender) Send() {
	// TODO SingleSend
}
