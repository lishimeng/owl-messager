package model

import "time"

type Pk struct {
	// ID
	Id int
}

type TableChangeInfo struct {
	// 状态
	Status int
	// 创建时间
	CreateTime time.Time
	// 修改时间
	UpdateTime time.Time
}

type TableInfo struct {
	// 创建时间
	CreateTime time.Time
}