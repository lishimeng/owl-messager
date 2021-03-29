package templateApi

import (
	"github.com/kataras/iris"
	"github.com/lishimeng/app-starter"
	"github.com/lishimeng/go-log"
	"github.com/lishimeng/owl/internal/api/common"
	"github.com/lishimeng/owl/internal/db/model"
	"github.com/lishimeng/owl/internal/db/repo"
	"github.com/lishimeng/owl/internal/db/service"
)

type Info struct {
	Id           int    `json:"id,omitempty"`
	TemplateCode string `json:"templateCode,omitempty"`
	TemplateBody string `json:"templateBody,omitempty"`
	Status       int    `json:"status,omitempty"`
	CreateTime   string `json:"createTime,omitempty"`
	UpdateTime   string `json:"updateTime,omitempty"`
}

type InfoWrapper struct {
	app.Response
	Info
}

func GetMailTemplateList(ctx iris.Context) {
	log.Debug("get mail template list")
	var resp app.PagerResponse
	var status = ctx.URLParamIntDefault("status", repo.ConditionIgnore)
	var pageSize = ctx.URLParamIntDefault("pageSize", repo.DefaultPageSize)
	var pageNo = ctx.URLParamIntDefault("pageNo", repo.DefaultPageNo)
	page := app.Pager{
		PageSize: pageSize,
		PageNum:  pageNo,
	}
	page, err := repo.GetMailTemplateList(status, page)
	if err != nil {
		log.Debug("get templates failed")
		log.Debug(err)
		resp.Code = -1
		resp.Message = "get templates failed"
		common.ResponseJSON(ctx, resp)
		return
	}

	resp.Pager = page
	common.ResponseJSON(ctx, resp)
}

func GetMailTemplateInfo(ctx iris.Context) {
	log.Debug("get mail template")
	var resp InfoWrapper
	id, err := ctx.Params().GetInt("id")
	if err != nil {
		log.Debug("id must be a int value")
		resp.Response.Code = common.RespCodeNotFound
		resp.Message = common.RespMsgIdNum
		common.ResponseJSON(ctx, resp)
		return
	}
	log.Debug("id:%d", id)
	tpl, err := repo.GetMailTemplateById(id)
	if err != nil {
		log.Debug("get mail template failed")
		log.Debug(err)
		resp.Response.Code = common.RespCodeNotFound
		resp.Message = common.RespMsgNotFount
		common.ResponseJSON(ctx, resp)
		return
	}

	var tmpInfo = Info{
		Id:           tpl.Id,
		TemplateCode: tpl.Code,
		TemplateBody: tpl.Body,
		Status:       tpl.Status,
		CreateTime:   common.FormatTime(tpl.CreateTime),
		UpdateTime:   common.FormatTime(tpl.UpdateTime),
	}
	resp.Info = tmpInfo

	common.ResponseJSON(ctx, resp)
}

type MailTemplateReq struct {
	Id          int    `json:"id,omitempty"`
	Code        string `json:"code,omitempty"`
	Name        string `json:"name,omitempty"`
	Body        string `json:"body,omitempty"`
	Description string `json:"description,omitempty"`
	Category    int    `json:"category,omitempty"`
	Status      int    `json:"status,omitempty"`
}

/**
@Summary add a new template for send email

@Accept  json

@Produce  json

@Router /api/mail_template [post]

@Example

http://localhost/api/mail_template

{"body":"fasfasgasd", "category":2,"name":"电量低超提醒"}
*/
func AddMailTemplate(ctx iris.Context) {
	log.Debug("add mail template")
	var req MailTemplateReq
	var resp InfoWrapper
	err := ctx.ReadJSON(&req)
	if err != nil {
		resp.Code = -1
		common.ResponseJSON(ctx, resp)
		return
	}

	// check params
	if len(req.Name) == 0 {
		log.Debug("param name nil")
		resp.Code = -1
		resp.Message = "name nil"
		common.ResponseJSON(ctx, resp)
		return
	}
	if len(req.Body) == 0 {
		log.Debug("param body nil")
		resp.Code = -1
		resp.Message = "body nil"
		common.ResponseJSON(ctx, resp)
		return
	}

	if req.Category == 0 {
		log.Debug("param category nil, use default")
		req.Category = model.MailTemplateCategoryDefault
	}

	code := common.GetRandomString(common.DefaultCodeLen)
	code = "tpl_" + code

	m, err := repo.CreateMailTemplate(code, req.Name, req.Body, req.Description, req.Category)
	if err != nil {
		log.Info("can't create template")
		log.Info(err)
		resp.Code = -1
		resp.Message = "create template failed"
		common.ResponseJSON(ctx, resp)
		return
	}

	log.Debug("create template success, id:%d", m.Id)
	resp.Id = m.Id

	var tmpInfo = Info{
		Id:           m.Id,
		TemplateCode: m.Code,
		TemplateBody: m.Body,
		Status:       m.Status,
		CreateTime:   common.FormatTime(m.CreateTime),
		UpdateTime:   common.FormatTime(m.UpdateTime),
	}
	resp.Info = tmpInfo
	resp.Code = common.RespCodeSuccess
	common.ResponseJSON(ctx, resp)
}

func UpdateMailTemplate(ctx iris.Context) {
	log.Debug("update mail template")
	var req MailTemplateReq
	var resp InfoWrapper
	err := ctx.ReadJSON(&req)
	if err != nil {
		log.Debug("req err")
		log.Debug(err)
		resp.Code = -1
		resp.Message = "req err"
		common.ResponseJSON(ctx, resp)
		return
	}

	// check params
	if req.Id == 0 {
		log.Debug("param id nil")
		resp.Code = -1
		resp.Message = "id nil"
		common.ResponseJSON(ctx, resp)
		return
	}
	if len(req.Body) == 0 {
		log.Debug("param body nil")
		resp.Code = -1
		resp.Message = "body nil"
		common.ResponseJSON(ctx, resp)
		return
	}

	m, err := service.UpdateMailTemplate(req.Id, req.Status, req.Body, req.Description)
	if err != nil {
		log.Info("can't update template")
		log.Info(err)
		resp.Code = -1
		resp.Message = "create update failed"
		common.ResponseJSON(ctx, resp)
		return
	}

	log.Debug("update template success, id:%d", m.Id)
	resp.Id = m.Id

	var tmpInfo = Info{
		Id:           m.Id,
		TemplateCode: m.Code,
		TemplateBody: m.Body,
		Status:       m.Status,
		CreateTime:   common.FormatTime(m.CreateTime),
		UpdateTime:   common.FormatTime(m.UpdateTime),
	}
	resp.Info = tmpInfo
	resp.Code = common.RespCodeSuccess
	common.ResponseJSON(ctx, resp)
}

func DeleteMailTemplate(ctx iris.Context) {

	log.Debug("delete mail template")
	var resp app.Response
	id, err := ctx.Params().GetInt("id")
	if err != nil {
		log.Debug("id must be a int value")
		resp.Code = -1
		resp.Message = "id must be a int value"
		common.ResponseJSON(ctx, resp)
		return
	}
	err = repo.DeleteMailTemplate(id)
	if err != nil {
		log.Info("delete mail template failed")
		resp.Code = -1
		resp.Message = "delete template failed"
		common.ResponseJSON(ctx, resp)
		return
	}

	common.ResponseJSON(ctx, resp)
}
