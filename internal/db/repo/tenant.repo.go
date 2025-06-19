package repo

import (
	"github.com/lishimeng/app-starter"
	"github.com/lishimeng/app-starter/persistence"
	"github.com/lishimeng/owl-messager/internal/db/model"
)

func GetTenant(code string) (t model.Tenant, err error) {
	err = app.GetOrm().Context.
		QueryTable(new(model.Tenant)).
		Filter("Code", code).
		Filter("Status", 1).One(&t)
	return
}

func GetTenantById(ctx persistence.TxContext, id int) (t model.Tenant, err error) {
	err = ctx.Context.
		QueryTable(new(model.Tenant)).
		Filter("Id", id).
		Filter("Status", 1).One(&t)
	return
}
