package repo

import (
	"crypto/sha256"
	"fmt"
	"strings"
	"time"

	"github.com/lishimeng/app-starter/persistence"
	"github.com/lishimeng/owl-messager/internal/db/model"
	"github.com/lishimeng/x/util"
)

func GetClientById(ctx persistence.OrmContext, id int) (c model.OpenClient, err error) {
	err = ctx.Model(&model.OpenClient{}).Equal("id", id).First(&c)
	return
}

func GetClientByAppId(appId string) (c model.OpenClient, err error) {
	err = orm().Model(&model.OpenClient{}).Equal("app_id", appId).First(&c)
	return
}

func GetClientByBasicAuth(token string) (c model.OpenClient, err error) {
	err = orm().Model(&model.OpenClient{}).Equal("basic_auth", token).First(&c)
	return
}

func GetClients(ctx persistence.OrmContext, key string) (c []model.OpenClient, err error) {
	q := ctx.Model(&model.OpenClient{})
	if len(key) > 0 {
		q = q.Equal("tenant_code", key)
	}
	err = q.Find(&c)
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

func genBasicAuth(tenant, appId string) (code string) {
	now := time.Now().Format(time.RFC3339Nano)
	var tmp = fmt.Sprintf("BasicAuth_%s_%s_%s", now, tenant, appId)
	sh := sha256.New()
	sh.Write([]byte(tmp))
	bs := sh.Sum(nil)
	code = strings.ToLower(util.BytesToHex(bs))
	return
}

func AddClient(ctx persistence.TxContext, tenant string, org int, name string) (client model.OpenClient, err error) {
	appId := genAppId(tenant)
	secret := genSecret(appId)
	basicAuth := genBasicAuth(tenant, appId)
	client = model.OpenClient{
		AppId:      appId,
		Secret:     secret,
		BasicAuth:  basicAuth,
		TenantCode: tenant,
		Name:       name,
	}
	client.Org = org
	err = ctx.Create(&client)
	return
}

func DeleteClient(ctx persistence.TxContext, tenant string, appId string) (err error) {
	var client model.OpenClient
	err = ctx.Model(&model.OpenClient{}).
		Equal("tenant_code", tenant).
		Equal("app_id", appId).
		First(&client)
	if err != nil {
		return
	}
	return ctx.Delete(&client)
}
