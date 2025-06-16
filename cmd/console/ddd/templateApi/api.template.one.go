package templateApi

import (
	"github.com/lishimeng/app-starter/server"
	"github.com/lishimeng/app-starter/tool"
	"github.com/lishimeng/owl-messager/internal/db/repo"
)

func GetTemplateInfo(ctx server.Context) {
	var resp respTemplate
	var code = ctx.C.URLParamDefault("code", "")
	//var category = ctx.C.URLParamDefault("category", "")
	if code == "" {
		resp.Code = tool.RespCodeNotFound
		resp.Message = "请求中code为空"
		ctx.Json(resp)
		return
	}

	tpl, err := repo.GetMessageTemplateByCode(code)
	if err != nil {
		resp.Code = tool.RespCodeNotFound
		resp.Message = "未查到记录"
		ctx.Json(resp)
		return
	}

	resp.Item = TemplateResp{
		Code:        tpl.Code,
		Name:        tpl.Name,
		Body:        tpl.Body,
		Description: tpl.Description,
		Provider:    string(tpl.Provider),
	}
	resp.Code = tool.RespCodeSuccess
	resp.Message = "成功"
	ctx.Json(resp)
}
