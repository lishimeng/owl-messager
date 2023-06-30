package open

import (
	"github.com/lishimeng/app-starter"
	"github.com/lishimeng/owl-messager/internal/db/model"
)

func getClientByAppId(appId string) (c model.OpenClient, err error) {
	err = app.GetOrm().Context.
		QueryTable(new(model.OpenClient)).
		Filter("AppId", appId).
		One(&c)
	return
}
