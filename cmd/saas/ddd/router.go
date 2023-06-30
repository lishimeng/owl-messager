package ddd

import (
	"github.com/kataras/iris/v12"
	"github.com/lishimeng/owl-messager/cmd/saas/ddd/apnsApi"
	"github.com/lishimeng/owl-messager/cmd/saas/ddd/mailApi"
	"github.com/lishimeng/owl-messager/cmd/saas/ddd/messageApi"
	"github.com/lishimeng/owl-messager/cmd/saas/ddd/sender"
	"github.com/lishimeng/owl-messager/cmd/saas/ddd/senderApi"
	"github.com/lishimeng/owl-messager/cmd/saas/ddd/smsApi"
	"github.com/lishimeng/owl-messager/cmd/saas/ddd/taskApi"
	"github.com/lishimeng/owl-messager/cmd/saas/ddd/templateApi"
)

func Route(app *iris.Application) {
	root := app.Party("/api")
	router(root)
	return
}

func router(root iris.Party) {
	task(root.Party("/task"))
	message(root.Party("/message"))
	template(root.Party("/template"))
	vendor(root.Party("/vendor"))
	mail(root.Party("/mail"))
	// send message
	sendMail(root.Party("/send/mail"))
	sms(root.Party("/send/sms"))
	apns(root.Party("/send/apns"))
	sender.Route(root.Party("/sender"))
}

// vendor /api/vendor/
func vendor(p iris.Party) {
	p.Get("/mail", templateApi.GetMailVendors)
	p.Get("/sms", templateApi.GetSmsVendors)
}

// template /api/template/
func template(p iris.Party) {
	mailTemplate(p.Party("/mail"))
	smsTemplate(p.Party("/sms"))
	p.Get("/getTemplateList", templateApi.GetTemplateListByPage)
	p.Post("/createTemplate", templateApi.CreateTemplate)
	p.Post("/updateTemplate", templateApi.UpdateTemplate)
	p.Get("/getTemplateInfo", templateApi.GetTemplateInfo)
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
	p.Post("/", senderApi.AddMailSender)
	p.Put("/{id}", senderApi.UpdateMailSender)
	p.Delete("/{id}", senderApi.DeleteMailSender)

	p.Get("/", senderApi.GetMailSenderList)
	p.Get("/{id}", senderApi.GetMailSenderInfo)
}

func smsSender(p iris.Party) {
	p.Post("/", senderApi.AddSmsSender)
	p.Put("/{id}", senderApi.UpdateSmsSender)
	p.Delete("/{id}", senderApi.DeleteSmsSender)

	p.Get("/", senderApi.GetSmsSenderList)
	p.Get("/{id}", senderApi.GetSmsSenderInfo)
}

func mailTemplate(p iris.Party) {
	p.Post("/", templateApi.AddMailTemplate)
	p.Put("/{id}", templateApi.UpdateMailTemplate)
	p.Put("/{id}/status", templateApi.ChangeMailTemplateStatus)
	p.Delete("/{id}", templateApi.DeleteMailTemplate)

	p.Get("/", templateApi.GetMailTemplateList)
	p.Get("/{id}", templateApi.GetMailTemplateInfo)
}

func smsTemplate(p iris.Party) {
	p.Post("/", templateApi.AddSmsTemplate)
	p.Put("/{id}", templateApi.UpdateMailTemplate)
	p.Put("/{id}/status", templateApi.ChangeSmsTemplateStatus)
	p.Delete("/{id}", templateApi.DeleteMailTemplate)

	p.Get("/", templateApi.GetMailTemplateList)
	p.Get("/{id}", templateApi.GetMailTemplateInfo)
}

func mail(p iris.Party) {
	p.Get("/message/mail/{id}", mailApi.GetByMessage)

	p.Get("/message/sms/{id}", smsApi.GetByMessage)
	p.Get("/message/apns/{id}", apnsApi.GetByMessage)
}
func sendMail(p iris.Party) {
	p.Post("/", mailApi.SendMail)
}

func sms(p iris.Party) {
	p.Post("/", smsApi.SendSms)
}

func apns(p iris.Party) {
	p.Post("/", apnsApi.SendApns)
}
