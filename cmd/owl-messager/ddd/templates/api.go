package templates

import (
	"github.com/beego/beego/v2/client/orm"
	"github.com/lishimeng/app-starter"
	"github.com/lishimeng/app-starter/midware/auth"
	"github.com/lishimeng/app-starter/persistence"
	"github.com/lishimeng/app-starter/server"
	"github.com/lishimeng/app-starter/tool"
	"github.com/lishimeng/go-log"
	"github.com/lishimeng/owl-messager/internal/db/model"
	"github.com/lishimeng/owl-messager/internal/db/repo"
	"github.com/lishimeng/owl-messager/pkg"
	"github.com/lishimeng/owl-messager/pkg/msg"
	"strings"
)

type resp struct {
	app.PagerResponse
	Data []pkg.TemplateInfo
}

func templates(ctx server.Context) {

	var err error
	var resp resp
	var tpls []pkg.TemplateInfo
	var org = ctx.C.GetHeader(auth.OrgKey)
	var pageNo = ctx.C.URLParamIntDefault("pageNo", 1)      // ?
	var pageSize = ctx.C.URLParamIntDefault("pageSize", 10) // ?
	var category = ctx.C.Params().GetStringDefault("category", "")
	category = strings.TrimSpace(category)

	valid := msg.IsValidCategory(msg.MessageCategory(category))
	if !valid { // 不支持的类型,数据列表为空
		resp.Code = tool.RespCodeSuccess
		ctx.Json(resp)
		return
	}
	tenant, err := repo.GetTenant(org)
	if err != nil {
		log.Debug("unknown tenant: %s", org)
		resp.Code = -1
		resp.Message = "unknown tenant"
		ctx.Json(resp)
		return
	}
	tpls, err = getTemplates(msg.MessageCategory(category), tenant.Id, pageNo, pageSize)

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

func getTemplates(category msg.MessageCategory, org int, pageNo, pageSize int) (tpls []pkg.TemplateInfo, err error) {

	var pager app.SimplePager[model.MessageTemplate, pkg.TemplateInfo]
	pager.PageSize = pageSize
	pager.PageNum = pageNo
	pager.Transform = func(src model.MessageTemplate, dst *pkg.TemplateInfo) {
		dst.Id = src.Id
		dst.Name = src.Name
		dst.Body = src.Body
		dst.CloudTemplate = src.CloudTemplate
		dst.Description = src.Description
		dst.Category = src.Category.String()
		dst.Params = src.Params
		dst.Provider = src.Provider.String()
		dst.Code = src.Code
	}
	pager.QueryBuilder = func(tx persistence.TxContext) any {
		cond := orm.NewCondition()
		cond = cond.And("org", org)
		if len(category) > 0 {
			cond = cond.And("message_category", category)
		}
		return tx.Context.QueryTable(new(model.MessageTemplate)).SetCond(cond)
	}
	pager.OrderByExp = append(pager.OrderByExp, "createTime")
	err = app.QueryPage(&pager)
	return
}
