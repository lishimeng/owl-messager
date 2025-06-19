package repo

import (
	"crypto/sha256"
	"fmt"
	"github.com/lishimeng/app-starter"
	"github.com/lishimeng/app-starter/persistence"
	"github.com/lishimeng/owl-messager/internal/db/model"
	"github.com/lishimeng/x/util"
	"strings"
	"time"
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

func genAppId(tenant string) (code string) {
	now := time.Now().Format(time.RFC3339Nano)
	var tmp = fmt.Sprintf("AppId_%s_%s", now, tenant)
	sh := sha256.New()
	sh.Write([]byte(tmp))
	bs := sh.Sum(nil)
	code = strings.ToLower(util.BytesToHex(bs))
	return
}

func genSecret(appId string) (code string) {
	now := time.Now().Format(time.RFC3339Nano)
	var tmp = fmt.Sprintf("Secret_%s_%s", now, appId)
	sh := sha256.New()
	sh.Write([]byte(tmp))
	bs := sh.Sum(nil)
	code = strings.ToLower(util.BytesToHex(bs))
	return
}

func AddClient(ctx persistence.TxContext, tenant string, org int, name string) (client model.OpenClient, err error) {
	appId := genAppId(tenant)
	secret := genSecret(appId)
	client = model.OpenClient{
		AppId:  appId,
		Secret: secret,
		Domain: tenant,
		Name:   name,
	}
	client.Org = org
	_, err = ctx.Context.Insert(&client)
	return
}

func DeleteClient(ctx persistence.TxContext, tenant string, appId string) (err error) {
	_, err = ctx.Context.QueryTable(new(model.OpenClient)).
		Filter("Domain", tenant).
		Filter("AppId", appId).
		Limit(1).
		Delete()
	return
}
