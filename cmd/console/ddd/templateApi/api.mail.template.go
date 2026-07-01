package templateApi

import (
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

// GetMailVendors 平台支持的 mail provider（来自 pkg/msg 目录）
func GetMailVendors(ctx server.Context) {
	var resp SmsVendors

	resp.Code = tool.RespCodeSuccess
	resp.Message = "Mail Vendors"
	resp.Data = msg.ListProviders(msg.MailMessage)
	ctx.Json(resp)
}

func GetMailTemplateList(ctx server.Context) {
	var resp respInfo
	var pageSize = ctx.C.URLParamIntDefault("pageSize", repo.DefaultPageSize)
	var pageNo = ctx.C.URLParamIntDefault("pageNo", repo.DefaultPageNo)
	var provider = ctx.C.URLParamDefault("provider", "")
	orgID := consoleorg.TenantFilter(ctx)

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
	pager.QueryBuilder = func(tx persistence.TxContext) persistence.Query {
		q := tx.Model(&model.MessageTemplate{}).Equal("category", msg.MailMessage)
		if orgID != "" {
			q = q.Equal("tenant_code", orgID)
		}
		if len(provider) > 0 {
			q = q.Equal("message_provider", provider)
		}
		return q
	}
	pager.OrderExp = append(pager.OrderExp, "ctime")
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
	if err != nil || tpl.Category != msg.MailMessage {
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
	TenantCode  string `json:"tenantCode,omitempty"`
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
	if req.TenantCode == "" {
		resp.Code = -1
		resp.Message = "tenantCode required"
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
		req.TenantCode,
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
		if err != nil {
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
	if err != nil || tpl.Category != msg.MailMessage {
		resp.Code = tool.RespCodeNotFound
		resp.Message = "not found"
		ctx.Json(resp)
		return
	}

	err = app.Transaction(func(tx persistence.TxContext) error {
		return tx.Delete(&model.MessageTemplate{}, id)
	})
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
	if err != nil {
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
