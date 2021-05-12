package api

import (
	"github.com/kataras/iris/v12"
	"github.com/lishimeng/owl/internal/api/mailApi"
	"github.com/lishimeng/owl/internal/api/messageApi"
	"github.com/lishimeng/owl/internal/api/openApi"
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

	mail(root.Party("/mail"))

	// send message
	sendMail(root.Party("/send/mail"))
	sms(root.Party("/send/sms"))

	oauth(root.Party("/oauth2"))
}

func oauth(p iris.Party) {
	p.Get("/authorize", openApi.Authorize)
	p.Get("/token", openApi.Token)
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
}

func mailSender(p iris.Party) {
	p.Post("/", senderApi.AddMailSender)
	p.Put("/{id}", senderApi.UpdateMailSender)
	p.Delete("/{id}", senderApi.DeleteMailSender)

	p.Get("/", senderApi.GetMailSenderList)
	p.Get("/{id}", senderApi.GetMailSenderInfo)
}

func mailTemplate(p iris.Party) {
	p.Post("/", templateApi.AddMailTemplate)
	p.Put("/{id}", templateApi.UpdateMailTemplate)
	p.Delete("/{id}", templateApi.DeleteMailTemplate)

	p.Get("/", templateApi.GetMailTemplateList)
	p.Get("/{id}", templateApi.GetMailTemplateInfo)
}

func mail(p iris.Party) {
	p.Get("/message/{id}", mailApi.GetByMessage)
}
func sendMail(p iris.Party) {
	p.Post("/", openapi.CheckAccessToken, mailApi.SendMail)
}

func sms(p iris.Party) {
	p.Post("/", smsApi.SendSms)
}
