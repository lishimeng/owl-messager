package smsApi

import (
	"github.com/kataras/iris"
	"github.com/lishimeng/app-starter"
	"github.com/lishimeng/go-log"
	"github.com/lishimeng/owl/internal/api/common"
)

//SendSms
/**
@Summary send a sms

@Router /api/send/sms [post]
*/
func SendSms(ctx iris.Context) {
	log.Debug("send sms api")
	var resp app.Response
	// TODO
	resp.Code = common.RespCodeSuccess
	common.ResponseJSON(ctx, resp)
}
