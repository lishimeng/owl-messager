package service

import (
	"fmt"
	"github.com/lishimeng/app-starter"
	"github.com/lishimeng/go-log"
	persistence "github.com/lishimeng/go-orm"
	"github.com/lishimeng/owl/internal/db/model"
	"github.com/lishimeng/owl/internal/db/repo"
	"github.com/lishimeng/owl/internal/messager/msg"
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

func OnTaskHandleFail(task model.MessageTask) (err error) {
	log.Debug("on task handle failed")
	err = app.GetOrm().Transaction(func(ctx persistence.OrmContext) (e error) {
		// message status -> fail
		_, e = repo.UpdateMessageStatus(task.MessageId, model.MessageSendFailed)
		if e != nil {
			return
		}
		// task status -> fail
		_, e = repo.UpdateTaskStatus(task.Id, model.MessageTaskSendFailed)
		if e != nil {
			return
		}
		// delete running task
		e = repo.DeleteRunningTaskByTaskId(task.Id)
		return
	})

	return
}

func OnTaskHandleSuccess(task model.MessageTask) (err error) {
	log.Debug("on task handle success")
	err = app.GetOrm().Transaction(func(ctx persistence.OrmContext) (e error) {
		// message status -> success
		_, e = repo.UpdateMessageStatus(task.MessageId, model.MessageSendSuccess)
		if e != nil {
			return
		}
		// task status -> success
		_, e = repo.UpdateTaskStatus(task.Id, model.MessageTaskSendSuccess)
		if e != nil {
			return
		}
		// delete running success
		e = repo.DeleteRunningTaskByTaskId(task.Id)
		return
	})

	return
}

func CreateMessageTask(message model.MessageInfo, messageInstanceId int) (task model.MessageTask, err error) {

	err = app.GetOrm().Transaction(func(ctx persistence.OrmContext) (e error) {
		// 创建task
		task, e = repo.AddMessageTask(message.Id, messageInstanceId)
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
		// message status->sending
		_, e = repo.UpdateMessageStatus(message.Id, model.MessageSending)
		if e != nil {
			log.Info("change message [status -> sending] failed")
			return
		}
		return
	})

	return
}

func GetMessageInstanceId(message model.MessageInfo) (id int, err error) {
	switch message.Category {
	case msg.Email:
		var mail model.MailMessageInfo
		mail, err = repo.GetMailByMessageId(message.Id)
		if err == nil {
			id = mail.Id
		}
	case msg.Sms:
		var sms model.SmsMessageInfo
		sms, err = repo.GetSmsByMessageId(message.Id)
		if err == nil {
			id = sms.Id
		}
	default:
		log.Info("known message category:%d[message id:%d]", message.Category, message.Id)
		err = fmt.Errorf("known message category:%d[message id:%d]", message.Category, message.Id)
	}
	return
}
