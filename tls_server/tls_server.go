package main

import (
	"crypto/tls"
	"log"
	"net/http"
	"golang-examples/tls_server/gen_cert"
	"k8s.io/client-go/util/cert"
)

func main() {
	if err := gen_cert.GenerateCertKey("server.crt", "server.key", "ca.crt", "ca.key"); err != nil {
		panic(err)
	}

	caCertPool, err := cert.NewPool("ca.crt")
	if err != nil {
		log.Fatal(err)
	}

	srv := &http.Server{
		Addr:    ":8443",
		Handler: &handler{},
		TLSConfig: &tls.Config{
			ClientAuth: tls.RequireAndVerifyClientCert,
			ClientCAs:  caCertPool,
		},
	}

	log.Println("TLS server running at port 8443")
	log.Fatal(srv.ListenAndServeTLS("server.crt", "server.key"))
}

type handler struct{}

func (h *handler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	log.Println(req)
	w.Write([]byte("PONG\n"))
}
