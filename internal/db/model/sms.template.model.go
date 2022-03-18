package model

import "encoding/json"

type SmsParam map[string]string

type SmsTemplateInfo struct {
	Pk
	Code             string   `orm:"column(code);unique"`        // owl中的唯一编码
	Name             string   `orm:"column(name)"`               // 模板名字
	Sender           int      `orm:"column(sender_id)"`          // 发送平台
	Body             string   `orm:"column(body)"`               // 发送的内容主体，可空
	SenderTemplateId string   `orm:"column(sender_template_id)"` // 发送平台的模板ID
	Signature        string   `orm:"column(signature)"`          // 在发送平台上预留的签名，根据不同平台保存签名文本或签名ID
	Description      string   `orm:"column(description);null"`
	Params           string   `orm:"column(params)"` // json:key--data_type
	paramList        SmsParam // 参数列表，不进入数据库
	TableChangeInfo
}

func (t *SmsTemplateInfo) AddParam(name string, dataType string) {
	if t.paramList == nil {
		t.paramList = make(map[string]string)
	}
	t.paramList[name] = dataType
}

func (t *SmsTemplateInfo) Build() {

	bs, err := json.Marshal(t)
	if err != nil {
		// TODO
		return
	}
	t.Params = string(bs)
}

func (t *SmsTemplateInfo) UnmarshallParam() {

	if t.paramList == nil {
		t.paramList = make(map[string]string)
	}
	err := json.Unmarshal([]byte(t.Params), &t.paramList)
	if err != nil {
		// TODO
	}
}

const (
	SmsTemplateEnable  = 1 // enable
	SmsTemplateDisable = 0 // disable
)
