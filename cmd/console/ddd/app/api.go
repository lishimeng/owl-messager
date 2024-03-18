package tenant

import (
	"crypto/sha256"
	"fmt"
	"github.com/lishimeng/app-starter"
	"github.com/lishimeng/app-starter/midware/auth"
	"github.com/lishimeng/app-starter/server"
	"github.com/lishimeng/app-starter/tool"
	"github.com/lishimeng/go-log"
	"github.com/lishimeng/owl-messager/internal/common"
	"github.com/lishimeng/owl-messager/internal/db/model"
	"github.com/lishimeng/x/util"
	"strings"
	"time"
)

type AddReq struct {
	Name string `json:"name,omitempty"`
}

type AddResp struct {
	app.Response
	AppName string `json:"appName,omitempty"`
	Org     string `json:"org,omitempty"`
	AppId   string `json:"appId,omitempty"`
	Secret  string `json:"secret,omitempty"`
}

func add(ctx server.Context) {

	var err error
	var req AddReq
	var resp AddResp

	err = ctx.C.ReadJSON(&req)
	if err != nil {
		log.Debug(err)
		resp.Code = tool.RespCodeError
		ctx.Json(resp)
		return
	}

	if len(req.Name) == 0 {
		common.RespLostParam("name", ctx)
		return
	}

	tenantCode := ctx.C.GetHeader(auth.OrgKey)

	appId := genAppId(tenantCode)
	secret := genSecret(appId)

	tenant, err := getTenant(tenantCode)
	if err != nil {
		log.Debug("can't find tenant: %s", tenantCode)
		resp.Code = tool.RespCodeNotFound
		ctx.Json(resp)
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
	ctx.Json(resp)
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
