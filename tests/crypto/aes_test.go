package crypto

import (
	"crypto/aes"
	"encoding/base64"
	"fmt"
	"testing"

	"github.com/matryer/is"
	"github.com/stretchr/testify/require"
)

func TestAES_1(t *testing.T) {
	is := is.New(t)

	//b := rand.Bytes(32)
	key := []byte("4x68zjqhg9lqlxbr4q8wjm7g79mtx6sw")
	c, err := aes.NewCipher(key)
	require.NoError(t, err)
	is.NoErr(err)

	plain := "Hello,World, OK,123456789"
	in := []byte(plain)
	out := make([]byte, len(plain))

	c.Encrypt(out, in)

	rein := make([]byte, len(out))
	c.Decrypt(rein, out)

	fmt.Println("out  : ", base64.StdEncoding.EncodeToString(out))
	fmt.Println("re in: ", string(rein))
}
