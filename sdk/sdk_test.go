package sdk

import (
	"encoding/json"
	"testing"
)

func TestSdk(t *testing.T) {
	Debug(true)
	params := make(map[string]string)
	params["content"] = "ABC123"
	resp, err := New(WithHost("http://localhost/"),
		WithAuth("aewfvsfvadv1", "bhnsasdvdzvdvs")).SendMail(MailRequest{
		Template:      "tpl_test001", // 测试模板
		CloudTemplate: false,
		TemplateParam: params,
		Title:         "验证码",
		Receiver:      "xxx@qq.com", // 收件人邮箱
	})
	if err != nil {
		t.Fatal(err)
		return
	}
	bs, _ := json.Marshal(resp)
	t.Logf("Response %s", string(bs))
}
