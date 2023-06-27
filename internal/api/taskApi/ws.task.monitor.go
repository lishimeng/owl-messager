package taskApi

import (
	"github.com/kataras/iris/v12/context"
	"github.com/kataras/iris/v12/websocket"
	"github.com/lishimeng/go-log"
	"github.com/lishimeng/owl-messager/internal/messager/task"
)

func TaskMonitorWs() context.Handler {

	ws := websocket.New(
		websocket.DefaultGorillaUpgrader,
		websocket.Events{},
	)

	ws.OnConnect = func(c *websocket.Conn) error {
		task.GetMonitor().Subscribe(c)
		return nil
	}

	ws.OnDisconnect = func(c *websocket.Conn) {
		task.GetMonitor().UnSubscribe(c)
		return
	}
	
	ws.OnUpgradeError = func(err error) {
		log.Info("Upgrade ws fail:%v", err)
	}

	return websocket.Handler(ws)
}
