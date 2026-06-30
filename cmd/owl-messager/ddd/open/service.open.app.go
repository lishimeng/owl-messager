package open

import (
	"github.com/lishimeng/app-starter"
	"github.com/lishimeng/app-starter/persistence"
	"github.com/lishimeng/owl-messager/internal/db/model"
)

func getAppInfo(appId string) (ai AppInfo, err error) {
	err = app.GetOrm().Transaction(func(ctx persistence.TxContext) (e error) {
		c, e := _getClientByAppId(appId, ctx)
		if e != nil {
			return
		}
		t, e := _getTenantById(c.Org, ctx)
		if e != nil {
			return
		}
		ai.AppId = c.AppId
		ai.Secret = c.Secret
		ai.Org = t.Code
		return
	})
	return
}

func _getClientByAppId(appId string, ctx persistence.TxContext) (c model.OpenClient, err error) {
	err = ctx.Model(&model.OpenClient{}).Equal("app_id", appId).First(&c)
	return
}

func _getTenantById(id int, ctx persistence.TxContext) (t model.Tenant, err error) {
	err = ctx.Model(&model.Tenant{}).Equal("id", id).First(&t)
	return
}
