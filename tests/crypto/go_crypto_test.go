package crypto

import (
	"fmt"
	"testing"

	"github.com/go-crypt/crypt"
	"github.com/go-crypt/crypt/algorithm"
	"github.com/go-crypt/crypt/algorithm/argon2"
	"github.com/stretchr/testify/require"
)

func TestGoCrypt_1(t *testing.T) {
	var (
		decoder *crypt.Decoder
		err     error
		digest  algorithm.Digest
	)

	decoder, err = NewDecoderArgon2idOnly()
	require.NoError(t, err)

	digest, err = decoder.Decode("$argon2id$v=19$m=2097152,t=1,p=4$BjVeoTI4ntTQc0WkFQdLWg$OAUnkkyx5STI0Ixl+OSpv4JnI6J1TYWKuCuvIbUGHTY")
	require.NoError(t, err)

	fmt.Printf("Digest Matches Password 'example': %t\n", digest.Match("example"))
	fmt.Printf("Digest Matches Password 'invalid': %t\n", digest.Match("invalid"))
}

// NewDecoderArgon2idOnly returns a decoder which can only decode argon2id encoded digests.
func NewDecoderArgon2idOnly() (decoder *crypt.Decoder, err error) {
	decoder = crypt.NewDecoder()

	if err = argon2.RegisterDecoderArgon2id(decoder); err != nil {
		return nil, err
	}

	return decoder, nil
}
