package main

import (
	"crypto/tls"
	"fmt"
	"golang-examples/tls_server/gen_cert"
	"io/ioutil"
	"log"
	"net/http"
	"k8s.io/client-go/util/cert"
)

func main() {
	if err := gen_cert.GenerateCertKey("client.crt", "client.key", "ca.crt", "ca.key"); err != nil {
		log.Fatal(err)
	}

	caCertPool, err := cert.NewPool("ca.crt")
	if err != nil {
		log.Fatal(err)
	}
	cert, err := tls.LoadX509KeyPair("client.crt", "client.key")
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
