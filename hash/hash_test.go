package hash

import (
	"encoding/base64"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSIP_1(t *testing.T) {
	key := []byte("world is under the dark")
	data := []byte("门前冷落鞍马息")
	sumB64StrExp := "NlmigqXyYyw="

	sum := SIP(key, data)
	sumB64Str := base64.StdEncoding.EncodeToString(sum)
	require.Equal(t, sumB64StrExp, sumB64Str)
}
