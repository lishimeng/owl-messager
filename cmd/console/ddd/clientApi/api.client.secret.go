package clientApi

import (
	"github.com/lishimeng/app-starter"
	"github.com/lishimeng/app-starter/server"
	"github.com/lishimeng/app-starter/tool"
	"github.com/lishimeng/owl-messager/internal/db/repo"
)

type respSecret struct {
	app.Response
	Secret string `json:"secret"`
}

func queryClientSecret(ctx server.Context) {
	var resp respSecret
	var appId = ctx.C.URLParamDefault("appId", "")
	if len(appId) <= 0 {
		resp.Code = tool.RespCodeNotFound
		resp.Message = "URL参数解析失败"
		ctx.Json(resp)
		return
	}

	client, err := repo.GetClientByAppId(appId)
	if err != nil {
		resp.Code = tool.RespCodeError
		resp.Message = "删除失败"
		ctx.Json(resp)
		return
	}

	resp.Secret = client.Secret
	resp.Code = tool.RespCodeSuccess
	resp.Message = "success"
	ctx.Json(resp)
}
