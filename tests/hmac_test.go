package tests

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestHMAC_1(t *testing.T) {
	str := "Hello, world, 人生需要点行动力"
	key := "password"
	macval := CreateMAC([]byte(str), []byte(key))

	fmt.Println("MAC Value in base64: ", base64.StdEncoding.EncodeToString(macval), "--- len: ", len(macval))

	ok := ValidMAC([]byte(str), macval, []byte(key))
	require.Equal(t, true, ok)
}

func ValidMAC(message, messageMAC, key []byte) bool {
	mac := hmac.New(sha256.New, key)
	mac.Write(message)
	expectedMAC := mac.Sum(nil)
	return hmac.Equal(messageMAC, expectedMAC)
}

func CreateMAC(message, key []byte) []byte {
	mac := hmac.New(sha256.New, key)
	mac.Write(message)
	macval := mac.Sum(nil)
	return macval
}
