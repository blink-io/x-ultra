package tls

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"os"
)

var (
	ErrEmptyTLSConfig = fmt.Errorf("empty TLS config")
)

func CreateBaseTLSConfig(caFile string, caOptional bool, certFile, keyFile string, insecureSkipVerify bool) (*tls.Config, error) {
	certPool := x509.NewCertPool()
	clientAuth := tls.NoClientCert

	caFileBytes, err := ReadMaybeFile(caFile)
	if err != nil {
		return nil, fmt.Errorf("unable to read CA file from %s", caFile)
	}
	certFileBytes, err := ReadMaybeFile(certFile)
	if err != nil {
		return nil, fmt.Errorf("unable to read CA file from %s", certFile)
	}
	keyFileBytes, err := ReadMaybeFile(keyFile)
	if err != nil {
		return nil, fmt.Errorf("unable to read CA file from %s", keyFile)
	}

	if len(caFileBytes) > 0 {
		if !certPool.AppendCertsFromPEM(caFileBytes) {
			return nil, fmt.Errorf("unable to parse CA file")
		}

		if caOptional {
			clientAuth = tls.VerifyClientCertIfGiven
		} else {
			clientAuth = tls.RequireAndVerifyClientCert
		}
	}

	if !insecureSkipVerify && (len(certFileBytes) == 0 || len(keyFileBytes) == 0) {
		return nil, fmt.Errorf("TLS Certificate or Key file must be set when TLS configuration is created")
	}

	cert := tls.Certificate{}
	if len(certFileBytes) > 0 && len(keyFileBytes) > 0 {
		var err error
		cert, err = tls.X509KeyPair(certFileBytes, keyFileBytes)
		if err != nil {
			return nil, fmt.Errorf("failed to load TLS keypair: %v", err)
		}
	}

	return &tls.Config{
		Certificates:       []tls.Certificate{cert},
		RootCAs:            certPool,
		InsecureSkipVerify: insecureSkipVerify,
		ClientAuth:         clientAuth,
	}, nil
}

func CreateClientTLSConfig(caFile string, caOptional bool, certFile, keyFile string, insecureSkipVerify bool) (*tls.Config, error) {
	c, err := CreateBaseTLSConfig(
		caFile,
		caOptional,
		certFile,
		keyFile,
		insecureSkipVerify,
	)
	return c, err
}

func CreateServerTLSConfig(caFile string, caOptional bool, certFile, keyFile, serverName string, insecureSkipVerify bool) (*tls.Config, error) {
	c, err := CreateBaseTLSConfig(
		caFile,
		caOptional,
		certFile,
		keyFile,
		insecureSkipVerify,
	)
	if err == nil {
		c.ServerName = serverName
	}
	return c, err
}

func ReadMaybeFile(maybeFile string) ([]byte, error) {
	var data []byte
	if _, errKey := os.Stat(maybeFile); errKey == nil {
		var err error
		data, err = os.ReadFile(maybeFile)
		if err != nil {
			return nil, err
		}
	} else {
		data = []byte(maybeFile)
	}
	return data, nil
}
