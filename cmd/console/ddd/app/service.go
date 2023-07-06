package tenant

import (
	"github.com/lishimeng/app-starter"
	"github.com/lishimeng/owl-messager/internal/db/model"
)

func getTenant(code string) (t model.Tenant, err error) {
	err = app.GetOrm().Context.
		QueryTable(new(model.Tenant)).
		Filter("Code", code).One(&t)
	return
}
