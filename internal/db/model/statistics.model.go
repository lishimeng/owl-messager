package model

import (
	"time"

	"github.com/lishimeng/app-starter"
	"github.com/lishimeng/owl-messager/pkg/msg"
)

type ProviderStats struct {
	TenantScope `gorm:"uniqueIndex:idx_provider_stat"`
	Category    msg.MessageCategory `gorm:"column:category;uniqueIndex:idx_provider_stat"`
	Provider    msg.MessageProvider `gorm:"column:provider;uniqueIndex:idx_provider_stat"`
	Value    int                 `gorm:"column:value"`
	app.TableChangeInfo
}

type DailySummary struct {
	TenantScope `gorm:"uniqueIndex:idx_daily_tenant_date"`
	Date        time.Time `gorm:"column:date;type:date;uniqueIndex:idx_daily_tenant_date"`
	Mail int       `gorm:"column:mail"`
	Sms  int       `gorm:"column:sms"`
	Im   int       `gorm:"column:im"`
	Apns int       `gorm:"column:apns"`
}
