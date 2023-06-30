package open

import (
	"fmt"
	"github.com/kataras/iris/v12"
	"github.com/lishimeng/app-starter"
	"github.com/lishimeng/app-starter/factory"
	"github.com/lishimeng/app-starter/token"
	"github.com/lishimeng/app-starter/tool"
	"github.com/lishimeng/go-log"
	"github.com/lishimeng/owl-messager/internal/common"
	"github.com/lishimeng/owl-messager/internal/etc"
	"github.com/pkg/errors"
)

type CredentialReq struct {
	AppId  string `json:"appId,omitempty"`
	Secret string `json:"secret,omitempty"`
}

type CredentialResp struct {
	app.Response
	Token string `json:"token,omitempty"`
}

type AppInfo struct {
	AppId  string
	Secret string
	Org    string
}

func genCredential(ctx iris.Context) {
	var err error
	var req CredentialReq
	var resp CredentialResp

	err = ctx.ReadJSON(&req)
	if err != nil {
		log.Debug(errors.Wrap(err, "读取参数错误"))
		resp.Code = tool.RespCodeError
		tool.ResponseJSON(ctx, resp)
		return
	}

	if len(req.AppId) == 0 {
		log.Debug("appId: %s, secret:%s", req.AppId, req.Secret)
		resp.Code = common.CodeAppNotFound
		resp.Message = common.MsgAppNotFound
		tool.ResponseJSON(ctx, resp)
		return
	}
	c, err := getAppInfo(req.AppId)
	if err != nil {
		log.Debug(errors.Wrap(err, fmt.Sprintf("appId not found:%s", req.AppId)))
		resp.Code = common.CodeAppNotFound
		resp.Message = common.MsgAppNotFound
		tool.ResponseJSON(ctx, resp)
		return
	}
	if c.Secret != req.Secret {
		log.Debug("appId: %s, secret not match", req.AppId)
		resp.Code = common.CodeSecretNotValid
		resp.Message = common.MsgSecretNotValid
		tool.ResponseJSON(ctx, resp)
		return
	}

	var provider token.JwtProvider
	err = factory.Get(&provider)
	if err != nil {
		log.Debug(errors.Wrap(err, "no jwt provider"))
		resp.Code = tool.RespCodeError
		tool.ResponseJSON(ctx, resp)
		return
	}

	p := token.JwtPayload{
		Org: c.Org,
		Uid: c.AppId,
	}
	bs, err := provider.GenWithTTL(p, etc.TokenTTL)
	if err != nil {
		log.Debug(errors.Wrap(err, "gen credential err"))
		resp.Code = tool.RespCodeError
		tool.ResponseJSON(ctx, resp)
		return
	}

	resp.Token = string(bs)
	resp.Code = tool.RespCodeSuccess
	tool.ResponseJSON(ctx, resp)
	return
}
