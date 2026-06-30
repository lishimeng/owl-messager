package repo

import (
	"github.com/lishimeng/app-starter/persistence"
	"github.com/lishimeng/owl-messager/internal/db/model"
)

func initTemplateChannel(
	d *model.TemplateChannelDetail,
	message model.MessageInfo,
	template model.MessageTemplate,
	templateParams, receiver string,
) {
	d.TenantCode = message.TenantCode
	d.MessageId = message.Id
	d.Template = template.Id
	d.Params = templateParams
	d.Receivers = receiver
}

func initPushChannel(
	d *model.PushChannelDetail,
	message model.MessageInfo,
	sender model.MessageSenderInfo,
	mode int,
	bundleId, params, receiver string,
) {
	d.TenantCode = message.TenantCode
	d.MessageId = message.Id
	d.SenderId = sender.Id
	d.ApnsMode = model.ApnsMode(mode)
	d.BundleId = bundleId
	d.Params = params
	d.Receivers = receiver
}

func createTemplateChannelMessage(
	ctx persistence.TxContext,
	dest interface{},
	d *model.TemplateChannelDetail,
	message model.MessageInfo,
	template model.MessageTemplate,
	templateParams, receiver string,
) error {
	initTemplateChannel(d, message, template, templateParams, receiver)
	return ctx.Create(dest)
}
