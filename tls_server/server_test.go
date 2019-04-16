package tls

import (
	"crypto/tls"
	"crypto/x509"
	"log"
	"net/http"
	"testing"

	"k8s.io/client-go/util/cert"
)

func TestTLSServer(t *testing.T) {
	var (
		caCertPath     = "/tmp/ca.crt"
		caKeyPath      = "/tmp/ca.key"
		serverCertPath = "/tmp/server.crt"
		serverKeyPath  = "/tmp/server.key"
	)

	// generate CA cert and key
	opt := CertGenerator{
		SelSigned:      true,
		OutputCertPath: caCertPath,
		OutputKeyPath:  caKeyPath,
		Config:         cert.Config{},
	}
	if err := opt.Generate(); err != nil {
		t.Fatal(err)
	}

	// generate server cert and key signed by CA
	opt = CertGenerator{
		CACertPath:     caCertPath,
		CAKeyPath:      caKeyPath,
		OutputCertPath: serverCertPath,
		OutputKeyPath:  serverKeyPath,
		Config: cert.Config{
			CommonName: "server",
			AltNames: cert.AltNames{
				DNSNames: []string{"localhost"},
			},
			Usages: []x509.ExtKeyUsage{x509.ExtKeyUsageAny},
		},
	}
	if err := opt.Generate(); err != nil {
		t.Fatal(err)
	}

	caCertPool, err := cert.NewPool(caCertPath)
	if err != nil {
		t.Fatal(err)
	}

	server := &http.Server{
		Addr:    ":8443",
		Handler: &handler{},
		TLSConfig: &tls.Config{
			ClientAuth: tls.RequireAndVerifyClientCert,
			ClientCAs:  caCertPool,
		},
	}

	t.Logf("TLS server running at port 8443")
	t.Fatal(server.ListenAndServeTLS(serverCertPath, serverKeyPath))
}

type handler struct{}

func (h *handler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	log.Println(req)
	w.Write([]byte("PONG\n"))
}
