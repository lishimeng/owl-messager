package repo

import (
	"github.com/lishimeng/app-starter"
	"github.com/lishimeng/owl/internal/db/model"
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
func GetMailTemplateList(status int, page app.Pager) (p app.Pager, err error) {
	var tpls []model.MailTemplateInfo
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
	if len(tpls) > 0 {
		for _, tpl := range tpls {
			page.Data = append(page.Data, tpl)
		}
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
