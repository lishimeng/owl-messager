package repo

import (
	"github.com/lishimeng/app-starter"
	"math"
)

type PagerReq struct {
	PageSize int
	PageNo   int
}

type PagerRes struct {
	// 页面大小
	PageSize int
	// 页号
	PageNo int
	// 总页数
	TotalPage int
	// 总条数
	Count int
	// 首页
	FirstPage bool
	// 尾页
	LastPage bool
}

func calcPageOffset(p app.Pager) int {
	return (p.PageNum - 1) * p.PageSize
}

func calcTotalPage(p app.Pager, count int64) int {
	t := math.Ceil(float64(count) / float64(p.PageSize))
	return int(t)
}

const (
	ConditionIgnore = 0
	DefaultPageNo   = 1
	DefaultPageSize = 10
)
