package clientApi

import (
	"github.com/lishimeng/app-starter"
	"github.com/lishimeng/app-starter/server"
	"github.com/lishimeng/app-starter/tool"
	"github.com/lishimeng/owl-messager/internal/db/repo"
)

type respBasicAuth struct {
	app.Response
	BasicAuth string `json:"basicAuth"`
}

func queryClientBasicAuth(ctx server.Context) {
	var resp respBasicAuth
	appId := ctx.C.URLParamDefault("appId", "")
	if len(appId) <= 0 {
		resp.Code = tool.RespCodeNotFound
		resp.Message = "appId required"
		ctx.Json(resp)
		return
	}

	client, err := repo.GetClientByAppId(appId)
	if err != nil {
		resp.Code = tool.RespCodeError
		resp.Message = "client not found"
		ctx.Json(resp)
		return
	}

	resp.BasicAuth = client.BasicAuth
	resp.Code = tool.RespCodeSuccess
	resp.Message = "success"
	ctx.Json(resp)
}
