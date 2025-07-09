package repo

import (
	"errors"
	"github.com/lishimeng/app-starter"
	"github.com/lishimeng/app-starter/persistence"
	"github.com/lishimeng/owl-messager/internal/db/model"
	"github.com/lishimeng/owl-messager/pkg/msg"
	"time"
)

func updateDailyStat(ctx persistence.TxContext, org int, category msg.MessageCategory) (err error) {
	var daily model.DailySummary
	newDate := false
	err = ctx.Context.QueryTable(new(model.DailySummary)).
		Filter("Org", org).
		Filter("Date", time.Now()).
		One(&daily)
	if err != nil {
		daily = model.DailySummary{Date: time.Now()}
		daily.Org = org
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
		_, err = ctx.Context.Update(&daily)
	} else {
		_, err = ctx.Context.Insert(&daily)
	}
	return
}

func updateProviderStat(ctx persistence.TxContext, org int, category msg.MessageCategory, provider msg.MessageProvider) (err error) {
	var stat model.ProviderStats
	err = ctx.Context.QueryTable(new(model.ProviderStats)).
		Filter("Org", org).
		Filter("Category", category).
		Filter("Provider", provider).
		One(&stat)
	if err != nil {
		stat = model.ProviderStats{
			Category: category,
			Provider: provider,
			Value:    1,
		}
		stat.Org = org
		_, err = ctx.Context.Insert(&stat)
		return
	}
	stat.Value = stat.Value + 1
	_, err = ctx.Context.Update(&stat)
	return
}

func UpdateStatistics(ctx persistence.TxContext, task model.MessageTask) (err error) {
	var info model.MessageInfo
	var templateId int
	err = ctx.Context.QueryTable(new(model.MessageInfo)).
		Filter("Id", task.MessageId).
		One(&info)
	if err != nil {
		return
	}
	org := info.Org
	category := info.Category
	err = updateDailyStat(ctx, org, category)
	if err != nil {
		return
	}
	switch category {
	case msg.MailMessage:
		var info model.MailMessageInfo
		err = ctx.Context.QueryTable(new(model.MailMessageInfo)).
			Filter("message_id", task.MessageInstanceId).
			One(&info)
		if err != nil {
			return
		}
		templateId = info.Template
	case msg.SmsMessage:
		var info model.SmsMessageInfo
		err = app.GetOrm().Context.QueryTable(new(model.SmsMessageInfo)).
			Filter("message_id", task.MessageInstanceId).
			One(&info)
		if err != nil {
			return
		}
		templateId = info.Template
	case msg.ImMessage:
		var info model.ImMessageInfo
		err = app.GetOrm().Context.QueryTable(new(model.ImMessageInfo)).
			Filter("message_id", task.MessageInstanceId).
			One(&info)
		if err != nil {
			return
		}
		templateId = info.Template
	default:
		err = errors.New("unknown category")
		return
	}
	// 查询template才能知道provider
	var tpl model.MessageTemplate
	err = app.GetOrm().Context.QueryTable(new(model.MessageTemplate)).
		Filter("id", templateId).
		One(&tpl)
	if err != nil {
		return
	}
	err = updateProviderStat(ctx, org, category, tpl.Provider)
	return
}
