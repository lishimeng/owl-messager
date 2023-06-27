package sender

import (
	"github.com/kataras/iris/v12"
	"github.com/lishimeng/app-starter"
	"github.com/lishimeng/app-starter/tool"
	"github.com/lishimeng/go-log"
	"github.com/lishimeng/owl-messager/internal/api/common"
	"github.com/lishimeng/owl-messager/internal/db/model"
	"github.com/lishimeng/owl-messager/internal/db/repo"
	"github.com/lishimeng/owl-messager/internal/db/service"
)

type req struct {
	DefaultSender int    `json:"defaultSender,omitempty"`
	Vendor        string `json:"vendor,omitempty"`
	Config        string `json:"config,omitempty"`
	Code          string `json:"code,omitempty"`
	Category      string `json:"category,omitempty"`
}

func SetMailSenderInfo(ctx iris.Context) {
	var resp app.Response
	var req req
	err := ctx.ReadJSON(&req)
	if err != nil {
		resp.Code = tool.RespCodeNotFound
		tool.ResponseJSON(ctx, resp)
		return
	}
	log.Debug("req：%s", req)
	code := tool.GetRandomString(common.DefaultCodeLen)
	code = "sender_" + code
	switch req.Category {
	case model.SenderCategoryMail:
		_, err = service.CreateMsi(code, req.Vendor, req.Config, req.DefaultSender)
		if err != nil {
			resp.Code = tool.RespCodeNotFound
			tool.ResponseJSON(ctx, resp)
			return
		}
	case model.SenderCategorySms:
		_, err = service.CreateSsi(code, req.Vendor, req.Config, req.DefaultSender)
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

func UpMailSenderInfo(ctx iris.Context) {
	var resp app.Response
	var req req
	err := ctx.ReadJSON(&req)
	if err != nil {
		resp.Code = tool.RespCodeNotFound
		tool.ResponseJSON(ctx, resp)
		return
	}
	log.Debug("req：%s", req)
	switch req.Category {
	case model.SenderCategoryMail:
		_, err = service.UpdateMsi(req.Code, req.Vendor, req.Config, req.DefaultSender)
		if err != nil {
			resp.Code = tool.RespCodeNotFound
			tool.ResponseJSON(ctx, resp)
			return
		}
	case model.SenderCategorySms:
		_, err = service.UpdateSsi(req.Code, req.Vendor, req.Config, req.DefaultSender)
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

type RespInfo struct {
	app.Response
	DefaultSender int    `json:"defaultSender,omitempty"`
	Vendor        string `json:"vendor,omitempty"`
	Config        string `json:"config,omitempty"`
	Scode         string `json:"scode,omitempty"`
}

func GetMailSenderInfo(ctx iris.Context) {
	var resp RespInfo
	var vendor = ctx.URLParamDefault("vendor", "")
	log.Debug("vendor:%s", vendor)
	m := model.MailSenderInfo{}
	err := app.GetOrm().Context.QueryTable(new(model.MailSenderInfo)).Filter("Vendor", vendor).One(&m)
	if err != nil {
		resp.Code = tool.RespCodeNotFound
		tool.ResponseJSON(ctx, resp)
		return
	}
	resp.Scode = m.Code
	resp.DefaultSender = m.Default
	resp.Vendor = string(m.Vendor)
	resp.Config = m.Config
	resp.Code = tool.RespCodeSuccess
	log.Debug("resp:%s", resp)
	tool.ResponseJSON(ctx, resp)
}

func ListByPage(ctx iris.Context) {
	var resp app.PagerResponse
	var pageNum = ctx.URLParamIntDefault("pageNum", 1)
	var pageSize = ctx.URLParamIntDefault("pageSize", 10)
	var category = ctx.URLParamDefault("category", "")
	page := app.Pager{
		PageSize: pageSize,
		PageNum:  pageNum,
	}
	log.Debug("category:%s", category)
	switch category {
	case model.SenderCategoryMail:
		page, list, err := repo.GetMailSenderList(1, page)
		if err != nil {
			resp.Code = tool.RespCodeNotFound
			tool.ResponseJSON(ctx, resp)
			return
		}
		if len(list) > 0 {
			for _, ms := range list {
				page.Data = append(page.Data, ms)
			}
		}
		resp.Pager = page
	case model.SenderCategorySms:
		page, list, err := repo.GetSmsSenderList(1, page)
		if err != nil {
			resp.Code = tool.RespCodeNotFound
			tool.ResponseJSON(ctx, resp)
			return
		}
		if len(list) > 0 {
			for _, ms := range list {
				page.Data = append(page.Data, ms)
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

type AllSenderInfo struct {
	app.Response
	Data interface{} `json:"item,omitempty"`
}

func GetSenderInfoByCategory(ctx iris.Context) {
	var resp AllSenderInfo
	var category = ctx.URLParamDefault("category", "")
	var code = ctx.URLParamDefault("code", "")
	switch category {
	case model.SenderCategoryMail:
		info, err := repo.GetMailSenderByCode(code)
		if err != nil {
			resp.Code = tool.RespCodeNotFound
			tool.ResponseJSON(ctx, resp)
			return
		}
		resp.Data = info
	case model.SenderCategorySms:
		info, err := repo.GetSmsSenderByCode(code)
		if err != nil {
			resp.Code = tool.RespCodeNotFound
			tool.ResponseJSON(ctx, resp)
			return
		}
		resp.Data = info
	default:
		resp.Code = tool.RespCodeNotFound
		tool.ResponseJSON(ctx, resp)
		return
	}
	resp.Code = tool.RespCodeSuccess
	tool.ResponseJSON(ctx, resp)
}
