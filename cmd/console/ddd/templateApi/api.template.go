package templateApi

import (
	"github.com/beego/beego/v2/client/orm"
	"github.com/lishimeng/app-starter"
	"github.com/lishimeng/app-starter/persistence"
	"github.com/lishimeng/app-starter/server"
	"github.com/lishimeng/app-starter/tool"
	"github.com/lishimeng/go-log"
	"github.com/lishimeng/owl-messager/internal/db/model"
	"github.com/lishimeng/owl-messager/internal/db/repo"
	"github.com/lishimeng/owl-messager/pkg/msg"
	"github.com/lishimeng/x/util"
	"time"
)

type respPager struct {
	app.PagerResponse
	app.BasePager
	Items []TemplateResp `json:"items"`
}

func GetTemplateListByPage(ctx server.Context) {
	var resp respPager
	var category = ctx.C.URLParamDefault("category", "")
	var pageNum = ctx.C.URLParamIntDefault("pageNum", 1)
	var pageSize = ctx.C.URLParamIntDefault("pageSize", 10)
	var pager app.SimplePager[model.MessageTemplate, TemplateResp]
	pager.PageSize = pageSize
	pager.PageNum = pageNum
	pager.Transform = func(src model.MessageTemplate, dst *TemplateResp) {
		dst.Id = src.Id
		dst.Name = src.Name
		dst.Body = src.Body
		dst.CloudTemplate = src.CloudTemplate
		dst.Description = src.Description
		dst.Category = string(src.Category)
		dst.Params = src.Params
		dst.Provider = string(src.Provider)
		dst.Status = src.Status
		dst.Code = src.Code
		dst.CreateTime = src.CreateTime.UTC().Format(time.RFC3339)
		dst.UpdateTime = src.UpdateTime.UTC().Format(time.RFC3339)
	}
	pager.QueryBuilder = func(tx persistence.TxContext) any {
		cond := orm.NewCondition()
		cond = cond.And("org", 1)
		cond = cond.And("org", 1)
		if len(category) > 0 {
			cond = cond.And("message_category", category)
		}
		//todo: 按 Provider SenderEnable 筛选
		return tx.Context.QueryTable(new(model.MessageTemplate)).SetCond(cond)
	}
	//pager.OrderByExp = append(pager.OrderByExp, "createTime")
	err := app.QueryPage(&pager)
	if err != nil {
		log.Info("GetTemplateListByPage failed: %s ", err)
		resp.Code = tool.RespCodeNotFound
		resp.Message = "not found"
		ctx.Json(resp)
		return
	}
	resp.Items = pager.Data
	resp.BasePager = pager.BasePager
	resp.BasePager.More = pager.TotalPage * pager.PageSize

	resp.Code = tool.RespCodeSuccess
	ctx.Json(resp)
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
