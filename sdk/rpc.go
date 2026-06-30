package sdk

import (
	"errors"
	"fmt"
	"github.com/lishimeng/owl-messager/utils"
)

type Rpc struct {
	host          string
	appId, secret string

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
	if r.logic == nil {
		return
	}
	if len(r.appId) == 0 || len(r.secret) == 0 {
		return errors.New("appId and secret required")
	}

	code, err := r.logic(utils.NewRest(r.host).BasicAuth(r.appId, r.secret))
	if err != nil {
		return
	}
	if code != CodeSuccess {
		err = errors.New(fmt.Sprintf("%d", code))
	}
	return
}

func (r *Rpc) Auth(appId, secret string) *Rpc {
	r.appId = appId
	r.secret = secret
	return r
}
