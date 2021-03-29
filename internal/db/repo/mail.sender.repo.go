package repo

import (
	"github.com/lishimeng/app-starter"
	"github.com/lishimeng/owl/internal/db/model"
)

// 查询邮件发送账号
func GetMailSenderByCode(code string) (s model.MailSenderInfo, err error) {
	err = app.GetOrm().Context.QueryTable(new(model.MailSenderInfo)).Filter("Code", code).One(&s)
	return
}

// 查询邮件发送账号
func GetMailSenderById(id int) (s model.MailSenderInfo, err error) {
	s.Id = id
	err = app.GetOrm().Context.Read(&s)
	return
}

// 查询邮件发送账号列表
func GetMailSenderList(status int, page app.Pager) (p app.Pager, err error) {
	var senders []model.MailSenderInfo
	var qs = app.GetOrm().Context.QueryTable(new(model.MailSenderInfo))
	if status > ConditionIgnore {
		qs = qs.Filter("Status", status)
	}
	sum, err := qs.Count()
	if err != nil {
		return
	}
	page.TotalPage =calcTotalPage(page, sum)
	_, err = qs.OrderBy("CreateTime").Offset(calcPageOffset(page)).Limit(page.PageSize).All(&senders)
	if err != nil {
		return
	}
	if len(senders) > 0 {
		for _, s := range senders {
			page.Data = append(page.Data, s)
		}
	}
	p = page
	return
}
