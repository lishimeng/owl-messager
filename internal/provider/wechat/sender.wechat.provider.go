package wechat

import (
	"fmt"
	"github.com/lishimeng/go-log"
)

const (
	wxHost = "https://api.weixin.qq.com/cgi-bin/token?grant_type=client_credential&appid=%s&secret=%s"
)

type Provider struct {
	Host      string // wechat公众号provider地址
	AppId     string
	AppSecret string
}

type PayloadItem struct {
	Content string `json:"content"`
	Color   string `json:"color,omitempty"`
}

type Sender struct {
	Host      string // wechat open platform sender server
	AppSecret string // app secret
	AppId     string // app id
}

type ClientCredentialToken struct {
	AccessToken  string `json:"access_token,omitempty"`
	ExpireIn     int64  `json:"expire_in,omitempty"`
	ErrorCode    int64  `json:"errcode,omitempty"`
	ErrorMessage string `json:"errmsg,omitempty"`
}
type To struct {
	OpenId     string                 // user's open id of wechat open platform
	Payload    map[string]PayloadItem // data
	TemplateId string                 // tpl id on wechat open platform
}

func (p *Provider) Send(t To) {
	// TODO send to receiver
}

func (p *Provider) GetAccessToken() (t ClientCredentialToken) {
	host := fmt.Sprintf(wxHost, p.AppId, p.AppSecret)
	log.Info("get wx access token: %s", host)
	return
}

func (p *Provider) TokenExpired() (b bool) {
	return
}
