package repo

import "github.com/lishimeng/owl/internal/db/model"

// 创建投送task
// 从未投送的message中取出一个
func AddMessageTask(messageId int) (task model.MessageTask,err error) {
	// 创建task
	// task添加进running task表
	// task提交给发送器
	return
}

// 取消超时的task
func CancelExpiredTask(taskId int) {
	// 从running删除
	// task表设置状态(超时)
	// message表设置状态(超时)
	// TODO 是否用数据库函数执行
}

// 超时的列表
// size:取出数据量
func GetExpiredTasks(size int) (tasks []model.MessageRunningTask, err error) {
	return
}