package model

import "github.com/lishimeng/app-starter"

type OpenClient struct {
	TenantScope
	AppId  string `gorm:"column:app_id;uniqueIndex"`
	Secret string `gorm:"column:secret"`
	Name   string `gorm:"column:name"`
	app.TableInfo
}

func (OpenClient) TableName() string {
	return "open_client"
}

func (t OpenClient) GetID() string {
	return t.AppId
}

func (t OpenClient) GetSecret() string {
	return t.Secret
}
