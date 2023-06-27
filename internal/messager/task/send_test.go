package task

import (
	"crypto/tls"
	"github.com/go-gomail/gomail"
	"testing"
)

func TestSend(t *testing.T) {
	m := gomail.NewMessage()

	t.Log("start")
	// 收件人
	m.SetHeader("To", "ryker@thingple.com")

	// 第三个参数为发件人别名，如"李大锤"，可以为空(此时则为邮箱名称)
	m.SetAddressHeader("From", "noreply@thingplecloud.com", "ThingpleCloud")

	// -----------------------------------
	// 主题
	m.SetHeader("Subject", "测试mail")
	// 正文
	m.SetBody("text/html", "<html><head><meta charset=\"utf-8\"></head><body>\t<h3>警告：以下设备库存不足或超限！</h3>\t<p>{{ .content }}</p></body></html>")

	d := gomail.NewDialer("smtp.mxhichina.com", 465, "noreply@thingplecloud.com", "N7oreply")
	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}
	// 发送
	err := d.DialAndSend(m)

	if err !=nil{
		t.Log("errrrrrrorrrrr")
		t.Log(err)
	}
	t.Log("success")
}
