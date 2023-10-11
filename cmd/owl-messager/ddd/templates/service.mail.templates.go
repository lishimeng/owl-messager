package templates

import (
	"github.com/lishimeng/app-starter"
	"github.com/lishimeng/owl-messager/internal/db/model"
	"github.com/lishimeng/owl-messager/internal/db/repo"
	"github.com/lishimeng/owl-messager/pkg/msg"
)

func getList(org, status int, category msg.MessageCategory, page app.Pager) (tpls []model.MessageTemplate, err error) {
	var qs = app.GetOrm().Context.QueryTable(new(model.MessageTemplate))
	if status > repo.ConditionIgnore {
		qs = qs.Filter("Status", status)
	}
	if err != nil {
		return
	}
	_, err = qs.Filter("Org", org).Filter("Category", category).
		OrderBy("CreateTime").
		Offset(repo.CalcPageOffset(page)).
		Limit(page.PageSize).
		All(&tpls)
	if err != nil {
		return
	}
	return
}
