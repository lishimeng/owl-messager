package certificate

import (
	"crypto"
	"crypto/x509"
	"encoding/pem"
	"errors"
	"github.com/lishimeng/go-log"
)

type PemHandler struct {
	Pem string
}

const (
	KeyTypePKCS1 = "PKCS1"
	KeyTypePKCS8 = "PKCS8"
	KeyTypeEC    = "EC"
)

func (h *PemHandler) ParseCrt() (crt *x509.Certificate, err error) {
	b, err := h.decodePem([]byte(h.Pem))
	if err != nil {
		return
	}
	crt, err = x509.ParseCertificate(b.Bytes)
	if err != nil {
		log.Info(err)
		return
	}
	return
}

func (h *PemHandler) ParseKey() (key crypto.PrivateKey, err error) {
	key, err = h.parseKey("")
	if err != nil {
		log.Info(err)
		return
	}
	return
}

func (h *PemHandler) parseKey(category string) (key crypto.PrivateKey, err error) {

	b, err := h.decodePem([]byte(h.Pem))
	if err != nil {
		return
	}

	switch category {
	case KeyTypePKCS1:
		key, err = x509.ParsePKCS1PrivateKey(b.Bytes)
	case KeyTypePKCS8:
		key, err = x509.ParsePKCS8PrivateKey(b.Bytes)
	case KeyTypeEC:
		key, err = x509.ParseECPrivateKey(b.Bytes)
	default:
		err = errors.New("unknown private key category")
	}

	if err != nil { // 尝试pkcs1
		key, err = x509.ParsePKCS1PrivateKey(b.Bytes)
	}
	if err != nil { // 尝试pkcs8
		key, err = x509.ParsePKCS8PrivateKey(b.Bytes)
	}
	if err != nil { // 尝试ec
		key, err = x509.ParseECPrivateKey(b.Bytes)
	}

	return
}

func (h *PemHandler) decodePem(content []byte) (block *pem.Block, err error) {
	block, _ = pem.Decode(content)
	if block == nil {
		err = errors.New("der nil")
		log.Info(err)
		return
	}
	return
}
