package repo

import (
	"github.com/lishimeng/app-starter"
	"github.com/lishimeng/go-log"
	"github.com/lishimeng/owl/internal/db/model"
	"time"
)

func GetMessageById(id int) (m model.MessageInfo, err error) {
	log.Debug("get message from db: %d", id)
	m.Id = id
	err = app.GetOrm().Context.Read(&m)
	return
}

// 查询需要发送的消息
func GetMessageToSend(size int) (messages []model.MessageInfo, err error) {
	_, err = app.GetOrm().Context.
		QueryTable(new(model.MessageInfo)).
		Filter("Status", model.MessageInit).
		OrderBy("UpdateTime").Limit(size).All(&messages)
	return
}

func GetMessages(status int, category int, page app.Pager) (p app.Pager, err error) {
	var messages []model.MessageInfo
	qs := app.GetOrm().Context.QueryTable(new(model.MessageInfo))
	if status > ConditionIgnore {
		qs = qs.Filter("Status", status)
	}
	if category > ConditionIgnore {
		qs = qs.Filter("Category", category)
	}
	sum, err := qs.Count()
	if err != nil {
		return
	}
	page.TotalPage = calcTotalPage(page, sum)
	_, err = qs.OrderBy("CreateTime").Offset(calcPageOffset(page)).Limit(page.PageSize).All(&messages)
	if err != nil {
		return
	}
	if len(messages) > 0 {
		for _, m := range messages {
			page.Data = append(page.Data, m)
		}
	}
	p = page
	return
}

func UpdateMessageStatus(id int, status int) (m model.MessageInfo, err error) {
	m.Id = id
	m.Status = status
	m.UpdateTime = time.Now()
	_, err = app.GetOrm().Context.Update(&m, "Status")
	return
}

func UpdateMessagePriority(id int, priority int) (m model.MessageInfo, err error) {
	m.Id = id
	m.Priority = priority
	m.UpdateTime = time.Now()
	_, err = app.GetOrm().Context.Update(&m, "Priority")
	return
}

func CreateMessage(subject string, category int) (m model.MessageInfo, err error) {
	log.Debug("create message %s[category:%d]", subject, category)
	m.Subject = subject
	m.Priority = model.MessagePriorityNormal
	m.Category = category
	m.Status = model.MessageInit
	_, err = app.GetOrm().Context.Insert(&m)
	return
}
