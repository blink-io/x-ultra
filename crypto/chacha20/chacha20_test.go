package chacha20

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestChacha20_1(t *testing.T) {
	keyb := []byte("adsfdsaenc,.dsafadafd14454646dasfasdf90789689dasfdas113232")
	sskey := [32]byte(keyb)

	plan := []byte("hello")
	edata, err := Encrypt(sskey, plan)
	require.NoError(t, err)

	ddata, err := Decrypt(sskey, edata)
	require.NoError(t, err)

	fmt.Println("Decrypted:  ", string(ddata))
}
