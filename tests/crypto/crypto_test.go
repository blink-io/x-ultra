package crypto

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestRSA_1(t *testing.T) {
	privkey, err := rsa.GenerateKey(rand.Reader, 2048)
	require.NoError(t, err)

	data := x509.MarshalPKCS1PrivateKey(privkey)
	dataB64 := base64.StdEncoding.EncodeToString(data)

	pemData := pem.EncodeToMemory(&pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: data,
	})

	fmt.Println(dataB64)
	fmt.Println(string(pemData))
}
