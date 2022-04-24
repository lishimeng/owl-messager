package wechat

import (
	"encoding/json"
	"errors"
	"fmt"
)

// 小程序-统一消息服务

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

type Resp struct {
	ErrorCode int    `json:"errcode,omitempty"`
	ErrorMsg  string `json:"errmsg,omitempty"`
}

func (p *Provider) Send(req Req) (err error) {
	var resp Resp
	err = p.refreshToken()
	if err != nil {
		return
	}
	host := fmt.Sprintf(uniformHostTpl, p.token.AccessToken)
	rsp, err := p.rest.R().SetBody(req).Post(host)
	if err != nil {
		return
	}
	if rsp.StatusCode() != 200 {
		err = errors.New(rsp.Status())
		return
	}
	err = json.Unmarshal(rsp.Body(), &resp)
	return
}
