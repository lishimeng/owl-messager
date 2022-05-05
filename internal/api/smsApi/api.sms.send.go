package smsApi

import (
	"encoding/json"
	"github.com/kataras/iris/v12"
	"github.com/lishimeng/app-starter"
	"github.com/lishimeng/go-log"
	"github.com/lishimeng/owl/internal/api/common"
	"github.com/lishimeng/owl/internal/db/model"
	"github.com/lishimeng/owl/internal/db/repo"
	"github.com/lishimeng/owl/internal/db/service"
)

type Req struct {
	Template      string      `json:"template,omitempty"` // template of this mail
	TemplateParam interface{} `json:"params,omitempty"`   // template params
	Sender        string      `json:"sender,omitempty"`   // sender空时，使用vendor
	Receiver      string      `json:"receiver,omitempty"` // receiver list(with comma if multi)
}

type Resp struct {
	app.Response
	MessageId int `json:"messageId,omitempty"`
}

func SendSms(ctx iris.Context) {
	log.Debug("send mail api")
	var req Req
	var resp Resp
	err := ctx.ReadJSON(&req)
	if err != nil {
		log.Info("read req fail")
		log.Info(err)
		resp.Code = -1
		resp.Message = "req error"
		common.ResponseJSON(ctx, resp)
		return
	}

	// check params
	log.Debug("check params")

	if len(req.Sender) == 0 {
		log.Debug("param sender code nil")
		resp.Code = -1
		resp.Message = "sender nil"
		common.ResponseJSON(ctx, resp)
		return
	}
	var sender *model.SmsSenderInfo
	if len(req.Sender) > 0 {
		s, err := repo.GetSmsSenderByCode(req.Sender)
		if err == nil {
			sender = &s
		}
	}

	if len(req.Template) == 0 {
		log.Debug("param template code nil")
		resp.Code = -1
		resp.Message = "template nil"
		common.ResponseJSON(ctx, resp)
		return
	}
	tpl, err := repo.GetSmsTemplateByCode(req.Template)
	if err != nil {
		log.Debug("param template not exist")
		resp.Code = -1
		resp.Message = "template not exist"
		common.ResponseJSON(ctx, resp)
		return
	}

	var tplParams string
	switch req.TemplateParam.(type) {
	case string:
		tplParams = (req.TemplateParam).(string)
	default:
		bs, e := json.Marshal(req.TemplateParam)
		if e == nil {
			tplParams = string(bs)
		}
	}

	m, err := service.CreateSmsMessage(
		sender,
		tpl,
		tplParams,
		req.Receiver)
	if err != nil {
		log.Info("can't create sms")
		log.Info(err)
		resp.Code = -1
		resp.Message = "create message failed"
		common.ResponseJSON(ctx, resp)
		return
	}

	log.Debug("create message success, id:%d", m.Id)
	resp.MessageId = m.Id

	resp.Code = common.RespCodeSuccess
	common.ResponseJSON(ctx, resp)
}
