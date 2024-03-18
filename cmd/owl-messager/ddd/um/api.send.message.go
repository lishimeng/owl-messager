package um

import (
	"encoding/json"
	"fmt"
	"github.com/lishimeng/app-starter"
	"github.com/lishimeng/app-starter/midware/auth"
	"github.com/lishimeng/app-starter/server"
	"github.com/lishimeng/app-starter/tool"
	"github.com/lishimeng/go-log"
	"github.com/lishimeng/owl-messager/internal/db/model"
	"github.com/lishimeng/owl-messager/internal/db/repo"
	"github.com/lishimeng/owl-messager/internal/etc"
	"github.com/lishimeng/owl-messager/internal/messager/task"
	"github.com/lishimeng/owl-messager/pkg/msg"
	"github.com/lishimeng/x/container"
)

const (
	DefaultTitle = ""
)

type Req struct {
	BundleId      string      `json:"bundleId,omitempty"` // bundle id
	Template      string      `json:"template"`           // 模板
	TemplateParam interface{} `json:"params"`             // 参数
	Title         string      `json:"subject,omitempty"`  // 标题
	Receiver      string      `json:"receiver"`           // 接收者，多个时用逗号分隔
}

type Resp struct {
	app.Response
	MessageId int `json:"messageId,omitempty"`
}

func sendMessage(ctx server.Context) {
	log.Info("Union message send function")
	var req Req
	var resp Resp
	var org = ctx.C.GetHeader(auth.OrgKey)
	err := ctx.C.ReadJSON(&req)
	if err != nil {
		log.Info("read req fail")
		log.Info(err)
		resp.Code = -1
		resp.Message = "req error"
		ctx.Json(resp)
		return
	}

	var category = ctx.C.Params().Get("category")
	validCategory := msg.IsValidCategory(msg.MessageCategory(category))
	if !validCategory {
		log.Debug("unknown category: %s", category)
		resp.Code = -1
		resp.Message = "unknown category"
		ctx.Json(resp)
		return
	}

	// 检查收信人
	if len(req.Receiver) == 0 {
		log.Debug("param receiver nil")
		resp.Code = -1
		resp.Message = "receiver nil"
		ctx.Json(resp)
		return
	}

	if len(req.Template) == 0 {
		log.Debug("param template code nil")
		resp.Code = -1
		resp.Message = "template nil"
		ctx.Json(resp)
		return
	}

	var params string
	switch req.TemplateParam.(type) {
	case string:
		params = (req.TemplateParam).(string)
	default:
		bs, e := json.Marshal(req.TemplateParam)
		if e == nil {
			params = string(bs)
		}
	}

	tenant, err := repo.GetTenant(org)
	if err != nil {
		log.Debug("unknown tenant: %s", org)
		resp.Code = -1
		resp.Message = "unknown tenant"
		ctx.Json(resp)
		return
	}

	// 检查消息类型(是否支持)
	var message model.MessageInfo
	switch msg.MessageCategory(category) {
	case msg.MailMessage:
		message, resp, err = createMail(tenant.Id, req, params)
	case msg.SmsMessage:
		message, resp, err = createSms(tenant.Id, req, params)
	case msg.ApnsMessage:
		message, resp, err = createApns(tenant.Id, req, params)
	default:
		err = fmt.Errorf("unkown message category")
		resp.Code = -1
	}

	if err != nil {
		log.Info("can't create message")
		log.Info(err)
		ctx.Json(resp)
		return
	}

	var senderStrategy = task.Strategy(etc.Config.Sender.Strategy)
	switch senderStrategy {
	case task.MemQueue:
		var handler task.MessageTask
		e := container.Get(&handler)
		if e != nil {
			log.Debug(e)
		} else {
			_ = handler.HandleMessage(message)
		}
	}

	log.Debug("create message success, id:%d", message.Id)
	resp.MessageId = message.Id

	resp.Code = tool.RespCodeSuccess
	resp.Message = "OK"
	ctx.Json(resp)
}

func createMail(org int, req Req, params string) (m model.MessageInfo, resp Resp, err error) {
	if len(req.Title) == 0 {
		log.Debug("no title, use default: %s", DefaultTitle)
		req.Title = DefaultTitle
	}

	m, err = serviceAddMail(org, req.Template, params, req.Title, req.Receiver)
	if err != nil {
		resp.Code = -1
		resp.Message = "create mail message failed"
	}
	return
}

func createSms(org int, req Req, params string) (m model.MessageInfo, resp Resp, err error) {
	m, err = serviceAddSms(org, req.Template, params, req.Receiver)
	if err != nil {
		resp.Code = -1
		resp.Message = "create sms message failed"
	}
	return
}

func createApns(org int, req Req, params string) (m model.MessageInfo, resp Resp, err error) {
	if len(req.Title) == 0 {
		log.Debug("no title, use default: %s", DefaultTitle)
		req.Title = DefaultTitle
	}

	if len(req.BundleId) == 0 {
		log.Debug("param bundle id nil")
		resp.Code = -1
		resp.Message = "bundleId nil"
		return
	}

	m, err = serviceAddApns(org, req.Template, params, req.Title, req.Receiver)
	if err != nil {
		resp.Code = -1
		resp.Message = "create sms message failed"
	}
	return
}
