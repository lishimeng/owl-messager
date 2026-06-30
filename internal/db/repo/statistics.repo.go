package repo

import (
	"errors"
	"time"

	"github.com/lishimeng/app-starter/persistence"
	"github.com/lishimeng/owl-messager/internal/db/model"
	"github.com/lishimeng/owl-messager/pkg/msg"
)

func updateDailyStat(ctx persistence.TxContext, tenantCode string, category msg.MessageCategory) (err error) {
	var daily model.DailySummary
	newDate := false
	err = ctx.Model(&model.DailySummary{}).
		Equal("tenant_code", tenantCode).
		Where("date = ?", time.Now().Format("2006-01-02")).
		First(&daily)
	if err != nil {
		daily = model.DailySummary{Date: time.Now()}
		daily.TenantCode = tenantCode
		newDate = true
	}
	switch category {
	case msg.MailMessage:
		daily.Mail++
	case msg.SmsMessage:
		daily.Sms++
	case msg.ImMessage:
		daily.Im++
	case msg.ApnsMessage:
		daily.Apns++
	default:
		err = errors.New("unknown category")
		return
	}
	if !newDate {
		err = ctx.Save(&daily)
	} else {
		err = ctx.Create(&daily)
	}
	return
}

func updateProviderStat(ctx persistence.TxContext, tenantCode string, category msg.MessageCategory, provider msg.MessageProvider) (err error) {
	var stat model.ProviderStats
	err = ctx.Model(&model.ProviderStats{}).
		Equal("tenant_code", tenantCode).
		Equal("category", category).
		Equal("provider", provider).
		First(&stat)
	if err != nil {
		stat = model.ProviderStats{
			Category: category,
			Provider: provider,
			Value:    1,
		}
		stat.TenantCode = tenantCode
		err = ctx.Create(&stat)
		return
	}
	stat.Value = stat.Value + 1
	err = ctx.Save(&stat)
	return
}

func UpdateStatistics(ctx persistence.TxContext, task model.MessageTask) (err error) {
	var info model.MessageInfo
	err = ctx.Model(&model.MessageInfo{}).Equal("id", task.MessageId).First(&info)
	if err != nil {
		return
	}
	tenantCode := info.TenantCode
	category := info.Category
	err = updateDailyStat(ctx, tenantCode, category)
	if err != nil {
		return
	}

	if category == msg.ApnsMessage {
		var apnsInfo model.ApnsMessageInfo
		err = ctx.Model(&model.ApnsMessageInfo{}).
			Equal("id", task.MessageInstanceId).
			First(&apnsInfo)
		if err != nil {
			return
		}
		var sender model.MessageSenderInfo
		err = ctx.Model(&model.MessageSenderInfo{}).
			Equal("id", apnsInfo.SenderId).
			First(&sender)
		if err != nil {
			return
		}
		return updateProviderStat(ctx, tenantCode, category, sender.Provider)
	}

	var templateId int
	switch category {
	case msg.MailMessage:
		var mailInfo model.MailMessageInfo
		err = ctx.Model(&model.MailMessageInfo{}).
			Equal("id", task.MessageInstanceId).
			First(&mailInfo)
		if err != nil {
			return
		}
		templateId = mailInfo.Template
	case msg.SmsMessage:
		var smsInfo model.SmsMessageInfo
		err = ctx.Model(&model.SmsMessageInfo{}).
			Equal("id", task.MessageInstanceId).
			First(&smsInfo)
		if err != nil {
			return
		}
		templateId = smsInfo.Template
	case msg.ImMessage:
		var imInfo model.ImMessageInfo
		err = ctx.Model(&model.ImMessageInfo{}).
			Equal("id", task.MessageInstanceId).
			First(&imInfo)
		if err != nil {
			return
		}
		templateId = imInfo.Template
	default:
		err = errors.New("unknown category")
		return
	}
	var tpl model.MessageTemplate
	err = ctx.Model(&model.MessageTemplate{}).Equal("id", templateId).First(&tpl)
	if err != nil {
		return
	}
	err = updateProviderStat(ctx, tenantCode, category, tpl.Provider)
	return
}
