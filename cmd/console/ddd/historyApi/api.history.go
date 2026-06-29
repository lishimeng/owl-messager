package historyApi

import (
	"github.com/beego/beego/v2/client/orm"
	"github.com/lishimeng/app-starter"
	"github.com/lishimeng/app-starter/persistence"
	"github.com/lishimeng/app-starter/server"
	"github.com/lishimeng/app-starter/tool"
	"github.com/lishimeng/owl-messager/cmd/console/ddd/consoleorg"
	"github.com/lishimeng/owl-messager/internal/db/model"
	"time"
)

type respHistoryMessage struct {
	Id         int    `json:"id,omitempty"`
	Subject    string `json:"subject,omitempty"`
	Status     int    `json:"status,omitempty"`
	CreateTime string `json:"createTime,omitempty"`
	Category   string `json:"category,omitempty"`
}

type respPager struct {
	app.PagerResponse
	app.BasePager
	Items []respHistoryMessage `json:"items"`
}

func GetHistoryList(ctx server.Context) {
	//	app.GetOrm().Context.Raw(`
	//SELECT mi.status, imi.template_id
	//FROM message_info mi
	//INNER JOIN im_message_info imi ON mi.id = imi.message_id AND mi.category = 'im'
	//INNER JOIN mail_message_info mmi ON mi.id = mmi.message_id AND mi.category = 'mail'
	//`)
	var resp respPager
	var category = ctx.C.URLParamDefault("category", "")
	var pageNum = ctx.C.URLParamIntDefault("pageNum", 1)
	var pageSize = ctx.C.URLParamIntDefault("pageSize", 10)
	var pager app.SimplePager[model.MessageInfo, respHistoryMessage]
	pager.PageSize = pageSize
	pager.PageNum = pageNum
	pager.Transform = func(src model.MessageInfo, dst *respHistoryMessage) {
		dst.Id = src.Id
		dst.Subject = src.Subject
		dst.Status = src.Status
		dst.Category = src.Category.String()
		dst.CreateTime = src.CreateTime.UTC().Format(time.RFC3339)
	}
	pager.QueryBuilder = func(tx persistence.TxContext) any {
		cond := orm.NewCondition()
		cond = cond.And("org", consoleorg.ID(ctx))
		if len(category) > 0 {
			cond = cond.And("category", category)
		}
		return tx.Context.QueryTable(new(model.MessageInfo)).SetCond(cond)
	}
	pager.OrderByExp = append(pager.OrderByExp, "-createTime")
	err := app.QueryPage(&pager)
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
