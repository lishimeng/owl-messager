package ms

import (
	"crypto"
	"crypto/x509"
	"encoding/json"
	"errors"
	"github.com/lishimeng/owl-messager/internal/certificate"
	"github.com/lishimeng/owl-messager/internal/db/model"
)

func New(config string) (p *AzureGraphProvider, err error) {

	var cc model.GraphConfig
	err = json.Unmarshal([]byte(config), &cc)
	if err != nil {
		err = errors.New("config err")
		return
	}

	var c AzureAuthConfig
	c.Sender = cc.Sender
	c.Tenant = cc.Tenant
	c.ClientId = cc.ClientId
	c.Scope = cc.Scope

	var crt *x509.Certificate
	h := certificate.PemHandler{Pem: cc.Certificate}
	crt, err = h.ParseCrt()
	if err != nil {
		return
	}
	c.Certificate = append(c.Certificate, crt)

	var certKey crypto.PrivateKey
	h = certificate.PemHandler{Pem: cc.CertificateKey}
	certKey, err = h.ParseKey()
	if err != nil {
		return
	}
	c.CertificateKey = certKey

	p = &AzureGraphProvider{Config: c}
	return
}
