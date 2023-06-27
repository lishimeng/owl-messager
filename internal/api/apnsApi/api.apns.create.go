package apnsApi

import (
	"encoding/json"
	"github.com/kataras/iris/v12"
	"github.com/lishimeng/app-starter"
	"github.com/lishimeng/app-starter/tool"
	"github.com/lishimeng/go-log"
	"github.com/lishimeng/owl-messager/internal/db/repo"
	"github.com/lishimeng/owl-messager/internal/db/service"
)

const (
	DefaultApnsSubject = "no title"
)

type Req struct {
	BundleId string      `json:"bundleId,omitempty"` // bundle id
	Param    interface{} `json:"params,omitempty"`   // params
	Mode     int         `json:"mode,omitempty"`     // mode
	Subject  string      `json:"subject,omitempty"`  // subject
	Sender   string      `json:"sender,omitempty"`   // mail send account on the platform
	Receiver string      `json:"receiver,omitempty"` // receiver list(with comma if multi)
}

type Resp struct {
	app.Response
	MessageId int `json:"messageId,omitempty"`
}

func SendApns(ctx iris.Context) {
	log.Debug("send apns api")
	var req Req
	var resp Resp
	err := ctx.ReadJSON(&req)
	if err != nil {
		log.Info("read req fail")
		log.Info(err)
		resp.Code = -1
		resp.Message = "req error"
		tool.ResponseJSON(ctx, resp)
		return
	}

	// check params
	log.Debug("check params")
	if len(req.Subject) == 0 {
		log.Debug("no subject, use default: %s", DefaultApnsSubject)
		req.Subject = DefaultApnsSubject
	}

	if len(req.Sender) == 0 {
		log.Debug("param sender code nil")
		resp.Code = -1
		resp.Message = "sender nil"
		tool.ResponseJSON(ctx, resp)
		return
	}
	sender, err := repo.GetApnsSenderByCode(req.Sender)
	if err != nil {
		log.Debug("param sender not exist")
		resp.Code = -1
		resp.Message = "sender not exist"
		tool.ResponseJSON(ctx, resp)
		return
	}

	if len(req.BundleId) == 0 {
		log.Debug("param bundle id nil")
		resp.Code = -1
		resp.Message = "bundleId nil"
		tool.ResponseJSON(ctx, resp)
		return
	}

	var params string
	switch req.Param.(type) {
	case string:
		params = (req.Param).(string)
	default:
		bs, e := json.Marshal(req.Param)
		if e == nil {
			params = string(bs)
		}
	}

	m, err := service.CreateApnsMessage(
		sender,
		req.Mode,
		req.BundleId,
		params,
		req.Subject, req.Receiver)
	if err != nil {
		log.Info("can't create apns")
		log.Info(err)
		resp.Code = -1
		resp.Message = "create message failed"
		tool.ResponseJSON(ctx, resp)
		return
	}

	log.Debug("create message success, id:%d", m.Id)
	resp.MessageId = m.Id

	resp.Code = tool.RespCodeSuccess
	tool.ResponseJSON(ctx, resp)
}
