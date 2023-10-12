package sdk

import (
	"errors"
	"fmt"
	"github.com/lishimeng/owl-messager/pkg"
	"github.com/lishimeng/owl-messager/utils"
)

type Rpc struct {
	host          string
	appId, secret string
	token         string

	logic func(rest *utils.RestClient) (code int, err error)
}

func NewRpc(host string) *Rpc {
	r := &Rpc{host: host}
	return r
}

func (r *Rpc) BuildReq(handler func(*utils.RestClient) (int, error)) *Rpc {
	r.logic = handler
	return r
}

func (r *Rpc) Exec() (err error) {

	var code int
	if r.logic == nil {
		return
	}
	if len(r.token) <= 0 { // 预先检查token长度, 初始化时token是空的
		code = CodeNotAllow
	} else {
		// 执行logic
		code, err = r.logic(utils.NewRest(r.host).Auth(r.token))
	}

	if err != nil {
		return
	}

	// 如果出现credentials无效, 登录(CodeNotAllow)
	if code == CodeNotAllow {
		err = r.credentials()
		if err != nil {
			return
		}
		// 再执行
		code, err = r.logic(utils.NewRest(r.host).Auth(r.token))
	}

	if code != CodeSuccess {
		err = errors.New(fmt.Sprintf("%d", code))
		return
	}
	if err != nil {
		return
	}
	return
}

func (r *Rpc) Auth(appId, secret string) *Rpc {
	r.appId = appId
	r.secret = secret
	return r
}

func (r *Rpc) credentials() (err error) {
	var resp pkg.CredentialResp
	req := pkg.CredentialReq{
		AppId:  r.appId,
		Secret: r.secret,
	}
	code, err := utils.NewRest(r.host).Path(ApiCredential).ResponseJson(&resp).Post(req)
	if err != nil { // 网络异常
		return
	}
	if code != CodeSuccess {
		err = errors.New(fmt.Sprintf("%d", resp.Code))
		return
	}
	if resp.Code != CodeSuccess { // 业务异常
		err = errors.New(fmt.Sprintf("%d[%s]", resp.Code, resp.Message))
		return
	}
	r.token = resp.Token
	return
}
