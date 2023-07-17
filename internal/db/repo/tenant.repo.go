package repo

import (
	"github.com/lishimeng/app-starter"
	"github.com/lishimeng/owl-messager/internal/db/model"
)

func GetTenant(code string) (t model.Tenant, err error) {
	err = app.GetOrm().Context.
		QueryTable(new(model.Tenant)).
		Filter("Code", code).
		Filter("Status", 1).One(&t)
	return
}
