package sdk

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/lishimeng/go-log"
	"github.com/lishimeng/owl-messager/cmd/owl-messager/ddd/open"
	"github.com/pkg/errors"
	"io"
	"net/http"
)

func getCredential(host string, appId, secret string) (response open.CredentialResp, err error) {

	client := &http.Client{}
	req := open.CredentialReq{
		AppId:  appId,
		Secret: secret,
	}
	bs, err := json.Marshal(req)
	if err != nil {
		return
	}
	resp, err := client.Post(host, "application/json", bytes.NewBuffer(bs))
	if err != nil {
		return
	}
	defer func() {
		_ = resp.Body.Close()
	}()
	if resp.StatusCode != CodeSuccess {
		err = errors.New(fmt.Sprintf("%d", resp.StatusCode))
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
	return
}
