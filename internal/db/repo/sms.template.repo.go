package repo

import (
	"github.com/lishimeng/app-starter"
	"github.com/lishimeng/owl/internal/db/model"
)

// GetSmsTemplateByCode 查询Template
func GetSmsTemplateByCode(code string) (s model.SmsTemplateInfo, err error) {
	err = app.GetOrm().Context.QueryTable(new(model.SmsTemplateInfo)).Filter("Code", code).One(&s)
	return
}

// GetSmsTemplateById 查询Template
func GetSmsTemplateById(id int) (s model.SmsTemplateInfo, err error) {
	s.Id = id
	err = app.GetOrm().Context.Read(&s)
	return
}

// GetSmsTemplateList 查询Template列表
func GetSmsTemplateList(status int, page app.Pager) (p app.Pager, tpls []model.SmsTemplateInfo, err error) {
	var qs = app.GetOrm().Context.QueryTable(new(model.SmsTemplateInfo))
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

func DeleteSmsTemplate(id int) (err error) {
	var t model.SmsTemplateInfo
	t.Id = id
	_, err = app.GetOrm().Context.Delete(&t)
	return
}

func CreateSmsTemplate(code, name, body, description string) (m model.SmsTemplateInfo, err error) {
	m = model.SmsTemplateInfo{
		Code: code,
		Name: name,
		Body: body,
		// TODO
	}
	if len(description) > 0 {
		m.Description = description
	}
	m.Status = model.MailTemplateEnable
	_, err = app.GetOrm().Context.Insert(&m)

	return
}

func UpdateSmsTemplate(ori model.SmsTemplateInfo, cols ...string) (m model.SmsTemplateInfo, err error) {
	_, err = app.GetOrm().Context.Update(&ori, cols...)
	m = ori
	return
}
