package templateApi

import (
	"github.com/lishimeng/app-starter"
	"github.com/lishimeng/app-starter/server"
	"github.com/lishimeng/app-starter/tool"
	"github.com/lishimeng/owl-messager/internal/db/repo"
	"github.com/lishimeng/owl-messager/pkg/msg"
	"github.com/lishimeng/x/util"
)

func GetTemplateListByPage(ctx server.Context) {
	var resp app.PagerResponse
	var category = ctx.C.URLParamDefault("category", "")
	var pageNum = ctx.C.URLParamIntDefault("pageNum", 1)
	var pageSize = ctx.C.URLParamIntDefault("pageSize", 10)
	page := app.Pager{
		PageSize: pageSize,
		PageNum:  pageNum,
	}
	switch msg.MessageCategory(category) {
	case msg.MailMessage:
		list, err := repo.GetMessageTemplates(1, msg.MailMessage, msg.Ali) // TODO
		if err != nil {
			resp.Code = tool.RespCodeNotFound
			ctx.Json(resp)
			return
		}
		if len(list) > 0 {
			for _, tl := range list {
				page.Data = append(page.Data, tl)
			}
		}
		resp.Pager = page
	case msg.SmsMessage:
		list, err := repo.GetMessageTemplates(1, msg.SmsMessage, msg.Ali) // TODO
		if err != nil {
			resp.Code = tool.RespCodeNotFound
			ctx.Json(resp)
			return
		}
		if len(list) > 0 {
			for _, tl := range list {
				page.Data = append(page.Data, tl)
			}
		}
		resp.Pager = page
	default:
		resp.Code = tool.RespCodeNotFound
		ctx.Json(resp)
		return
	}
	resp.Code = tool.RespCodeSuccess
	ctx.Json(resp)
}

type TemplateReq struct {
	Name          string `json:"name,omitempty"`
	Body          string `json:"body,omitempty"`
	CloudTemplate string `json:"cloudTemplate,omitempty"`
	Description   string `json:"description,omitempty"`
	Category      string `json:"category,omitempty"`
	Params        string `json:"params,omitempty"`
	Provider      string `json:"provider,omitempty"`
	Status        int    `json:"status,omitempty"`
	Code          string `json:"code,omitempty"`
}
type respTemplate struct {
	app.Response
	Item TemplateReq `json:"item"`
}

func GetTemplateInfo(ctx server.Context) {
	var resp respTemplate
	var code = ctx.C.URLParamDefault("code", "")
	var category = ctx.C.URLParamDefault("category", "")
	switch msg.MessageCategory(category) {
	case msg.MailMessage:
		info, err := repo.GetTemplateByCode(code, msg.MailMessage)
		if err != nil {
			resp.Code = tool.RespCodeNotFound
			resp.Message = "未查到记录"
			ctx.Json(resp)
			return
		}
		resp.Item = TemplateReq{
			Code:        info.Code,
			Name:        info.Name,
			Body:        info.Body,
			Description: info.Description,
			Provider:    string(info.Provider),
		}
	case msg.SmsMessage:
		info, err := repo.GetTemplateByCode(code, msg.SmsMessage)
		if err != nil {
			resp.Code = tool.RespCodeNotFound
			resp.Message = "未查到记录"
			ctx.Json(resp)
			return
		}
		resp.Item = TemplateReq{
			Code:        info.Code,
			Name:        info.Name,
			Body:        info.Body,
			Description: info.Description,
			Params:      info.Params,
			Provider:    string(info.Provider),
		}
	default:
		resp.Code = tool.RespCodeNotFound
		resp.Message = "失败,无此类型"
		ctx.Json(resp)
		return
	}
	resp.Code = tool.RespCodeSuccess
	resp.Message = "成功"
	ctx.Json(resp)
}

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
		_, err := repo.CreateMessageTemplate(
			code, req.Name, req.Body, req.CloudTemplate, req.Params, req.Description,
			msg.MailMessage, msg.MessageProvider(req.Provider),
		)
		if err != nil {
			resp.Code = tool.RespCodeNotFound
			resp.Message = "添加失败"
			ctx.Json(resp)
			return
		}
	case msg.SmsMessage:
		code = "tl_sms_" + code
		_, err := repo.CreateMessageTemplate(
			code, req.Name, req.Body, req.CloudTemplate, req.Params, req.Description,
			msg.SmsMessage, msg.MessageProvider(req.Provider),
		)
		if err != nil {
			resp.Code = tool.RespCodeNotFound
			resp.Message = "失败"
			ctx.Json(resp)
			return
		}
	default:
		resp.Code = tool.RespCodeNotFound
		resp.Message = "失败,无此类型"
		ctx.Json(resp)
		return
	}
	resp.Code = tool.RespCodeSuccess
	resp.Message = "成功"
	ctx.Json(resp)
}

func UpdateTemplate(ctx server.Context) {
	var resp app.Response
	var req TemplateReq
	err := ctx.C.ReadJSON(&req)
	if err != nil {
		resp.Code = tool.RespCodeNotFound
		ctx.Json(resp)
		return
	}
	switch msg.MessageCategory(req.Category) {
	case msg.MailMessage: // TODO
		_, err := repo.UpdateMessageTemplate(req.Status, req.Code, req.Name, req.Body, req.Description)
		if err != nil {
			resp.Code = tool.RespCodeNotFound
			ctx.Json(resp)
			return
		}
	case msg.SmsMessage: // TODO
		_, err := repo.UpdateMessageTemplate(req.Status, req.Code, req.Name, req.Body, req.Description)
		if err != nil {
			resp.Code = tool.RespCodeNotFound
			ctx.Json(resp)
			return
		}
	default:
		resp.Code = tool.RespCodeNotFound
		ctx.Json(resp)
		return
	}
	resp.Code = tool.RespCodeSuccess
	ctx.Json(resp)
}
