package sdk

import (
	"bytes"
	"encoding/json"
	"github.com/lishimeng/app-starter/tool"
	"github.com/lishimeng/go-log"
	"github.com/pkg/errors"
	"io"
	"net/http"
	"time"
)

// err不空说明执行失败了, err为nil 再检查code 和 response
func _send(credential string, url string, req []byte) (code int, response Response, err error) {
	client := &http.Client{Timeout: 8 * time.Second}
	request, err := http.NewRequest("POST", url, bytes.NewBuffer(req))
	if err != nil {
		err = errors.Wrap(err, "create http request err")
		if debugEnable {
			log.Debug(err)
		}
		return
	}
	request.Header.Set(tool.AuthHeader, tool.Realm+credential)
	request.Header.Set("Content-Type", "application/json")
	resp, err := client.Do(request)
	if err != nil {
		err = errors.Wrap(err, "client Post err")
		if debugEnable {
			log.Debug(err)
		}
		return
	}
	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(resp.Body)

	if resp.StatusCode != CodeSuccess { // http不成功
		code = resp.StatusCode
		return
	}

	result, _ := io.ReadAll(resp.Body)
	err = json.Unmarshal(result, &response)
	if err != nil {
		err = errors.Wrap(err, "response json unmarshal err")
		if debugEnable {
			log.Debug(err)
		}
		return
	}
	if response.Code == float64(CodeNotAllow) {
		code = CodeNotAllow
		return
	}
	return
}
