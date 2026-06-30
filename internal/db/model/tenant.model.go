package model

import "github.com/lishimeng/app-starter"

// Tenant 租户
type Tenant struct {
	app.Pk
	Code string `gorm:"column:code;uniqueIndex"`
	Name string `gorm:"column:name"`
	app.TableChangeInfo
}
