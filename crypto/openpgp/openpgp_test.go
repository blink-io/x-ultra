package openpgp

import (
	"fmt"
	"testing"

	"github.com/ProtonMail/gopenpgp/v2/crypto"
	"github.com/ProtonMail/gopenpgp/v2/helper"
	"github.com/stretchr/testify/require"
)

func TestGenKey(t *testing.T) {
	const (
		name    = "Max Mustermann"
		email   = "max.mustermann@example.com"
		rsaBits = 2048
	)
	//passphrase := []byte("LongSecret")

	// Curve25519, string
	ecKey, err := crypto.GenerateKey(name, email, "x25519", 0)
	ecpubKey, err := ecKey.GetArmoredPublicKey()
	require.NoError(t, err)
	ecpriKey, err := ecKey.Armor()
	require.NoError(t, err)

	fmt.Println("ecpubKey: ", ecpubKey)
	fmt.Println("ecpriKey: ", ecpriKey)
}

func TestC1(t *testing.T) {
	fmt.Println(crypto.GetUnixTime())

	pwd := "hello, world，我是密码"
	var password = []byte(pwd)

	plain := "my message"
	// Encrypt data with password
	armor, err := helper.EncryptMessageWithPassword(password, plain)
	require.NoError(t, err)

	// Decrypt data with password
	message, err := helper.DecryptMessageWithPassword(password, armor)
	require.NoError(t, err)

	require.Equal(t, plain, message)
}
