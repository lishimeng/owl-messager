package model

import (
	"github.com/lishimeng/app-starter"
	"github.com/lishimeng/owl-messager/pkg/msg"
	"time"
)

type ProviderStats struct {
	app.TenantPk
	app.TableChangeInfo
	Category msg.MessageCategory `orm:"column(category)"`
	Provider msg.MessageProvider `orm:"column(provider)"`
	Value    int                 `orm:"column(value)"`
}

type DailySummary struct {
	app.TenantPk
	Date time.Time `orm:"column(date);type(date);unique"`
	Mail int       `orm:"column(mail)"`
	Sms  int       `orm:"column(sms)"`
	Im   int       `orm:"column(im)"`
	Apns int       `orm:"column(apns)"`
}
