package service

import (
	"github.com/lishimeng/app-starter"
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
