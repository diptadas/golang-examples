package tls

import (
	"crypto/tls"
	"crypto/x509"
	"io/ioutil"
	"net/http"
	"testing"

	"k8s.io/client-go/util/cert"
)

func TestTLSClient(t *testing.T) {
	var (
		caCertPath     = "/tmp/ca.crt"
		caKeyPath      = "/tmp/ca.key"
		clientCertPath = "/tmp/client.crt"
		clientKeyPath  = "/tmp/client.key"
	)

	// generate client cert and key signed by CA
	opt := CertGenerator{
		CACertPath:     caCertPath,
		CAKeyPath:      caKeyPath,
		OutputCertPath: clientCertPath,
		OutputKeyPath:  clientKeyPath,
		Config: cert.Config{
			CommonName: "client",
			Usages:     []x509.ExtKeyUsage{x509.ExtKeyUsageAny},
		},
	}
	if err := opt.Generate(); err != nil {
		t.Fatal(err)
	}

	caCertPool, err := cert.NewPool(caCertPath)
	if err != nil {
		t.Fatal(err)
	}

	cert, err := tls.LoadX509KeyPair(clientCertPath, clientKeyPath)
	if err != nil {
		t.Fatal(err)
	}

	client := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				RootCAs:      caCertPool,
				Certificates: []tls.Certificate{cert},
			},
		},
	}

	resp, err := client.Get("https://localhost:8443")
	if err != nil {
		t.Fatal(err)
	}

	htmlData, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatal(err)
	}
	defer resp.Body.Close()

	t.Logf("%v\n", resp.Status)
	t.Logf(string(htmlData))
}
