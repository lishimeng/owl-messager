package fastmsg

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/lishimeng/go-log"
	"github.com/lishimeng/owl-messager/internal/messager"
	"github.com/lishimeng/owl-messager/internal/provider/template"
	"github.com/lishimeng/owl-messager/pkg/msg"
	"net/http"
	"net/url"
	"resty.dev/v3"
	"strings"
)

// fast_msg

type SDK interface {
}

type FastMsgSdkImpl struct {
	config    *msg.FastMsgConfig
	sessionId string
	SDK
}

type AuthParams struct {
	Name  string
	Value string
}

func New(config msg.FastMsgConfig) (sdk *FastMsgSdkImpl, err error) {

	var oto FastMsgSdkImpl
	oto.config = &config

	// 检查host合法性
	if !strings.HasPrefix(config.Host, "http") {
		err = errors.New("invalid host")
		return
	}
	sdk = &oto
	return
}

func (sdk *FastMsgSdkImpl) Auth() (err error) {
	log.Info("login...")
	const action = "chatserver/api/user/auth"
	if sdk.config == nil {
		err = errors.New("nil FastMsgConfig")
		return
	}
	resource, err := url.JoinPath(sdk.config.Host, action)
	if err != nil {
		return
	}
	var req = make(map[string]string)

	req["token"] = sdk.config.Token
	req["username"] = sdk.config.Username
	req["password"] = sdk.config.Password
	var resp = make(map[string]any)
	err = request(resource, req, "application/x-www-form-urlencoded", &resp)
	if err != nil {
		log.Info(err)
		return
	}
	sessionId, err := getStrField(resp, "sessionid")
	if err != nil {
		log.Info(err)
		return
	}
	sdk.sessionId = sessionId
	log.Info("session: %s", sdk.sessionId)
	return
}

func (sdk *FastMsgSdkImpl) Send(req messager.Request) (err error) {
	content, err := template.Rend(req.Params, req.Template.Body)
	if err != nil {
		return
	}

	receivers := strings.Split(req.Receivers, ",")
	for _, part := range receivers {
		receiver := strings.TrimSpace(part)
		if receiver == "" {
			continue
		}
		err = sdk.SendRobotMessage(receiver, content)
		if err != nil {
			return
		}
	}

	return
}

func (sdk *FastMsgSdkImpl) SendRobotMessage(receiver string, content string) (err error) {
	const action = "chatserver/api/message/new"
	err = sdk.Auth() // 每次登录

	if err != nil {
		return
	}
	resource, err := url.JoinPath(sdk.config.Host, action)
	if err != nil {
		return
	}
	var resp = make(map[string]any)
	var params = make(map[string]any)
	params["username"] = receiver
	params["content"] = content // TODO urlencoding时是否需要转换, 以支持空格/UTF-8字符
	params["msgtype"] = "0"     // 需要是 map[string]string
	err = request(resource, params, "application/x-www-form-urlencoded", &resp, AuthParams{
		Name:  "session-id",
		Value: sdk.sessionId,
	})
	if err != nil {
		return
	}
	errorCode, err := getStrField(resp, "errorcode")
	if err != nil {
		log.Info(err)
		return
	}
	if errorCode != "OK" {
		log.Info("fast_err: %s", errorCode)
		return errors.New("fast_msg error: " + errorCode)
	}
	msgId, err := getMapValue(resp, "msgid")
	if err != nil {
		log.Info(err)
		return
	}
	log.Info("fast_msg: ", msgId)
	bs, err := json.Marshal(resp)
	if err != nil {
		log.Info(err)
		return
	}
	log.Info(string(bs))
	return
}

func request(url string, params any, contentType string, respPtr any, auth ...AuthParams) (err error) {
	var tmp = strings.ToLower(contentType)
	var code int

	switch tmp {
	case "application/x-www-form-urlencoded":
		var data map[string]string
		var resp *resty.Response
		data, err = json2Map(params)
		if err != nil {
			return
		}
		var r = resty.New().R()

		if len(auth) > 0 {
			//r = r.AuthCookie(auth[0].Name, auth[0].Value)
			r = r.SetCookies([]*http.Cookie{{
				Name:  auth[0].Name,
				Value: auth[0].Value,
			}})
		}
		//code, body, err = r.Form(url, data, nil)
		resp, err = r.SetFormData(data).
			SetResult(respPtr).
			Post(url)
		if err != nil {
			return
		}
		code = resp.StatusCode()
		if code != 200 {
			err = errors.New(fmt.Sprintf("invalid response code %d", code))
		}
		//err = json.Unmarshal([]byte(body), respPtr)
	default:
		err = errors.New("invalid content type:" + tmp)
		return
	}
	return err
}

func json2Map(p any) (m map[string]string, err error) {
	bs, err := json.Marshal(p)
	if err != nil {
		return
	}
	m = make(map[string]string)
	err = json.Unmarshal(bs, &m)
	return
}

func getMapValue(m map[string]any, key string) (v any, err error) {
	if m == nil {
		err = errors.New("nil map")
		return
	}
	v, ok := m[key]
	if !ok {
		err = errors.New("key not found: " + key)
		return
	}
	return
}

func getStrField(m map[string]any, key string) (str string, err error) {
	v, err := getMapValue(m, key)
	if err != nil {
		return
	}
	str, ok := v.(string)
	if !ok {
		err = errors.New("value is not string: " + key)
		return
	}
	return
}

func getNumField(m map[string]any, key string) (num int64, err error) {
	v, err := getMapValue(m, key)
	if err != nil {
		return
	}
	num, ok := v.(int64)
	if !ok {
		err = errors.New("value is not number: " + key)
		return
	}
	return
}
