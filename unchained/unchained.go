package unchained

import (
	"errors"
	"strings"

	"github.com/blink-io/x/unchained/argon2"
	"github.com/blink-io/x/unchained/bcrypt"
	"github.com/blink-io/x/unchained/pbkdf2"
)

// hasher defines Django hasher identifiers.
type hasher int

const (
	UnknownHasher = iota
	Argon2Hasher
	BCryptHasher
	BCryptSHA256Hasher
	PBKDF2SHA256Hasher
	PBKDF2SHA512Hasher
)

func (h hasher) String() string {
	switch h {
	case Argon2Hasher:
		return argon2.AlgorithmArgon2
	case BCryptHasher:
		return bcrypt.AlgorithmBCrypt
	case BCryptSHA256Hasher:
		return bcrypt.AlgorithmBCryptSHA256
	case PBKDF2SHA256Hasher:
		return pbkdf2.AlgorithmPBKDF2SHA256
	case PBKDF2SHA512Hasher:
		return pbkdf2.AlgorithmPBKDF2SHA512
	default:
		return ""
	}
}

func hasherToString(h hasher) string {
	return h.String()
}

func hasherFromString(s string) hasher {
	switch s {
	case argon2.AlgorithmArgon2:
		return Argon2Hasher
	case bcrypt.AlgorithmBCrypt:
		return BCryptHasher
	case bcrypt.AlgorithmBCryptSHA256:
		return BCryptSHA256Hasher
	case pbkdf2.AlgorithmPBKDF2SHA256:
		return PBKDF2SHA256Hasher
	case pbkdf2.AlgorithmPBKDF2SHA512:
		return PBKDF2SHA512Hasher
	}
	return UnknownHasher
}

const (
	// UnusablePasswordPrefix is used in unusable passwords.
	UnusablePasswordPrefix = "!"
	// UnusablePasswordSuffixLength defines the length of unusable passwords after the prefix.
	UnusablePasswordSuffixLength = 40
	// DefaultHasher defines the default hasher used in Django.
	DefaultHasher hasher = PBKDF2SHA256Hasher
	// DefaultSaltSize defines the default salt size used in Django.
	DefaultSaltSize = 12
)

var (
	// ErrInvalidHasher is returned if the hasher is invalid or unknown.
	ErrInvalidHasher = errors.New("unchained: invalid hasher")
)

// IsValidHasher returns true if the hasher
// is supported by Django, or false otherwise.
func IsValidHasher(hasher hasher) bool {
	return hasher != UnknownHasher
}

// IdentifyHasher returns the hasher used in the encoded password.
func IdentifyHasher(encoded string) hasher {
	s := strings.SplitN(encoded, "$", 2)[0]
	h := hasherFromString(s)
	return h
}

// IsPasswordUsable returns true if encoded password
// is usable, or false otherwise.
func IsPasswordUsable(encoded string) bool {
	return encoded != "" && !strings.HasPrefix(encoded, UnusablePasswordPrefix)
}

// CheckPassword validates if the raw password matches the encoded digest.
//
// This is a shortcut that discovers the hasher used in the encoded digest
// to perform the correct validation.
func CheckPassword(password, encoded string) (bool, error) {
	if !IsPasswordUsable(encoded) {
		return false, nil
	}

	hasher := IdentifyHasher(encoded)

	if !IsValidHasher(hasher) {
		return false, ErrInvalidHasher
	}

	switch hasher {
	case Argon2Hasher:
		return argon2.NewArgon2Hasher().Verify(password, encoded)
	case BCryptHasher:
		return bcrypt.NewBCryptHasher().Verify(password, encoded)
	case BCryptSHA256Hasher:
		return bcrypt.NewBCryptSHA256Hasher().Verify(password, encoded)
	case PBKDF2SHA256Hasher:
		return pbkdf2.NewPBKDF2SHA256Hasher().Verify(password, encoded)
	}

	return false, ErrInvalidHasher
}

// MakePassword turns a plain-text password into a hash.
//
// If password is empty, then return a concatenation
// of UnusablePasswordPrefix and a random string.
// If salt is empty, then a random string is generated.
// BCrypt algorithm ignores salt parameter.
// If hasher is "default", encode using default hasher.
func MakePassword(password, salt string, hasher hasher) (string, error) {
	if password == "" {
		return UnusablePasswordPrefix + GetRandomString(UnusablePasswordSuffixLength), nil
	}

	if salt == "" {
		salt = GetRandomString(DefaultSaltSize)
	}

	if !IsValidHasher(hasher) {
		return "", ErrInvalidHasher
	}

	switch hasher {
	case Argon2Hasher:
		return argon2.NewArgon2Hasher().Encode(password, salt)
	case BCryptHasher:
		return bcrypt.NewBCryptHasher().Encode(password, salt)
	case BCryptSHA256Hasher:
		return bcrypt.NewBCryptSHA256Hasher().Encode(password, salt)
	case PBKDF2SHA256Hasher:
		return pbkdf2.NewPBKDF2SHA256Hasher().Encode(password, salt, 0)
	}

	return "", ErrInvalidHasher
}
