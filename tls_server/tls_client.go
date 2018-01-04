package main

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"golang-examples/tls_server/gen_cert"
	"io/ioutil"
	"log"
	"net/http"

	"k8s.io/client-go/util/cert"
)

const (
	caCertPath     = "/tmp/ca.crt"
	caKeyPath      = "/tmp/ca.key"
	clientCertPath = "/tmp/client.crt"
	clientKeyPath  = "/tmp/client.key"
)

func main() {
	// generate client cert and key signed by CA
	opt := gen_cert.Options{
		CACertPath: caCertPath,
		CAKeyPath:  caKeyPath,
		Config: cert.Config{
			CommonName: "client",
			Usages:     []x509.ExtKeyUsage{x509.ExtKeyUsageAny},
		},
	}
	if err := opt.Generate(clientCertPath, clientKeyPath); err != nil {
		log.Fatal(err)
	}

	caCertPool, err := cert.NewPool(caCertPath)
	if err != nil {
		log.Fatal(err)
	}

	cert, err := tls.LoadX509KeyPair(clientCertPath, clientKeyPath)
	if err != nil {
		log.Fatal(err)
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
		log.Fatal(err)
	}

	htmlData, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	fmt.Printf("%v\n", resp.Status)
	fmt.Printf(string(htmlData))
}
