package repo

import (
	"github.com/lishimeng/app-starter"
	"github.com/lishimeng/owl-messager/internal/db/model"
)

// GetMailSenderByCode 查询邮件发送账号
func GetMailSenderByCode(code string) (s model.MailSenderInfo, err error) {
	err = app.GetOrm().Context.QueryTable(new(model.MailSenderInfo)).Filter("Code", code).One(&s)
	return
}

// GetMailSenderById 查询邮件发送账号
func GetMailSenderById(id int) (s model.MailSenderInfo, err error) {
	s.Id = id
	err = app.GetOrm().Context.Read(&s)
	return
}

func GetDefaultMailSender(org string) (s model.MailSenderInfo, err error) {
	err = app.GetOrm().Context.
		QueryTable(new(model.MailSenderInfo)).
		Filter("Default", model.DefaultSenderEnable).
		One(&s)
	return
}

func DeleteMailSender(id int) (err error) {
	var t model.MailSenderInfo
	t.Id = id
	_, err = app.GetOrm().Context.Delete(&t)
	return
}

// GetMailSenderList 查询邮件发送账号列表
func GetMailSenderList(status int, page app.Pager) (p app.Pager, senders []model.MailSenderInfo, err error) {
	var qs = app.GetOrm().Context.QueryTable(new(model.MailSenderInfo))
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

func GetMailSenders(org int) (senders []model.MailSenderInfo, err error) {
	var qs = app.GetOrm().Context.QueryTable(new(model.MailSenderInfo))

	if err != nil {
		return
	}
	_, err = qs.OrderBy("CreateTime").
		// ORG filter
		All(&senders)
	if err != nil {
		return
	}
	return
}

// 创建邮件箱配置内容
func CreateMailSenderInfo(code, vendor, config string, defaultSender int) (m model.MailSenderInfo, err error) {
	m.Status = 1
	m.Vendor = model.MailVendor(vendor)
	m.Config = config
	m.Code = code
	m.Default = defaultSender
	_, err = app.GetOrm().Context.Insert(&m)
	return
}

// 编辑邮件箱配置内容
func UpdateMailSenderInfo(ori model.MailSenderInfo, cols ...string) (m model.MailSenderInfo, err error) {
	_, err = app.GetOrm().Context.Update(&ori, cols...)
	m = ori
	return
}
