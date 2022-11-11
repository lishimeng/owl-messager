package templateApi

import (
	"fmt"
	"github.com/kataras/iris/v12"
	"github.com/lishimeng/app-starter"
	"github.com/lishimeng/app-starter/tool"
	"github.com/lishimeng/go-log"
	"github.com/lishimeng/owl/internal/api/common"
	"github.com/lishimeng/owl/internal/db/model"
	"github.com/lishimeng/owl/internal/db/repo"
	"github.com/lishimeng/owl/internal/db/service"
)

type Info struct {
	Id           int    `json:"id,omitempty"`
	Codes string `json:"code,omitempty"`
	Body        string `json:"body,omitempty"`
	Name        string `json:"name,omitempty"`
	Status       int    `json:"status,omitempty"`
	Category    int    `json:"category,omitempty"`
	Description string `json:"description,omitempty"`
	CreateTime   string `json:"createTime,omitempty"`
	UpdateTime   string `json:"updateTime,omitempty"`
}

type InfoWrapper struct {
	app.Response
	Info
}

type InfoWrapperVo struct {
	app.Response
	Data interface{}  `json:"data"`
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
	page, tpls, err := repo.GetMailTemplateList(status, page)
	if err != nil {
		log.Debug("get templates failed")
		log.Debug(err)
		resp.Code = -1
		resp.Message = "get templates failed"
		tool.ResponseJSON(ctx, resp)
		return
	}

	for _, tpl := range tpls {
		var tmpInfo = Info{
			Id:           tpl.Id,
			Codes: tpl.Code,
			Body: tpl.Body,
			Status:       tpl.Status,
			Description: tpl.Description,
			Name: tpl.Name,
			Category: tpl.Category,
			CreateTime:   tool.FormatTime(tpl.CreateTime),
			UpdateTime:   tool.FormatTime(tpl.UpdateTime),
		}
		page.Data = append(page.Data, tmpInfo)
	}

	resp.Pager = page
	resp.Code = tool.RespCodeSuccess
	tool.ResponseJSON(ctx, resp)
}

func GetMailTemplateInfo(ctx iris.Context) {
	log.Debug("get mail template")
	var resp InfoWrapperVo
	id, err := ctx.Params().GetInt("id")
	if err != nil {
		log.Debug("id must be a int value")
		resp.Response.Code = tool.RespCodeNotFound
		resp.Message = tool.RespMsgIdNum
		tool.ResponseJSON(ctx, resp)
		return
	}
	log.Debug("id:%d", id)
	tpl, err := repo.GetMailTemplateById(id)
	if err != nil {
		log.Debug("get mail template failed")
		log.Debug(err)
		resp.Response.Code = tool.RespCodeNotFound
		resp.Message = tool.RespMsgNotFount
		tool.ResponseJSON(ctx, resp)
		return
	}

	var tmpInfo = Info{
		Id:           tpl.Id,
		Codes: tpl.Code,
		Body: tpl.Body,
		Status:       tpl.Status,
		Description: tpl.Description,
		Name: tpl.Name,
		Category: tpl.Category,
		CreateTime:   tool.FormatTime(tpl.CreateTime),
		UpdateTime:   tool.FormatTime(tpl.UpdateTime),
	}
	resp.Data = tmpInfo
	resp.Code = tool.RespCodeSuccess
	tool.ResponseJSON(ctx, resp)
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

// AddMailTemplate
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
		fmt.Println(err)

		tool.ResponseJSON(ctx, resp)
		return
	}

	// check params
	if len(req.Name) == 0 {
		log.Debug("param name nil")

		resp.Message = "name nil"
		tool.ResponseJSON(ctx, resp)
		return
	}
	if len(req.Body) == 0 {
		log.Debug("param body nil")

		resp.Message = "body nil"
		tool.ResponseJSON(ctx, resp)
		return
	}

	if req.Category == 0 {
		log.Debug("param category nil, use default")
		req.Category = model.MailTemplateCategoryDefault
	}

	code := tool.GetRandomString(common.DefaultCodeLen)
	code = "tpl_" + code

	m, err := repo.CreateMailTemplate(code, req.Name, req.Body, req.Description, req.Category)
	if err != nil {
		log.Info("can't create template")
		log.Info(err)

		resp.Message = "create template failed"
		tool.ResponseJSON(ctx, resp)
		return
	}

	log.Debug("create template success, id:%d", m.Id)
	resp.Id = m.Id

	var tmpInfo = Info{
		Id:           m.Id,
		Codes: m.Code,
		Body: m.Body,
		Status:       m.Status,
		Description: m.Description,
		Name: m.Name,
		Category: m.Category,
		CreateTime:   tool.FormatTime(m.CreateTime),
		UpdateTime:   tool.FormatTime(m.UpdateTime),
	}
	resp.Info = tmpInfo
	resp.Code = tool.RespCodeSuccess
	tool.ResponseJSON(ctx, resp)
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
		tool.ResponseJSON(ctx, resp)
		return
	}

	// check params
	if req.Id == 0 {
		log.Debug("param id nil")
		resp.Code = -1
		resp.Message = "id nil"
		tool.ResponseJSON(ctx, resp)
		return
	}
	if len(req.Body) == 0 {
		log.Debug("param body nil")
		resp.Code = -1
		resp.Message = "body nil"
		tool.ResponseJSON(ctx, resp)
		return
	}

	m, err := service.UpdateMailTemplate(req.Id, req.Status, req.Body, req.Description,req.Name)
	if err != nil {
		log.Info("can't update template")
		log.Info(err)
		resp.Code = -1
		resp.Message = "create update failed"
		tool.ResponseJSON(ctx, resp)
		return
	}

	log.Debug("update template success, id:%d", m.Id)
	resp.Id = m.Id

	var tmpInfo = Info{
		Id:           m.Id,
		Codes: m.Code,
		Body: m.Body,
		Status:       m.Status,
		Description: m.Description,
		Name: m.Name,
		Category: m.Category,
		CreateTime:   tool.FormatTime(m.CreateTime),
		UpdateTime:   tool.FormatTime(m.UpdateTime),
	}
	resp.Info = tmpInfo
	resp.Code = tool.RespCodeSuccess
	tool.ResponseJSON(ctx, resp)
}

func DeleteMailTemplate(ctx iris.Context) {

	log.Debug("delete mail template")
	var resp app.Response
	id, err := ctx.Params().GetInt("id")
	if err != nil {
		log.Debug("id must be a int value")
		resp.Code = tool.RespCodeNotFound
		resp.Message = tool.RespMsgIdNum
		tool.ResponseJSON(ctx, resp)
		return
	}
	err = repo.DeleteMailTemplate(id)
	if err != nil {
		log.Info("delete mail template failed")
		resp.Code = -1
		resp.Message = "delete template failed"
		tool.ResponseJSON(ctx, resp)
		return
	}
	resp.Code = tool.RespCodeSuccess
	tool.ResponseJSON(ctx, resp)
}
