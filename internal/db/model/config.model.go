package model

import "github.com/lishimeng/app-starter"

const ConfigCodeTaskChannel = "task.channel"

// TaskChannelSettings controls how messages are dispatched (memqueue vs db polling).
type TaskChannelSettings struct {
	Channel      string `json:"channel"`
	ScanInterval int    `json:"scanInterval"`
}

type Config struct {
	app.Pk
	Code    string `gorm:"column:code;uniqueIndex"`
	Content string `gorm:"column:content"`
	app.TableChangeInfo
}
