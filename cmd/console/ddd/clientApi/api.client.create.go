package clientApi

import (
	"github.com/lishimeng/app-starter"
	"github.com/lishimeng/app-starter/persistence"
	"github.com/lishimeng/app-starter/server"
	"github.com/lishimeng/app-starter/tool"
	"github.com/lishimeng/owl-messager/cmd/console/ddd/consoleorg"
	"github.com/lishimeng/owl-messager/internal/db/repo"
)

type respCreate struct {
	respSecret
	AppId     string `json:"appId,omitempty"`
	BasicAuth string `json:"basicAuth,omitempty"`
}

func createClient(ctx server.Context) {
	var req reqClient
	var resp respCreate

	err := ctx.C.ReadJSON(&req)
	if err != nil {
		resp.Code = tool.RespCodeNotFound
		resp.Message = "json参数解析失败"
		ctx.Json(resp)
		return
	}
	tenantCode := req.TenantCode
	if tenantCode == "" {
		tenantCode = consoleorg.Code(ctx)
	}
	if tenantCode == "" {
		resp.Code = tool.RespCodeError
		resp.Message = "tenantCode required"
		ctx.Json(resp)
		return
	}

	err = app.GetOrm().Transaction(func(ctx persistence.TxContext) (e error) {
		_, e = repo.GetTenant(tenantCode)
		if e != nil {
			return
		}
		client, e := repo.AddClient(ctx, tenantCode, req.Name)
		if e != nil {
			return
		}
		resp.AppId = client.AppId
		resp.Secret = client.Secret
		resp.BasicAuth = client.BasicAuth
		return
	})

	if err != nil {
		resp.Code = tool.RespCodeError
		resp.Message = "创建失败"
		ctx.Json(resp)
		return
	}

	resp.Code = tool.RespCodeSuccess
	resp.Message = "success"
	ctx.Json(resp)
}
