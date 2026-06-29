package util

import (
	"crypto/tls"
	"io/ioutil"
	"resty.dev/v3"
)

type Rest interface {
	Get(uri string) (code int, body string, err error)
	Form(url string, data map[string]string, headers map[string]string) (code int, body string, err error)
	Post(uri string) (code int, body string, err error)
	PostJson(uri string, params any, respPtr any) (code int, err error)
}

type RestHandler struct {
	proxy *resty.Client
}

func New() (r Rest) {

	h := RestHandler{proxy: resty.New()}
	h.proxy.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: false})
	r = &h
	return
}

func (r *RestHandler) Form(url string, data map[string]string, headers map[string]string) (code int, body string, err error) {
	req := r.proxy.NewRequest()
	req = req.SetDebug(true)
	resp, err := req.SetFormData(data).SetHeaders(headers).Post(url)
	if err != nil {
		return
	}
	code = resp.StatusCode()
	defer resp.Body.Close()
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}
	body = string(bodyBytes)
	return
}

func (r *RestHandler) FormUrlEncoded(url string, data map[string]string, headers map[string]string) (code int, body string, err error) {
	req := r.proxy.NewRequest()
	resp, err := req.SetFormData(data).SetHeaders(headers).Post(url)
	if err != nil {
		return
	}
	code = resp.StatusCode()
	defer resp.Body.Close()
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}
	body = string(bodyBytes)
	return
}

func (r *RestHandler) Get(uri string) (code int, body string, err error) {

	resp, err := r.proxy.R().Get(uri)
	if err != nil {
		return
	}

	code = resp.StatusCode()
	defer resp.Body.Close()
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}
	body = string(bodyBytes)
	return
}

func (r *RestHandler) GetJson(uri string, body interface{}) (code int, err error) {

	resp, err := r.proxy.R().
		SetResult(body).
		Get(uri)
	if err != nil {
		return
	}

	code = resp.StatusCode()
	return
}

func (r *RestHandler) Post(uri string) (code int, body string, err error) {
	resp, err := r.proxy.R().Post(uri)
	if err != nil {
		return
	}
	code = resp.StatusCode()
	defer resp.Body.Close()
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}
	body = string(bodyBytes)
	return
}

func (r *RestHandler) PostJson(uri string, params any, respPtr any) (code int, err error) {
	var request = r.proxy.R()
	if params != nil {
		request = request.SetBody(params)
	}
	if respPtr != nil {
		request = request.SetResult(respPtr)
	}
	resp, err := request.Post(uri)
	if err != nil {
		return
	}
	code = resp.StatusCode()
	return
}
