package templateApi

import (
	"github.com/lishimeng/app-starter"
	"github.com/lishimeng/app-starter/server"
	"github.com/lishimeng/app-starter/tool"
	"github.com/lishimeng/go-log"
	"github.com/lishimeng/owl-messager/internal/db/repo"
	"github.com/lishimeng/owl-messager/pkg/msg"
	"github.com/lishimeng/x/util"
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

// GetMailVendors 平台支持的mail类型
func GetMailVendors(ctx server.Context) {
	var resp SmsVendors

	resp.Code = tool.RespCodeSuccess
	resp.Message = "Mail Vendors"
	for key, _ := range msg.MailProviders {
		resp.Data = append(resp.Data, key)
	}
	ctx.Json(resp)
}

func GetMailTemplateList(ctx server.Context) {
	log.Debug("get mail template list")
	var resp app.PagerResponse
	//var status = ctx.URLParamIntDefault("status", repo.ConditionIgnore)
	var pageSize = ctx.C.URLParamIntDefault("pageSize", repo.DefaultPageSize)
	var pageNo = ctx.C.URLParamIntDefault("pageNo", repo.DefaultPageNo)
	page := app.Pager{ // TODO
		PageSize: pageSize,
		PageNum:  pageNo,
	}
	tpls, err := repo.GetMessageTemplates(1, msg.MailMessage, msg.Ali)
	if err != nil {
		log.Debug("get templates failed")
		log.Debug(err)
		resp.Code = -1
		resp.Message = "get templates failed"
		ctx.Json(resp)
		return
	}

	for _, tpl := range tpls {
		var tmpInfo = Info{
			Id:           tpl.Id,
			TemplateCode: tpl.Code,
			TemplateBody: tpl.Body,
			Status:       tpl.Status,
			CreateTime:   util.FormatTime(tpl.CreateTime),
			UpdateTime:   util.FormatTime(tpl.UpdateTime),
		}
		page.Data = append(page.Data, tmpInfo)
	}

	resp.Pager = page
	resp.Code = tool.RespCodeSuccess
	ctx.Json(resp)
}

func GetMailTemplateInfo(ctx server.Context) {
	//log.Debug("get mail template") TODO
	//var resp InfoWrapper
	//id, err := ctx.Params().GetInt("id")
	//if err != nil {
	//	log.Debug("id must be a int value")
	//	resp.Response.Code = tool.RespCodeNotFound
	//	resp.Message = tool.RespMsgIdNum
	//	tool.ResponseJSON(ctx, resp)
	//	return
	//}
	//log.Debug("id:%d", id)
	//tpl, err := repo.GetMailTemplateById(id)
	//if err != nil {
	//	log.Debug("get mail template failed")
	//	log.Debug(err)
	//	resp.Response.Code = tool.RespCodeNotFound
	//	resp.Message = tool.RespMsgNotFount
	//	tool.ResponseJSON(ctx, resp)
	//	return
	//}
	//
	//var tmpInfo = Info{
	//	Id:           tpl.Id,
	//	TemplateCode: tpl.Code,
	//	TemplateBody: tpl.Body,
	//	Status:       tpl.Status,
	//	CreateTime:   tool.FormatTime(tpl.CreateTime),
	//	UpdateTime:   tool.FormatTime(tpl.UpdateTime),
	//}
	//resp.Info = tmpInfo
	//resp.Code = tool.RespCodeSuccess
	//tool.ResponseJSON(ctx, resp)
}

type MailTemplateReq struct {
	Id          int    `json:"id,omitempty"`
	Code        string `json:"code,omitempty"`
	Name        string `json:"name,omitempty"`
	Body        string `json:"body,omitempty"`
	Description string `json:"description,omitempty"`
	Category    string `json:"category,omitempty"`
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
func AddMailTemplate(ctx server.Context) {
	log.Debug("add mail template")
	var req MailTemplateReq
	var resp InfoWrapper
	err := ctx.C.ReadJSON(&req)
	if err != nil {
		resp.Code = -1
		ctx.Json(resp)
		return
	}

	// check params
	if len(req.Name) == 0 {
		log.Debug("param name nil")
		resp.Code = -1
		resp.Message = "name nil"
		ctx.Json(resp)
		return
	}
	if len(req.Body) == 0 {
		log.Debug("param body nil")
		resp.Code = -1
		resp.Message = "body nil"
		ctx.Json(resp)
		return
	}

	if !msg.IsValidCategory(msg.MessageCategory(req.Category)) {
		log.Debug("param category nil, use default")
		req.Category = msg.MailMessage.String()
	}

	//code := tool.UUIDString()

	//m, err := repo.CreateMailTemplate(code, req.Name, req.Body, req.Description, req.Category)
	//if err != nil {
	//	log.Info("can't create template")
	//	log.Info(err)
	//	resp.Code = -1
	//	resp.Message = "create template failed"
	//	tool.ResponseJSON(ctx, resp)
	//	return
	//}

	//log.Debug("create template success, id:%d", m.Id)
	//resp.Id = m.Id
	//
	//var tmpInfo = Info{
	//	Id:           m.Id,
	//	TemplateCode: m.Code,
	//	TemplateBody: m.Body,
	//	Status:       m.Status,
	//	CreateTime:   tool.FormatTime(m.CreateTime),
	//	UpdateTime:   tool.FormatTime(m.UpdateTime),
	//}
	//resp.Info = tmpInfo
	//resp.Code = tool.RespCodeSuccess
	ctx.Json(resp)
}

func UpdateMailTemplate(ctx server.Context) {
	//log.Debug("update mail template") TODO
	//var req MailTemplateReq
	//var resp InfoWrapper
	//err := ctx.ReadJSON(&req)
	//if err != nil {
	//	log.Debug("req err")
	//	log.Debug(err)
	//	resp.Code = -1
	//	resp.Message = "req err"
	//	tool.ResponseJSON(ctx, resp)
	//	return
	//}
	//
	//// check params
	//if req.Id == 0 {
	//	log.Debug("param id nil")
	//	resp.Code = -1
	//	resp.Message = "id nil"
	//	tool.ResponseJSON(ctx, resp)
	//	return
	//}
	//if len(req.Body) == 0 {
	//	log.Debug("param body nil")
	//	resp.Code = -1
	//	resp.Message = "body nil"
	//	tool.ResponseJSON(ctx, resp)
	//	return
	//}
	//
	//m, err := service.UpdateMailTemplate(req.Id, req.Status, req.Body, req.Description)
	//if err != nil {
	//	log.Info("can't update template")
	//	log.Info(err)
	//	resp.Code = -1
	//	resp.Message = "create update failed"
	//	tool.ResponseJSON(ctx, resp)
	//	return
	//}
	//
	//log.Debug("update template success, id:%d", m.Id)
	//resp.Id = m.Id
	//
	//var tmpInfo = Info{
	//	Id:           m.Id,
	//	TemplateCode: m.Code,
	//	TemplateBody: m.Body,
	//	Status:       m.Status,
	//	CreateTime:   tool.FormatTime(m.CreateTime),
	//	UpdateTime:   tool.FormatTime(m.UpdateTime),
	//}
	//resp.Info = tmpInfo
	//resp.Code = tool.RespCodeSuccess
	//tool.ResponseJSON(ctx, resp)
}

func DeleteMailTemplate(ctx server.Context) {

	//log.Debug("delete mail template")
	//var resp app.Response
	//id, err := ctx.Params().GetInt("id")
	//if err != nil {
	//	log.Debug("id must be a int value")
	//	resp.Code = tool.RespCodeNotFound
	//	resp.Message = tool.RespMsgIdNum
	//	tool.ResponseJSON(ctx, resp)
	//	return
	//}
	//err = repo.DeleteMailTemplate(id)
	//if err != nil {
	//	log.Info("delete mail template failed")
	//	resp.Code = -1
	//	resp.Message = "delete template failed"
	//	tool.ResponseJSON(ctx, resp)
	//	return
	//}
	//resp.Code = tool.RespCodeSuccess
	//tool.ResponseJSON(ctx, resp) TODO
}

type MailStatusReq struct {
	Status int `json:"status,omitempty"`
	Id     int `json:"id,omitempty"`
}

func ChangeMailTemplateStatus(ctx server.Context) {

	//var req SmsStatusReq
	//var resp app.Response
	//var err error
	//
	//err = ctx.ReadJSON(&req)
	//if err != nil {
	//	resp.Code = tool.RespCodeError
	//	tool.ResponseJSON(ctx, resp)
	//	return
	//}
	//
	//if req.Id <= 0 {
	//	log.Debug("param id nil")
	//	resp.Code = tool.RespCodeError
	//	resp.Message = "id nil"
	//	tool.ResponseJSON(ctx, resp)
	//	return
	//}
	//
	//if !util.StatusIn(req.Status, model.MailTemplateStatus) {
	//	log.Debug("param unknown status: %d", req.Status)
	//	resp.Code = tool.RespCodeError
	//	resp.Message = fmt.Sprintf("unknown status:%d", req.Status)
	//	tool.ResponseJSON(ctx, resp)
	//	return
	//}
	//
	//tpl, err := repo.GetMailTemplateById(req.Id)
	//if err != nil {
	//	log.Debug("template not found")
	//	resp.Code = tool.RespCodeNotFound
	//	resp.Message = "template not found"
	//	tool.ResponseJSON(ctx, resp)
	//	return
	//}
	//
	//tpl.Status = req.Status
	//
	//_, err = repo.UpdateMailTemplateInfo(tpl, "status")
	//if err != nil {
	//	log.Debug(err)
	//	resp.Code = tool.RespCodeError
	//	resp.Message = err.Error()
	//	tool.ResponseJSON(ctx, resp)
	//	return
	//}
	//resp.Code = tool.RespCodeSuccess
	//tool.ResponseJSON(ctx, resp) TODO
}
