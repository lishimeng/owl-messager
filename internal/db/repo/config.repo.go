package repo

import (
	"encoding/base64"
	"encoding/json"

	"github.com/lishimeng/go-log"
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

func DecodeConfigContent(encoded string, dest interface{}) error {
	js, err := base64.StdEncoding.DecodeString(encoded)
	if err != nil {
		return err
	}
	return json.Unmarshal(js, dest)
}

// LoadTaskChannelSettings reads task channel config from DB; missing row returns zero value (memqueue).
func LoadTaskChannelSettings() (settings model.TaskChannelSettings) {
	cfg, err := GetOneConfig(model.ConfigCodeTaskChannel)
	if err != nil {
		return
	}
	if err = DecodeConfigContent(cfg.Content, &settings); err != nil {
		log.Debug("decode task channel config: %v", err)
	}
	return
}

// LoadConsoleTokenSettings reads console management token from config table.
func LoadConsoleTokenSettings() (settings model.ConsoleTokenSettings) {
	cfg, err := GetOneConfig(model.ConfigCodeConsoleToken)
	if err != nil {
		return
	}
	if err = DecodeConfigContent(cfg.Content, &settings); err != nil {
		log.Debug("decode console token config: %v", err)
	}
	return
}
