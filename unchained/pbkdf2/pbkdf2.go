package pbkdf2

import (
	"crypto/hmac"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/base64"
	"errors"
	"fmt"
	"hash"
	"strconv"
	"strings"

	"golang.org/x/crypto/pbkdf2"
)

const (
	AlgorithmPBKDF2SHA256 = "pbkdf2_sha256"
	AlgorithmPBKDF2SHA512 = "pbkdf2_sha512"
)

// Errors returned by PBKDF2Hasher.
var (
	ErrHashComponentUnreadable = errors.New("unchained/pbkdf2: unreadable component in hashed password")
	ErrHashComponentMismatch   = errors.New("unchained/pbkdf2: hashed password components mismatch")
	ErrAlgorithmMismatch       = errors.New("unchained/pbkdf2: algorithm mismatch")
	ErrSaltContainsDollarSing  = errors.New("unchained/pbkdf2: salt contains dollar sign ($)")
)

// PBKDF2Hasher implements PBKDF2 password hasher.
type PBKDF2Hasher struct {
	// Algorithm identifier.
	Algorithm string
	// Defines the number of rounds used to encode the password.
	Iterations int
	// Defines the length of the hash in bytes.
	Size int
	// Defines the hash function used to encode the password.
	Digest func() hash.Hash
}

// Encode turns a plain-text password into a hash.
func (h *PBKDF2Hasher) Encode(password string, salt string, iterations int) (string, error) {
	if strings.Contains(salt, "$") {
		return "", ErrSaltContainsDollarSing
	}

	if iterations <= 0 {
		iterations = h.Iterations
	}

	hash := pbkdf2.Key([]byte(password), []byte(salt), iterations, h.Size, h.Digest)
	b64Hash := base64.StdEncoding.EncodeToString(hash)
	return fmt.Sprintf("%s$%d$%s$%s", h.Algorithm, iterations, salt, b64Hash), nil
}

// Verify if a plain-text password matches the encoded digest.
func (h *PBKDF2Hasher) Verify(password string, encoded string) (bool, error) {
	s := strings.Split(encoded, "$")

	if len(s) != 4 {
		return false, ErrHashComponentMismatch
	}

	algorithm, iterations, salt := s[0], s[1], s[2]

	if algorithm != h.Algorithm {
		return false, ErrAlgorithmMismatch
	}

	i, err := strconv.Atoi(iterations)

	if err != nil {
		return false, ErrHashComponentUnreadable
	}

	newencoded, err := h.Encode(password, salt, i)

	if err != nil {
		return false, err
	}

	return hmac.Equal([]byte(newencoded), []byte(encoded)), nil
}

// NewPBKDF2SHA256Hasher secures password hashing using the PBKDF2 algorithm.
//
// Configured to use PBKDF2 + HMAC + SHA256.
// The result is a 64 byte binary string.
func NewPBKDF2SHA256Hasher() *PBKDF2Hasher {
	return &PBKDF2Hasher{
		Algorithm:  AlgorithmPBKDF2SHA256,
		Iterations: 216000,
		Size:       sha256.Size,
		Digest:     sha256.New,
	}
}

// NewPBKDF2SHA512Hasher secures password hashing using the PBKDF2 algorithm.
//
// Configured to use PBKDF2 + HMAC + SHA512.
// The result is a 64 byte binary string.
func NewPBKDF2SHA512Hasher() *PBKDF2Hasher {
	return &PBKDF2Hasher{
		Algorithm:  AlgorithmPBKDF2SHA512,
		Iterations: 216000,
		Size:       sha512.Size,
		Digest:     sha512.New,
	}
}
