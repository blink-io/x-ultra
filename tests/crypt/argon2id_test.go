package crypt

import (
	"fmt"
	"testing"

	"github.com/alexedwards/argon2id"
	"github.com/matthewhartstonge/argon2"
	"github.com/stretchr/testify/require"
)

func TestArgon2id(t *testing.T) {
	hash, err := argon2id.CreateHash("pa$$word", argon2id.DefaultParams)
	require.NoError(t, err)
	// ComparePasswordAndHash performs a constant-time comparison between a
	// plain-text password and Argon2id hash, using the parameters and salt
	// contained in the hash. It returns true if they match, otherwise it returns
	// false.
	match, err := argon2id.ComparePasswordAndHash("pa$$word", hash)
	require.NoError(t, err)

	fmt.Printf("Match: %v", match)
}

func TestArgon2(t *testing.T) {
	argon := argon2.DefaultConfig()
	// Waaahht??! It includes magic salt generation for me ! Yasss...
	encoded, err := argon.HashEncoded([]byte("p@ssw0rd"))
	require.NoError(t, err)

	fmt.Println(string(encoded))
	// > $argon2id$v=19$m=65536,t=1,p=4$WXJGqwIB2qd+pRmxMOw9Dg$X4gvR0ZB2DtQoN8vOnJPR2SeFdUhH9TyVzfV98sfWeE

	ok, err := argon2.VerifyEncoded([]byte("p@ssw0rd"), encoded)
	require.NoError(t, err)

	matches := "no 🔒"
	if ok {
		matches = "yes 🔓"
	}
	fmt.Printf("Password Matches: %s\n", matches)
}
