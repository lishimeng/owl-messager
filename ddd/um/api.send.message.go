package um

import (
	"github.com/kataras/iris/v12"
	"github.com/lishimeng/app-starter"
	"github.com/lishimeng/go-log"
	"github.com/lishimeng/owl/internal/api/common"
	"github.com/lishimeng/owl/internal/db/model"
	"github.com/lishimeng/owl/internal/messager/msg"
)

type Req struct {
	Template      string      `json:"template"`          // 模板
	TemplateParam interface{} `json:"params"`            // 参数
	Title         string      `json:"subject,omitempty"` // 标题
	Receiver      string      `json:"receiver"`          // 接收者，多个时用逗号分隔
	Category      int         `json:"category"`          // 消息类型
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
		common.ResponseJSON(ctx, resp)
		return
	}

	// 检查收信人
	if len(req.Receiver) == 0 {
		log.Debug("param receiver nil")
		resp.Code = -1
		resp.Message = "receiver nil"
		common.ResponseJSON(ctx, resp)
		return
	}

	if len(req.Template) == 0 {
		log.Debug("param template code nil")
		resp.Code = -1
		resp.Message = "template nil"
		common.ResponseJSON(ctx, resp)
		return
	}

	// 检查消息类型(是否支持)
	var message model.MessageInfo
	switch req.Category {
	case msg.Email:
		message, resp, err = createMail(req)
	case msg.Sms:
		message, err = serviceAddSms(req.Template, req.Template, req.Receiver)
	case msg.Apns:
	default:

	}

	if err != nil {
		log.Info("can't create message")
		log.Info(err)
		common.ResponseJSON(ctx, resp)
		return
	}

	log.Debug("create message success, id:%d", message.Id)
	resp.MessageId = message.Id

	resp.Code = common.RespCodeSuccess
	common.ResponseJSON(ctx, resp)
}

func createMail(req Req) (m model.MessageInfo, errResponse Resp, err error) {

	m, err = serviceAddMail(req.Template, req.Template, req.Template, req.Receiver)
	return
}
