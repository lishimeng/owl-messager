package repo

import (
	"github.com/lishimeng/app-starter"
	persistence "github.com/lishimeng/go-orm"
	"github.com/lishimeng/owl-messager/internal/db/model"
	"time"
)

// 创建投送task
// 从未投送的message中取出一个
func AddMessageTask(ctx persistence.TxContext, messageId int, messageInstanceId int) (task model.MessageTask, err error) {
	task = model.MessageTask{
		MessageId:         messageId,
		MessageInstanceId: messageInstanceId,
	}
	task.Status = model.MessageTaskInit
	_, err = ctx.Context.Insert(&task)
	return
}

func GetMessageTask(id int) (t model.MessageTask, err error) {
	t.Id = id
	err = app.GetOrm().Context.Read(&t)
	return
}

func GetTaskByMessage(messageId int) (t model.MessageTask, err error) {

	err = app.GetOrm().Context.QueryTable(new(model.MessageTask)).
		Filter("MessageId", messageId).
		OrderBy("-CreateTime").
		Limit(1).
		One(&t)
	return
}

func TaskSendFail(messageId int) {

}

// 取消超时的task
func CancelExpiredTask(taskId int) {
	// 从running删除
	// task表设置状态(超时)
	// message表设置状态(超时)
	// TODO 是否用数据库函数执行
}

func UpdateTaskStatus(ctx persistence.TxContext, taskId int, status int) (task model.MessageTask, err error) {
	task.Id = taskId
	task.Status = status
	task.UpdateTime = time.Now()
	_, err = ctx.Context.Update(&task, "Status")
	return
}

// 超时的列表
// size:取出数据量
func GetExpiredTasks(size int, timeLatest time.Time) (tasks []model.MessageRunningTask, err error) {
	_, err = app.GetOrm().Context.
		QueryTable(new(model.MessageRunningTask)).
		Filter("CreateTime__lt", timeLatest).
		Limit(size).
		All(&tasks)
	return
}

func AddRunningTask(ctx persistence.TxContext, task model.MessageTask) (runningTask model.MessageRunningTask, err error) {
	runningTask = model.MessageRunningTask{
		TaskId: task.Id,
	}
	_, err = ctx.Context.Insert(&runningTask)
	return
}

func DeleteRunningTaskByTaskId(ctx persistence.TxContext, taskId int) (err error) {
	var runningTask model.MessageRunningTask
	err = ctx.Context.QueryTable(new(model.MessageRunningTask)).Filter("TaskId", taskId).One(&runningTask)
	if err != nil {
		return
	}
	_, err = ctx.Context.Delete(&runningTask)
	return
}

func DeleteRunningTask(ctx persistence.TxContext, id int) (err error) {
	var runningTask model.MessageRunningTask
	runningTask.Id = id
	_, err = ctx.Context.Delete(runningTask)
	return
}

func GetTaskList(status int, page app.Pager) (p app.Pager, list []model.MessageTask, err error) {
	var qs = app.GetOrm().Context.QueryTable(new(model.MessageTask))
	if status > ConditionIgnore {
		qs = qs.Filter("Status", status)
	}
	all, err := qs.Count()
	if err != nil {
		return
	}
	page.TotalPage = calcTotalPage(page, all)
	qs = qs.Offset(calcPageOffset(page)).Limit(page.PageSize)
	qs = qs.OrderBy("CreateTime")

	_, err = qs.All(&list)
	p = page
	return
}
