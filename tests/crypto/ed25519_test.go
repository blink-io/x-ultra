package crypto

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
	"golang.org/x/crypto/ed25519"
)

func TestEd25519_1(t *testing.T) {
	pubkey, prikey, err := ed25519.GenerateKey(rand.Reader)
	require.NoError(t, err)

	plain := "Hello, World"
	plainData := []byte(plain)
	signedData := ed25519.Sign(prikey, plainData)

	fmt.Println("prikey: ", base64.StdEncoding.EncodeToString(prikey))
	fmt.Println("pubkey: ", base64.StdEncoding.EncodeToString(pubkey))

	fmt.Println("signed data in base64: ", base64.StdEncoding.EncodeToString(signedData))

	valid := ed25519.Verify(pubkey, plainData, signedData)
	require.Equal(t, true, valid)
	if valid {
		fmt.Println("Your passcode is valid")
	} else {
		fmt.Println("Oh, WTH...")
	}
}

func TestEcdsa_1(t *testing.T) {
	privkey, err := ecdsa.GenerateKey(elliptic.P384(), rand.Reader)
	require.NoError(t, err)
	require.NotNil(t, privkey)

}
