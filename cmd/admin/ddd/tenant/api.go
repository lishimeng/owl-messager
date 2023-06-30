package tenant

import (
	"crypto/sha256"
	"fmt"
	"github.com/kataras/iris/v12"
	"github.com/lishimeng/app-starter"
	"github.com/lishimeng/app-starter/tool"
	"github.com/lishimeng/go-log"
	"github.com/lishimeng/owl-messager/internal/common"
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

	tenantCode := genTenantCode()

	// TODO create tenant --> db(check duplicate:name and code)

	resp.Code = tenantCode
	resp.OrgName = req.Name
	resp.Code = tool.RespCodeSuccess
	tool.ResponseJSON(ctx, resp)
}

func genTenantCode() (code string) {
	now := time.Now().Format(time.RFC3339Nano)
	var tmp = fmt.Sprintf("tenant_%s", now)
	sh := sha256.New()
	sh.Write([]byte(tmp))
	bs := sh.Sum(nil)
	code = strings.ToLower(tool.BytesToHex(bs))
	return
}
