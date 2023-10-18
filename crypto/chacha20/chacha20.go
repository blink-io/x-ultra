package chacha20

import (
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"errors"

	"github.com/hashicorp/golang-lru/v2"
	"golang.org/x/crypto/chacha20poly1305"
)

const (
	MaxSize = 100_000
	KeySize = 32
)

var cc *lru.Cache[string, cipher.AEAD]

func init() {
	cc, _ = lru.New[string, cipher.AEAD](MaxSize)
}

func Encrypt(key [KeySize]byte, data []byte) ([]byte, error) {
	aead, err := getOrCreateAEADFromCache(key)
	if err != nil {
		return nil, err
	}

	nonce := make([]byte, aead.NonceSize(), aead.NonceSize()+len(data)+aead.Overhead())
	if _, err := rand.Read(nonce); err != nil {
		return nil, err
	}

	return aead.Seal(nonce, nonce, data, nil), nil
}

func Decrypt(key [KeySize]byte, data []byte) ([]byte, error) {
	aead, err := getOrCreateAEADFromCache(key)
	if err != nil {
		return nil, err
	}
	if len(data) < aead.NonceSize() {
		return nil, errors.New("cipher data is too short")
	}
	// Split nonce and ciphertext.
	nonce, ciphertext := data[:aead.NonceSize()], data[aead.NonceSize():]
	// Decrypt the message and check it wasn't tampered with.
	plain, err := aead.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		return nil, err
	}
	return plain, nil
}

func getOrCreateAEADFromCache(key [KeySize]byte) (cipher.AEAD, error) {
	sskey := key[:]
	skey := base64.StdEncoding.EncodeToString(sskey)
	if aead, ok := cc.Get(skey); ok {
		return aead.(cipher.AEAD), nil
	}
	aead, err := chacha20poly1305.New(sskey)
	if err != nil {
		return nil, err
	}
	_ = cc.Add(skey, aead)
	return aead, nil
}
