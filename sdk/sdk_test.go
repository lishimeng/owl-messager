package sdk

import "testing"

func TestSdk(t *testing.T) {
	params := make(map[string]string)
	params["content"] = "ABC123"
	resp, err := New(WithHost("http://localhost/api"),
		WithDebug(true),
		WithAuth("aewfvsfvadv", "bhnsasdvdzvdvs")).SendMail(MailRequest{
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
	t.Logf("Response code:%v , %+v", resp.Code, resp)
}
