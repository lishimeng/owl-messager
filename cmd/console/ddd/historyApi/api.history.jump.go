package historyApi

import (
	"github.com/lishimeng/app-starter"
	"github.com/lishimeng/app-starter/server"
	"github.com/lishimeng/app-starter/tool"
	"github.com/lishimeng/owl-messager/internal/db/model"
	"time"
)

type respCount struct {
	app.Response
	Count int `json:"count"`
}

func GetHistoryCount(ctx server.Context) {
	var resp respCount
	var category = ctx.C.URLParamDefault("category", "")
	var timeStr = ctx.C.URLParamDefault("time", "")
	timeLatest, err := time.Parse(time.RFC3339, timeStr)
	if err != nil {
		resp.Code = tool.RespCodeNotFound
		resp.Message = "illegal time"
		ctx.Json(resp)
		return
	}
	q := app.GetOrm().Context.QueryTable(new(model.MessageInfo))
	if category != "" {
		q = q.Filter("category", category)
	}
	count, err := q.
		Filter("ctime__gt", timeLatest.Add(time.Hour*24)).
		Count()
	if err != nil {
		resp.Code = tool.RespCodeNotFound
		resp.Message = err.Error()
		ctx.Json(resp)
		return
	}
	resp.Code = tool.RespCodeSuccess
	resp.Count = int(count)
	ctx.Json(resp)
}
