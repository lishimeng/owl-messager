package templateApi

import (
	"github.com/lishimeng/app-starter"
	"github.com/lishimeng/app-starter/server"
	"github.com/lishimeng/app-starter/tool"
	"github.com/lishimeng/owl-messager/internal/db/repo"
	"github.com/lishimeng/owl-messager/pkg/msg"
	"github.com/lishimeng/x/util"
)

func CreateTemplate(ctx server.Context) {
	var resp app.Response
	var req TemplateReq
	err := ctx.C.ReadJSON(&req)
	if err != nil {
		resp.Code = tool.RespCodeNotFound
		resp.Message = "json参数解析失败"
		ctx.Json(resp)
		return
	}

	code := util.UUIDString()
	switch msg.MessageCategory(req.Category) {
	case msg.MailMessage:
		code = "tl_mail_" + code
	case msg.SmsMessage:
		code = "tl_sms_" + code
	case msg.ImMessage:
		code = "tl_im_" + code
	default:
		resp.Code = tool.RespCodeNotFound
		resp.Message = "失败：未知通讯方式"
		ctx.Json(resp)
		return
	}
	// 处理param
	params, err := paramsToMap(req.Params)
	if err != nil {
		resp.Code = tool.RespCodeNotFound
		ctx.Json(resp)
		return
	}
	if req.TenantCode == "" {
		resp.Code = tool.RespCodeNotFound
		resp.Message = "tenantCode required"
		ctx.Json(resp)
		return
	}
	_, err = repo.CreateMessageTemplate(
		req.TenantCode,
		code, req.Name, req.Body, req.CloudTemplate, params, req.Description,
		msg.MessageCategory(req.Category), msg.MessageProvider(req.Provider),
	)
	if err != nil {
		resp.Code = tool.RespCodeError
		resp.Message = "添加失败"
		ctx.Json(resp)
		return
	}

	resp.Code = tool.RespCodeSuccess
	resp.Message = "成功"
	ctx.Json(resp)
}
