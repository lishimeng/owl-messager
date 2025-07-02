package repo

import (
	"encoding/base64"
	"encoding/json"
	"github.com/lishimeng/app-starter"
	"github.com/lishimeng/owl-messager/internal/db/model"
)

func SaveConfig(code string, content interface{}) error {
	js, err := json.Marshal(content)
	if err != nil {
		return err
	}
	encoded := base64.StdEncoding.EncodeToString(js)

	cfg, err := GetOneConfig(code)
	if err == nil {
		// 查询得到结果，已经存在此条目，改为update
		cfg.Content = encoded
		_, err = app.GetOrm().Context.Update(&cfg)
	} else {
		// 不存在
		cfg = model.Config{
			Content: encoded,
			Code:    code,
		}
		cfg.Status = 10
		_, err = app.GetOrm().Context.Insert(&cfg)
	}
	return err
}

func GetOneConfig(code string) (config model.Config, err error) {
	err = app.GetOrm().Context.QueryTable(new(model.Config)).
		Filter("code", code).
		One(&config)
	return
}
