package sdk

import (
	"encoding/json"
	"github.com/lishimeng/owl-messager/pkg/msg"
	"testing"
)

const (
	appId  = "aewfvsfvadv"
	secret = "bhnsasdvdzvdvs"
)

func TestSdk(t *testing.T) {
	Debug(true)
	params := make(map[string]string)
	params["content"] = "ABC123"
	resp, err := New(WithHost("http://localhost/"),
		WithAuth("aewfvsfvadv1", "bhnsasdvdzvdvs")).SendMail(MailRequest{
		Template:      "tpl_test001", // 测试模板
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

func TestSendMail001(t *testing.T) {
	Debug(true)
	params := make(map[string]string)
	params["code"] = "ABC123"
	resp, err := New(WithHost("http://localhost/"),
		WithAuth(appId, secret)).SendMail(MailRequest{
		Template:      "tl_mail_bb5aa5904acabd1fbceda57009152c95", // 测试模板
		TemplateParam: params,
		Title:         "验证码",
		Receiver:      "alex@thingple.com", // 收件人邮箱
	})
	if err != nil {
		t.Fatal(err)
		return
	}
	bs, _ := json.Marshal(resp)
	t.Logf("Response %s", string(bs))
}

func TestTemplateList(t *testing.T) {
	Debug(true)
	resp, err := New(WithHost("http://localhost/"),
		WithAuth(appId, secret),
	).Templates(TemplateRequest{
		Category: msg.SmsMessage,
	})
	if err != nil {
		t.Fatal(err)
		return
	}
	t.Log(resp)
}
