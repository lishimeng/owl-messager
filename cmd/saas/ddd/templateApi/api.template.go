package templateApi

import (
	"github.com/kataras/iris/v12"
	"github.com/lishimeng/app-starter"
	"github.com/lishimeng/app-starter/tool"
	"github.com/lishimeng/owl-messager/internal/api/common"
	"github.com/lishimeng/owl-messager/internal/db/model"
	"github.com/lishimeng/owl-messager/internal/db/repo"
	"github.com/lishimeng/owl-messager/internal/db/service"
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
	switch category {
	case model.SenderCategoryMail:
		page, list, err := repo.GetMailTemplateList(1, page)
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
	case model.SenderCategorySms:
		page, list, err := repo.GetSmsTemplateList(1, page)
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
	Name        string `json:"name,omitempty"`
	Body        string `json:"body,omitempty"`
	Description string `json:"description,omitempty"`
	Category    string `json:"category,omitempty"`
	TemplateId  string `json:"templateId,omitempty"`
	Params      string `json:"params,omitempty"`
	Sender      int    `json:"sender,omitempty"`
	Signature   string `json:"signature,omitempty"`
	Vendor      string `json:"vendor,omitempty"`
	Status      int    `json:"status,omitempty"`
	Code        string `json:"code,omitempty"`
}
type respTemplate struct {
	app.Response
	Item TemplateReq `json:"item"`
}

func GetTemplateInfo(ctx iris.Context) {
	var resp respTemplate
	var code = ctx.URLParamDefault("code", "")
	var category = ctx.URLParamDefault("category", "")
	switch category {
	case model.SenderCategoryMail:
		info, err := repo.GetMailTemplateByCode(code)
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
			Vendor:      info.Vendor,
		}
	case model.SenderCategorySms:
		info, err := repo.GetSmsTemplateByCode(code)
		if err != nil {
			resp.Code = tool.RespCodeNotFound
			resp.Message = "未查到记录"
			tool.ResponseJSON(ctx, resp)
			return
		}
		resp.Item = TemplateReq{
			Code:        info.Code,
			Name:        info.Name,
			Sender:      info.Sender,
			Body:        info.Body,
			TemplateId:  info.CloudTemplateId,
			Signature:   info.Signature,
			Description: info.Description,
			Params:      info.Params,
			Vendor:      info.Vendor,
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
	code := tool.GetRandomString(common.DefaultCodeLen)
	switch req.Category {
	case model.SenderCategoryMail:
		code = "tl_mail_" + code
		_, err := repo.CreateMailTemplateNew(code, req.Name, req.Body, req.Description, req.Vendor, 2)
		if err != nil {
			resp.Code = tool.RespCodeNotFound
			resp.Message = "添加失败"
			tool.ResponseJSON(ctx, resp)
			return
		}
	case model.SenderCategorySms:
		code = "tl_sms_" + code
		_, err := repo.CreateSmsTemplateNew(code, req.Name, req.TemplateId, req.Params, req.Description, req.Vendor, req.Signature, req.Sender)
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
	switch req.Category {
	case model.SenderCategoryMail:
		_, err := service.UpdateMailTemplateByCode(req.Status, req.Code, req.Name, req.Body, req.Description)
		if err != nil {
			resp.Code = tool.RespCodeNotFound
			tool.ResponseJSON(ctx, resp)
			return
		}
	case model.SenderCategorySms:
		_, err := service.UpdateSmsTemplateByCode(req.Status, req.Sender, req.Code, req.Name,
			req.Body, req.TemplateId, req.Signature, req.Description, req.Params, req.Vendor)
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
