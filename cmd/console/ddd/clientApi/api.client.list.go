package clientApi

import (
	"github.com/lishimeng/app-starter"
	"github.com/lishimeng/app-starter/persistence"
	"github.com/lishimeng/app-starter/server"
	"github.com/lishimeng/app-starter/tool"
	"github.com/lishimeng/owl-messager/internal/db/model"
	"github.com/lishimeng/owl-messager/internal/db/repo"
	"time"
)

type respPager struct {
	app.PagerResponse
	app.BasePager
	Items []respClient `json:"items"`
}

func getClientByPage(ctx server.Context) {
	var resp respPager
	var org = ctx.C.URLParamIntDefault("org", 0)
	var pageNum = ctx.C.URLParamIntDefault("pageNum", 1)
	var pageSize = ctx.C.URLParamIntDefault("pageSize", 10)
	var pager app.SimplePager[model.OpenClient, respClient]
	pager.PageSize = pageSize
	pager.PageNum = pageNum
	pager.Transform = func(src model.OpenClient, dst *respClient) {
		dst.Id = src.Id
		dst.AppId = src.AppId
		dst.Org = src.Org
		dst.Name = src.Name
		dst.CreateTime = src.CreateTime.UTC().Format(time.RFC3339)
	}
	var tenant string
	err := app.GetOrm().Transaction(func(ctx persistence.TxContext) (e error) {
		t, e := repo.GetTenantById(ctx, org)
		if e != nil {
			return
		}
		tenant = t.Code
		return
	})
	if err != nil {
		resp.Code = tool.RespCodeNotFound
		resp.Message = "unknown tenant"
		ctx.Json(resp)
		return
	}
	pager.QueryBuilder = func(tx persistence.TxContext) persistence.Query {
		q := tx.Model(&model.OpenClient{})
		if org > repo.ConditionIgnore {
			q = q.Equal("tenant_code", tenant)
		}
		return q
	}
	pager.OrderExp = append(pager.OrderExp, "ctime")
	err = app.QueryPage(&pager)
	if err != nil {
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
