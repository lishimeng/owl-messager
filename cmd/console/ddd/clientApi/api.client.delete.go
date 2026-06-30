package clientApi

import (
	"github.com/lishimeng/app-starter"
	"github.com/lishimeng/app-starter/persistence"
	"github.com/lishimeng/app-starter/server"
	"github.com/lishimeng/app-starter/tool"
	"github.com/lishimeng/owl-messager/internal/db/repo"
)

func deleteClient(ctx server.Context) {
	var req reqClient
	var resp app.Response

	err := ctx.C.ReadJSON(&req)
	if err != nil {
		resp.Code = tool.RespCodeNotFound
		resp.Message = "json参数解析失败"
		ctx.Json(resp)
		return
	}
	if req.TenantCode == "" || req.AppId == "" {
		resp.Code = tool.RespCodeError
		resp.Message = "tenantCode and appId required"
		ctx.Json(resp)
		return
	}

	err = app.GetOrm().Transaction(func(ctx persistence.TxContext) (e error) {
		_, e = repo.GetTenant(req.TenantCode)
		if e != nil {
			return
		}
		e = repo.DeleteClient(ctx, req.TenantCode, req.AppId)
		return
	})

	if err != nil {
		resp.Code = tool.RespCodeError
		resp.Message = "删除失败"
		ctx.Json(resp)
		return
	}

	resp.Code = tool.RespCodeSuccess
	resp.Message = "success"
	ctx.Json(resp)
}
