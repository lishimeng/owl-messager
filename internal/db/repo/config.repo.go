package repo

import (
	"encoding/base64"
	"encoding/json"

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
		cfg.Content = encoded
		err = orm().Model(&cfg).Select("Content").Updates(&cfg)
	} else {
		cfg = model.Config{
			Content: encoded,
			Code:    code,
		}
		cfg.Status = 10
		err = create(&cfg)
	}
	return err
}

func GetOneConfig(code string) (config model.Config, err error) {
	err = orm().Model(&model.Config{}).Equal("code", code).First(&config)
	return
}
