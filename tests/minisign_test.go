package tests

import (
	"crypto/rand"
	"fmt"
	"testing"

	"aead.dev/minisign"
	"github.com/stretchr/testify/require"
)

func TestMinisign_1(t *testing.T) {
	pubkey, privkey, err := minisign.GenerateKey(rand.Reader)
	require.NoError(t, err)
	require.NotNil(t, privkey)

	fmt.Println("pubkey : ", pubkey.String())
}
