package api

import (
	"github.com/kataras/iris"
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
	mail(root.Party("/mail"))
	sms(root.Party("/sms"))
}

func message(p iris.Party) {
	p.Get("/", GetMessageList)
	p.Get("/{message_id}", GetMessageInfo)
}

func task(p iris.Party) {
	p.Get("/", GetTaskList)
	p.Get("/{task_id}", GetTaskInfo)
}

func mailSender(p iris.Party) {
	p.Post("/", AddMailSender)
	p.Put("/{id}", UpdateMailSender)
	p.Delete("/{id}", DeleteMailSender)

	p.Get("/", GetMailSenderList)
	p.Get("/{id}", GetMailSenderInfo)
}

func mailTemplate(p iris.Party) {
	p.Post("/", AddMailTemplate)
	p.Put("/{id}", UpdateMailTemplate)
	p.Delete("/{id}", DeleteMailTemplate)

	p.Get("/", GetMailTemplateList)
	p.Get("/{id}", GetMailTemplateInfo)
}

func mail(p iris.Party) {
	p.Post("/", SendMail)
}

func sms(p iris.Party) {
	p.Post("/", SendSms)
}