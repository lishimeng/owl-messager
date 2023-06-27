package mailApi

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

const (
	DefaultMailSubject = "no title"
)

type Req struct {
	Template      string      `json:"template"`           // template of this mail
	TemplateParam interface{} `json:"params,omitempty"`   // template params
	Subject       string      `json:"subject,omitempty"`  // mail's subject
	Sender        string      `json:"sender,omitempty"`   // mail send account on the platform
	Receiver      string      `json:"receiver,omitempty"` // receiver list(with comma if multi)
	Cc            string      `json:"cc,omitempty"`       // cc list(with comma if multi)
}

type Resp struct {
	app.Response
	MessageId int `json:"messageId,omitempty"`
}

// SendMail
/**
@Summary send a email

@Router /api/send/mail [post]

@Example
http://localhost/api/send/mail

{
	"template":"b7411049bbfe8068",
	"params":{"content":"O35A0001"},
	"subject":"电量低超提醒",
	"sender":"e949ae24481a9527",
	"receiver":"xxxx@qq.com"
}

*/
func SendMail(ctx iris.Context) {
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
	if len(req.Subject) == 0 {
		log.Debug("no subject, use default: %s", DefaultMailSubject)
		req.Subject = DefaultMailSubject
	}

	if len(req.Receiver) == 0 {
		log.Debug("param receiver nil")
		resp.Code = -1
		resp.Message = "receiver nil"
		tool.ResponseJSON(ctx, resp)
		return
	}

	var sender *model.MailSenderInfo
	if len(req.Sender) > 0 { // 如果指定sender，应检查sender是否配置了
		*sender, err = repo.GetMailSenderByCode(req.Sender)
		if err != nil {
			log.Debug("mail sender not exist: %s", req.Sender)
			resp.Code = -1
			resp.Message = "sender not exist"
			tool.ResponseJSON(ctx, resp)
			return
		}
	}

	if len(req.Template) == 0 {
		log.Debug("param template code nil")
		resp.Code = -1
		resp.Message = "template nil"
		tool.ResponseJSON(ctx, resp)
		return
	}
	tpl, err := repo.GetMailTemplateByCode(req.Template)
	if err != nil {
		log.Debug("param template not exist")
		resp.Code = -1
		resp.Message = "template not exist"
		tool.ResponseJSON(ctx, resp)
		return
	}

	var templateParams string
	switch req.TemplateParam.(type) {
	case string:
		templateParams = (req.TemplateParam).(string)
	default:
		bs, e := json.Marshal(req.TemplateParam)
		if e == nil {
			templateParams = string(bs)
		}
	}

	m, err := service.CreateMailMessage(
		sender,
		tpl,
		templateParams,
		req.Subject, req.Receiver, req.Cc)
	if err != nil {
		log.Info("can't create mail")
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
