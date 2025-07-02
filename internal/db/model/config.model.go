package model

import "github.com/lishimeng/app-starter"

type Config struct {
	app.Pk
	Code    string `orm:"column(code)"`
	Content string `orm:"column(content)"`
	app.TableChangeInfo
}
