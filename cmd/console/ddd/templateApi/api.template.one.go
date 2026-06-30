package templateApi

import (
	"github.com/lishimeng/app-starter/server"
	"github.com/lishimeng/app-starter/tool"
	"github.com/lishimeng/owl-messager/cmd/console/ddd/consoleorg"
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

	tpl, err := repo.GetMessageTemplateByCode(code, consoleorg.Code(ctx))
	if err != nil {
		resp.Code = tool.RespCodeNotFound
		resp.Message = "未查到记录"
		ctx.Json(resp)
		return
	}

	params, err := mapToParams(tpl.Params)
	if err != nil {
		resp.Code = tool.RespCodeError
		resp.Message = "参数异常"
		ctx.Json(resp)
		return
	}

	resp.Item = TemplateResp{
		Code:          tpl.Code,
		Name:          tpl.Name,
		Body:          tpl.Body,
		Params:        params,
		CloudTemplate: tpl.CloudTemplate,
		Description:   tpl.Description,
		Provider:      string(tpl.Provider),
	}
	resp.Code = tool.RespCodeSuccess
	resp.Message = "成功"
	ctx.Json(resp)
}
