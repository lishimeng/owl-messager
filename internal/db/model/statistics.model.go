package model

import (
	"time"

	"github.com/lishimeng/app-starter"
	"github.com/lishimeng/owl-messager/pkg/msg"
)

type ProviderStats struct {
	app.TenantPk
	app.TableChangeInfo
	Category msg.MessageCategory `gorm:"column:category"`
	Provider msg.MessageProvider `gorm:"column:provider"`
	Value    int                 `gorm:"column:value"`
}

type DailySummary struct {
	app.TenantPk
	Date time.Time `gorm:"column:date;type:date;uniqueIndex"`
	Mail int       `gorm:"column:mail"`
	Sms  int       `gorm:"column:sms"`
	Im   int       `gorm:"column:im"`
	Apns int       `gorm:"column:apns"`
}
