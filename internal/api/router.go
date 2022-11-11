package api

import (
	"github.com/kataras/iris/v12"
	"github.com/lishimeng/owl/ddd"
	"github.com/lishimeng/owl/internal/api/apnsApi"
	"github.com/lishimeng/owl/internal/api/mailApi"
	"github.com/lishimeng/owl/internal/api/messageApi"
	"github.com/lishimeng/owl/internal/api/senderApi"
	"github.com/lishimeng/owl/internal/api/smsApi"
	"github.com/lishimeng/owl/internal/api/taskApi"
	"github.com/lishimeng/owl/internal/api/templateApi"
	"github.com/lishimeng/owl/internal/openapi"
)

func Route(app *iris.Application) {
	root := app.Party("/api")
	router(root)
	return
}

func router(root iris.Party) {
	task(root.Party("/task"))
	message(root.Party("/message"))

	mailSender(root.Party("/mail_sender"))
	mailTemplate(root.Party("/mail_template"))

	smsSender(root.Party("/sms_sender"))

	mail(root.Party("/mail"))

	// send message
	sendMail(root.Party("/send/mail"))
	sms(root.Party("/send/sms"))
	apns(root.Party("/send/apns"))

	// api v2
	ddd.Router(root.Party("/v2"))
}

func message(p iris.Party) {
	p.Get("/", messageApi.GetMessageList)
	p.Get("/{id}", messageApi.GetMessageInfo)
	p.Post("/send/{id}", messageApi.Send)
}

func task(p iris.Party) {
	p.Get("/", taskApi.GetTaskList)
	p.Get("/{id}", taskApi.GetTaskInfo)
	p.Get("/message/{id}", taskApi.GetByMessage)

	p.Get("/send/monitor", taskApi.TaskMonitorWs())
}

func mailSender(p iris.Party) {
	p.Post("/add", senderApi.AddMailSender)
	p.Post("/update", senderApi.UpdateMailSender)
	p.Get("/del/{id}", senderApi.DeleteMailSender)

	p.Get("/", senderApi.GetMailSenderListNow)
	p.Get("/one/{id}", senderApi.GetMailSenderInfoNow)
}

func smsSender(p iris.Party) {
	p.Post("/", senderApi.AddSmsSender)
	p.Put("/{id}", senderApi.UpdateSmsSender)
	p.Delete("/{id}", senderApi.DeleteSmsSender)

	p.Get("/", senderApi.GetSmsSenderList)
	p.Get("/{id}", senderApi.GetSmsSenderInfo)
}

func mailTemplate(p iris.Party) {
	p.Post("/add", templateApi.AddMailTemplate)
	p.Post("/update", templateApi.UpdateMailTemplate)
	p.Get("/del/{id}", templateApi.DeleteMailTemplate)

	p.Get("/", templateApi.GetMailTemplateList)
	p.Get("/one/{id}", templateApi.GetMailTemplateInfo)
}

func mail(p iris.Party) {
	p.Get("/message/mail/{id}", mailApi.GetByMessage)

	p.Get("/message/sms/{id}", smsApi.GetByMessage)
	p.Get("/message/apns/{id}", apnsApi.GetByMessage)
}
func sendMail(p iris.Party) {
	p.Post("/", openapi.CheckAccessToken, mailApi.SendMail)
}

func sms(p iris.Party) {
	p.Post("/", smsApi.SendSms)
}

func apns(p iris.Party) {
	p.Post("/", apnsApi.SendApns)
}
