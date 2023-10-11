package utils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/lishimeng/app-starter/tool"
	"github.com/lishimeng/go-log"
	"github.com/pkg/errors"
	"io"
	"net/http"
	"net/url"
	"time"
)

var DebugEnable bool

type RestClient struct {
	host     string
	proxy    *http.Client
	headers  map[string]string
	path     []string
	query    map[string]string
	jsonResp any
}

func init() {
	DebugEnable = false
}

func NewRest(host string) *RestClient {
	rc := &RestClient{host: host}
	rc.proxy = &http.Client{Timeout: 8 * time.Second}
	rc.headers = make(map[string]string)
	return rc
}

func (rc *RestClient) build(method string, body []byte) (code int, err error) {
	var requestUri string
	var request *http.Request

	log.Debug("path:", rc.path)
	requestUri, err = url.JoinPath(rc.host, rc.path...) // 拼接uri
	if len(rc.query) > 0 {                              // 拼接query部分
		var content = ""
		for key, value := range rc.query {
			content = content + fmt.Sprintf("%s=%s&", key, value)
		}
		requestUri = fmt.Sprintf("%s?%s", requestUri, content)
	}
	if DebugEnable {
		log.Debug("request: [%s]%s", method, requestUri)
	}

	request, err = http.NewRequest(method, requestUri, bytes.NewBuffer(body))
	if err != nil {
		err = errors.Wrap(err, "create http request err")
		if DebugEnable {
			log.Debug(err)
		}
		return
	}
	for key, value := range rc.headers { // 设置header
		if DebugEnable {
			log.Debug("header: [%s:%s]", key, value)
		}
		request.Header.Set(key, value)
	}

	if DebugEnable {
		log.Debug("start http request...")
	}
	resp, err := rc.proxy.Do(request)
	if err != nil {
		err = errors.Wrap(err, "client Post err")
		if DebugEnable {
			log.Debug(err)
		}
		return
	}
	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(resp.Body)

	code = resp.StatusCode
	if DebugEnable {
		log.Debug("response code: %d", code)
	}
	if code != 200 { // http不成功
		return
	}

	if rc.jsonResp != nil {
		result, _ := io.ReadAll(resp.Body)
		err = json.Unmarshal(result, rc.jsonResp)
		if err != nil {
			err = errors.Wrap(err, "response json unmarshal err")
			if DebugEnable {
				log.Debug(err)
			}
			return
		}
	}

	return
}

func (rc *RestClient) Auth(token string) *RestClient {
	rc.headers[tool.AuthHeader] = tool.Realm + token
	return rc
}

func (rc *RestClient) Path(name ...string) *RestClient {
	rc.path = append(rc.path, name...)
	return rc
}

func (rc *RestClient) Header(name, value string) *RestClient {
	rc.headers[name] = value
	return rc
}

func (rc *RestClient) Post(payload any) (code int, err error) {
	var body []byte
	if payload != nil {
		rc.headers["Content-Type"] = "application/json"
		body, err = json.Marshal(payload)
		if err != nil {
			return
		}
	}
	code, err = rc.build("POST", body)

	return
}

func (rc *RestClient) Form(req map[string]string) (code int, err error) {
	var content = ""
	for key, value := range req {
		content = content + fmt.Sprintf("%s=%s&", key, value)
	}
	code, err = rc.build("POST", []byte(content))
	return
}

func (rc *RestClient) FormUrl(req map[string]string) (code int, err error) {
	rc.query = req
	code, err = rc.build("POST", []byte(""))
	return
}

func (rc *RestClient) Get(query map[string]string) (code int, err error) {
	rc.query = query
	code, err = rc.build("GET", []byte(""))
	return
}

func (rc *RestClient) ResponseJson(respPtr any) *RestClient {
	rc.jsonResp = respPtr
	return rc
}
