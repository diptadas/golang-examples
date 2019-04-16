package tls

import (
	"crypto/x509"
	"flag"
	"net"
	"strings"
	"testing"

	"k8s.io/client-go/util/cert"
)

func TestGenCerts(t *testing.T) {
	var (
		caCertPath = "certs/ca.crt"
		caKeyPath  = "certs/ca.key"
		certPath   = "certs/tls.crt"
		keyPath    = "certs/tls.key"

		reuseCA      = false
		caCommonName = "CA"
		domains      = ""
		ips          = ""
	)

	flag.BoolVar(&reuseCA, "reuse-ca", false, "Use existing CA (ca.crt and ca.key)")
	flag.StringVar(&domains, "domains", "", "Alternative Domain names")
	flag.StringVar(&ips, "ips", "127.0.0.1", "Alternative IP addresses")
	flag.Parse()

	if !reuseCA { // create new CA certs first
		opt := CertGenerator{
			SelSigned:      true,
			OutputCertPath: caCertPath,
			OutputKeyPath:  caKeyPath,
			Config: cert.Config{
				CommonName: caCommonName,
				Usages:     []x509.ExtKeyUsage{x509.ExtKeyUsageAny},
			},
		}
		if err := opt.Generate(); err != nil {
			t.Fatal(err)
		}
	}

	// sign new certs using CA
	opt := CertGenerator{
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
		t.Fatal(err)
	}

	t.Logf("Successfully created certs")
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
