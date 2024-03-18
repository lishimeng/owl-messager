package tenant

import (
	"crypto/sha256"
	"fmt"
	"github.com/lishimeng/app-starter"
	"github.com/lishimeng/app-starter/server"
	"github.com/lishimeng/app-starter/tool"
	"github.com/lishimeng/go-log"
	"github.com/lishimeng/owl-messager/internal/common"
	"github.com/lishimeng/x/util"
	"strings"
	"time"
)

type AddReq struct {
	Name string `json:"name,omitempty"`
}

type AddResp struct {
	app.Response
	OrgName string `json:"orgName,omitempty"`
	OrgCode string `json:"orgCode,omitempty"`
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

	tenantCode := genTenantCode()

	// TODO create tenant --> db(check duplicate:name and code)

	resp.OrgCode = tenantCode
	resp.OrgName = req.Name
	resp.Code = tool.RespCodeSuccess
	ctx.Json(resp)
}

func genTenantCode() (code string) {
	now := time.Now().Format(time.RFC3339Nano)
	var tmp = fmt.Sprintf("tenant_%s", now)
	sh := sha256.New()
	sh.Write([]byte(tmp))
	bs := sh.Sum(nil)
	code = strings.ToLower(util.BytesToHex(bs))
	return
}
