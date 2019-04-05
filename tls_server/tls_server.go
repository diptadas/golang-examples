package main

import (
	"crypto/tls"
	"crypto/x509"
	"log"
	"net/http"

	"github.com/diptadas/golang-examples/tls_server/gen_cert"
	"k8s.io/client-go/util/cert"
)

const (
	caCertPath     = "/tmp/ca.crt"
	caKeyPath      = "/tmp/ca.key"
	serverCertPath = "/tmp/server.crt"
	serverKeyPath  = "/tmp/server.key"
)

func main() {
	// generate CA cert and key
	opt := gen_cert.Options{
		SelSigned:      true,
		OutputCertPath: caCertPath,
		OutputKeyPath:  caKeyPath,
		Config:         cert.Config{},
	}
	if err := opt.Generate(); err != nil {
		log.Fatal(err)
	}

	// generate server cert and key signed by CA
	opt = gen_cert.Options{
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
		log.Fatal(err)
	}

	caCertPool, err := cert.NewPool(caCertPath)
	if err != nil {
		log.Fatal(err)
	}

	server := &http.Server{
		Addr:    ":8443",
		Handler: &handler{},
		TLSConfig: &tls.Config{
			ClientAuth: tls.RequireAndVerifyClientCert,
			ClientCAs:  caCertPool,
		},
	}

	log.Println("TLS server running at port 8443")
	log.Fatal(server.ListenAndServeTLS(serverCertPath, serverKeyPath))
}

type handler struct{}

func (h *handler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	log.Println(req)
	w.Write([]byte("PONG\n"))
}
