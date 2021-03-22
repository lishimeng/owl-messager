package model

import "time"

type Pk struct {
	// ID
	Id int `orm:"pk;column(id)"`
}

type TableChangeInfo struct {
	// 状态
	Status int `orm:"column(status)"`
	// 创建时间
	CreateTime time.Time `orm:"column(ctime)"`
	// 修改时间
	UpdateTime time.Time `orm:"column(mtime)"`
}

type TableInfo struct {
	// 创建时间
	CreateTime time.Time `orm:"column(ctime)"`
}