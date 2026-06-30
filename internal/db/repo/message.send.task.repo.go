package repo

import (
	"github.com/lishimeng/app-starter/persistence"
	"github.com/lishimeng/owl-messager/internal/db/model"
	"github.com/lishimeng/owl-messager/pkg/msg"
	"time"
)

func AddMessageTask(ctx persistence.TxContext, messageId int, messageInstanceId int, category msg.MessageCategory) (task model.MessageTask, err error) {
	task = model.MessageTask{
		MessageId:         messageId,
		MessageInstanceId: messageInstanceId,
		Category:          category,
	}
	task.Status = model.MessageTaskInit
	err = ctx.Create(&task)
	return
}

func GetMessageTask(id int) (t model.MessageTask, err error) {
	err = orm().Model(&model.MessageTask{}).Equal("id", id).First(&t)
	return
}

func GetTaskByMessage(messageId int) (t model.MessageTask, err error) {
	err = orm().Model(&model.MessageTask{}).
		Equal("message_id", messageId).
		Order("-ctime").
		Limit(1).
		First(&t)
	return
}

func TaskSendFail(messageId int) {}

func CancelExpiredTask(taskId int) {}

func UpdateTaskStatus(ctx persistence.TxContext, taskId int, status int) (task model.MessageTask, err error) {
	task.Id = taskId
	task.Status = status
	task.UpdateTime = time.Now()
	err = updateSelect(ctx, &task, "Status", "UpdateTime")
	return
}

func GetExpiredTasks(size int, timeLatest time.Time) (tasks []model.MessageRunningTask, err error) {
	err = orm().Model(&model.MessageRunningTask{}).
		Where("ctime < ?", timeLatest).
		Limit(size).
		Find(&tasks)
	return
}

func AddRunningTask(ctx persistence.TxContext, task model.MessageTask) (runningTask model.MessageRunningTask, err error) {
	runningTask = model.MessageRunningTask{TaskId: task.Id}
	err = ctx.Create(&runningTask)
	return
}

func DeleteRunningTaskByTaskId(ctx persistence.TxContext, taskId int) (err error) {
	var runningTask model.MessageRunningTask
	err = ctx.Model(&model.MessageRunningTask{}).Equal("task_id", taskId).First(&runningTask)
	if err != nil {
		return
	}
	return ctx.Delete(&runningTask)
}

func DeleteRunningTask(ctx persistence.TxContext, id int) (err error) {
	return ctx.Delete(&model.MessageRunningTask{}, id)
}
