package templateApi

import (
	"github.com/kataras/iris/v12"
	"github.com/lishimeng/app-starter"
	"github.com/lishimeng/app-starter/tool"
	"github.com/lishimeng/owl-messager/internal/db/repo"
	"github.com/lishimeng/owl-messager/pkg/msg"
)

func GetTemplateListByPage(ctx iris.Context) {
	var resp app.PagerResponse
	var category = ctx.URLParamDefault("category", "")
	var pageNum = ctx.URLParamIntDefault("pageNum", 1)
	var pageSize = ctx.URLParamIntDefault("pageSize", 10)
	page := app.Pager{
		PageSize: pageSize,
		PageNum:  pageNum,
	}
	switch msg.MessageCategory(category) {
	case msg.MailMessage:
		list, err := repo.GetMessageTemplates(1, msg.MailMessage, msg.Ali) // TODO
		if err != nil {
			resp.Code = tool.RespCodeNotFound
			tool.ResponseJSON(ctx, resp)
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
			tool.ResponseJSON(ctx, resp)
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
		tool.ResponseJSON(ctx, resp)
		return
	}
	resp.Code = tool.RespCodeSuccess
	tool.ResponseJSON(ctx, resp)
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

func GetTemplateInfo(ctx iris.Context) {
	var resp respTemplate
	var code = ctx.URLParamDefault("code", "")
	var category = ctx.URLParamDefault("category", "")
	switch msg.MessageCategory(category) {
	case msg.MailMessage:
		info, err := repo.GetTemplateByCode(code, msg.MailMessage)
		if err != nil {
			resp.Code = tool.RespCodeNotFound
			resp.Message = "未查到记录"
			tool.ResponseJSON(ctx, resp)
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
			tool.ResponseJSON(ctx, resp)
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
		tool.ResponseJSON(ctx, resp)
		return
	}
	resp.Code = tool.RespCodeSuccess
	resp.Message = "成功"
	tool.ResponseJSON(ctx, resp)
}

func CreateTemplate(ctx iris.Context) {
	var resp app.Response
	var req TemplateReq
	err := ctx.ReadJSON(&req)
	if err != nil {
		resp.Code = tool.RespCodeNotFound
		resp.Message = "json参数解析失败"
		tool.ResponseJSON(ctx, resp)
		return
	}
	code := tool.UUIDString()
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
			tool.ResponseJSON(ctx, resp)
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
			tool.ResponseJSON(ctx, resp)
			return
		}
	default:
		resp.Code = tool.RespCodeNotFound
		resp.Message = "失败,无此类型"
		tool.ResponseJSON(ctx, resp)
		return
	}
	resp.Code = tool.RespCodeSuccess
	resp.Message = "成功"
	tool.ResponseJSON(ctx, resp)
}

func UpdateTemplate(ctx iris.Context) {
	var resp app.Response
	var req TemplateReq
	err := ctx.ReadJSON(&req)
	if err != nil {
		resp.Code = tool.RespCodeNotFound
		tool.ResponseJSON(ctx, resp)
		return
	}
	switch msg.MessageCategory(req.Category) {
	case msg.MailMessage: // TODO
		_, err := repo.UpdateMessageTemplate(req.Status, req.Code, req.Name, req.Body, req.Description)
		if err != nil {
			resp.Code = tool.RespCodeNotFound
			tool.ResponseJSON(ctx, resp)
			return
		}
	case msg.SmsMessage: // TODO
		_, err := repo.UpdateMessageTemplate(req.Status, req.Code, req.Name, req.Body, req.Description)
		if err != nil {
			resp.Code = tool.RespCodeNotFound
			tool.ResponseJSON(ctx, resp)
			return
		}
	default:
		resp.Code = tool.RespCodeNotFound
		tool.ResponseJSON(ctx, resp)
		return
	}
	resp.Code = tool.RespCodeSuccess
	tool.ResponseJSON(ctx, resp)
}
