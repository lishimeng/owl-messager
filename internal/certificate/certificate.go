package certificate

import (
	"crypto"
	"crypto/tls"
	"crypto/x509"
	"encoding/pem"
	"errors"
	"strings"
)

var (
	ErrFailedToDecryptKey       = errors.New("failed to decrypt private key")
	ErrFailedToParsePrivateKey  = errors.New("failed to parse private key")
	ErrFailedToParseCertificate = errors.New("failed to parse certificate PEM data")
	ErrNoPrivateKey             = errors.New("no private key")
	ErrNoCertificate            = errors.New("no certificate")
)

func FromPemBytes(bytes []byte, password string) (tls.Certificate, error) {
	var cert tls.Certificate
	var block *pem.Block
	for {
		block, bytes = pem.Decode(bytes)
		if block == nil {
			break
		}
		if block.Type == "CERTIFICATE" {
			cert.Certificate = append(cert.Certificate, block.Bytes)
		}
		if strings.HasSuffix(block.Type, "PRIVATE KEY") {
			key, err := unencryptPrivateKey(block)
			if err != nil {
				return tls.Certificate{}, err
			}
			cert.PrivateKey = key
		}
	}
	if len(cert.Certificate) == 0 {
		return tls.Certificate{}, ErrNoCertificate
	}
	if cert.PrivateKey == nil {
		return tls.Certificate{}, ErrNoPrivateKey
	}
	if c, e := x509.ParseCertificate(cert.Certificate[0]); e == nil {
		cert.Leaf = c
	}
	return cert, nil
}

func unencryptPrivateKey(block *pem.Block) (crypto.PrivateKey, error) {
	return ParsePrivateKey(block.Bytes)
}

func ParsePrivateKey(bytes []byte) (key crypto.PrivateKey, err error) {
	key, err = x509.ParsePKCS8PrivateKey(bytes)
	if err == nil {
		return
	}
	return
}

func ParseCertificateFromPem(pemContent []byte) (crt *x509.Certificate, err error) {
	var rest = pemContent
	var block *pem.Block
	block, rest = pem.Decode(rest)
	crt, err = x509.ParseCertificate(block.Bytes)
	return
}

func ParsePrivateKeyFromPem(pemContent []byte) (crt crypto.PrivateKey, err error) {
	var rest = pemContent
	var block *pem.Block
	block, rest = pem.Decode(rest)
	crt, err = x509.ParsePKCS8PrivateKey(block.Bytes)
	return
}

func ParseCertificatesFromPem(pemContent []byte) (crt *x509.Certificate, err error) {

	var rest = pemContent
	var block *pem.Block
	var blocks []*pem.Block
	for {
		block, rest = pem.Decode(rest)
		if block == nil {
			break
		} else {
			blocks = append(blocks, block)
		}
	}

	for _, b := range blocks {
		if strings.Contains(b.Type, "CERTIFICATE") {
			crt, err = x509.ParseCertificate(b.Bytes)
		} else if strings.Contains(b.Type, "KEY") {

		}
	}

	return
}
