package model

import "github.com/lishimeng/app-starter"

// TenantScope 租户隔离字段，使用 tenant.code（便于迁移，不依赖自增 org id）
type TenantScope struct {
	app.Pk
	TenantCode string `gorm:"column:tenant_code;index"`
}
