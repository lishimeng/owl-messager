package senderApi

import (
	"encoding/json"

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
	Id         int    `json:"id,omitempty"`
	SenderCode string `json:"senderCode,omitempty"`
	Config     string `json:"config,omitempty"`
	Status     int    `json:"itemStatus,omitempty"`
	CreateTime string `json:"createTime,omitempty"`
	UpdateTime string `json:"updateTime,omitempty"`
}

type InfoWrapper struct {
	app.Response
	Info
}

type respList struct {
	app.PagerResponse
	app.BasePager
	Items []Info `json:"items"`
}

func GetMailSenderList(ctx server.Context) {
	var resp respList
	orgID := consoleorg.Code(ctx)
	pageSize := ctx.C.URLParamIntDefault("pageSize", repo.DefaultPageSize)
	pageNo := ctx.C.URLParamIntDefault("pageNo", repo.DefaultPageNo)

	var pager app.SimplePager[model.MessageSenderInfo, Info]
	pager.PageSize = pageSize
	pager.PageNum = pageNo
	pager.Transform = func(src model.MessageSenderInfo, dst *Info) {
		dst.Id = src.Id
		dst.SenderCode = src.Code
		dst.Config = string(src.Config)
		dst.Status = src.Status
		dst.CreateTime = util.FormatTime(src.CreateTime)
		dst.UpdateTime = util.FormatTime(src.UpdateTime)
	}
	pager.QueryBuilder = func(tx persistence.TxContext) persistence.Query {
		return tx.Model(&model.MessageSenderInfo{}).
			Equal("tenant_code", orgID).
			Equal("message_category", msg.MailMessage)
	}
	pager.OrderExp = append(pager.OrderExp, "ctime")
	if err := app.QueryPage(&pager); err != nil {
		log.Debug("get senders failed: %v", err)
		resp.Code = -1
		resp.Message = "get senders failed"
		ctx.Json(resp)
		return
	}
	resp.Items = pager.Data
	resp.BasePager = pager.BasePager
	resp.BasePager.More = pager.TotalPage * pager.PageSize
	resp.Code = tool.RespCodeSuccess
	ctx.Json(resp)
}

func GetMailSenderInfo(ctx server.Context) {
	var resp InfoWrapper
	id, err := ctx.C.Params().GetInt("id")
	if err != nil || id <= 0 {
		resp.Code = tool.RespCodeNotFound
		resp.Message = "id must be a int value"
		ctx.Json(resp)
		return
	}

	var ms model.MessageSenderInfo
	err = app.GetOrm().Model(&model.MessageSenderInfo{}).
		Equal("id", id).
		Equal("tenant_code", consoleorg.Code(ctx)).
		Equal("message_category", msg.MailMessage).
		First(&ms)
	if err != nil {
		resp.Code = tool.RespCodeNotFound
		resp.Message = "not found"
		ctx.Json(resp)
		return
	}

	resp.Info = Info{
		Id:         ms.Id,
		SenderCode: ms.Code,
		Config:     string(ms.Config),
		Status:     ms.Status,
		CreateTime: util.FormatTime(ms.CreateTime),
		UpdateTime: util.FormatTime(ms.UpdateTime),
	}
	resp.Code = tool.RespCodeSuccess
	ctx.Json(resp)
}

type addMailSenderReq struct {
	Vendor  msg.MessageProvider `json:"vendor,omitempty"`
	Config  json.RawMessage     `json:"config,omitempty"`
	Default int                 `json:"default,omitempty"`
}

func AddMailSender(ctx server.Context) {
	var req addMailSenderReq
	var resp InfoWrapper
	if err := ctx.C.ReadJSON(&req); err != nil {
		resp.Code = -1
		resp.Message = "req error"
		ctx.Json(resp)
		return
	}
	if len(req.Vendor) == 0 {
		resp.Code = -1
		resp.Message = "vendor nil"
		ctx.Json(resp)
		return
	}

	code := "sender_mail_" + util.UUIDString()
	orgID := consoleorg.Code(ctx)
	isDefault := 0
	if req.Default == 1 {
		isDefault = 1
	}

	var cfg msg.SenderConfig
	if len(req.Config) > 0 {
		cfg = msg.SenderConfig(req.Config)
	}
	ms, err := repo.CreateMessageSender(orgID, msg.MailMessage, req.Vendor, isDefault, code, cfg)
	if err != nil {
		resp.Code = -1
		resp.Message = "create sender failed"
		ctx.Json(resp)
		return
	}

	resp.Info = Info{
		Id:         ms.Id,
		SenderCode: ms.Code,
		Config:     string(ms.Config),
		Status:     ms.Status,
		CreateTime: util.FormatTime(ms.CreateTime),
		UpdateTime: util.FormatTime(ms.UpdateTime),
	}
	resp.Code = tool.RespCodeSuccess
	ctx.Json(resp)
}

func UpdateMailSender(ctx server.Context) {
	var req addMailSenderReq
	var resp app.Response
	id, err := ctx.C.Params().GetInt("id")
	if err != nil || id <= 0 {
		resp.Code = tool.RespCodeNotFound
		resp.Message = "id must be a int value"
		ctx.Json(resp)
		return
	}
	if err := ctx.C.ReadJSON(&req); err != nil {
		resp.Code = -1
		resp.Message = "req error"
		ctx.Json(resp)
		return
	}

	var ms model.MessageSenderInfo
	err = app.GetOrm().Model(&model.MessageSenderInfo{}).
		Equal("id", id).
		Equal("tenant_code", consoleorg.Code(ctx)).
		Equal("message_category", msg.MailMessage).
		First(&ms)
	if err != nil {
		resp.Code = tool.RespCodeNotFound
		resp.Message = "not found"
		ctx.Json(resp)
		return
	}

	if len(req.Config) > 0 {
		_, err = repo.UpdateMessageSender(ms.Code, msg.SenderConfig(req.Config))
		if err != nil {
			resp.Code = -1
			resp.Message = "update sender failed"
			ctx.Json(resp)
			return
		}
	}
	resp.Code = tool.RespCodeSuccess
	ctx.Json(resp)
}

func DeleteMailSender(ctx server.Context) {
	var resp app.Response
	id, err := ctx.C.Params().GetInt("id")
	if err != nil || id <= 0 {
		resp.Code = tool.RespCodeNotFound
		resp.Message = "id must be a int value"
		ctx.Json(resp)
		return
	}

	err = app.Transaction(func(tx persistence.TxContext) error {
		return tx.Delete(&model.MessageSenderInfo{}, id)
	})
	if err != nil {
		resp.Code = -1
		resp.Message = "delete sender failed"
		ctx.Json(resp)
		return
	}
	resp.Code = tool.RespCodeSuccess
	ctx.Json(resp)
}
