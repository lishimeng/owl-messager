package repo

import (
	"github.com/lishimeng/app-starter/persistence"
	"github.com/lishimeng/owl-messager/internal/db/model"
)

func GetTenant(code string) (t model.Tenant, err error) {
	err = orm().Model(&model.Tenant{}).
		Equal("code", code).
		Equal("status", 1).
		First(&t)
	return
}

func GetTenantById(ctx persistence.TxContext, id int) (t model.Tenant, err error) {
	err = ctx.Model(&model.Tenant{}).
		Equal("id", id).
		Equal("status", 1).
		First(&t)
	return
}
