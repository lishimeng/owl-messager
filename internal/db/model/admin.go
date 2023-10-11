package model

import "github.com/lishimeng/app-starter"

// Admin
// 管理员系统账号
type Admin struct {
	app.Pk
	Username string
	Password string
}
