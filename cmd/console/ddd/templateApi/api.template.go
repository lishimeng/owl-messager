package templateApi

import (
	"github.com/lishimeng/app-starter"
	"github.com/lishimeng/app-starter/persistence"
	"github.com/lishimeng/app-starter/server"
	"github.com/lishimeng/app-starter/tool"
	"github.com/lishimeng/go-log"
	"github.com/lishimeng/owl-messager/cmd/console/ddd/consoleorg"
	"github.com/lishimeng/owl-messager/internal/db/model"
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
	var provider = ctx.C.URLParamDefault("provider", "")
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
	pager.QueryBuilder = func(tx persistence.TxContext) persistence.Query {
		q := tx.Model(&model.MessageTemplate{}).Equal("tenant_code", consoleorg.Code(ctx))
		if len(category) > 0 {
			q = q.Equal("message_category", category)
		}
		if len(provider) > 0 {
			q = q.Equal("message_provider", provider)
		}
		return q
	}
	pager.OrderExp = append(pager.OrderExp, "ctime")
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
