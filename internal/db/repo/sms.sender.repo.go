package repo

import (
	"github.com/lishimeng/app-starter"
	persistence "github.com/lishimeng/go-orm"
	"github.com/lishimeng/owl-messager/internal/db/model"
)

// GetSmsSenderByCode 查询短信发送账号
func GetSmsSenderByCode(code string) (s model.SmsSenderInfo, err error) {
	err = app.GetOrm().Context.QueryTable(new(model.SmsSenderInfo)).Filter("Code", code).One(&s)
	return
}

// GetSmsSenderById 查询短信发送账号
func GetSmsSenderById(id int) (s model.SmsSenderInfo, err error) {
	s.Id = id
	err = app.GetOrm().Context.Read(&s)
	return
}

func GetDefaultSmsSender(org int) (s model.SmsSenderInfo, err error) {

	// TODO org
	err = app.GetOrm().Context.
		QueryTable(new(model.SmsSenderInfo)).
		Filter("Default", model.DefaultSenderEnable).
		One(&s)
	return
}

func DeleteSmsSender(id int) (err error) {
	var t model.SmsSenderInfo
	t.Id = id
	_, err = app.GetOrm().Context.Delete(&t)
	return
}

// GetSmsSenderList 查询短信发送账号列表
func GetSmsSenderList(status int, page app.Pager) (p app.Pager, senders []model.SmsSenderInfo, err error) {
	var qs = app.GetOrm().Context.QueryTable(new(model.SmsSenderInfo))
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

func GetSmsSenders(status int) (senders []model.SmsSenderInfo, err error) {
	_, err = app.GetOrm().Context.
		QueryTable(new(model.SmsSenderInfo)).
		Filter("Status", status).
		All(&senders)
	return
}

func SmsSenderEnable(id int) (err error) {
	var s model.SmsSenderInfo
	s.Id = id
	err = app.GetOrm().Transaction(func(context persistence.TxContext) (e error) {
		e = context.Context.Read(&s)
		if e != nil {
			return
		}
		s.Status = model.SmsSenderEnable
		_, e = context.Context.Update(&s, "Status")
		return
	})

	return
}

func SmsSenderDisable(id int) (err error) {
	var s model.SmsSenderInfo
	s.Id = id
	err = app.GetOrm().Transaction(func(context persistence.TxContext) (e error) {
		e = context.Context.Read(&s)
		if e != nil {
			return
		}
		s.Status = model.SmsSenderDisable
		_, e = context.Context.Update(&s, "Status")
		return
	})

	return
}

// 创建配置内容
func CreateSmslSenderInfo(code, vendor, config string, defaultSender int) (m model.SmsSenderInfo, err error) {
	m.Status = 1
	m.Vendor = model.SmsVendor(vendor)
	m.Config = config
	m.Code = code
	m.Default = defaultSender
	_, err = app.GetOrm().Context.Insert(&m)
	return
}

// 编辑配置内容
func UpdateSmsSenderInfo(ori model.SmsSenderInfo, cols ...string) (m model.SmsSenderInfo, err error) {
	_, err = app.GetOrm().Context.Update(&ori, cols...)
	m = ori
	return
}
