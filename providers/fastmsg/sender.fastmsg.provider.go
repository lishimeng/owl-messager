package fastmsg

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/baidubce/bce-sdk-go/util/log"
	"github.com/lishimeng/owl-messager/internal/messager"
	"github.com/lishimeng/owl-messager/internal/provider/template"
	"github.com/lishimeng/owl-messager/internal/util"
	"github.com/lishimeng/owl-messager/pkg/msg"
	"path"
	"strings"
)

// fast_msg

type SDK interface {
}

type fastMsgSdkImpl struct {
	config    *msg.FastMsgConfig
	sessionId string
}

func New(config msg.FastMsgConfig) (sdk messager.ImProvider, err error) {

	var oto fastMsgSdkImpl
	oto.config = &config

	// 检查host合法性
	if !strings.HasPrefix(config.Host, "http") {
		err = errors.New("invalid host")
		return
	}
	sdk = &oto
	return
}

func (sdk *fastMsgSdkImpl) Auth() (err error) {
	var resource = joinPath(sdk.config.Host, "chatserver/api/user/auth")
	var req = make(map[string]string)
	if sdk.config == nil {
		err = errors.New("nil FastMsgConfig")
		return
	}
	req["token"] = sdk.config.Token
	req["username"] = sdk.config.Username
	req["password"] = sdk.config.Password
	var resp = make(map[string]string)
	err = request(resource, req, "application/x-www-form-urlencoded", &resp)
	if err != nil {
		return
	}
	if sessionId, ok := resp["sessionid"]; ok {
		sdk.sessionId = sessionId
	} else {
		err = errors.New("no sessionid")
	}
	return
}

func (sdk *fastMsgSdkImpl) SendRobotMessage(receiver string, content string) (err error) {
	err = sdk.Auth() // 每次登录

	if err != nil {
		return
	}
	var resource = joinPath(sdk.config.Host, "chatserver/api/message/new")
	var resp = make(map[string]string)
	var params = make(map[string]any)
	// TODO session的使用
	params["username"] = receiver
	params["content"] = content // TODO urlencoding时是否需要转换, 以支持空格/UTF-8字符
	params["msgtype"] = "0"     // 需要是 map[string]string
	err = request(resource, params, "application/x-www-form-urlencoded", &resp)
	if err != nil {
		return
	}
	if errorCode, ok := resp["errorcode"]; ok {
		if errorCode != "OK" {
			log.Info("fast_err: %s", errorCode)
			return errors.New("fast_msg error: " + errorCode)
		}
	}
	if msgId, ok := resp["msgid"]; ok {
		log.Info("fast_msg: %s", msgId)
	}
	return
}

func (sdk *fastMsgSdkImpl) Send(req messager.Request) (err error) {
	receiver := req.Receivers
	content, err := template.Rend(req.Params, req.Template.Body)
	if err != nil {
		return
	}
	return sdk.SendRobotMessage(receiver, content)
}

func request(url string, params any, contentType string, respPtr any) (err error) {
	var tmp = strings.ToLower(contentType)
	var code int

	switch tmp {
	case "application/x-www-form-urlencoded":
		var data map[string]string
		var body string
		data, err = json2Map(params)
		if err != nil {
			return
		}
		code, body, err = util.New().Form(url, data, nil)
		if err != nil {
			return
		}
		if code != 200 {
			err = errors.New(fmt.Sprintf("invalid response code %d", code))
		}
		err = json.Unmarshal([]byte(body), respPtr)
	case "application/json":
		code, err = util.New().PostJson(url, params, respPtr)
		if err != nil {
			return
		}
		if code != 200 {
			err = errors.New(fmt.Sprintf("invalid response code %d", code))
		}
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

func joinPath(host string, p string) string {
	index := strings.Index(host, "://")
	if index == -1 {
		return path.Join(host, p)
	}
	protocol := host[:index+3]
	rest := host[index+3:]
	return protocol + path.Join(rest, p)
}
