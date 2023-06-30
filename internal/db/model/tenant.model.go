package model

import "github.com/lishimeng/app-starter"

// Tenant 租户
type Tenant struct {
	app.Pk
	Code string `orm:"column(code);unique"`
	Name string `orm:"column(name)"`
	app.TableChangeInfo
}
