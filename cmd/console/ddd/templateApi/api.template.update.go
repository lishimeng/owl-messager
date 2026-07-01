package templateApi

import (
	"github.com/lishimeng/app-starter"
	"github.com/lishimeng/app-starter/server"
	"github.com/lishimeng/app-starter/tool"
	"github.com/lishimeng/owl-messager/internal/db/repo"
)

func UpdateTemplate(ctx server.Context) {
	var resp app.Response
	var req TemplateReq
	err := ctx.C.ReadJSON(&req)
	if err != nil {
		resp.Code = tool.RespCodeNotFound
		ctx.Json(resp)
		return
	}
	params, err := normalizeParams(req.Params)
	if err != nil {
		resp.Code = tool.RespCodeNotFound
		resp.Message = err.Error()
		ctx.Json(resp)
		return
	}
	_, err = repo.UpdateMessageTemplate(req.Status, req.Code, req.Name, req.Body, params, req.Description, req.Provider)
	if err != nil {
		resp.Code = tool.RespCodeNotFound
		ctx.Json(resp)
		return
	}
	resp.Code = tool.RespCodeSuccess
	ctx.Json(resp)
}
