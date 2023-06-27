package repo

import (
	"github.com/lishimeng/app-starter"
	"github.com/lishimeng/owl-messager/internal/db/model"
)

// GetApnsSenderByCode 查询发送账号
func GetApnsSenderByCode(code string) (s model.ApnsSenderInfo, err error) {
	err = app.GetOrm().Context.QueryTable(new(model.ApnsSenderInfo)).Filter("Code", code).One(&s)
	return
}

// GetApnsSenderById 查询发送账号
func GetApnsSenderById(id int) (s model.ApnsSenderInfo, err error) {
	s.Id = id
	err = app.GetOrm().Context.Read(&s)
	return
}

func DeleteApnsSender(id int) (err error) {
	var t model.ApnsSenderInfo
	t.Id = id
	_, err = app.GetOrm().Context.Delete(&t)
	return
}

// GetApnsSenderList 查询发送账号列表
func GetApnsSenderList(status int, page app.Pager) (p app.Pager, senders []model.ApnsSenderInfo, err error) {
	var qs = app.GetOrm().Context.QueryTable(new(model.ApnsSenderInfo))
	if status > ConditionIgnore {
		qs = qs.Filter("Status", status)
	}
	sum, err := qs.Count()
	if err != nil {
		return
	}
	page.TotalPage = calcTotalPage(page, sum)
	_, err = qs.OrderBy("CreateTime").Offset(calcPageOffset(page)).Limit(page.PageSize).All(&senders)
	if err != nil {
		return
	}
	p = page
	return
}
