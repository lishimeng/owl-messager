package sender

import (
	"github.com/lishimeng/app-starter"
	"github.com/lishimeng/app-starter/persistence"
	"github.com/lishimeng/app-starter/server"
	"github.com/lishimeng/app-starter/tool"
	"github.com/lishimeng/owl-messager/cmd/console/ddd/consoleorg"
	"github.com/lishimeng/owl-messager/internal/db/model"
)

func _setDefault(code string, category string, tenantCode string, provider string) (err error) {
	err = app.GetOrm().Transaction(func(ctx persistence.TxContext) (e error) {
		e = ctx.Model(&model.MessageSenderInfo{}).
			Equal("message_category", category).
			Equal("message_provider", provider).
			Equal("tenant_code", tenantCode).
			Updates(map[string]any{"default_sender": 0})
		if e != nil {
			return
		}
		e = ctx.Model(&model.MessageSenderInfo{}).
			Equal("code", code).
			Update("default_sender", 1)
		return
	})
	return
}

type reqSetDefault struct {
	Code       string `json:"code"`
	Category   string `json:"category"`
	Provider   string `json:"provider"`
	TenantCode string `json:"tenantCode"`
}

func setDefaultSender(ctx server.Context) {

	var resp app.Response
	var req reqSetDefault

	err := ctx.C.ReadJSON(&req)

	if err != nil || req.Category == "" || req.Provider == "" || req.Code == "" {
		resp.Code = tool.RespCodeNotFound
		ctx.Json(resp)
		return
	}
	tenantCode := req.TenantCode
	if tenantCode == "" {
		tenantCode = consoleorg.Code(ctx)
	}
	err = _setDefault(req.Code, req.Category, tenantCode, req.Provider)
	if err != nil {
		resp.Code = tool.RespCodeNotFound
		ctx.Json(resp)
		return
	}
	resp.Code = tool.RespCodeSuccess
	ctx.Json(resp)
}
