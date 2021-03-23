package repo

import (
	"github.com/lishimeng/app-starter"
	"github.com/lishimeng/owl/internal/db/model"
	"time"
)

// 创建投送task
// 从未投送的message中取出一个
func AddMessageTask(messageId int, messageInstanceId int) (task model.MessageTask, err error) {
	task = model.MessageTask{
		MessageId:         messageId,
		MessageInstanceId: messageInstanceId,
		TableChangeInfo: model.TableChangeInfo{
			Status:     model.MessageTaskInit, // TODO
			CreateTime: time.Now(),
			UpdateTime: time.Now(),
		},
	}
	_, err = app.GetOrm().Context.Insert(&task)
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

func UpdateTaskStatus(taskId int, status int) (task model.MessageTask, err error) {
	task.Id = taskId
	task.Status = status
	task.UpdateTime = time.Now()
	_, err = app.GetOrm().Context.Update(&task, "Status")
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

func AddRunningTask(task model.MessageTask) (runningTask model.MessageRunningTask, err error) {
	runningTask = model.MessageRunningTask{
		TaskId: task.Id,
		TableInfo: model.TableInfo{
			CreateTime: time.Now(),
		},
	}
	_, err = app.GetOrm().Context.Insert(&runningTask)
	return
}

func DeleteRunningTaskByTaskId(taskId int) (err error) {
	var runningTask model.MessageRunningTask
	err = app.GetOrm().Context.QueryTable(new(model.MessageRunningTask)).Filter("TaskId", taskId).One(&runningTask)
	if err != nil {
		return
	}
	_, err = app.GetOrm().Context.Delete(&runningTask)
	return
}

func DeleteRunningTask(id int) (err error) {
	var runningTask model.MessageRunningTask
	runningTask.Id = id
	_, err = app.GetOrm().Context.Delete(runningTask)
	return
}
