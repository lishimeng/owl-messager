package model

import "github.com/lishimeng/app-starter"

type Config struct {
	app.Pk
	Code    string `gorm:"column:code"`
	Content string `gorm:"column:content"`
	app.TableChangeInfo
}
