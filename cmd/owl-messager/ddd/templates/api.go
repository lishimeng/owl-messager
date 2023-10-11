package templates

import (
	"github.com/kataras/iris/v12"
	"github.com/lishimeng/app-starter"
	"github.com/lishimeng/app-starter/midware/auth"
	"github.com/lishimeng/app-starter/tool"
	"github.com/lishimeng/go-log"
	"github.com/lishimeng/owl-messager/internal/db/repo"
	"github.com/lishimeng/owl-messager/pkg"
	"github.com/lishimeng/owl-messager/pkg/msg"
	"strings"
)

func templates(ctx iris.Context) {

	var err error
	var resp app.PagerResponse
	var tpls []pkg.TemplateInfo
	var org = ctx.GetHeader(auth.OrgKey)
	var pageNo = ctx.URLParamIntDefault("pageNo", 1)      // ?
	var pageSize = ctx.URLParamIntDefault("pageSize", 10) // ?
	var category = ctx.Params().GetStringDefault("category", "")
	category = strings.TrimSpace(category)

	valid := msg.IsValidCategory(msg.MessageCategory(category))
	if !valid { // 不支持的类型,数据列表为空
		resp.Code = iris.StatusOK
		tool.ResponseJSON(ctx, resp)
		return
	}
	tenant, err := repo.GetTenant(org)
	if err != nil {
		log.Debug("unknown tenant: %s", org)
		resp.Code = -1
		resp.Message = "unknown tenant"
		tool.ResponseJSON(ctx, resp)
		return
	}
	tpls, err = getTemplates(msg.MessageCategory(category), tenant.Id, pageNo, pageSize)

	if err != nil {
		resp.Code = iris.StatusOK
		resp.Message = err.Error()
		tool.ResponseJSON(ctx, resp)
		return
	} else {
		for _, tpl := range tpls {
			resp.Data = append(resp.Data, tpl)
		}
	}
	resp.Code = iris.StatusOK
	resp.Message = "OK"
	tool.ResponseJSON(ctx, resp)
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
