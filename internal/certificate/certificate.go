package certificate

import (
	"crypto"
	"crypto/tls"
	"crypto/x509"
	"encoding/pem"
	"errors"
	"golang.org/x/crypto/pkcs12"
	"strings"
)

var (
	ErrFailedToDecryptKey       = errors.New("failed to decrypt private key")
	ErrFailedToParsePrivateKey  = errors.New("failed to parse private key")
	ErrFailedToParseCertificate = errors.New("failed to parse certificate PEM data")
	ErrNoPrivateKey             = errors.New("no private key")
	ErrNoCertificate            = errors.New("no certificate")
)

func FromP12(p12 []byte, password string) (tls.Certificate, error) {
	key, cert, err := pkcs12.Decode(p12, password)
	if err != nil {
		return tls.Certificate{}, err
	}
	return tls.Certificate{
		Certificate: [][]byte{cert.Raw},
		PrivateKey:  key,
		Leaf:        cert,
	}, nil
}

func FromPemBytes(bytes []byte) (tls.Certificate, error) {
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

func CertFromPem(pemBytes []byte, password string) (certs []*x509.Certificate, err error) {
	var cert tls.Certificate
	var block *pem.Block
	for {
		block, _ = pem.Decode(pemBytes)
		if block == nil {
			break
		}
		if block.Type == "CERTIFICATE" {
			cert.Certificate = append(cert.Certificate, block.Bytes)
		}
		if strings.HasSuffix(block.Type, "PRIVATE KEY") {
			key, err := unencryptPrivateKey(block)
			if err != nil {
				return
			}
			cert.PrivateKey = key
		}
	}

	if len(cert.Certificate) == 0 {
		err = ErrNoCertificate
		return
	}
	if cert.PrivateKey == nil {
		err = ErrNoPrivateKey
		return
	}
	for _, bs := range cert.Certificate {
		var c *x509.Certificate
		c, err = x509.ParseCertificate(bs)
		if err != nil {
			return
		}
		certs = append(certs, c)
	}
	return
}

func unencryptPrivateKey(block *pem.Block) (crypto.PrivateKey, error) {
	return parsePrivateKey(block.Bytes)
}

func parsePrivateKey(bytes []byte) (crypto.PrivateKey, error) {
	var key crypto.PrivateKey
	key, err := x509.ParsePKCS1PrivateKey(bytes)
	if err == nil {
		return key, nil
	}
	key, err = x509.ParsePKCS8PrivateKey(bytes)
	if err == nil {
		return key, nil
	}
	return nil, ErrFailedToParsePrivateKey
}
