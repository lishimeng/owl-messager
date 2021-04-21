package api

import (
	"github.com/kataras/iris"
	"github.com/lishimeng/owl/internal/api/mailApi"
	"github.com/lishimeng/owl/internal/api/messageApi"
	mailSender2 "github.com/lishimeng/owl/internal/api/senderApi"
	"github.com/lishimeng/owl/internal/api/smsApi"
	"github.com/lishimeng/owl/internal/api/taskApi"
	"github.com/lishimeng/owl/internal/api/templateApi"
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

	// send message
	mail(root.Party("/send/mail"))
	sms(root.Party("/send/sms"))
}

func message(p iris.Party) {
	p.Get("/", messageApi.GetMessageList)
	p.Get("/{id}", messageApi.GetMessageInfo)
	p.Post("/send/{id}", messageApi.Send)
}

func task(p iris.Party) {
	p.Get("/", taskApi.GetTaskList)
	p.Get("/{id}", taskApi.GetTaskInfo)
}

func mailSender(p iris.Party) {
	p.Post("/", mailSender2.AddMailSender)
	p.Put("/{id}", mailSender2.UpdateMailSender)
	p.Delete("/{id}", mailSender2.DeleteMailSender)

	p.Get("/", mailSender2.GetMailSenderList)
	p.Get("/{id}", mailSender2.GetMailSenderInfo)
}

func mailTemplate(p iris.Party) {
	p.Post("/", templateApi.AddMailTemplate)
	p.Put("/{id}", templateApi.UpdateMailTemplate)
	p.Delete("/{id}", templateApi.DeleteMailTemplate)

	p.Get("/", templateApi.GetMailTemplateList)
	p.Get("/{id}", templateApi.GetMailTemplateInfo)
}

func mail(p iris.Party) {
	p.Post("/", mailApi.SendMail)
	p.Get("/message/{id}", mailApi.GetByMessage)
}

func sms(p iris.Party) {
	p.Post("/", smsApi.SendSms)
}
