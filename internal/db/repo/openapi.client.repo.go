package repo

import (
	persistence "github.com/lishimeng/go-orm"
	"github.com/lishimeng/owl-messager/internal/db/model"
)

func GetClientById(ctx persistence.OrmContext, id int) (c model.OpenClient, err error) {
	return
}

func GetClientByAppKey(ctx persistence.OrmContext, key string) (c model.OpenClient, err error) {
	return
}

func GetClients(ctx persistence.OrmContext, key string) (c []model.OpenClient, err error) {
	return
}

