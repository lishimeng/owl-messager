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
	p.Post("/{mail_sender_id}", UpdateMailSender)
	p.Post("/{mail_sender_id}", DeleteMailSender)

	p.Get("/", GetMailSenderList)
	p.Get("/", GetMailSenderInfo)
}

func mail(p iris.Party) {
	p.Post("/", SendMail)
}

func sms(p iris.Party) {
	p.Post("/", SendSms)
}