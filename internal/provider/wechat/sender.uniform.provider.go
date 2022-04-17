package wechat

// 小程序-模板消息

const (
	uniformHostTpl = "https://api.weixin.qq.com/cgi-bin/message/wxopen/template/uniform_send?access_token=%s"
)

type MaProvider struct {
	token string
}

type Req struct {
	ToUser string   `json:"touser"`
	MeMsg  MeMsgReq `json:"weapp_template_msg"`
	MpMsg  MpMsgReq `json:"mp_template_msg"`
}

type MeMsgReq struct {
	Tpl     string `json:"template_id"`
	Page    string `json:"page"`
	FormId  string `json:"form_id"`
	Data    string `json:"data"`
	KeyWord string `json:"emphasis_keyword"`
}

type MpMsgReq struct {
	AppId       string `json:"appid"`
	TemplateId  string `json:"template_id"`
	Url         string `json:"url"`
	Data        string `json:"data"`
	MiniProgram string `json:"miniprogram"`
}

func (m *MaProvider) ChangeToken() {

}

func (m *MaProvider) Send() {

}
