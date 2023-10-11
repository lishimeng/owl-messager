package msg

import (
	"encoding/json"
	"github.com/lishimeng/go-log"
)

// MessageParams 消息参数
type MessageParams map[string][]string

func (mp MessageParams) AddParam(name string, data string) MessageParams {
	var v []string
	v = mp[name]
	v = append(v, data)
	mp[name] = v
	return mp
}

func BuildMessageParam(content string) (mp MessageParams, err error) {
	err = json.Unmarshal([]byte(content), &mp)
	return
}

func (mp MessageParams) Marshal() string {
	b, _ := json.Marshal(mp)
	return string(b)
}

func HandleMessageParams(input string, mapping string) (params map[string]any, err error) {
	var templateParams = make(map[string]any)
	params = make(map[string]any)
	err = json.Unmarshal([]byte(input), &templateParams)
	if err != nil {
		log.Info("params of mail template is not json format:%s", input)
		return
	}

	// TODO 处理参数
	var paraMappings MessageParams
	if len(mapping) > 0 {
		paraMappings, err = BuildMessageParam(mapping)
		if err != nil {
			return
		}
	}

	for k, innerParams := range paraMappings {
		if value, ok := templateParams[k]; ok {
			for _, innerParam := range innerParams {
				params[innerParam] = value
			}
		}
	}
	return
}
