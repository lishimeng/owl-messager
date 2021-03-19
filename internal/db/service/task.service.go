package service

import (
	"github.com/lishimeng/app-starter"
	"github.com/lishimeng/go-log"
	persistence "github.com/lishimeng/go-orm"
	"github.com/lishimeng/owl/internal/db/model"
	"github.com/lishimeng/owl/internal/db/repo"
)

func HandleExpiredTask(runningTask model.MessageRunningTask) (err error) {
	// start transmission
	err = app.GetOrm().Transaction(func(ctx persistence.OrmContext) (e error) {
		// remove from running runningTask
		e = repo.DeleteRunningTask(runningTask.Id)
		if e != nil {
			return
		}
		// set runningTask status expired
		task, e := repo.UpdateTaskStatus(runningTask.TaskId, model.MessageTaskSendExpired)
		if e != nil {
			return
		}
		// set message status expired
		_, e = repo.UpdateMessageStatus(task.MessageId, model.MessageSendExpired)
		// TODO add log

		return e
	})
	return
}

func OnTaskHandleFail(taskId int) (err error) {
	return
}

func CreateMessageTask(message model.MessageInfo, messageInstanceId int) (task model.MessageTask, err error) {

	err = app.GetOrm().Transaction(func(ctx persistence.OrmContext) (e error) {
		// 创建task
		task, e := repo.AddMessageTask(message.Id, messageInstanceId)
		if e != nil {
			log.Info("create message task failed")
			return
		}
		// task添加进running task表
		runningTask, e := repo.AddRunningTask(task)
		if e != nil {
			log.Info("create message running task failed")
			return
		}
		log.Info("running task create success [%d]", runningTask.Id)
		return
	})

	return
}
