package templates

import (
	"github.com/lishimeng/app-starter"
	"github.com/lishimeng/app-starter/midware/auth"
	"github.com/lishimeng/app-starter/server"
	"github.com/lishimeng/app-starter/tool"
	"github.com/lishimeng/go-log"
	"github.com/lishimeng/owl-messager/internal/db/repo"
	"github.com/lishimeng/owl-messager/pkg"
	"github.com/lishimeng/owl-messager/pkg/msg"
	"strings"
)

func templates(ctx server.Context) {

	var err error
	var resp app.PagerResponse
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
	data, err := getList(org, 0, category, app.Pager{PageSize: pageSize, PageNum: pageNo})
	if err != nil {
		return
	}
	for _, d := range data {
		tpls = append(tpls, pkg.TemplateInfo{
			Id:            d.Id,
			Code:          d.Code,
			Name:          d.Name,
			Category:      d.Category.String(),
			Body:          d.Body,
			Params:        d.Params,
			Provider:      d.Provider.String(),
			CloudTemplate: d.CloudTemplate,
			Description:   d.Description,
		})
	}
	return
}
