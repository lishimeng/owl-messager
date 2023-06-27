package sms

import (
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"
	uuid "github.com/iris-contrib/go.uuid"
	"github.com/lishimeng/go-log"
	"github.com/lishimeng/owl-messager/internal/db/model"
	"github.com/lishimeng/owl-messager/internal/messager"
	"github.com/lishimeng/owl-messager/internal/util"
	"net/url"
	"strings"
	"time"
)

// 华为云SMS

type HuaweiSdk struct {
	config model.HuaweiSmsConfig
	client util.Rest
}

const (
	huaweiHost       = "https://smsapi.cn-north-4.myhuaweicloud.com:443/sms/batchSendSms/v1"
	WsseHeaderFormat = "UsernameToken Username=\"%s\",PasswordDigest=\"%s\",Nonce=\"%s\",Created=\"%s\""
	AuthHeaderValue  = "WSSE realm=\"SDP\",profile=\"UsernameToken\",type=\"Appkey\""
)

func NewHuawei(conf model.HuaweiSmsConfig) (sdk messager.SmsProvider) {

	h := &HuaweiSdk{
		config: conf,
		client: util.New(),
	}
	sdk = h
	return
}

func (sdk *HuaweiSdk) Send(message messager.Request) (resp messager.Response, err error) {
	signature := sdk.config.SignName // TODO 暂时只支持一个账号唯一签名
	receiver := message.Receivers
	statusCallBack := ""
	var m = make(map[string]interface{})
	err = json.Unmarshal([]byte(message.Params), &m)
	if err != nil {
		// TODO
		return
	}
	params := map2array(m)
	templateParas := buildHuaweiTemplateParams(params)

	body := sdk.buildRequestBody(sdk.config.Sender, receiver, message.Template, templateParas, statusCallBack, signature)

	headers := sdk.buildHeader()
	err = sdk._send(body, headers)
	return
}

func (sdk *HuaweiSdk) buildTemplateParam() (s string) {
	return
}

func (sdk *HuaweiSdk) buildWsseHeader() string {
	var header = ""
	var appKey = sdk.config.AppId
	var appSecret = sdk.config.AppKey
	var cTime = time.Now().Format("2006-01-02T15:04:05Z")
	var nonceUUID, err = uuid.NewV4()
	var nonce = nonceUUID.String()
	if err != nil {
		return header
	}
	nonce = strings.ReplaceAll(nonce, "-", "")

	h := sha256.New()
	h.Write([]byte(nonce + cTime + appSecret))
	passwordDigestBase64Str := base64.StdEncoding.EncodeToString(h.Sum(nil))

	header = fmt.Sprintf(WsseHeaderFormat, appKey, passwordDigestBase64Str, nonce, cTime)
	return header
}

func (sdk *HuaweiSdk) buildHeader() map[string]string {
	var headers = make(map[string]string)
	headers["Content-Type"] = "application/x-www-form-urlencoded"
	headers["Authorization"] = AuthHeaderValue
	headers["X-WSSE"] = sdk.buildWsseHeader()
	return headers
}

func (sdk *HuaweiSdk) buildRequestBody(sender, receiver, templateId, templateParas, statusCallBack, signature string) map[string]string {
	var params = make(map[string]string)
	params["from"] = url.QueryEscape(sender)
	params["to"] = url.QueryEscape(receiver)
	params["templateId"] = url.QueryEscape(templateId)
	params["templateParas"] = url.QueryEscape(templateParas)
	params["statusCallback"] = url.QueryEscape(statusCallBack)
	params["signature"] = url.QueryEscape(signature)
	return params
}

func (sdk *HuaweiSdk) _send(data map[string]string, headers map[string]string) (err error) {
	code, body, err := sdk.client.Form(huaweiHost, data, headers)
	if err != nil {
		return
	}
	log.Debug(code)
	log.Debug(body)
	return
}
