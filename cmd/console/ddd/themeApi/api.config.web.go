package themeApi

import (
	"encoding/base64"
	"encoding/json"
	"github.com/lishimeng/app-starter"
	"github.com/lishimeng/app-starter/server"
	"github.com/lishimeng/app-starter/tool"
	"github.com/lishimeng/owl-messager/internal/db/repo"
)

type respWebConfig struct {
	app.Response
	Config themeConfig `json:"themeConfig"`
}

func saveConfig() error {
	return repo.SaveConfig(defaultTheme, defaultConfig)
}

func GetThemeConfig(ctx server.Context) {
	var resp respWebConfig
	var config themeConfig

	cfg, err := repo.GetOneConfig(defaultTheme)
	if err != nil {
		resp.Message = err.Error()
		resp.Code = tool.RespCodeNotFound
		ctx.Json(resp)
		return
	}

	js, err := base64.StdEncoding.DecodeString(cfg.Content)
	if err != nil {
		resp.Message = err.Error()
		resp.Code = tool.RespCodeError
		ctx.Json(resp)
		return
	}

	err = json.Unmarshal(js, &config)
	if err != nil {
		resp.Message = err.Error()
		resp.Code = tool.RespCodeError
		ctx.Json(resp)
		return
	}
	resp.Code = tool.RespCodeSuccess
	resp.Config = config
	ctx.Json(resp)
}
