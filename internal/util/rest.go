package util

import (
	"crypto/tls"
	"encoding/json"
	"github.com/go-resty/resty/v2"
)

type Rest interface {
	Get(uri string) (code int, body string, err error)
	Form(url string, data map[string]string, headers map[string]string) (code int, body string, err error)
	Post(uri string) (code int, body string, err error)
	PostJson(uri string, body interface{}) (code int, err error)
}

type RestHandler struct {
	proxy *resty.Client
}

func New() (r Rest) {

	h := RestHandler{proxy: resty.New()}
	h.proxy.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	r = &h
	return
}

func (r *RestHandler) Form(url string, data map[string]string, headers map[string]string) (code int, body string, err error) {
	req := r.proxy.NewRequest()
	resp, err := req.SetFormData(data).SetHeaders(headers).Post(url)
	if err != nil {
		return
	}
	code = resp.StatusCode()
	txt := resp.Body()
	body = string(txt)
	return
}

func (r *RestHandler) Get(uri string) (code int, body string, err error) {

	resp, err := r.proxy.R().Get(uri)
	if err != nil {
		return
	}

	code = resp.StatusCode()
	bodyBs := resp.Body()
	body = string(bodyBs)
	return
}

func (r *RestHandler) GetJson(uri string, body interface{}) (code int, err error) {

	resp, err := r.proxy.R().Get(uri)
	if err != nil {
		return
	}

	code = resp.StatusCode()
	txt := resp.Body()
	err = json.Unmarshal(txt, body)
	return
}

func (r *RestHandler) Post(uri string) (code int, body string, err error) {
	resp, err := r.proxy.R().Post(uri)
	if err != nil {
		return
	}
	code = resp.StatusCode()
	body = string(resp.Body())
	return
}

func (r *RestHandler) PostJson(uri string, body interface{}) (code int, err error) {
	resp, err := r.proxy.R().Post(uri)
	if err != nil {
		return
	}
	code = resp.StatusCode()
	txt := resp.Body()
	err = json.Unmarshal(txt, body)
	return
}
