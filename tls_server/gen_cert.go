package tls

import (
	"crypto/rsa"
	"crypto/x509"
	"log"

	"k8s.io/client-go/util/cert"
)

type CertGenerator struct {
	SelSigned  bool
	CACertPath string // required when selfSigned false
	CAKeyPath  string // required when selfSigned false

	OutputCertPath string
	OutputKeyPath  string

	Config cert.Config
}

func (op CertGenerator) Generate() error {
	var (
		newCert *x509.Certificate
		newKey  *rsa.PrivateKey
		err     error
	)

	log.Printf("Creating certs with options: %#v\n", op)

	// generate and write key
	if newKey, err = cert.NewPrivateKey(); err != nil {
		return err
	} else if err = cert.WriteKey(op.OutputKeyPath, cert.EncodePrivateKeyPEM(newKey)); err != nil {
		return err
	}

	// generate and write cert
	if op.SelSigned {
		if newCert, err = cert.NewSelfSignedCACert(op.Config, newKey); err != nil {
			return err
		}
	} else {
		if caCert, caKey, err := loadCertPair(op.CACertPath, op.CAKeyPath); err != nil {
			return err
		} else if newCert, err = cert.NewSignedCert(op.Config, newKey, caCert, caKey); err != nil {
			return err
		}
	}
	return cert.WriteCert(op.OutputCertPath, cert.EncodeCertPEM(newCert))
}

func loadCertPair(certPath, keyPath string) (*x509.Certificate, *rsa.PrivateKey, error) {
	certs, err := cert.CertsFromFile(certPath)
	if err != nil {
		return nil, nil, err
	}
	key, err := cert.PrivateKeyFromFile(keyPath)
	if err != nil {
		return nil, nil, err
	}
	return certs[0], key.(*rsa.PrivateKey), err
}
