package main

import (
	"crypto/x509"
	"flag"
	"log"
	"net"
	"strings"

	"github.com/diptadas/golang-examples/tls_server/gen_cert"
	"k8s.io/client-go/util/cert"
)

func main() {
	var (
		reuseCA      = false
		caCommonName = "CA"
		caCertPath   = "ca.crt"
		caKeyPath    = "ca.key"
		certPath     = "tls.crt"
		keyPath      = "tls.key"
		domains      = ""
		ips          = ""
	)

	flag.BoolVar(&reuseCA, "reuse-ca", false, "Use existing CA (ca.crt and ca.key)")
	flag.StringVar(&domains, "domains", "", "Alternative Domain names")
	flag.StringVar(&ips, "ips", "127.0.0.1", "Alternative IP addresses")
	flag.Parse()

	if !reuseCA { // create new CA certs first
		opt := gen_cert.Options{
			SelSigned:      true,
			OutputCertPath: caCertPath,
			OutputKeyPath:  caKeyPath,
			Config: cert.Config{
				CommonName: caCommonName,
				Usages:     []x509.ExtKeyUsage{x509.ExtKeyUsageAny},
			},
		}
		if err := opt.Generate(); err != nil {
			log.Fatal(err)
		}
	}

	// sign new certs using CA
	opt := gen_cert.Options{
		CACertPath:     caCertPath,
		CAKeyPath:      caKeyPath,
		OutputCertPath: certPath,
		OutputKeyPath:  keyPath,
		Config: cert.Config{
			CommonName: parseCommonName(domains),
			AltNames:   parseAltNames(domains, ips),
			Usages:     []x509.ExtKeyUsage{x509.ExtKeyUsageAny},
		},
	}
	if err := opt.Generate(); err != nil {
		log.Fatal(err)
	}

	log.Println("Successfully created certs")
}

func parseCommonName(domains string) string {
	commonName := strings.Split(domains, ",")[0]
	if commonName == "" {
		commonName = "server"
	}
	return commonName
}

func parseAltNames(domains, ips string) cert.AltNames {
	altNames := cert.AltNames{}
	altNames.DNSNames = strings.Split(domains, ",")
	for _, ip := range strings.Split(ips, ",") {
		altNames.IPs = append(altNames.IPs, net.ParseIP(ip))
	}
	return altNames
}
