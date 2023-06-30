package model

import "github.com/lishimeng/app-starter"

type OpenClient struct {
	app.TenantPk
	AppId  string `orm:"column(app_id);unique"`
	Secret string `orm:"column(secret)"`
	Domain string `orm:"column(domain)"`
	app.TableChangeInfo
}

func (t OpenClient) GetID() string {
	return t.AppId
}

func (t OpenClient) GetSecret() string {
	return t.Secret
}

func (t OpenClient) GetDomain() string {
	return t.Domain
}
