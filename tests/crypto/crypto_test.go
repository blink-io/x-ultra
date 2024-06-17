package crypto

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"fmt"
	"testing"

	"github.com/blink-io/x/misc/random"
	"github.com/stretchr/testify/require"
	"golang.org/x/crypto/pbkdf2"
)

func TestRSA_1(t *testing.T) {
	privkey, err := rsa.GenerateKey(rand.Reader, 2048)
	require.NoError(t, err)

	data := x509.MarshalPKCS1PrivateKey(privkey)
	//dataB64 := base64.StdEncoding.EncodeToString(data)
	pemData := pem.EncodeToMemory(&pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: data,
	})
	//fmt.Println(dataB64)
	fmt.Println(string(pemData))

	data2, err2 := x509.MarshalPKCS8PrivateKey(privkey)
	require.NoError(t, err2)
	pemData2 := pem.EncodeToMemory(&pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: data2,
	})
	fmt.Println(string(pemData2))
}

type MyStrSlices []string

func TestKKK(t *testing.T) {
	mss := MyStrSlices([]string{"1", "2", "3", "4"})
	for _, i := range mss {
		fmt.Println("i: ", i)
	}
}

func TestPBKDF_2(t *testing.T) {
	salt := random.String(64)
	pwd := "Hello"

	data := pbkdf2.Key([]byte(pwd), []byte(salt), 4096, 32, sha256.New)
	data2 := pbkdf2.Key([]byte(pwd), []byte(salt), 4096, 32, sha256.New)
	dataB64 := base64.StdEncoding.EncodeToString(data)
	data2B64 := base64.StdEncoding.EncodeToString(data2)

	fmt.Println("len: ", len(data), ", :----> ", dataB64)
	fmt.Println("len: ", len(data2), ", :----> ", data2B64)
}
