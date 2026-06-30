package model

import "github.com/lishimeng/app-starter"

type OpenClient struct {
	app.TenantPk
	AppId      string `gorm:"column:app_id;uniqueIndex"`
	Secret     string `gorm:"column:secret"`
	TenantCode string `gorm:"column:domain"`
	Name       string `gorm:"column:name"`
	app.TableInfo
}

func (t OpenClient) GetID() string {
	return t.AppId
}

func (t OpenClient) GetSecret() string {
	return t.Secret
}

func (t OpenClient) GetDomain() string {
	return t.TenantCode
}
