package repo

import (
	"github.com/lishimeng/app-starter"
	persistence "github.com/lishimeng/go-orm"
	"github.com/lishimeng/owl-messager/internal/db/model"
)

// 查询邮件Template
func GetMailTemplateByCode(code string) (s model.MailTemplateInfo, err error) {
	err = app.GetOrm().Context.QueryTable(new(model.MailTemplateInfo)).Filter("Code", code).One(&s)
	return
}

// 查询邮件Template
func GetMailTemplateById(id int) (s model.MailTemplateInfo, err error) {
	s.Id = id
	err = app.GetOrm().Context.Read(&s)
	return
}

// 查询邮件Template列表
func GetMailTemplateList(status int, page app.Pager) (p app.Pager, tpls []model.MailTemplateInfo, err error) {
	var qs = app.GetOrm().Context.QueryTable(new(model.MailTemplateInfo))
	if status > ConditionIgnore {
		qs = qs.Filter("Status", status)
	}
	sum, err := qs.Count()
	if err != nil {
		return
	}
	page.TotalPage = calcTotalPage(page, sum)
	_, err = qs.OrderBy("CreateTime").Offset(calcPageOffset(page)).Limit(page.PageSize).All(&tpls)
	if err != nil {
		return
	}
	p = page
	return
}

func DeleteMailTemplate(id int) (err error) {
	var t model.MailTemplateInfo
	t.Id = id
	_, err = app.GetOrm().Context.Delete(&t)
	return
}

func CreateMailTemplate(code, name, body, description string, category int) (m model.MailTemplateInfo, err error) {
	m = model.MailTemplateInfo{
		Code:     code,
		Name:     name,
		Body:     body,
		Category: category,
	}
	if len(description) > 0 {
		m.Description = description
	}
	m.Status = model.MailTemplateEnable
	_, err = app.GetOrm().Context.Insert(&m)

	return
}
func CreateMailTemplateNew(code, name, body, description, vendor string, category int) (m model.MailTemplateInfo, err error) {
	m = model.MailTemplateInfo{
		Code:     code,
		Name:     name,
		Body:     body,
		Category: category,
		Vendor:   vendor,
	}
	if len(description) > 0 {
		m.Description = description
	}
	m.Status = model.MailTemplateEnable
	_, err = app.GetOrm().Context.Insert(&m)

	return
}

func UpdateMailTemplate(ctx persistence.TxContext, ori model.MailTemplateInfo, cols ...string) (m model.MailTemplateInfo, err error) {
	_, err = ctx.Context.Update(&ori, cols...)
	m = ori
	return
}

func UpdateMailTemplateInfo(ori model.MailTemplateInfo, cols ...string) (m model.MailTemplateInfo, err error) {
	_, err = app.GetOrm().Context.Update(&ori, cols...)
	m = ori
	return
}
