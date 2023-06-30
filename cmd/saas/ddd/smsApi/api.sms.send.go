package smsApi

import (
	"encoding/json"
	"github.com/kataras/iris/v12"
	"github.com/lishimeng/app-starter"
	"github.com/lishimeng/app-starter/tool"
	"github.com/lishimeng/go-log"
	"github.com/lishimeng/owl-messager/internal/db/model"
	"github.com/lishimeng/owl-messager/internal/db/repo"
	"github.com/lishimeng/owl-messager/internal/db/service"
)

type Req struct {
	Template      string      `json:"template"`         // template of this mail
	TemplateParam interface{} `json:"params"`           // template params
	Sender        string      `json:"sender,omitempty"` // sms send account on the platform
	Signature     string      `json:"signature,omitempty"`
	Receiver      string      `json:"receiver"` // receiver list(with comma if multi)
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
		tool.ResponseJSON(ctx, resp)
		return
	}

	// check params
	log.Debug("check params")

	if len(req.Receiver) == 0 {
		log.Debug("param receiver nil")
		resp.Code = -1
		resp.Message = "receiver nil"
		tool.ResponseJSON(ctx, resp)
		return
	}

	if len(req.Template) == 0 {
		log.Debug("param template code nil")
		resp.Code = -1
		resp.Message = "template nil"
		tool.ResponseJSON(ctx, resp)
		return
	}

	var sender *model.SmsSenderInfo
	if len(req.Sender) > 0 { // 如果指定sender，应检查sender是否配置了
		*sender, err = repo.GetSmsSenderByCode(req.Sender)
		if err != nil {
			log.Debug("sms sender not exist: %s", req.Sender)
			resp.Code = -1
			resp.Message = "sender not exist"
			tool.ResponseJSON(ctx, resp)
			return
		}
	}

	tpl, err := repo.GetSmsTemplateByCode(req.Template)
	if err != nil {
		log.Debug("param template not exist")
		resp.Code = -1
		resp.Message = "template not exist"
		tool.ResponseJSON(ctx, resp)
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
		tpl,
		tplParams,
		req.Receiver,
	)
	if err != nil {
		log.Info("can't create sms")
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
