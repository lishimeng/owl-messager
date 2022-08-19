package ms

import (
	"crypto"
	"crypto/x509"
	"errors"
	"github.com/lishimeng/owl/internal/certificate"
)

const (
	clientIdKey    = "clientId"
	tenantKey      = "tenant"
	scopeKey       = "scope"
	certificateKey = "certificate"
	privateKey     = "certificateKey"
	senderKey      = "sender"
)

func New(config map[string]string) (p *AzureGraphProvider, err error) {

	var c AzureAuthConfig

	if sender, ok := config[senderKey]; ok && len(sender) > 0 {
		c.Sender = sender
	} else {
		err = errors.New("sender nil")
		return
	}

	if tenant, ok := config[tenantKey]; ok && len(tenant) > 0 {
		c.Tenant = tenant
	} else {
		err = errors.New("tenant nil")
		return
	}

	if clientId, ok := config[clientIdKey]; ok && len(clientId) > 0 {
		c.ClientId = clientId
	} else {
		err = errors.New("client nil")
		return
	}

	if scope, ok := config[scopeKey]; ok && len(scope) > 0 {
		c.Scope = scope
	} else {
		err = errors.New("scope nil")
		return
	}

	if certPem, ok := config[certificateKey]; ok && len(certPem) > 0 {

		var crt *x509.Certificate
		h := certificate.PemHandler{Pem: certPem}
		crt, err = h.ParseCrt()
		if err != nil {
			return
		}
		c.Certificate = append(c.Certificate, crt)
	} else {
		err = errors.New("cert nil")
		return
	}

	if certKeyPem, ok := config[privateKey]; ok && len(certKeyPem) > 0 {

		var certKey crypto.PrivateKey
		h := certificate.PemHandler{Pem: certKeyPem}
		certKey, err = h.ParseKey()
		if err != nil {
			return
		}
		c.CertificateKey = certKey
	} else {
		err = errors.New("tenant nil")
		return
	}

	p = &AzureGraphProvider{Config: c}
	return
}
