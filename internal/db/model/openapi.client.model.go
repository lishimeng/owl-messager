package model

import "github.com/lishimeng/app-starter"

type OpenClient struct {
	app.TenantPk
	AppId      string `gorm:"column:app_id;uniqueIndex"`
	Secret     string `gorm:"column:secret"`
	BasicAuth  string `gorm:"column:basic_auth;uniqueIndex"`
	TenantCode string `gorm:"column:tenant_code"`
	Name       string `gorm:"column:name"`
	app.TableInfo
}

func (t OpenClient) GetID() string {
	return t.AppId
}

func (t OpenClient) GetSecret() string {
	return t.Secret
}
