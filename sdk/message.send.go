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

func _send(credential string, url string, req []byte) (code int, response Response, err error) {
	client := &http.Client{Timeout: 8 * time.Second}
	request, err := http.NewRequest("POST", url, bytes.NewBuffer(req))
	if err != nil {
		log.Debug(errors.Wrap(err, "create http request err"))
		return
	}
	request.Header.Set(tool.AuthHeader, tool.Realm+credential)
	request.Header.Set("Content-Type", "application/json")
	resp, err := client.Do(request)
	if err != nil {
		log.Debug(errors.Wrap(err, "client Post err"))
		return
	}
	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(resp.Body)

	if resp.StatusCode == CodeNotAllow {
		code = CodeNotAllow
		return
	}
	result, _ := io.ReadAll(resp.Body)
	err = json.Unmarshal(result, &response)
	if err != nil {
		log.Debug(errors.Wrap(err, "response json unmarshal err"))
		return
	}
	if response.Code == float64(CodeNotAllow) {
		code = CodeNotAllow
		return
	}
	return
}
