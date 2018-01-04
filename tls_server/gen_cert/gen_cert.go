package gen_cert

import (
	"crypto/rsa"
	"crypto/x509"
	"k8s.io/client-go/util/cert"
)

func GenerateCA(caCertPath, caKeyPath string) error {
	// check CA already exists
	if ok, err := cert.CanReadCertAndKey(caCertPath, caKeyPath); err == nil && ok {
		return nil
	}

	// generate CA key
	caKey, err := cert.NewPrivateKey()
	if err != nil {
		return err
	}
	if err := cert.WriteKey(caKeyPath, cert.EncodePrivateKeyPEM(caKey)); err != nil {
		return err
	}

	// generate CA cert
	caCert, err := cert.NewSelfSignedCACert(
		cert.Config{
			CommonName:   "ac",
			Organization: []string{"ac"},
		},
		caKey,
	)
	if err != nil {
		return err
	}
	if err := cert.WriteCert(caCertPath, cert.EncodeCertPEM(caCert)); err != nil {
		return err
	}

	return nil
}

func GenerateCertKey(certPath, keyPath, caCertPath, caKeyPath string) error {
	if err := GenerateCA(caCertPath, caKeyPath); err != nil {
		return err
	}
	caCerts, err := cert.CertsFromFile(caCertPath)
	if err != nil {
		return err
	}
	caKey, err := cert.PrivateKeyFromFile(caKeyPath)
	if err != nil {
		return err
	}

	// generate user key
	userKey, err := cert.NewPrivateKey()
	if err != nil {
		return err
	}
	if err := cert.WriteKey(keyPath, cert.EncodePrivateKeyPEM(userKey)); err != nil {
		return err
	}

	// generate user cert signed by CA
	userCert, err := cert.NewSignedCert(
		cert.Config{
			CommonName:   "ac",
			Organization: []string{"ac"},
			AltNames:     cert.AltNames{
				DNSNames: []string{"localhost"},
			},
			Usages:       []x509.ExtKeyUsage{x509.ExtKeyUsageAny},
		},
		userKey,
		caCerts[0],
		caKey.(*rsa.PrivateKey),
	)
	if err != nil {
		return err
	}
	if err := cert.WriteCert(certPath, cert.EncodeCertPEM(userCert)); err != nil {
		return err
	}

	return nil
}
