package tenant

import (
	"github.com/lishimeng/app-starter"
	"github.com/lishimeng/owl-messager/internal/db/model"
)

func getTenant(code string) (t model.Tenant, err error) {
	err = app.GetOrm().Model(&model.Tenant{}).Equal("code", code).First(&t)
	return
}
