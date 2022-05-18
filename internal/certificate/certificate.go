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
			key, err := unencryptPrivateKey(block, password)
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

func unencryptPrivateKey(block *pem.Block, password string) (crypto.PrivateKey, error) {
	if x509.IsEncryptedPEMBlock(block) {
		bytes, err := x509.DecryptPEMBlock(block, []byte(password))
		if err != nil {
			return nil, ErrFailedToDecryptKey
		}
		return parsePrivateKey(bytes)
	}
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
