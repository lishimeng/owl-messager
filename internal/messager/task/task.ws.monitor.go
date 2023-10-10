package task

import (
	"context"
	"encoding/json"
	"github.com/kataras/iris/v12/websocket"
	"github.com/lishimeng/go-log"
	"github.com/lishimeng/owl-messager/pkg/msg"
	"sync"
	"time"
)

type MonitorData struct {
	TaskId            int                 `json:"taskId,omitempty"`
	MessageId         int                 `json:"messageId,omitempty"`
	MessageCategory   msg.MessageCategory `json:"messageCategory,omitempty"`
	Subject           string              `json:"subject,omitempty"`
	MessageInstanceId int                 `json:"messageInstanceId,omitempty"`
	Status            int                 `json:"status,omitempty"`
}

type SendMonitor interface {
	Subscribe(c *websocket.Conn)
	UnSubscribe(c *websocket.Conn)
	Pub(msg MonitorData)
}

type sendMonitor struct {
	subscribers map[string]*websocket.Conn
	messageChan chan MonitorData
	ctx         context.Context

	mux sync.Locker
}

var singleton SendMonitor

func InitMonitor(ctx context.Context) {
	s := new(sendMonitor)
	s.ctx = ctx
	s.messageChan = make(chan MonitorData, 10)
	s.subscribers = make(map[string]*websocket.Conn)
	s.mux = new(sync.Mutex)
	singleton = s
	go s.pubLoop()
}

func GetMonitor() SendMonitor {
	return singleton
}

func (s *sendMonitor) Subscribe(c *websocket.Conn) {
	s.mux.Lock()
	defer s.mux.Unlock()
	s.subscribers[c.ID()] = c
}

func (s *sendMonitor) UnSubscribe(c *websocket.Conn) {
	s.mux.Lock()
	defer s.mux.Unlock()
	key := c.ID()
	if _, ok := s.subscribers[key]; ok {
		delete(s.subscribers, key)
	}
}

func (s *sendMonitor) pubLoop() {
	for {
		select {
		case <-s.ctx.Done():
			return
		case d := <-s.messageChan:
			s.pub(d)
		}
	}
}

func (s *sendMonitor) pub(msg MonitorData) {
	s.mux.Lock()
	defer s.mux.Unlock()
	bs, err := json.Marshal(msg)
	if err != nil {
		return
	}
	for id, c := range s.subscribers {
		log.Info("task monitor pub to: %s", id)
		if c.IsClosed() {
			log.Info("task monitor client[%s] closed skip", id)
			continue
		}
		_ = c.Socket().WriteText(bs, time.Second*8)
	}
}

func (s *sendMonitor) Pub(msg MonitorData) {
	s.messageChan <- msg
}
