package templates

import (
	"strings"

	app "github.com/lishimeng/app-starter"
	"github.com/lishimeng/app-starter/persistence"
	"github.com/lishimeng/app-starter/server"
	"github.com/lishimeng/app-starter/tool"
	"github.com/lishimeng/go-log"
	"github.com/lishimeng/owl-messager/cmd/owl-messager/midware"
	"github.com/lishimeng/owl-messager/internal/db/model"
	"github.com/lishimeng/owl-messager/pkg"
	"github.com/lishimeng/owl-messager/pkg/msg"
)

type resp struct {
	app.PagerResponse
	Data []pkg.TemplateInfo `json:"items,omitempty"`
}

func templates(ctx server.Context) {

	var err error
	var resp resp
	var tpls []pkg.TemplateInfo
	var org string
	org, err = midware.TenantCodeFromContext(ctx)
	if err != nil || org == "" {
		log.Debug("unknown tenant: %s", org)
		resp.Code = -1
		resp.Message = "unknown tenant"
		ctx.Json(resp)
		return
	}
	var pageNo = ctx.C.URLParamIntDefault("pageNo", 1)
	var pageSize = ctx.C.URLParamIntDefault("pageSize", 10)
	var category = ctx.C.Params().GetStringDefault("category", "")
	category = strings.TrimSpace(category)

	valid := msg.IsValidCategory(msg.MessageCategory(category))
	if !valid {
		resp.Code = tool.RespCodeSuccess
		ctx.Json(resp)
		return
	}
	tpls, err = getTemplates(msg.MessageCategory(category), org, pageNo, pageSize)

	if err != nil {
		resp.Code = tool.RespCodeSuccess
		resp.Message = err.Error()
		ctx.Json(resp)
		return
	} else {
		for _, tpl := range tpls {
			resp.Data = append(resp.Data, tpl)
		}
	}
	resp.Code = tool.RespCodeSuccess
	resp.Message = "OK"
	ctx.Json(resp)
}

func getTemplates(category msg.MessageCategory, tenantCode string, pageNo, pageSize int) (tpls []pkg.TemplateInfo, err error) {

	var pager app.SimplePager[model.MessageTemplate, pkg.TemplateInfo]
	pager.PageSize = pageSize
	pager.PageNum = pageNo
	pager.Transform = func(src model.MessageTemplate, dst *pkg.TemplateInfo) {
		*dst = pkg.TemplateInfoFromModel(src)
	}
	pager.QueryBuilder = func(tx persistence.TxContext) persistence.Query {
		q := tx.Model(&model.MessageTemplate{}).Equal("tenant_code", tenantCode)
		if len(category) > 0 {
			q = q.Equal("category", category)
		}
		return q
	}
	pager.OrderExp = append(pager.OrderExp, "ctime")
	err = app.QueryPage(&pager)
	if err != nil {
		return
	}
	tpls = pager.Data
	return
}
