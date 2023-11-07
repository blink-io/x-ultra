package testdata

import (
	"crypto/tls"
	"crypto/x509"
	"log"
)

func CreateClientTLSConfig() *tls.Config {
	pool, err := x509.SystemCertPool()
	if err != nil {
		log.Fatal(err)
	}
	AddRootCA(pool)

	tlsConf := &tls.Config{
		RootCAs:            pool,
		InsecureSkipVerify: true,
		//KeyLogWriter:       keyLog,
		MinVersion: tls.VersionTLS13,
	}
	return tlsConf
}
