package templates

import (
	"github.com/lishimeng/app-starter"
	"github.com/lishimeng/owl-messager/internal/db/model"
	"github.com/lishimeng/owl-messager/internal/db/repo"
)

func mailTemplates(org, status int, page app.Pager) (tpls []model.MessageTemplate, err error) {
	var qs = app.GetOrm().Context.QueryTable(new(model.MessageTemplate))
	if status > repo.ConditionIgnore {
		qs = qs.Filter("Status", status)
	}
	if err != nil {
		return
	}
	_, err = qs.Filter("Org", org).
		OrderBy("CreateTime").
		Offset(repo.CalcPageOffset(page)).
		Limit(page.PageSize).
		All(&tpls)
	if err != nil {
		return
	}
	return
}
