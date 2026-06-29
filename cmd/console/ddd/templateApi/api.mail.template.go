package templateApi

import (
	"github.com/beego/beego/v2/client/orm"
	"github.com/lishimeng/app-starter"
	"github.com/lishimeng/app-starter/persistence"
	"github.com/lishimeng/app-starter/server"
	"github.com/lishimeng/app-starter/tool"
	"github.com/lishimeng/go-log"
	"github.com/lishimeng/owl-messager/cmd/console/ddd/consoleorg"
	"github.com/lishimeng/owl-messager/internal/db/model"
	"github.com/lishimeng/owl-messager/internal/db/repo"
	"github.com/lishimeng/owl-messager/pkg/msg"
	"github.com/lishimeng/x/util"
)

type Info struct {
	Id           int    `json:"id,omitempty"`
	TemplateCode string `json:"templateCode,omitempty"`
	TemplateBody string `json:"templateBody,omitempty"`
	Status       int    `json:"itemStatus,omitempty"`
	CreateTime   string `json:"createTime,omitempty"`
	UpdateTime   string `json:"updateTime,omitempty"`
}

type respInfo struct {
	app.PagerResponse
	app.BasePager
	Items []Info `json:"items"`
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
	for key := range msg.MailProviders {
		resp.Data = append(resp.Data, key)
	}
	ctx.Json(resp)
}

func GetMailTemplateList(ctx server.Context) {
	var resp respInfo
	var pageSize = ctx.C.URLParamIntDefault("pageSize", repo.DefaultPageSize)
	var pageNo = ctx.C.URLParamIntDefault("pageNo", repo.DefaultPageNo)
	var provider = ctx.C.URLParamDefault("provider", "")
	orgID := consoleorg.ID(ctx)

	var pager app.SimplePager[model.MessageTemplate, Info]
	pager.PageSize = pageSize
	pager.PageNum = pageNo
	pager.Transform = func(src model.MessageTemplate, dst *Info) {
		dst.Id = src.Id
		dst.TemplateCode = src.Code
		dst.TemplateBody = src.Body
		dst.Status = src.Status
		dst.CreateTime = util.FormatTime(src.CreateTime)
		dst.UpdateTime = util.FormatTime(src.UpdateTime)
	}
	pager.QueryBuilder = func(tx persistence.TxContext) any {
		cond := orm.NewCondition()
		cond = cond.And("org", orgID)
		cond = cond.And("message_category", msg.MailMessage)
		if len(provider) > 0 {
			cond = cond.And("message_provider", provider)
		}
		return tx.Context.QueryTable(new(model.MessageTemplate)).SetCond(cond)
	}
	pager.OrderByExp = append(pager.OrderByExp, "createTime")
	err := app.QueryPage(&pager)
	if err != nil {
		log.Debug("get templates failed: %v", err)
		resp.Code = -1
		resp.Message = "get templates failed"
		ctx.Json(resp)
		return
	}
	resp.Items = pager.Data
	resp.BasePager = pager.BasePager
	resp.BasePager.More = pager.TotalPage * pager.PageSize
	resp.Code = tool.RespCodeSuccess
	ctx.Json(resp)
}

func GetMailTemplateInfo(ctx server.Context) {
	var resp InfoWrapper
	id, err := ctx.C.Params().GetInt("id")
	if err != nil || id <= 0 {
		resp.Code = tool.RespCodeNotFound
		resp.Message = "id must be a int value"
		ctx.Json(resp)
		return
	}

	tpl, err := repo.GetMessageTemplateById(id)
	if err != nil || tpl.Org != consoleorg.ID(ctx) || tpl.Category != msg.MailMessage {
		resp.Code = tool.RespCodeNotFound
		resp.Message = "not found"
		ctx.Json(resp)
		return
	}

	resp.Info = Info{
		Id:           tpl.Id,
		TemplateCode: tpl.Code,
		TemplateBody: tpl.Body,
		Status:       tpl.Status,
		CreateTime:   util.FormatTime(tpl.CreateTime),
		UpdateTime:   util.FormatTime(tpl.UpdateTime),
	}
	resp.Code = tool.RespCodeSuccess
	ctx.Json(resp)
}

type MailTemplateReq struct {
	Id          int    `json:"id,omitempty"`
	Code        string `json:"code,omitempty"`
	Name        string `json:"name,omitempty"`
	Body        string `json:"body,omitempty"`
	Description string `json:"description,omitempty"`
	Category    string `json:"category,omitempty"`
	Provider    string `json:"provider,omitempty"`
	Status      int    `json:"status,omitempty"`
}

