package wechat

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/go-resty/resty/v2"
	"github.com/lishimeng/go-log"
	"time"
)

const (
	tokenHost      = "https://api.weixin.qq.com/cgi-bin/token?grant_type=client_credential&appid=%s&secret=%s"
	uniformHostTpl = "https://api.weixin.qq.com/cgi-bin/message/wxopen/template/uniform_send?access_token=%s"
)

type Provider struct {
	Host      string // wechat公众号provider地址
	AppId     string
	AppSecret string
	token     *ClientCredentialToken

	rest *resty.Client
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
	Timestamp    int64  `json:"timestamp"`
}
type To struct {
	OpenId     string                 // user's open id of wechat open platform
	Payload    map[string]PayloadItem // data
	TemplateId string                 // tpl id on wechat open platform
}

func (p *Provider) Send(t To) (err error) {
	err = p.refreshToken()
	if err != nil {
		return
	}

	// TODO send to receiver
	return
}

func (p *Provider) getAccessToken() (t ClientCredentialToken, err error) {
	host := fmt.Sprintf(tokenHost, p.AppId, p.AppSecret)
	log.Info("get wx access token: %s", host)
	// TODO
	resp, err := p.rest.R().Get(host)
	if err != nil {
		return
	}
	if resp.StatusCode() != 200 {
		err = errors.New(resp.Status())
		return
	}
	err = json.Unmarshal(resp.Body(), &t)
	if err != nil {
		return
	}
	t.Timestamp = time.Now().Unix()
	return
}

func (p *Provider) refreshToken() (err error) {
	var token ClientCredentialToken
	if p.token == nil || p.token.Expired() {
		// get token
		token, err = p.getAccessToken()
		if err != nil {
			return
		}
		p.token = &token
	}
	return
}

func (t ClientCredentialToken) Expired() bool {
	return t.Timestamp <= time.Now().Unix()
}
