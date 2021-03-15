package repo

type PagerReq struct {
	PageSize int
	PageNo   int
}

type PagerRes struct {
	// 页面大小
	PageSize int
	// 页号
	PageNo   int
	// 总页数
	TotalPage int
	// 总条数
	Count int
	// 首页
	FirstPage bool
	// 尾页
	LastPage bool
}