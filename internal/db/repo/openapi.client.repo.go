package repo

import (
	"github.com/lishimeng/app-starter"
	persistence "github.com/lishimeng/go-orm"
	"github.com/lishimeng/owl-messager/internal/db/model"
)

func GetClientById(ctx persistence.OrmContext, id int) (c model.OpenClient, err error) {
	return
}

func GetClientByAppId(appId string) (c model.OpenClient, err error) {
	err = app.GetOrm().Context.
		QueryTable(new(model.OpenClient)).
		Filter("AppId", appId).
		One(&c)
	return
}

func GetClients(ctx persistence.OrmContext, key string) (c []model.OpenClient, err error) {
	return
}
