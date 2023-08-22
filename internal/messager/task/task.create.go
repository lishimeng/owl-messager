package task

import (
	"context"
	"github.com/lishimeng/go-log"
	"github.com/lishimeng/owl-messager/internal/db/model"
	"github.com/lishimeng/owl-messager/internal/db/repo"
	"github.com/lishimeng/owl-messager/internal/db/service"
	"github.com/lishimeng/owl-messager/internal/messager/sender"
	"math"
	"time"
)

type MessageTask interface {
	Run()
	HandleMessage(message model.MessageInfo) (err error)
}

type Strategy int

const (
	MemQueue Strategy = iota
	Db
)

const (
	DefaultQueueSize      = 1024
	DefaultDbScanInterval = 10
)

type messageTask struct {
	running        bool
	ctx            context.Context
	executor       sender.TaskExecutor
	ch             chan model.MessageInfo
	Strategy       Strategy
	queueSize      int // 队列长度
	scanDbInterval int // 查询数据库频率(秒)
}

type Option func(task *messageTask)

func WithQueue(queueSize int) Option {
	return func(mt *messageTask) {
		log.Debug("Strategy:%s[%d]", "MemQueue", queueSize)
		mt.queueSize = queueSize
		mt.Strategy = MemQueue
		if mt.queueSize <= 0 {
			log.Debug("use default queue size:%d", DefaultQueueSize)
			mt.queueSize = DefaultQueueSize
		}
	}
}

func WithDb(scanInternal int) Option {
	return func(mt *messageTask) {
		log.Debug("Strategy:%s[%d]", "Db", scanInternal)
		mt.scanDbInterval = scanInternal
		mt.Strategy = Db
		if mt.scanDbInterval <= 0 {
			log.Debug("use default db scan interval:%d", DefaultDbScanInterval)
			mt.scanDbInterval = DefaultDbScanInterval
		}
	}
}

func New(ctx context.Context, executor sender.TaskExecutor, opts ...Option) (t MessageTask, err error) {
	handler := &messageTask{ctx: ctx, executor: executor}
	for _, opt := range opts {
		opt(handler)
	}
	t = handler
	return
}

func (t *messageTask) Run() {
	if t.running {
		return
	}
	go t.loop()
}

func (t *messageTask) HandleMessage(message model.MessageInfo) (err error) {

	switch t.Strategy {
	case MemQueue:
		t.ch <- message
	}
	return
}

func (t *messageTask) loop() {
	for {
		select {
		case <-t.ctx.Done():
			return
		default:
			switch t.Strategy {
			case Db:
				t.sendFromDb()
			case MemQueue:
				t.sendFromQueue()
			default:
				t.sendFromQueue()
			}
		}
	}
}

func (t *messageTask) sendFromDb() {
	messages, err := t.getMessages(1)
	if err != nil {
		log.Info("get message failed")
		log.Info(err)
	} else {
		t.handleMessages(messages...)
	}
	if len(messages) == 0 {
		time.Sleep(time.Second * time.Duration(t.scanDbInterval))
	} else {
		time.Sleep(time.Second * time.Duration(math.Round(float64(t.scanDbInterval))))
	}
}

func (t *messageTask) sendFromQueue() {
	log.Debug("init message queue")

	t.ch = make(chan model.MessageInfo, t.queueSize)
	for {
		select {
		case <-t.ctx.Done():
			return
		case m, ok := <-t.ch:
			if !ok {
				log.Info("message queue closed")
				return
			}
			t.handleMessages(m)
		}
	}
}

func (t *messageTask) getMessages(size int) (messages []model.MessageInfo, err error) {
	messages, err = repo.GetMessageToSend(size)
	return
}

func (t *messageTask) handleMessages(messages ...model.MessageInfo) {
	for _, message := range messages {
		e := t.handleMessage(message)
		if e != nil {
			log.Info("handle message failed")
			log.Info(e)
		}
	}
}

func (t *messageTask) handleMessage(message model.MessageInfo) (err error) {
	instanceId, err := service.GetMessageInstanceId(message)
	if err != nil {
		log.Info("get message instance id failed")
		return
	}
	task, err := service.CreateMessageTask(message, instanceId)
	if err != nil {
		log.Info("create message task failed")
		log.Info(err)
		return
	}
	GetMonitor().Pub(MonitorData{
		TaskId:            task.Id,
		MessageId:         message.Id,
		MessageCategory:   message.Category,
		Subject:           message.Subject,
		MessageInstanceId: task.MessageInstanceId,
		Status:            task.Status,
	})
	// 提交给发送器
	err = t.executor.Execute(task)
	AfterHandleTask(task, message, err)
	return
}

func AfterHandleTask(task model.MessageTask, message model.MessageInfo, err error) {
	pubData := MonitorData{
		TaskId:            task.Id,
		MessageId:         message.Id,
		MessageCategory:   message.Category,
		Subject:           message.Subject,
		MessageInstanceId: task.MessageInstanceId,
	}
	if err != nil {
		err = service.OnTaskHandleFail(task)
		pubData.Status = model.MessageTaskSendFailed
	} else {
		err = service.OnTaskHandleSuccess(task)
		pubData.Status = model.MessageTaskSendSuccess
	}

	GetMonitor().Pub(pubData)
}
