package tenant

import (
	"crypto/sha256"
	"fmt"
	"github.com/kataras/iris/v12"
	"github.com/lishimeng/app-starter"
	"github.com/lishimeng/app-starter/tool"
	"github.com/lishimeng/go-log"
	"github.com/lishimeng/owl-messager/internal/common"
	"github.com/lishimeng/owl-messager/internal/db/model"
	"strings"
	"time"
)

type AddReq struct {
	Name string `json:"name,omitempty"`
	Org  string `json:"org,omitempty"` // 指定组织
}

type AddResp struct {
	app.Response
	AppName string `json:"appName,omitempty"`
	Org     string `json:"org,omitempty"`
	AppId   string `json:"appId,omitempty"`
	Secret  string `json:"secret,omitempty"`
}

func add(ctx iris.Context) {

	var err error
	var req AddReq
	var resp AddResp

	err = ctx.ReadJSON(&req)
	if err != nil {
		log.Debug(err)
		resp.Code = tool.RespCodeError
		tool.ResponseJSON(ctx, resp)
		return
	}

	if len(req.Name) == 0 {
		common.RespLostParam("name", ctx)
		return
	}
	if len(req.Org) == 0 {
		common.RespLostParam("org", ctx)
		return
	}

	tenantCode := req.Org

	appId := genAppId(tenantCode)
	secret := genSecret(appId)

	tenant, err := getTenant(tenantCode)
	if err != nil {
		log.Debug("can't find tenant: %s", tenantCode)
		resp.Code = tool.RespCodeNotFound
		tool.ResponseJSON(ctx, resp)
	}

	var appInfo = model.OpenClient{
		AppId:  appId,
		Secret: secret,
		Domain: tenantCode,
	}
	appInfo.Org = tenant.Id
	// TODO create app --> db(check duplicate:appid and name_in_tenant)

	resp.Code = tool.RespCodeSuccess
	resp.AppName = req.Name
	resp.AppId = appId
	resp.Secret = secret
	tool.ResponseJSON(ctx, resp)
}

func genAppId(tenant string) (code string) {
	now := time.Now().Format(time.RFC3339Nano)
	var tmp = fmt.Sprintf("AppId_%s_%s", now, tenant)
	sh := sha256.New()
	sh.Write([]byte(tmp))
	bs := sh.Sum(nil)
	code = strings.ToLower(tool.BytesToHex(bs))
	return
}

func genSecret(appId string) (code string) {
	now := time.Now().Format(time.RFC3339Nano)
	var tmp = fmt.Sprintf("Secret_%s_%s", now, appId)
	sh := sha256.New()
	sh.Write([]byte(tmp))
	bs := sh.Sum(nil)
	code = strings.ToLower(tool.BytesToHex(bs))
	return
}