func AddMailTemplate(ctx server.Context) {
	var req MailTemplateReq
	var resp InfoWrapper
	if err := ctx.C.ReadJSON(&req); err != nil {
		resp.Code = -1
		resp.Message = "req error"
		ctx.Json(resp)
		return
	}
	if len(req.Name) == 0 {
		resp.Code = -1
		resp.Message = "name nil"
		ctx.Json(resp)
		return
	}
	if len(req.Body) == 0 {
		resp.Code = -1
		resp.Message = "body nil"
		ctx.Json(resp)
		return
	}
	if !msg.IsValidCategory(msg.MessageCategory(req.Category)) {
		req.Category = msg.MailMessage.String()
	}
	provider := msg.MessageProvider(req.Provider)
	if len(provider) == 0 {
		provider = msg.Smtp
	}

	code := "tl_mail_" + util.UUIDString()
	m, err := repo.CreateMessageTemplate(
		consoleorg.ID(ctx),
		code, req.Name, req.Body, "", "{}", req.Description,
		msg.MailMessage, provider,
	)
	if err != nil {
		resp.Code = -1
		resp.Message = "create template failed"
		ctx.Json(resp)
		return
	}

	resp.Info = Info{
		Id:           m.Id,
		TemplateCode: m.Code,
		TemplateBody: m.Body,
		Status:       m.Status,
		CreateTime:   util.FormatTime(m.CreateTime),
		UpdateTime:   util.FormatTime(m.UpdateTime),
	}
	resp.Code = tool.RespCodeSuccess
	ctx.Json(resp)
}

func UpdateMailTemplate(ctx server.Context) {
	var req MailTemplateReq
	var resp InfoWrapper
	if err := ctx.C.ReadJSON(&req); err != nil {
		resp.Code = -1
		resp.Message = "req error"
		ctx.Json(resp)
		return
	}
	if id, err := ctx.C.Params().GetInt("id"); err == nil && id > 0 {
		req.Id = id
	}
	if req.Id <= 0 && len(req.Code) == 0 {
		resp.Code = -1
		resp.Message = "id nil"
		ctx.Json(resp)
		return
	}
	if len(req.Body) == 0 {
		resp.Code = -1
		resp.Message = "body nil"
		ctx.Json(resp)
		return
	}

	code := req.Code
	if code == "" {
		tpl, err := repo.GetMessageTemplateById(req.Id)
		if err != nil || tpl.Org != consoleorg.ID(ctx) {
			resp.Code = tool.RespCodeNotFound
			resp.Message = "not found"
			ctx.Json(resp)
			return
		}
		code = tpl.Code
	}

	m, err := repo.UpdateMessageTemplate(req.Status, code, req.Name, req.Body, "{}", req.Description, req.Provider)
	if err != nil {
		resp.Code = -1
		resp.Message = "update template failed"
		ctx.Json(resp)
		return
	}

	resp.Info = Info{
		Id:           m.Id,
		TemplateCode: m.Code,
		TemplateBody: m.Body,
		Status:       m.Status,
		CreateTime:   util.FormatTime(m.CreateTime),
		UpdateTime:   util.FormatTime(m.UpdateTime),
	}
	resp.Code = tool.RespCodeSuccess
	ctx.Json(resp)
}

func DeleteMailTemplate(ctx server.Context) {
	var resp app.Response
	id, err := ctx.C.Params().GetInt("id")
	if err != nil || id <= 0 {
		resp.Code = tool.RespCodeNotFound
		resp.Message = "id must be a int value"
		ctx.Json(resp)
		return
	}

	tpl, err := repo.GetMessageTemplateById(id)
	if err != nil || tpl.Org != consoleorg.ID(ctx) || tpl.Category != msg.MailMessage {
		resp.Code = tool.RespCodeNotFound
		resp.Message = "not found"
		ctx.Json(resp)
		return
	}

	_, err = app.GetOrm().Context.QueryTable(new(model.MessageTemplate)).
		Filter("Id", id).
		Delete()
	if err != nil {
		resp.Code = -1
		resp.Message = "delete template failed"
		ctx.Json(resp)
		return
	}
	resp.Code = tool.RespCodeSuccess
	ctx.Json(resp)
}

type MailStatusReq struct {
	Status int `json:"status,omitempty"`
	Id     int `json:"id,omitempty"`
}

func ChangeMailTemplateStatus(ctx server.Context) {
	var req MailStatusReq
	var resp app.Response
	if err := ctx.C.ReadJSON(&req); err != nil {
		resp.Code = tool.RespCodeError
		ctx.Json(resp)
		return
	}
	if req.Id <= 0 {
		resp.Code = tool.RespCodeError
		resp.Message = "id nil"
		ctx.Json(resp)
		return
	}

	tpl, err := repo.GetMessageTemplateById(req.Id)
	if err != nil || tpl.Org != consoleorg.ID(ctx) {
		resp.Code = tool.RespCodeNotFound
		resp.Message = "template not found"
		ctx.Json(resp)
		return
	}

	_, err = repo.UpdateMessageTemplate(req.Status, tpl.Code, "", "", "", "", "")
	if err != nil {
		resp.Code = tool.RespCodeError
		resp.Message = err.Error()
		ctx.Json(resp)
		return
	}
	resp.Code = tool.RespCodeSuccess
	ctx.Json(resp)
}
