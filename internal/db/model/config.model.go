package model

import "github.com/lishimeng/app-starter"

const ConfigCodeTaskChannel = "task.channel"
const ConfigCodeConsoleToken = "console.token"

// TaskChannelSettings controls how messages are dispatched (memqueue vs db polling).
type TaskChannelSettings struct {
	Channel      string `json:"channel"`
	ScanInterval int    `json:"scanInterval"`
}

// ConsoleTokenSettings holds the management console Bearer token.
type ConsoleTokenSettings struct {
	Token string `json:"token"`
}

type Config struct {
	app.Pk
	Code    string `gorm:"column:code;uniqueIndex"`
	Content string `gorm:"column:content"`
	app.TableChangeInfo
}

func (Config) TableName() string {
	return "config"
}
