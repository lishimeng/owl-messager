package um

import (
	"encoding/json"
	"fmt"
	"github.com/kataras/iris/v12"
	"github.com/lishimeng/app-starter"
	"github.com/lishimeng/app-starter/tool"
	"github.com/lishimeng/go-log"
	"github.com/lishimeng/owl-messager/internal/db/model"
	"github.com/lishimeng/owl-messager/internal/messager/msg"
)

const (
	DefaultTitle = ""
)

type Req struct {
	BundleId      string      `json:"bundleId,omitempty"`      // bundle id
	Template      string      `json:"template"`                // 模板
	CloudTemplate bool        `json:"cloudTemplate,omitempty"` // 云端模板
	TemplateParam interface{} `json:"params"`                  // 参数
	Title         string      `json:"subject,omitempty"`       // 标题
	Receiver      string      `json:"receiver"`                // 接收者，多个时用逗号分隔
}

type Resp struct {
	app.Response
	MessageId int `json:"messageId,omitempty"`
}

func sendMessage(ctx iris.Context) {
	log.Info("Union message send function")
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

	var category = ctx.Params().Get("category")

	// 检查收信人
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

	// 检查消息类型(是否支持)
	var message model.MessageInfo
	switch category {
	case msg.EmailCategory:
		message, resp, err = createMail(req, params)
	case msg.SmsCategory:
		message, resp, err = createSms(req, params)
	case msg.ApnsCategory:
		message, resp, err = createApns(req, params)
	default:
		err = fmt.Errorf("unkown message category")
		resp.Code = -1
	}

	if err != nil {
		log.Info("can't create message")
		log.Info(err)
		tool.ResponseJSON(ctx, resp)
		return
	}

	log.Debug("create message success, id:%d", message.Id)
	resp.MessageId = message.Id

	resp.Code = tool.RespCodeSuccess
	tool.ResponseJSON(ctx, resp)
}

func createMail(req Req, params string) (m model.MessageInfo, resp Resp, err error) {
	if len(req.Title) == 0 {
		log.Debug("no title, use default: %s", DefaultTitle)
		req.Title = DefaultTitle
	}
	m, err = serviceAddMail(req.Template, req.CloudTemplate, req.Template, params, req.Title, req.Receiver)
	if err != nil {
		resp.Code = -1
		resp.Message = "create mail message failed"
	}
	return
}

func createSms(req Req, params string) (m model.MessageInfo, resp Resp, err error) {
	m, err = serviceAddSms(req.Template, params, req.Receiver)
	if err != nil {
		resp.Code = -1
		resp.Message = "create sms message failed"
	}
	return
}

func createApns(req Req, params string) (m model.MessageInfo, resp Resp, err error) {
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

	m, err = serviceAddApns(req.Template, params, req.Title, req.Receiver)
	if err != nil {
		resp.Code = -1
		resp.Message = "create sms message failed"
	}
	return
}
