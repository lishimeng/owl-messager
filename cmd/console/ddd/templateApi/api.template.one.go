package templateApi

import (
	"github.com/lishimeng/app-starter/server"
	"github.com/lishimeng/app-starter/tool"
	"github.com/lishimeng/owl-messager/cmd/console/ddd/consoleorg"
	"github.com/lishimeng/owl-messager/internal/db/model"
	"github.com/lishimeng/owl-messager/internal/db/repo"
)

func GetTemplateInfo(ctx server.Context) {
	var resp respTemplate
	var code = ctx.C.URLParamDefault("code", "")
	if code == "" {
		resp.Code = tool.RespCodeNotFound
		resp.Message = "请求中code为空"
		ctx.Json(resp)
		return
	}

	tenantCode := consoleorg.TenantFilter(ctx)
	var tpl model.MessageTemplate
	var err error
	if tenantCode != "" {
		tpl, err = repo.GetMessageTemplateByCode(code, tenantCode)
	} else {
		tpl, err = repo.GetMessageTemplateByCodeAny(code)
	}
	if err != nil {
		resp.Code = tool.RespCodeNotFound
		resp.Message = "未查到记录"
		ctx.Json(resp)
		return
	}

	params, err := formatParamsForDisplay(tpl.Params)
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
