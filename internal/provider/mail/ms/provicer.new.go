package ms

import (
	"crypto"
	"crypto/x509"
	"errors"
)

const (
	clientIdKey  = "clientId"
	tenantKey    = "tenant"
	scopeKey     = "scope"
	certKey      = "certificate"
	certKetKey   = "certificateKey"
	certCategory = "certCategory"
	certPassword = "certPassword"
)

func New(config map[string]string) (p *AzureGraphProvider, err error) {

	// TODO build provider
	var c AzureAuthConfig

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

	if certPem, ok := config[certKey]; ok && len(certPem) > 0 {

		var cert x509.Certificate
		// TODO
		c.Certificate = append(c.Certificate, &cert)
	} else {
		err = errors.New("cert nil")
		return
	}

	if certKeyPem, ok := config[certKetKey]; ok && len(certKeyPem) > 0 {

		var certKey crypto.PrivateKey
		c.CertificateKey = certKey
	} else {
		err = errors.New("tenant nil")
		return
	}

	p = &AzureGraphProvider{Config: c}
	return
}
