package ddd

import (
	"github.com/lishimeng/app-starter/server"
	"github.com/lishimeng/owl-messager/cmd/console/ddd/apnsApi"
	"github.com/lishimeng/owl-messager/cmd/console/ddd/mailApi"
	"github.com/lishimeng/owl-messager/cmd/console/ddd/messageApi"
	"github.com/lishimeng/owl-messager/cmd/console/ddd/sender"
	"github.com/lishimeng/owl-messager/cmd/console/ddd/senderApi"
	"github.com/lishimeng/owl-messager/cmd/console/ddd/smsApi"
	"github.com/lishimeng/owl-messager/cmd/console/ddd/taskApi"
	"github.com/lishimeng/owl-messager/cmd/console/ddd/templateApi"
)

func Route(app server.Router) {
	root := app.Path("/api")
	router(root)
	return
}

func router(root server.Router) {
	task(root.Path("/task"))
	message(root.Path("/message"))
	template(root.Path("/template"))
	vendor(root.Path("/vendor"))
	mail(root.Path("/mail"))

	sender.Route(root.Path("/sender"))
}

// vendor /api/vendor/
func vendor(p server.Router) {
	p.Get("/mail", templateApi.GetMailVendors)
	p.Get("/sms", templateApi.GetSmsVendors)
}

// template /api/template/
func template(p server.Router) {
	mailTemplate(p.Path("/mail"))
	smsTemplate(p.Path("/sms"))
	p.Get("/getTemplateList", templateApi.GetTemplateListByPage)
	p.Post("/createTemplate", templateApi.CreateTemplate)
	p.Post("/updateTemplate", templateApi.UpdateTemplate)
	p.Get("/getTemplateInfo", templateApi.GetTemplateInfo)
}

func message(p server.Router) {
	p.Get("/", messageApi.GetMessageList)
	p.Get("/{id}", messageApi.GetMessageInfo)
	p.Post("/send/{id}", messageApi.Send)
}

func task(p server.Router) {
	p.Get("/", taskApi.GetTaskList)
	p.Get("/{id}", taskApi.GetTaskInfo)
	p.Get("/message/{id}", taskApi.GetByMessage)

	//p.Get("/send/monitor", taskApi.TaskMonitorWs()) // TODO
}

func mailSender(p server.Router) {
	p.Post("/", senderApi.AddMailSender)
	//p.Put("/{id}", senderApi.UpdateMailSender)
	//p.Delete("/{id}", senderApi.DeleteMailSender) // TODO

	p.Get("/", senderApi.GetMailSenderList)
	p.Get("/{id}", senderApi.GetMailSenderInfo)
}

func smsSender(p server.Router) {
	p.Post("/", senderApi.AddSmsSender)
	//p.Put("/{id}", senderApi.UpdateSmsSender)
	//p.Delete("/{id}", senderApi.DeleteSmsSender) // TODO

	p.Get("/", senderApi.GetSmsSenderList)
	p.Get("/{id}", senderApi.GetSmsSenderInfo)
}

func mailTemplate(p server.Router) {
	p.Post("/", templateApi.AddMailTemplate)
	//p.Put("/{id}", templateApi.UpdateMailTemplate) // TODO
	//p.Put("/{id}/status", templateApi.ChangeMailTemplateStatus)
	//p.Delete("/{id}", templateApi.DeleteMailTemplate)

	p.Get("/", templateApi.GetMailTemplateList)
	p.Get("/{id}", templateApi.GetMailTemplateInfo)
}

func smsTemplate(p server.Router) {
	p.Post("/", templateApi.AddSmsTemplate)
	//p.Put("/{id}", templateApi.UpdateMailTemplate) // TODO
	//p.Put("/{id}/status", templateApi.ChangeSmsTemplateStatus)
	//p.Delete("/{id}", templateApi.DeleteMailTemplate)

	p.Get("/", templateApi.GetMailTemplateList)
	p.Get("/{id}", templateApi.GetMailTemplateInfo)
}

func mail(p server.Router) {
	p.Get("/message/mail/{id}", mailApi.GetByMessage)

	p.Get("/message/sms/{id}", smsApi.GetByMessage)
	p.Get("/message/apns/{id}", apnsApi.GetByMessage)
}
