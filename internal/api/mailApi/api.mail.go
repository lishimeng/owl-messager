package mailApi

import (
	"github.com/kataras/iris"
	"github.com/lishimeng/app-starter"
	"github.com/lishimeng/go-log"
	"github.com/lishimeng/owl/internal/api/common"
	"github.com/lishimeng/owl/internal/db/repo"
	"github.com/lishimeng/owl/internal/db/service"
)

const (
	DefaultMailSubject = "no title"
)

type Req struct {
	Template string `json:"template"` // template of this mail
	TemplateParam string `json:"params"` // template params
	Subject string `json:"subject"` // mail's subject
	Sender string `json:"sender"` // mail send account on the platform
	Receiver string `json:"receiver"` // receiver list(with comma if multi)
	Cc string `json:"cc,omitempty"` // cc list(with comma if multi)
}

type Resp struct {
	app.Response
	MessageId int
}

func SendMail(ctx iris.Context) {
	log.Debug("send mail api")
	var req Req
	var resp Resp
	err := ctx.ReadJSON(&req)
	if err != nil {
		resp.Code = -1
		common.ResponseJSON(ctx, resp)
		return
	}

	// check params
	log.Debug("check params")
	if len(req.Subject) == 0 {
		log.Debug("no subject, use default: %s", DefaultMailSubject)
		req.Subject = DefaultMailSubject
	}

	if len(req.Sender) == 0 {
		log.Debug("param sender code nil")
		resp.Code = -1
		resp.Message = "sender nil"
		common.ResponseJSON(ctx, resp)
		return
	}
	sender, err := repo.GetMailSenderByCode(req.Sender)
	if err != nil {
		log.Debug("param sender not exist")
		resp.Code = -1
		resp.Message = "sender not exist"
		common.ResponseJSON(ctx, resp)
		return
	}

	if len(req.Template) == 0{
		log.Debug("param template code nil")
		resp.Code = -1
		resp.Message = "template nil"
		common.ResponseJSON(ctx, resp)
		return
	}
	tpl, err := repo.GetMailTemplateByCode(req.Template)
	if err != nil {
		log.Debug("param template not exist")
		resp.Code = -1
		resp.Message = "template not exist"
		common.ResponseJSON(ctx, resp)
		return
	}

	m, err := service.CreateMailMessage(
		sender,
		tpl,
		req.TemplateParam,
		req.Subject, req.Receiver, req.Cc)
	if err != nil {
		log.Info("can't create mail")
		log.Info(err)
		resp.Code = -1
		resp.Message = "create message failed"
		common.ResponseJSON(ctx, resp)
		return
	}

	log.Debug("create message success, id:%d", m.Id)
	resp.MessageId = m.Id

	resp.Code = 0
	common.ResponseJSON(ctx, resp)
}